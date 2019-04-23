package abstract

// 工厂
type IFactory interface {
	GetFruit(catalog uint32) IFruit
	GetColor(catalog uint32) IColor
}
