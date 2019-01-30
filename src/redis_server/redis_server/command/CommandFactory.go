package command

//工厂类
type CommandsFactory struct {
}

//创建命令方法
func (*CommandsFactory) CreateCommands(command string) Commands {
	switch   command {
	case "set":
		return &SetCommands{}
	case "get":
		return &GetCommands{}
	default:
		return &NotFoundCommands{}
	}
}

//创建工厂类方法
func NewCommandsFactory() *CommandsFactory {
	return &CommandsFactory{}
}
