package bridge

// 普通的消息
type CommonMessage struct {
	*AbstractMessage
}

func (cm *CommonMessage) Send(user, msg string) {
	cm.AbstractMessage.Send(user, msg)
}
