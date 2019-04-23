package abstract

const (
	COLOR_RED    = 1 // 红色
	COLOR_YELLOW = 2 // 黄色
)

type ColorFactory struct {
}

func (cf *ColorFactory) GetColor(catalog uint32) IColor {
	switch catalog {
	case COLOR_RED:
		return &Red{}
	case COLOR_YELLOW:
		return &Yellow{}
	}
	return nil
}
