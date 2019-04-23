package doublecheck

import (
	"fmt"
	"sync"
)

type user struct {
	name int
}

var u *user
var mutex sync.Mutex

func GetInstance() *user {
	if u == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if u == nil {
			u = &user{}
		}
	}
	return u
}

func (u *user) SetName(name int) {
	u.name = name
}

func (u *user) Say() {
	fmt.Printf("name:%v\n", u.name)
}
