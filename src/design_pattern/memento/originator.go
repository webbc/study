package memento

// 备忘录发起者
type Originator struct {
	state uint32
}

func (o *Originator) SetState(state uint32) {
	o.state = state
}

func (o *Originator) GetState() uint32 {
	return o.state
}

func (o *Originator) Save() *Memento {
	return NewMemento(o.state)
}
