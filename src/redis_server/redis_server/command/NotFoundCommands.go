package command

//不存在的命令类
type NotFoundCommands struct {
}

//NotFound命令类实现接口
func (*NotFoundCommands) Do(commands []string, redisMap map[string]string) (replyContent string) {
	return "not found commands"
}
