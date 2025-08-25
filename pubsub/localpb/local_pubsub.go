package localpb

import (
	"context"
	"log"
	"shopfood/common"
	"shopfood/pubsub"
	"sync"
)

// localPubSub chạy cục bộ (in-memory)
// Hệ thống sử dụng một queue (buffer channel) làm lõi và quản lý nhiều nhóm subscribers
// Cho phép gửi message với topic cụ thể đến nhiều subscribers trong cùng một nhóm
type localPubSub struct {
	messageQueue chan *pubsub.Message                    // Channel có buffer để lưu trữ message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message // Map các channels theo topic
	locker       *sync.RWMutex                           // Khóa read-write để đồng bộ hóa
}

// NewPubSub tạo và khởi chạy một instance mới của localPubSub
func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000), // Channel với buffer 10,000 message
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}
	pb.run() // Khởi chạy goroutine xử lý message
	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	// Thiết lập channel/topic cho message
	data.SetChannel(topic)

	// Gửi message vào queue thông qua goroutine để không block caller
	go func() {
		defer common.AppRecover() // Xử lý panic nếu có

		ps.messageQueue <- data // Gửi message vào channel queue

		log.Println("New event published:",
			data.String(),
			"with data",
			data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscriber(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	// Tạo channel mới cho subscriber
	c := make(chan *pubsub.Message)

	// Khóa để đảm bảo thread-safe khi thao tác với map
	ps.locker.Lock()

	// Thêm channel vào danh sách subscribers của topic
	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	// Trả về channel và hàm unsubscribe
	return c, func() {
		log.Println("Unsubscribing from topic:", topic)

		if chans, ok := ps.mapChannel[topic]; ok {
			// Tìm và xóa channel khỏi danh sách
			for i := range chans {
				if chans[i] == c {
					// Xóa element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}

		}
	}
}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess.String())

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
