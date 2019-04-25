package bridge

import "fmt"

// 手机消息
type MessageMobile struct {
}

func (mm *MessageMobile) Send(user, msg string) {
	fmt.Println("使用手机发送内容：" + msg + "给" + user)
}
