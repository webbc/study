package main

import (
	"sensitive/trie"
	"fmt"
	"strings"
)

func main() {
	var words = []string{"CNM", "操", "尼玛", "SB", "傻逼", "菜鸡", "TM"}
	tree := trie.NewSensitiveTree()
	for _, word := range words {
		tree.Insert(word)
	}

	texts := []string{"你TM是傻逼么？", "操a尼玛", "你是菜鸡吗？"}

	for _, text := range texts {

		t := []rune(text)
		if len(t) <= 0 {
			continue
		}

		p1 := tree.GetRoot()
		p2 := t[0]
		var p2Index, p3Index int
		newText := text

		for i := 0; i < len(t); i++ {
			char := int32(t[i])

			if !p1.Contains(p2) {
				if i < len(t)-1 {
					p2 = int32(t[i+1])
				}
				p2Index = i + 1
				p3Index = i + 1
				p1 = tree.GetRoot()
			} else {
				p1 = p1.GetChildNode(char)
				if i < len(t)-1 {
					p2 = int32(t[i+1])
				}
				p2Index = i + 1

				// 没有子节点了，说明字符串中含有敏感词
				if !p1.HasChild() {
					p1 = tree.GetRoot()

					// 将敏感词替换成*
					oldStr := string(t[p3Index:p2Index])
					newStr := strings.Repeat("*", p2Index-p3Index)
					newText = strings.Replace(newText, oldStr, newStr, -1)
				}
			}
			if i == len(t)-1 {
				fmt.Println(newText)
			}
		}
	}

}
