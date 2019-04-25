package main

import "design_pattern/bridge"

func main() {
	mm := &bridge.MessageMobile{}
	mq := &bridge.MessageQQ{}
	m1 := &bridge.CommonMessage{AbstractMessage: &bridge.AbstractMessage{IMessage: mm}}
	m2 := &bridge.CommonMessage{AbstractMessage: &bridge.AbstractMessage{IMessage: mq}}
	m1.Send("tom", "hello")
	m2.Send("tom", "hello")
}
