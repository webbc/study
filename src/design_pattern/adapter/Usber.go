package adapter

import "fmt"

type Usber struct {
	Usb
}

func (u *Usber) IsUsb() {
	fmt.Println("USB接口")
}
