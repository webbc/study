// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package pb

import codec "proto/codec"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = codec.Marshal
var _ = math.Inf

type FOO int32

const (
	FOO_X FOO = 17
	FOO_Y FOO = 20
)

var FOO_name = map[int32]string{
	17: "X",
	20: "Y",
}
var FOO_value = map[string]int32{
	"X": 17,
	"Y": 20,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return codec.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := codec.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

type Test struct {
	Name             *string `protobuf:"bytes,1,req,name=name,def=123" json:"name,omitempty"`
	Age              *uint32 `protobuf:"varint,2,req,name=age" json:"age,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return codec.CompactTextString(m) }
func (*Test) ProtoMessage()    {}

const Default_Test_Name string = "123"

func (m *Test) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_Test_Name
}

func (m *Test) GetAge() uint32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return 0
}

func init() {
	codec.RegisterEnum("pb.FOO", FOO_name, FOO_value)
}
