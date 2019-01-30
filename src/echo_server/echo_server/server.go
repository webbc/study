package main

import(
	"net"
	"fmt"
)
func main(){
	startServer()
}

func startServer(){
	ln, err := net.Listen("tcp", ":8080")//绑定端口监听器
	//错误处理
	if err != nil{
		fmt.Println("服务器开启失败")
		return
	}
	for {
		conn, err := ln.Accept()//获取客户端连接
		//错误处理
		if err != nil {
			fmt.Println("客户端连接失败")
			continue
		}
		go handleConnection(conn)//启动线程为客户端服务
	}
}

//处理连接
func handleConnection(conn net.Conn)  {
	for{
		var buffer = make([]byte,1024)//创建缓存区
		_,err := conn.Read(buffer)//从客户端读取数据
		//错误处理
		if err != nil{
			fmt.Println("读取数据失败")
		}
		//没有发送数据
		if buffer == nil{
			continue
		}
		_,err = conn.Write(buffer)//将数据写回给客户端
		//错误处理
		if err != nil{
			fmt.Println("写入数据失败")
		}
	}
}
























