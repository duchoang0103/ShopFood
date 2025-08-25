package pubsub

import (
	"fmt"
	"time"
)

// Message định nghĩa cấu trúc dữ liệu cho một message trong hệ thống
type Message struct {
	id        string      // ID duy nhất của message
	channel   Topic       // Channel/Topic mà message thuộc về (có thể bỏ qua)
	data      interface{} // Dữ liệu message (kiểu động)
	createdAt time.Time   // Thời điểm tạo message
}

// NewMessage tạo một message mới với dữ liệu được cung cấp
func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()), // ID là timestamp nano giây
		data:      data,
		createdAt: now,
	}
}

// String triển khai Stringer interface để in thông tin message
func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channel)
}

// Channel trả về topic/channel của message
func (evt *Message) Channel() Topic {
	return evt.channel
}

// SetChannel thiết lập topic/channel cho message
func (evt *Message) SetChannel(channel Topic) {
	evt.channel = channel
}

// Data trả về dữ liệu của message
func (evt *Message) Data() interface{} {
	return evt.data
}
