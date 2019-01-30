package command

import (
	"fmt"
	"sync"
)

var muRw *sync.RWMutex = new(sync.RWMutex)

//get命令类
type GetCommands struct {
}

//get命令类实现接口
func (*GetCommands) Do(commands []string, redisMap map[string]string) (replyContent string) {
	fmt.Printf("map=%v, find=%v.....\n", redisMap, commands[1])
	muRw.RLock() //加读锁
	value, ok := redisMap[commands[1]]
	muRw.RUnlock() //释放读锁
	if ok { //判断读取的key是否存在
		return value //返回指定的key的value值
	} else {
		return "not found"
	}
}
