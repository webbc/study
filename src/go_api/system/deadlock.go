package system

import (
	"sync"
	"fmt"
	"time"
)

var mux1 sync.Mutex
var mux2 sync.Mutex
var waitGroup sync.WaitGroup

func main()  {
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		mux1.Lock()
		fmt.Println("[thread 1]mux1 lock")
		time.Sleep(time.Second)
		mux2.Lock()
		fmt.Println("[thread 1]mux2 lock")
		mux2.Unlock()
		mux1.Unlock()
	}()
	go func() {
		defer waitGroup.Done()
		mux2.Lock()
		fmt.Println("[thread 2]mux2 lock")
		time.Sleep(time.Second)
		mux1.Lock()
		fmt.Println("[thread 2]mux1 lock")
		mux1.Unlock()
		mux2.Unlock()
	}()
	waitGroup.Wait()
}
