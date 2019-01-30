package adapter

type Adapter struct {
	Ps2
	usber Usb
}

func (a *Adapter) IsPs2() {
	a.usber.IsUsb()
}

func NewAdapter(usber Usb) *Adapter {
	return &Adapter{
		usber: usber,
	}
}
