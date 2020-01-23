// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/response/databags.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DataBags struct {
	DataBags             []*DataBagListItem `protobuf:"bytes,2,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty" toml:"data_bags,omitempty" mapstructure:"data_bags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32              `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBags) Reset()         { *m = DataBags{} }
func (m *DataBags) String() string { return proto.CompactTextString(m) }
func (*DataBags) ProtoMessage()    {}
func (*DataBags) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{0}
}

func (m *DataBags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBags.Unmarshal(m, b)
}
func (m *DataBags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBags.Marshal(b, m, deterministic)
}
func (m *DataBags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBags.Merge(m, src)
}
func (m *DataBags) XXX_Size() int {
	return xxx_messageInfo_DataBags.Size(m)
}
func (m *DataBags) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBags.DiscardUnknown(m)
}

var xxx_messageInfo_DataBags proto.InternalMessageInfo

func (m *DataBags) GetDataBags() []*DataBagListItem {
	if m != nil {
		return m.DataBags
	}
	return nil
}

type DataBagListItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBagListItem) Reset()         { *m = DataBagListItem{} }
func (m *DataBagListItem) String() string { return proto.CompactTextString(m) }
func (*DataBagListItem) ProtoMessage()    {}
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{1}
}

func (m *DataBagListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBagListItem.Unmarshal(m, b)
}
func (m *DataBagListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBagListItem.Marshal(b, m, deterministic)
}
func (m *DataBagListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBagListItem.Merge(m, src)
}
func (m *DataBagListItem) XXX_Size() int {
	return xxx_messageInfo_DataBagListItem.Size(m)
}
func (m *DataBagListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBagListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DataBagListItem proto.InternalMessageInfo

func (m *DataBagListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataBag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBag) Reset()         { *m = DataBag{} }
func (m *DataBag) String() string { return proto.CompactTextString(m) }
func (*DataBag) ProtoMessage()    {}
func (*DataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{2}
}

func (m *DataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBag.Unmarshal(m, b)
}
func (m *DataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBag.Marshal(b, m, deterministic)
}
func (m *DataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBag.Merge(m, src)
}
func (m *DataBag) XXX_Size() int {
	return xxx_messageInfo_DataBag.Size(m)
}
func (m *DataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBag.DiscardUnknown(m)
}

var xxx_messageInfo_DataBag proto.InternalMessageInfo

func (m *DataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*DataBags)(nil), "chef.automate.domain.infra_proxy.response.DataBags")
	proto.RegisterType((*DataBagListItem)(nil), "chef.automate.domain.infra_proxy.response.DataBagListItem")
	proto.RegisterType((*DataBag)(nil), "chef.automate.domain.infra_proxy.response.DataBag")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/response/databags.proto", fileDescriptor_bd99e1b7099a19b5)
}

var fileDescriptor_bd99e1b7099a19b5 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x39, 0x15, 0xbd, 0x5b, 0x0b, 0x61, 0xab, 0x6b, 0x84, 0x10, 0x10, 0x62, 0x33, 0x0b,
	0x6a, 0x65, 0x21, 0x12, 0x6c, 0x04, 0xab, 0x34, 0x82, 0x4d, 0x98, 0x6c, 0x26, 0xc9, 0x16, 0xfb,
	0x83, 0xdd, 0x89, 0xe8, 0x7f, 0x2f, 0x09, 0x09, 0x88, 0x58, 0xa4, 0x1b, 0x78, 0xf3, 0x3e, 0x1e,
	0x9f, 0x78, 0xc0, 0x60, 0x94, 0x71, 0x4c, 0x31, 0x51, 0xfc, 0x34, 0x9a, 0x94, 0x71, 0x5d, 0xc4,
	0x3a, 0x44, 0xff, 0xf5, 0xad, 0x22, 0xa5, 0xe0, 0x5d, 0x22, 0xd5, 0x22, 0x63, 0x83, 0x7d, 0x82,
	0x10, 0x3d, 0x7b, 0x79, 0xab, 0x07, 0xea, 0x00, 0x47, 0xf6, 0x16, 0x99, 0xa0, 0xf5, 0x16, 0x8d,
	0x83, 0x5f, 0x4d, 0x58, 0x9b, 0xb9, 0x16, 0xfb, 0x17, 0x64, 0x2c, 0xb1, 0x4f, 0xf2, 0x5d, 0x1c,
	0x26, 0x50, 0x3d, 0x91, 0x8e, 0x27, 0xd9, 0x69, 0x71, 0x79, 0xf7, 0x08, 0x9b, 0x51, 0xb0, 0x70,
	0xde, 0x4c, 0xe2, 0x57, 0x26, 0x5b, 0xed, 0xdb, 0x05, 0x9c, 0xdf, 0x88, 0xab, 0x3f, 0xa1, 0x94,
	0xe2, 0xcc, 0xa1, 0xa5, 0xe3, 0x2e, 0xdb, 0x15, 0x87, 0x6a, 0xbe, 0xf3, 0x6b, 0x71, 0xb1, 0xbc,
	0xfd, 0x17, 0x97, 0xcf, 0x1f, 0x4f, 0xbd, 0xe1, 0x61, 0x6c, 0x40, 0x7b, 0xab, 0xa6, 0x5d, 0x6a,
	0xdd, 0xa5, 0x36, 0x69, 0x6a, 0xce, 0x67, 0x3d, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99,
	0xb4, 0x8f, 0xf2, 0x56, 0x01, 0x00, 0x00,
}
