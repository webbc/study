package command

import (
	"fmt"
	"sync"
)

var mu *sync.Mutex = new(sync.Mutex)

//set命令类
type SetCommands struct {
}

//set命令类实现接口
func (*SetCommands) Do(commands []string,redisMap map[string]string) (replyContent string) {
	if len(commands) < 3 { //判断set命令是否合法
		return "set command error"
	} else {
		mu.Lock()                          //加写锁
		redisMap[commands[1]] = commands[2] //将redis的键存入map中
		mu.Unlock()                         //释放写锁

		fmt.Printf("map=%v\n", redisMap)
		return "ok"
	}
}