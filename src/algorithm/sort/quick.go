package main

import "fmt"

var num [10]int

func main() {
	// 输入10个数
	var n = 10
	for i := 0; i < n; i++ {
		var t int
		fmt.Scanf("%d", &t)
		num[i] = t
	}
	fmt.Println(num)
	quickSort(0, n-1)
	fmt.Println(num)
}

func quickSort(left int, right int) {
	var i, j, temp int
	if left > right {
		return
	}

	temp = num[left] // 基准数
	i = left
	j = right

	for i != j {

		for num[j] >= temp && i < j {
			j--
		}

		for num[i] <= temp && i < j {
			i++
		}

		if i < j {
			num[i], num[j] = num[j], num[i]
		}

		fmt.Println(num)
	}
	// 重调基准数
	num[left], num[i] = num[i], num[left]
	fmt.Println(num)

	quickSort(left, i-1)
	quickSort(i+1, right)
}
