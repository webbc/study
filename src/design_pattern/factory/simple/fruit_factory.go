package simple

const (
	FRUIT_APPLE  = 1 // 苹果
	FRUIT_BANANA = 2 // 香蕉
)

// 水果工厂
type FruitFactory struct {
}

func (ff *FruitFactory) Get(catalog uint32) IFruit {
	switch catalog {
	case FRUIT_APPLE:
		return &Apple{}
	case FRUIT_BANANA:
		return &Banana{}
	}
	return nil
}
