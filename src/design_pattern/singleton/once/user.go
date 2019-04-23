package once

import (
	"fmt"
	"sync"
)

type user struct {
	name int
}

var u *user
var once sync.Once

func GetInstance() *user {
	once.Do(func() {
		u = &user{}
	})
	return u
}

func (u *user) SetName(name int) {
	u.name = name
}

func (u *user) Say() {
	fmt.Printf("name:%v\n", u.name)
}
