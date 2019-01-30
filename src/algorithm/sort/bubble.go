package main

import "fmt"

func main() {
	// 输入10个数
	var num [5]int
	for i := 0; i < 5; i++ {
		var t int
		fmt.Scanf("%d", &t)
		num[i] = t
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
		fmt.Println(num)
	}
	for i := 0; i < 5; i++ {
		fmt.Print(num[i])
		fmt.Print(" ")
	}
}
