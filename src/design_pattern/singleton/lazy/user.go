package lazy

import (
	"fmt"
)

type user struct {
	name int
}

var u *user

func GetInstance() *user {
	if u == nil {
		u = &user{}
	}
	return u
}

func (u *user) SetName(name int) {
	u.name = name
}

func (u *user) Say() {
	fmt.Printf("name:%v\n", u.name)
}
