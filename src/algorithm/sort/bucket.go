package main

import "fmt"

func main() {
	bucket := new([11]int)

	fmt.Println("请输入5个数：")
	for i := 0; i < 5; i++ {
		var score int
		fmt.Scanf("%d", &score)
		bucket[score]++
	}

	fmt.Println("排序后的结果：")
	for i := len(bucket)-1; i >= 0; i-- {
		for j := 1; j <= bucket[i]; j++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}

}
