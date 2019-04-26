package prototype

type ICar interface {
	SetName(string)
	GetName() string
	Clone() ICar
}
