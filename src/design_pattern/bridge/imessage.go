package bridge

// 消息
type IMessage interface {
	Send(user, msg string)
}
