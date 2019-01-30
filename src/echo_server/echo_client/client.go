package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8080")//连接到服务端
	//错误处理
	if err != nil {
		fmt.Println("连接服务器失败")
	}
	for{
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()//从控制台中读取一行数据
		if err != nil{
			fmt.Println("读取失败")
		}
		_,err = conn.Write([]byte(line))//将数据传递给服务端
		//错误处理
		if err != nil{
			fmt.Println("客户端写入数据失败")
		}
		var buffer = make([]byte,1024)//设置缓存
		_,err = conn.Read(buffer)//从服务端读取数据
		//错误处理
		if err != nil{
			fmt.Println("客户端读取数据失败")
		}
		fmt.Println(string(buffer))//打印从服务器传来的数据
	}

}
