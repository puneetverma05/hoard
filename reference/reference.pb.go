// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reference.proto

package reference // import "github.com/monax/hoard/reference"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Ref struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	SecretKey            []byte   `protobuf:"bytes,2,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	Salt                 []byte   `protobuf:"bytes,3,opt,name=Salt,proto3" json:"Salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_reference_efa25a11fac9b38a, []int{0}
}
func (m *Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ref.Unmarshal(m, b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
}
func (dst *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(dst, src)
}
func (m *Ref) XXX_Size() int {
	return xxx_messageInfo_Ref.Size(m)
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

func (m *Ref) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Ref) GetSecretKey() []byte {
	if m != nil {
		return m.SecretKey
	}
	return nil
}

func (m *Ref) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func init() {
	proto.RegisterType((*Ref)(nil), "reference.Ref")
}
func (m *Ref) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReference(uint64(l))
	}
	l = len(m.SecretKey)
	if l > 0 {
		n += 1 + l + sovReference(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovReference(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovReference(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReference(x uint64) (n int) {
	return sovReference(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("reference.proto", fileDescriptor_reference_efa25a11fac9b38a) }

var fileDescriptor_reference_efa25a11fac9b38a = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x4d, 0x4b,
	0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0x55, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa9, 0x14, 0xc8,
	0xc5, 0x1c, 0x94, 0x9a, 0x26, 0x24, 0xc1, 0xc5, 0xee, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe3, 0x0a, 0xc9, 0x70, 0x71, 0x06, 0xa7, 0x26, 0x17,
	0xa5, 0x96, 0x78, 0xa7, 0x56, 0x4a, 0x30, 0x81, 0xe5, 0x10, 0x02, 0x42, 0x42, 0x5c, 0x2c, 0xc1,
	0x89, 0x39, 0x25, 0x12, 0xcc, 0x60, 0x09, 0x30, 0xdb, 0x49, 0x6d, 0xc1, 0x63, 0x39, 0xc6, 0x28,
	0x05, 0x24, 0x77, 0xe4, 0xe6, 0xe7, 0x25, 0x56, 0xe8, 0x67, 0xe4, 0x27, 0x16, 0xa5, 0xe8, 0xc3,
	0x5d, 0x9a, 0xc4, 0x06, 0x76, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xfe, 0x9d, 0x2b,
	0xce, 0x00, 0x00, 0x00,
}