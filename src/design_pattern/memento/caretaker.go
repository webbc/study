package memento

// 备忘录管理者
type CareTaker struct {
	mementos []*Memento
}

func (ct *CareTaker) Append(memento *Memento) {
	ct.mementos = append(ct.mementos, memento)
}

func (ct *CareTaker) Get(index int) *Memento {
	if ct.mementos[index] != nil {
		return ct.mementos[index]
	}
	return nil
}
