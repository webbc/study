package builder

import "bytes"

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

func (sb *StringBuilder) Append(str string) *StringBuilder {
	sb.buffer.WriteString(str)
	return sb
}

func (sb *StringBuilder) String() string {
	return sb.buffer.String()
}
