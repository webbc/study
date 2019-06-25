package test

import (
	"testing"
	"sensitive/trie"
	"fmt"
)

func Benchmark_SensitiveCheck(b *testing.B) {
	// 加载屏蔽字树
	e := trie.CreateSensitiveTree()
	if e != nil {
		fmt.Println(e)
	}

	for i := 0; i < b.N; i++ {
		trie.SensitiveCheck("本土无码公司待办列宁测试GM")
	}
}
