package main

import (
	"design_pattern/builder"
	"fmt"
)

func main() {
	sb := builder.NewStringBuilder()
	fmt.Println(sb.Append("a").Append("b").Append("c").String())
}
