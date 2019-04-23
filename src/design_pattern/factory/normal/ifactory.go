package normal

// 工厂
type IFactory interface {
	GetFruit(catalog uint32) IFruit
}
