package main

import (
	"sync"
	"fmt"
	"design_pattern/singleton/doublecheck"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			instance := doublecheck.GetInstance()
			fmt.Printf("%p\n", instance)
			wg.Done()
		}()
	}
	wg.Wait()
}
