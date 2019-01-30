package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	cmd "redis_server/redis_server/command"
)

var redisMap = make(map[string]string) //定义redis的键值对

//主方法
func main() {
	startServer() //开启服务端
}

//开启服务端
func startServer() {
	ln, err := net.Listen("tcp", ":8080") //创建端口的监听器
	//错误处理
	if err != nil {
		fmt.Println("服务器开启失败")
		os.Exit(0)
	}
	//让服务端永不断开
	for {
		conn, err := ln.Accept() //获取客户端连接
		//错误处理
		if err != nil {
			fmt.Println("客户端连接失败")
			continue
		}
		fmt.Printf("接收到客户端连接: %v", conn.RemoteAddr()) //输出客户端的地址
		go handleConn(conn)                           //启动线程去为客户端服务
	}
}

//处理客户端连接
func handleConn(conn net.Conn) {
	//让线程不断开
	for {
		var buffer = make([]byte, 1024) //创建缓存
		n, err := conn.Read(buffer)     //读取客户端发来的数据
		//错误处理
		if err != nil {
			fmt.Println("读取数据失败")
		}
		//如果数据为空，则断开此线程
		if buffer == nil {
			continue
		}
		var command = string(buffer[:n])                           //将byte数据转换为string数据
		command = strings.Replace(command, "\n", "", -1)           //删除掉数据中的换行
		command = strings.Replace(command, "\r", "", -1)           //删除掉数据中的换行
		commands := strings.Split(strings.TrimSpace(command), " ") //按照空格进行截取成数组

		fmt.Printf("recv %v\n", []byte(command))

		Commands := cmd.NewCommandsFactory().CreateCommands(commands[0]) //创建命令对象
		replyContent := Commands.Do(commands, redisMap)                  //命令对象做操作，得到返回给客户端的数据

		_, err = conn.Write([]byte(replyContent)) //发送给客户端

		//错误处理
		if err != nil {
			fmt.Println("发送数据错误")
		}
	}
}
