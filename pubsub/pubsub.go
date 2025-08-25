package pubsub

import "context"

type Topic string

// Publish là interface cho hệ thống publish-subscribe
type Pubsub interface {
	// Publish gửi message đến một channel cụ thể
	Publish(ctx context.Context, channel Topic, data *Message) error

	// Subscribe đăng ký nhận message từ một channel
	// Trả về:
	// - channel để nhận message (<-chan *Message)
	// - hàm close để hủy đăng ký
	Subscriber(ctx context.Context, channel Topic) (ch <-chan *Message, close func())

	// UnSubscribe đã bị comment lại, có thể sẽ được thêm sau
	// UnSubscribe(ctx context.Context, channel Channel) error
}
