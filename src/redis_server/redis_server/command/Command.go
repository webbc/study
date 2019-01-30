package command

//命令接口
type Commands interface {
	//处理方法
	Do(command []string, redisMap map[string]string) (replyContent string)
}
