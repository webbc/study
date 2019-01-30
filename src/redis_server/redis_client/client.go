package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main()  {
	conn,err := net.Dial("tcp","localhost:8080")//连接服务端
	//错误处理
	if err != nil{
		fmt.Println("连接服务器失败")
	}
	//客户端永不断开
	for{
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()//从控制台中读取一行数据
		//错误处理
		if err != nil{
			fmt.Println("读取数据失败")
		}
		conn.Write(line)//将数据写给服务端

		var buffer = make([]byte,1024)//创建缓存
		_,err = conn.Read(buffer)//从服务端读取数据并写入缓存
		//错误处理
		if err!= nil{
			fmt.Println("读取数据错误")
		}
		var value = string (buffer)
		fmt.Println(value)//打印服务端传来的数据
	}
}












