package main

import (
	"sync"
	"fmt"
	"design_pattern/singleton/lock"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			instance := lock.GetInstance()
			fmt.Printf("%p\n", instance)
			wg.Done()
		}()
	}
	wg.Wait()
}
