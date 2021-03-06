// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-grpc/protoc-gen-policy/iam/annotations.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Policy = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*PolicyInfo)(nil),
	Field:         50001,
	Name:          "chef.automate.api.iam.policy",
	Tag:           "bytes,50001,opt,name=policy",
	Filename:      "components/automate-grpc/protoc-gen-policy/iam/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Policy)
}

func init() {
	proto.RegisterFile("components/automate-grpc/protoc-gen-policy/iam/annotations.proto", fileDescriptor_fdc3d017cc513fc9)
}

var fileDescriptor_fdc3d017cc513fc9 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x29, 0x05, 0x0f, 0xee, 0x66, 0x28, 0x14, 0x0f, 0xc5, 0xdd, 0xba, 0xf8, 0x04, 0xed,
	0x96, 0x2c, 0x21, 0x4b, 0xf0, 0x10, 0x12, 0x32, 0x26, 0x93, 0x2c, 0xcb, 0xb2, 0xc0, 0xd2, 0x09,
	0xf9, 0x3c, 0xe4, 0x05, 0xf2, 0x7e, 0x79, 0xa3, 0x10, 0x29, 0xce, 0x94, 0xc5, 0xdb, 0x71, 0x7c,
	0xff, 0xc7, 0xff, 0xa7, 0x2b, 0x81, 0xc6, 0xa1, 0x95, 0x96, 0x06, 0xc6, 0x47, 0x42, 0xc3, 0x49,
	0x96, 0xca, 0x3b, 0xc1, 0x9c, 0x47, 0x42, 0x51, 0x2a, 0x69, 0x4b, 0x87, 0xbd, 0x16, 0x67, 0xa6,
	0xb9, 0x61, 0xdc, 0x5a, 0x24, 0x4e, 0x1a, 0xed, 0x00, 0x81, 0xc8, 0x3e, 0x45, 0x27, 0x5b, 0x98,
	0xb2, 0xc0, 0x9d, 0x06, 0xcd, 0x4d, 0xbe, 0x9c, 0x29, 0x8e, 0x67, 0x74, 0xe6, 0x85, 0x42, 0x54,
	0xbd, 0x8c, 0x68, 0x3d, 0xb6, 0xac, 0x91, 0x83, 0xf0, 0xda, 0x11, 0xfa, 0x48, 0x2c, 0x4e, 0x69,
	0x12, 0x13, 0xd9, 0x37, 0x44, 0x18, 0x26, 0x18, 0xb6, 0x92, 0x3a, 0x6c, 0x76, 0x2e, 0xb4, 0xfc,
	0xba, 0x5e, 0xde, 0x8b, 0xb7, 0xdf, 0x8f, 0xbf, 0x1f, 0x78, 0x59, 0x14, 0xf6, 0x41, 0x53, 0xd9,
	0x16, 0x0f, 0x0f, 0xe5, 0xba, 0x3a, 0x6e, 0x94, 0xa6, 0x6e, 0xac, 0x41, 0xa0, 0x61, 0xf7, 0xd8,
	0x73, 0x02, 0x9b, 0x37, 0xab, 0x4e, 0xc2, 0xfb, 0xff, 0x16, 0x00, 0x00, 0xff, 0xff, 0x32, 0x4e,
	0x19, 0x15, 0x68, 0x01, 0x00, 0x00,
}
