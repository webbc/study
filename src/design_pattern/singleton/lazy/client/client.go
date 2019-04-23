package main

import (
	"design_pattern/singleton/lazy"
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			instance := lazy.GetInstance()
			fmt.Printf("%p\n", instance)
			wg.Done()
		}()
	}
	wg.Wait()
}
