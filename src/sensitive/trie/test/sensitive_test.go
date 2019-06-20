package test

import (
	"testing"
	"sensitive/trie"
)

func Benchmark_SensitiveCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie.SensitiveCheck("本土无码公司待办列宁测试GM")
	}
}
