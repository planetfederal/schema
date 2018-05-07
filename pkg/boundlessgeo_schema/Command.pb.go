// Code generated by protoc-gen-go. DO NOT EDIT.
// source: boundlessgeo_schema/Command.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	boundlessgeo_schema/Command.proto
	boundlessgeo_schema/Metadata.proto

It has these top-level messages:
	Command
	Metadata
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Command_ContextType int32

const (
	Command_DESKTOP Command_ContextType = 0
	Command_WEB     Command_ContextType = 1
	Command_MOBILE  Command_ContextType = 2
	Command_SERVICE Command_ContextType = 3
)

var Command_ContextType_name = map[int32]string{
	0: "DESKTOP",
	1: "WEB",
	2: "MOBILE",
	3: "SERVICE",
}
var Command_ContextType_value = map[string]int32{
	"DESKTOP": 0,
	"WEB":     1,
	"MOBILE":  2,
	"SERVICE": 3,
}

func (x Command_ContextType) String() string {
	return proto.EnumName(Command_ContextType_name, int32(x))
}
func (Command_ContextType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Command from a client.
type Command struct {
	Id       string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Action   string              `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Data     []byte              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Context  Command_ContextType `protobuf:"varint,4,opt,name=context,enum=Command_ContextType" json:"context,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Command) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Command) GetContext() Command_ContextType {
	if m != nil {
		return m.Context
	}
	return Command_DESKTOP
}

func (m *Command) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterEnum("Command_ContextType", Command_ContextType_name, Command_ContextType_value)
}

func init() { proto.RegisterFile("boundlessgeo_schema/Command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xca, 0x2f, 0xcd,
	0x4b, 0xc9, 0x49, 0x2d, 0x2e, 0x4e, 0x4f, 0xcd, 0x8f, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4,
	0x77, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x52, 0xc2,
	0xa6, 0xc4, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x11, 0xa2, 0x46, 0xe9, 0x16, 0x23, 0x17,
	0x3b, 0x54, 0x97, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x53, 0x66, 0x8a, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58,
	0x0c, 0xca, 0x13, 0x12, 0xe2, 0x62, 0x01, 0x99, 0x20, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0x04,
	0x66, 0x0b, 0xe9, 0x71, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0x56, 0x94, 0x48, 0xb0, 0x28, 0x30,
	0x6a, 0xf0, 0x19, 0x89, 0xe8, 0xc1, 0x1c, 0xe3, 0x0c, 0x11, 0x0f, 0xa9, 0x2c, 0x48, 0x0d, 0x82,
	0x29, 0x12, 0x52, 0xe5, 0xe2, 0xc8, 0x85, 0xba, 0x44, 0x82, 0x55, 0x81, 0x51, 0x83, 0xdb, 0x88,
	0x53, 0x0f, 0xe6, 0xb4, 0x20, 0xb8, 0x94, 0x92, 0x0d, 0x17, 0x37, 0x92, 0x76, 0x21, 0x6e, 0x2e,
	0x76, 0x17, 0xd7, 0x60, 0xef, 0x10, 0xff, 0x00, 0x01, 0x06, 0x21, 0x76, 0x2e, 0xe6, 0x70, 0x57,
	0x27, 0x01, 0x46, 0x21, 0x2e, 0x2e, 0x36, 0x5f, 0x7f, 0x27, 0x4f, 0x1f, 0x57, 0x01, 0x26, 0x90,
	0x8a, 0x60, 0xd7, 0xa0, 0x30, 0x4f, 0x67, 0x57, 0x01, 0x66, 0x27, 0x8e, 0x28, 0x36, 0x88, 0xaf,
	0x93, 0xd8, 0xc0, 0xbe, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xb4, 0x75, 0xc4, 0x36,
	0x01, 0x00, 0x00,
}
