package data_struct

import (
	"strings"
	"fmt"
)

func main() {
	str := "\u0041"
	fmt.Println(strings.EqualFold(str, "\u0062"))

	str = "abcab"
	fmt.Println(strings.Count(str, "ab"))

	str = "abcab"
	fmt.Println(strings.Index(str, "f"))

	str = "abcab"
	fmt.Println(strings.Index(str, "f"))

	str = "aacdefabc"
	fmt.Println(strings.Replace(str,"a","1",-1))
}
