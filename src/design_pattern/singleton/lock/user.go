package lock

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
	mutex.Lock()
	if u == nil {
		u = &user{}
	}
	mutex.Unlock()
	return u
}

func (u *user) SetName(name int) {
	u.name = name
}

func (u *user) Say() {
	fmt.Printf("name:%v\n", u.name)
}
