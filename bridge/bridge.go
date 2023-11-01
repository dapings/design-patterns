package bridge

import (
	"fmt"
)

// AbstractMessage 高层次的抽象接口
type AbstractMessage interface {
	SendMessage(txt, to string)
}

// ConcreteAbstractionMessage 具体的抽象部分
type ConcreteAbstractionMessage struct {
	message MessageImplementor
}

func (c *ConcreteAbstractionMessage) SendMessage(txt, to string) {
	c.message.Send(txt, to)
}

func NewConcreteAbstractionMessage(message MessageImplementor) *ConcreteAbstractionMessage {
	return &ConcreteAbstractionMessage{
		message: message,
	}
}

type UrgencyAbstractionMessage struct {
	message MessageImplementor
}

func (m *UrgencyAbstractionMessage) SendMessage(txt, to string) {
	m.message.Send(fmt.Sprintf("[Urgency] %s", txt), to)
}

func NewUrgencyAbstractionMessage(message MessageImplementor) *UrgencyAbstractionMessage {
	return &UrgencyAbstractionMessage{message: message}
}

// MessageImplementor 低层次的实现接口
type MessageImplementor interface {
	Send(txt, to string)
}

// MessageSMS sms 的具体实现
type MessageSMS struct{}

func UseSMS() MessageImplementor {
	return &MessageSMS{}
}

func (*MessageSMS) Send(txt, to string) {
	fmt.Printf("send %s to %s by sms", txt, to)
}

// MessageEmail email 的具体实现
type MessageEmail struct{}

func UseEmail() MessageImplementor {
	return &MessageEmail{}
}

func (*MessageEmail) Send(txt, to string) {
	fmt.Printf("send %s to %s by email", txt, to)
}
