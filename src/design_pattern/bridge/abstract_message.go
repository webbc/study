package bridge

// 抽象的消息
type AbstractMessage struct {
	IMessage
}

func (am *AbstractMessage) Send(user, msg string) {
	am.IMessage.Send(user, msg)
}
