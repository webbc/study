package observer

type Observer interface {
	Update(subject ISubject, args ...interface{})
}
