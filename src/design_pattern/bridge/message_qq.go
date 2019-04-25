package bridge

import "fmt"

// QQ消息
type MessageQQ struct {
}

func (mq *MessageQQ) Send(user, msg string) {
	fmt.Println("使用QQ发送内容：" + msg + "给" + user)
}
