package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf_test/pb"
)

func main() {
	// 创建一个消息 Test
	test := &pb.Test{
		Name: proto.String("webbc"),
		Age:  proto.Uint32(18),
	}

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// 进行解码
	newTest := &pb.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// 测试结果
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

}
