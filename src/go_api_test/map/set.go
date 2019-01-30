package _map

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func (hs *HashSet) Add(e interface{}) bool {
	if hs.m[e] {
		return false
	}
	hs.m[e] = true
	return true
}

func (hs *HashSet) Remove(e interface{}) bool {
	if hs.m[e] {
		delete(hs.m, e)
		return true
	}
	return false
}

func (hs *HashSet) Clear() {
	hs.m = make(map[interface{}]bool)
}

func (hs *HashSet) Contains(e interface{}) bool {
	return hs.m[e]
}

func (hs *HashSet) Len() int {
	return len(hs.m)
}

func (hs *HashSet) Elements() []interface{} {
	length := len(hs.m)
	snapshot := make([]interface{}, length)
	i := 0
	for key := range hs.m {
		if i < length {
			snapshot[i] = key
		} else {
			snapshot = append(snapshot, key)
		}
		i ++
	}
	if i < length {
		snapshot = snapshot[:i]
	}
	return snapshot
}

func (hs *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	for key := range hs.m {
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func main() {
	hs := &HashSet{
		m: make(map[interface{}]bool),
	}
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	hs.Add(1)
	hs.Add(1)

	for k, v := range hs.Elements() {
		fmt.Println(fmt.Sprintf("k=%v,v=%v", k, v))
	}
}
