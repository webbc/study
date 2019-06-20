package trie

import "strings"

var tree *SensitiveTree

// 构建屏蔽字树
func CreateSensitiveTree() {
	var words = []string{"CNM", "操", "尼玛", "SB", "傻逼", "菜鸡", "TM", "狗"}
	tree = NewSensitiveTree()
	for _, word := range words {
		tree.Insert(word)
	}
}

func SensitiveCheck(word string) string {

	// 加载屏蔽字树
	CreateSensitiveTree()

	t := []rune(word)        // 转换为unicode数组
	p1 := tree.GetRoot()     // p1指向树的根节点
	p2 := t[0]               // p2是过滤内容的第一个字符
	var p2Index, p3Index int // p2的位置,p3的位置,默认是过滤内容的第一个字符的位置
	newWord := word          // 过滤后的字符串

	// p3指向最后一个字符，则结束
	for ; p3Index < len(t); {

		// 遍历从p3开始到结尾的字符串
		for i := p3Index; i < len(t); i++ {
			char := int32(t[i])

			if !p1.Contains(p2) {

				// 如果p2不是p1子节点，说明以p3开始的内容不需要屏蔽，p3指向下一个位置，p2指向p3，并将p1重置到root节点，终止循环
				p2Index = p3Index + 1
				p3Index = p3Index + 1

				if p2Index < len(t)-1 {
					p2 = int32(t[p2Index])
				}
				p1 = tree.GetRoot()

				break
			} else {

				// 如果p2是p1的子节点，p1节点移动到p2内容对应的节点上，p2移动到下个位置

				// 获取p1的子节点
				p1 = p1.GetChildNode(char)

				// 移动P2到下个位置
				p2Index = p2Index + 1
				if p2Index < len(t)-1 {
					p2 = int32(t[p2Index])
				}

				if !p1.HasChild() {

					// p1没有子节点，树的某条路径遍历完了，说明字符串中含有敏感词
					p1 = tree.GetRoot()

					// 将敏感词替换成*
					oldStr := string(t[p3Index:p2Index])
					newStr := strings.Repeat("*", p2Index-p3Index)
					newWord = strings.Replace(newWord, oldStr, newStr, -1)

					// p3移动到p2的位置
					p3Index = p2Index
				}
			}
		}
	}

	return newWord
}
