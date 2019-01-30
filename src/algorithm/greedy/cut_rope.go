// 割绳子问题
package main

import (
	"math"
	"fmt"
)

func main() {
	n := 8
	max := getResult(n)
	fmt.Println(max)
}

func getResult(n int) int {

	if n < 2 {
		panic("n is lower 2")
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	//尽可能多地去减长度为3的绳子段
	i := n / 3

	// 当绳子最后剩下的长度为4的时候，不能再去剪去长度为3的绳子段，而是将4除2
	if n-i*3 == 1 {
		i--
	}
	j := (n - i*3) / 2

	return int(math.Pow(3, float64(i)) * math.Pow(2, float64(j)))
}
