package main

import (
	"fmt"
)

func main() {
	num := "631758924"
	qq := ""
	length := len(num)
	for i := 0; i < length; i++ {
		data := string(num[i])
		if i%2 == 0 {
			qq += data
			num = num[1:]
		} else {
			num = num + data
		}
		length = len(num)
	}
	fmt.Println(qq)
}
