package error

import "fmt"

func main() {
	echoArr()
	fmt.Println("main exec end")
}

func echoArr() {
	defer func() {
		if v := recover(); v != nil {
			//fmt.Printf("%v\n", v)
			fmt.Println("recover a panic")
		}
	}()
	data := []string{"A", "B", "C"}
	echoItem(data, 0)
}

func echoItem(data []string, index int) {
	if index >= len(data) {
		panic("arr index out of range")
	}
	fmt.Printf("index:%v\n", index)
	elem := data[index]
	defer fmt.Printf("elem:%v\n", elem)
	echoItem(data, index+1)
}
