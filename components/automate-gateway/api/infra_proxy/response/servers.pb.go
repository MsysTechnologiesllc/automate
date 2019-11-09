// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/infra_proxy/response/servers.proto

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

type CreateServer struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServer) Reset()         { *m = CreateServer{} }
func (m *CreateServer) String() string { return proto.CompactTextString(m) }
func (*CreateServer) ProtoMessage()    {}
func (*CreateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{0}
}

func (m *CreateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServer.Unmarshal(m, b)
}
func (m *CreateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServer.Marshal(b, m, deterministic)
}
func (m *CreateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServer.Merge(m, src)
}
func (m *CreateServer) XXX_Size() int {
	return xxx_messageInfo_CreateServer.Size(m)
}
func (m *CreateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServer.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServer proto.InternalMessageInfo

func (m *CreateServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type DeleteServer struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServer) Reset()         { *m = DeleteServer{} }
func (m *DeleteServer) String() string { return proto.CompactTextString(m) }
func (*DeleteServer) ProtoMessage()    {}
func (*DeleteServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{1}
}

func (m *DeleteServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServer.Unmarshal(m, b)
}
func (m *DeleteServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServer.Marshal(b, m, deterministic)
}
func (m *DeleteServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServer.Merge(m, src)
}
func (m *DeleteServer) XXX_Size() int {
	return xxx_messageInfo_DeleteServer.Size(m)
}
func (m *DeleteServer) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServer.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServer proto.InternalMessageInfo

func (m *DeleteServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type UpdateServer struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateServer) Reset()         { *m = UpdateServer{} }
func (m *UpdateServer) String() string { return proto.CompactTextString(m) }
func (*UpdateServer) ProtoMessage()    {}
func (*UpdateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{2}
}

func (m *UpdateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServer.Unmarshal(m, b)
}
func (m *UpdateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServer.Marshal(b, m, deterministic)
}
func (m *UpdateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServer.Merge(m, src)
}
func (m *UpdateServer) XXX_Size() int {
	return xxx_messageInfo_UpdateServer.Size(m)
}
func (m *UpdateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServer proto.InternalMessageInfo

func (m *UpdateServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type GetServers struct {
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetServers) Reset()         { *m = GetServers{} }
func (m *GetServers) String() string { return proto.CompactTextString(m) }
func (*GetServers) ProtoMessage()    {}
func (*GetServers) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{3}
}

func (m *GetServers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServers.Unmarshal(m, b)
}
func (m *GetServers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServers.Marshal(b, m, deterministic)
}
func (m *GetServers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServers.Merge(m, src)
}
func (m *GetServers) XXX_Size() int {
	return xxx_messageInfo_GetServers.Size(m)
}
func (m *GetServers) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServers.DiscardUnknown(m)
}

var xxx_messageInfo_GetServers proto.InternalMessageInfo

func (m *GetServers) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type GetServer struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServer) Reset()         { *m = GetServer{} }
func (m *GetServer) String() string { return proto.CompactTextString(m) }
func (*GetServer) ProtoMessage()    {}
func (*GetServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{4}
}

func (m *GetServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServer.Unmarshal(m, b)
}
func (m *GetServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServer.Marshal(b, m, deterministic)
}
func (m *GetServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServer.Merge(m, src)
}
func (m *GetServer) XXX_Size() int {
	return xxx_messageInfo_GetServer.Size(m)
}
func (m *GetServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServer.DiscardUnknown(m)
}

var xxx_messageInfo_GetServer proto.InternalMessageInfo

func (m *GetServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type Server struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Fqdn                 string   `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	IpAddress            string   `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	OrgsCount            int32    `protobuf:"varint,6,opt,name=orgs_count,json=orgsCount,proto3" json:"orgs_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_423311f0ee37174d, []int{5}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Server) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Server) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Server) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Server) GetOrgsCount() int32 {
	if m != nil {
		return m.OrgsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateServer)(nil), "chef.automate.api.infra_proxy.response.CreateServer")
	proto.RegisterType((*DeleteServer)(nil), "chef.automate.api.infra_proxy.response.DeleteServer")
	proto.RegisterType((*UpdateServer)(nil), "chef.automate.api.infra_proxy.response.UpdateServer")
	proto.RegisterType((*GetServers)(nil), "chef.automate.api.infra_proxy.response.GetServers")
	proto.RegisterType((*GetServer)(nil), "chef.automate.api.infra_proxy.response.GetServer")
	proto.RegisterType((*Server)(nil), "chef.automate.api.infra_proxy.response.Server")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/infra_proxy/response/servers.proto", fileDescriptor_423311f0ee37174d)
}

var fileDescriptor_423311f0ee37174d = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x99, 0xfe, 0xcc, 0xc7, 0x9c, 0x7e, 0xb8, 0xc8, 0x2a, 0x1b, 0x61, 0xe8, 0x42, 0xba,
	0x31, 0x01, 0xbd, 0x02, 0x5b, 0x51, 0x37, 0x22, 0xb4, 0xe8, 0xc2, 0x4d, 0x49, 0x27, 0xa7, 0x6d,
	0xc0, 0x49, 0x62, 0x92, 0xaa, 0xbd, 0x1b, 0x2f, 0x55, 0x92, 0xe9, 0x94, 0x2e, 0xa5, 0xd8, 0xdd,
	0xe1, 0x3d, 0x3c, 0x0f, 0x27, 0xe1, 0x85, 0x71, 0x65, 0x6a, 0x6b, 0x34, 0xea, 0xe0, 0xb9, 0xd8,
	0x04, 0x53, 0x8b, 0x80, 0x97, 0x2b, 0x11, 0xf0, 0x53, 0x6c, 0xb9, 0xb0, 0x8a, 0x2b, 0xbd, 0x74,
	0x62, 0x6e, 0x9d, 0xf9, 0xda, 0x72, 0x87, 0xde, 0x1a, 0xed, 0x91, 0x7b, 0x74, 0x1f, 0xe8, 0x3c,
	0xb3, 0xce, 0x04, 0x43, 0x2e, 0xaa, 0x35, 0x2e, 0x59, 0x4b, 0x33, 0x61, 0x15, 0x3b, 0xa0, 0x58,
	0x4b, 0x0d, 0x5f, 0xe0, 0xff, 0xc4, 0xa1, 0x08, 0x38, 0x4b, 0x38, 0xb9, 0x83, 0xbc, 0x11, 0xd1,
	0xac, 0xcc, 0x46, 0x83, 0x2b, 0xc6, 0x7e, 0x27, 0x62, 0x0d, 0x3f, 0xdd, 0xd1, 0xd1, 0x7b, 0x8b,
	0x6f, 0x78, 0x0a, 0xef, 0xb3, 0x95, 0xa7, 0xb8, 0x17, 0xee, 0x31, 0x34, 0xa1, 0x27, 0x0f, 0xf0,
	0x6f, 0xf7, 0x9d, 0x34, 0x2b, 0xbb, 0x47, 0x68, 0x5b, 0x7c, 0x38, 0x83, 0x62, 0xef, 0xfd, 0xb3,
	0x63, 0xbf, 0x33, 0xc8, 0x77, 0xca, 0x33, 0xe8, 0x28, 0x99, 0x74, 0xc5, 0xb4, 0xa3, 0x24, 0x21,
	0xd0, 0xd3, 0xa2, 0x46, 0xda, 0x49, 0x49, 0x9a, 0x49, 0x09, 0x03, 0x89, 0xbe, 0x72, 0xca, 0x06,
	0x65, 0x34, 0xed, 0xa6, 0xd5, 0x61, 0x14, 0xa9, 0xe5, 0xbb, 0xd4, 0xb4, 0xd7, 0x50, 0x71, 0x26,
	0xe7, 0x00, 0xca, 0xce, 0x85, 0x94, 0x0e, 0xbd, 0xa7, 0xfd, 0xb4, 0x29, 0x94, 0xbd, 0x69, 0x82,
	0xb8, 0x36, 0x6e, 0xe5, 0xe7, 0x95, 0xd9, 0xe8, 0x40, 0xf3, 0x32, 0x1b, 0xf5, 0xa7, 0x45, 0x4c,
	0x26, 0x31, 0x18, 0x3f, 0xbd, 0x3e, 0xae, 0x54, 0x58, 0x6f, 0x16, 0xac, 0x32, 0x35, 0x8f, 0xcf,
	0xdc, 0x57, 0x99, 0x1f, 0x53, 0xef, 0x45, 0x9e, 0x7a, 0x7d, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x83, 0xa4, 0x38, 0x44, 0x1d, 0x03, 0x00, 0x00,
}
