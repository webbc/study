package memento

// 备忘录
type Memento struct {
	state uint32
}

func (m *Memento) GetState() uint32 {
	return m.state
}

func NewMemento(state uint32) *Memento {
	return &Memento{state: state}
}
