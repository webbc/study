package main

import (
	"sort"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type ByAge []User

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age > a[j].Age
}

func main() {
	num := []int{5, 1, 34, 52, 134, 6, 9}
	sort.Ints(num)
	fmt.Println(num)

	if sort.IntsAreSorted(num) {
		fmt.Println("ints are sorted")
	}

	num = []int{5, 1, 34, 52, 134, 6, 9}

	str := []string{"1", "b", "abc"}

	sort.Strings(str)

	fmt.Println(str)

	user := []User{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(user)
	sort.Sort(ByAge(user))
	fmt.Println(user)

	arr := []int{3, 1, 5, 6, 4, 1, 123, 3, 34456, 5, 7}
	sort.Ints(arr)
	fmt.Println(arr)
	i := sort.SearchInts(arr, 6)
	fmt.Println(i)
}
