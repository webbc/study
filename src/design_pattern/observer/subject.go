package observer

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Detach(observer Observer) {
	for k, v := range s.observers {
		if v == observer {
			s.observers = append(s.observers[:k], s.observers[k+1:]...)
		}
	}
}

func (s *Subject) NotifyObservers(self ISubject, args ...interface{}) {
	for _, v := range s.observers {
		v.Update(self, args...)
	}
}
