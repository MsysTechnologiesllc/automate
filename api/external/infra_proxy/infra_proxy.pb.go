// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	context "context"
	fmt "fmt"
	version "github.com/chef/automate/api/external/common/version"
	request "github.com/chef/automate/api/external/infra_proxy/request"
	response "github.com/chef/automate/api/external/infra_proxy/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("api/external/infra_proxy/infra_proxy.proto", fileDescriptor_8898095cc3dd9190)
}

var fileDescriptor_8898095cc3dd9190 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0xbf, 0x8f, 0xdc, 0x44,
	0x18, 0x95, 0x37, 0x68, 0xc9, 0x0d, 0x07, 0x81, 0x2f, 0x97, 0x8b, 0xcf, 0xc9, 0x25, 0x64, 0x20,
	0x07, 0x1c, 0xac, 0x7d, 0x7b, 0x7b, 0xa2, 0xd8, 0x06, 0xee, 0x0e, 0x14, 0xa5, 0x01, 0x04, 0x82,
	0x82, 0x82, 0x93, 0xcf, 0x37, 0xe7, 0x58, 0xb7, 0xf6, 0x38, 0xb6, 0xf7, 0x94, 0x25, 0x4a, 0xb3,
	0xe5, 0x82, 0x68, 0xa0, 0xa2, 0x40, 0x02, 0x0a, 0x28, 0xd3, 0xf8, 0xaf, 0xa0, 0x8c, 0x90, 0xa8,
	0x28, 0x10, 0x20, 0x0a, 0x0a, 0x1a, 0x2a, 0x1a, 0x34, 0x1e, 0xdb, 0x3b, 0xb6, 0x77, 0xd7, 0x9e,
	0x50, 0x90, 0x6a, 0x57, 0xf6, 0x7b, 0x33, 0xdf, 0x7b, 0xdf, 0x8f, 0xb1, 0x8d, 0x36, 0x4d, 0xdf,
	0x31, 0xc8, 0x9d, 0x88, 0x04, 0x9e, 0x39, 0x30, 0x1c, 0xef, 0x38, 0x30, 0x0f, 0xfc, 0x80, 0xde,
	0x19, 0x89, 0xff, 0x75, 0x3f, 0xa0, 0x11, 0x85, 0x75, 0xeb, 0x16, 0x39, 0xd6, 0xcd, 0x61, 0x44,
	0x5d, 0x33, 0x22, 0xba, 0xe9, 0x3b, 0xba, 0x00, 0xd2, 0x2e, 0xdb, 0x94, 0xda, 0x03, 0x62, 0xb0,
	0x15, 0x4d, 0xcf, 0xa3, 0x91, 0x19, 0x39, 0xd4, 0x0b, 0x39, 0x59, 0x7b, 0x79, 0xee, 0x46, 0x01,
	0xb9, 0x3d, 0x24, 0x61, 0x64, 0xd0, 0xc0, 0xce, 0xc0, 0x7a, 0x2d, 0x38, 0x24, 0xc1, 0x29, 0x09,
	0x32, 0xfc, 0x56, 0x2d, 0xde, 0xa2, 0xf4, 0xe4, 0x90, 0xd2, 0x93, 0x8c, 0xf1, 0xca, 0x02, 0x46,
	0xe8, 0x53, 0x2f, 0x24, 0x62, 0x3c, 0x46, 0x3d, 0xba, 0x18, 0x50, 0xb7, 0x9e, 0x50, 0x8e, 0xe8,
	0xa5, 0x02, 0xc5, 0xa2, 0xae, 0x4b, 0x3d, 0x83, 0x2d, 0xe9, 0x4c, 0x7f, 0x53, 0xe8, 0xeb, 0x16,
	0x75, 0x7d, 0xea, 0x11, 0x2f, 0x0a, 0x8d, 0x2c, 0x1d, 0x1d, 0x3b, 0xf0, 0x2d, 0x23, 0xb9, 0x6f,
	0x75, 0x6c, 0xe2, 0x75, 0x7c, 0x3a, 0x70, 0xac, 0xd1, 0x9c, 0x6c, 0xc8, 0xac, 0xe0, 0x98, 0x6e,
	0x75, 0x85, 0xed, 0xfb, 0x18, 0xa1, 0x9b, 0x4c, 0xd7, 0x3b, 0x4c, 0x16, 0xfc, 0xa1, 0x20, 0x74,
	0x83, 0x44, 0x1f, 0xf0, 0x38, 0x61, 0x47, 0xaf, 0xd6, 0x0a, 0x97, 0xa4, 0x67, 0x52, 0x52, 0xe8,
	0x4d, 0xef, 0x98, 0xbe, 0xcb, 0x93, 0xa4, 0x75, 0xa4, 0x58, 0x78, 0x38, 0x8e, 0xd5, 0x55, 0xb4,
	0xc2, 0x9c, 0x77, 0x2c, 0x72, 0xe0, 0x78, 0xc7, 0xb4, 0x9f, 0xe2, 0xc6, 0xb1, 0xda, 0x86, 0xc7,
	0x02, 0x62, 0x1e, 0x4d, 0x62, 0x55, 0x45, 0xab, 0xe1, 0x28, 0x8c, 0x88, 0xdb, 0x4f, 0xa1, 0x19,
	0x6a, 0x12, 0xab, 0x97, 0x60, 0xad, 0x78, 0x2f, 0xdd, 0xa0, 0x6f, 0x93, 0x68, 0xfc, 0xe0, 0xd7,
	0xcf, 0x5b, 0x4f, 0xc3, 0x53, 0x3c, 0x7f, 0x59, 0x0e, 0xe0, 0x01, 0x97, 0xfa, 0x1e, 0x4f, 0x38,
	0x74, 0xf5, 0x85, 0x6d, 0xa1, 0xa7, 0x15, 0xa8, 0x4f, 0x29, 0xda, 0x76, 0x2d, 0x85, 0xd7, 0x88,
	0xc0, 0xc1, 0x1f, 0x8d, 0x63, 0xf5, 0x1c, 0x7a, 0x32, 0x81, 0xf5, 0xd3, 0x62, 0x2b, 0xa8, 0x2c,
	0xdf, 0x9c, 0xc4, 0xea, 0x0a, 0x40, 0xe1, 0x52, 0x7f, 0xe0, 0x84, 0x65, 0x55, 0xe9, 0x2d, 0xf8,
	0x59, 0x41, 0x4b, 0xf9, 0x76, 0xb0, 0x25, 0x2b, 0x4a, 0xeb, 0x4a, 0x6b, 0xc2, 0x83, 0x71, 0xac,
	0xae, 0xa0, 0x52, 0x88, 0x77, 0x9d, 0xa3, 0x7b, 0x05, 0x5d, 0x33, 0x11, 0x93, 0x58, 0x3d, 0x0f,
	0xcf, 0x14, 0xaf, 0x67, 0x19, 0xbb, 0x00, 0xe7, 0x8b, 0xda, 0x0c, 0xc6, 0x80, 0xbf, 0x14, 0x74,
	0x2e, 0xdf, 0x7b, 0x6f, 0xf4, 0x96, 0xe9, 0x12, 0x78, 0x55, 0x56, 0x26, 0xe7, 0x3d, 0x8c, 0xd8,
	0x80, 0x17, 0x6b, 0x49, 0x8a, 0x67, 0xba, 0xa4, 0x28, 0x77, 0x0e, 0x66, 0x91, 0xe0, 0x8b, 0x70,
	0xa1, 0x2c, 0x38, 0xe1, 0xc0, 0xef, 0x0a, 0x5a, 0xde, 0x0f, 0x88, 0x19, 0x91, 0x34, 0xad, 0xbd,
	0x86, 0x7a, 0x45, 0x92, 0xb6, 0xd3, 0x54, 0xac, 0xc8, 0xc2, 0x27, 0xb3, 0xeb, 0xf5, 0x2c, 0xb4,
	0xad, 0x04, 0x36, 0xbb, 0x62, 0x57, 0xa1, 0xa4, 0x9e, 0x83, 0x79, 0xcd, 0xe2, 0x52, 0xcd, 0x26,
	0x57, 0xcf, 0xf4, 0x95, 0x4d, 0xf8, 0x5b, 0x41, 0xcb, 0xef, 0xfb, 0x47, 0xf2, 0x42, 0x45, 0x52,
	0x73, 0xa1, 0x22, 0x0b, 0x7f, 0xbc, 0xa0, 0x8a, 0xcf, 0x42, 0x7b, 0x98, 0x60, 0x17, 0xd4, 0x71,
	0x45, 0x32, 0x67, 0xf0, 0x52, 0xd6, 0x66, 0x95, 0xf2, 0x54, 0xf7, 0x9f, 0x0a, 0x5a, 0x7e, 0x83,
	0x0c, 0x88, 0xb4, 0x6e, 0x91, 0xd4, 0x5c, 0xb7, 0xc8, 0xc2, 0xb7, 0x17, 0xeb, 0x3e, 0x4a, 0xb0,
	0x32, 0xba, 0x39, 0x83, 0xeb, 0xde, 0x9c, 0xd9, 0xc2, 0xff, 0x28, 0xe8, 0xf1, 0x1b, 0x24, 0x7a,
	0x3b, 0xb0, 0x43, 0xd0, 0x9b, 0xb7, 0x2e, 0xc3, 0x6b, 0x86, 0x44, 0xcb, 0x32, 0x02, 0xfe, 0x4c,
	0x19, 0xc7, 0xea, 0x25, 0xb4, 0x56, 0x0a, 0x9e, 0xff, 0x39, 0x28, 0x4f, 0xa9, 0x67, 0xd1, 0x95,
	0xb9, 0xc0, 0x3e, 0x7b, 0x92, 0xa8, 0x2a, 0x66, 0xbb, 0xe4, 0x3d, 0x7c, 0x0d, 0xae, 0x96, 0x15,
	0x4f, 0x17, 0x48, 0x1e, 0x45, 0xe0, 0xd3, 0x16, 0x6a, 0xf3, 0xe0, 0xa0, 0x23, 0x25, 0x5e, 0xd3,
	0xe5, 0xb4, 0xe3, 0x6f, 0x98, 0xf4, 0xeb, 0xe8, 0xb9, 0xc5, 0x8a, 0xaa, 0xa3, 0xba, 0x19, 0xa5,
	0xc6, 0x89, 0x0d, 0x78, 0xbe, 0xc6, 0x09, 0x5e, 0x0c, 0x5f, 0xb5, 0xd0, 0x32, 0x8f, 0x37, 0x1d,
	0xe6, 0x3d, 0x29, 0x53, 0xd2, 0x49, 0x2e, 0x6b, 0xcd, 0xf7, 0xcc, 0x9a, 0x17, 0xd0, 0xf5, 0x3a,
	0x9d, 0xd5, 0xc1, 0xde, 0x94, 0x54, 0x63, 0xcf, 0x8b, 0xb0, 0x51, 0x6b, 0x0f, 0x9f, 0xfe, 0x5f,
	0xb4, 0xd0, 0x12, 0x1f, 0xc9, 0xac, 0x64, 0xb6, 0xa4, 0x46, 0x3f, 0xab, 0x9a, 0xae, 0xdc, 0xdc,
	0x67, 0xee, 0x7c, 0xcd, 0xdc, 0xa9, 0x6d, 0x85, 0xd2, 0x31, 0xd0, 0xa4, 0x75, 0xd6, 0xe0, 0x62,
	0xc5, 0x11, 0xe1, 0x68, 0xb8, 0x86, 0xeb, 0xba, 0x67, 0x3a, 0x33, 0xbf, 0x6d, 0xa1, 0x25, 0x3e,
	0xc0, 0x65, 0x6c, 0xc9, 0x19, 0xcd, 0x6d, 0xc9, 0x29, 0xf8, 0xbe, 0x54, 0x3f, 0x89, 0x87, 0x46,
	0xe3, 0x8e, 0x9a, 0x65, 0x90, 0x70, 0x90, 0x6c, 0x68, 0x8d, 0x9a, 0x6a, 0xea, 0xd2, 0x97, 0x2d,
	0xb4, 0xc4, 0xc7, 0xbd, 0x8c, 0x4b, 0x39, 0xa3, 0xb9, 0x4b, 0x39, 0x05, 0x7f, 0x27, 0xeb, 0x52,
	0x7e, 0xc4, 0xfc, 0x27, 0x97, 0x84, 0x63, 0x67, 0x63, 0xb3, 0xd9, 0xe8, 0xf9, 0x81, 0x8f, 0x9e,
	0xfd, 0xec, 0x0d, 0xae, 0x79, 0x73, 0x65, 0x0c, 0x89, 0xe6, 0xca, 0x28, 0xf8, 0x47, 0xe6, 0x4f,
	0x0f, 0x75, 0xeb, 0xa4, 0xd2, 0xc0, 0x4e, 0xe2, 0xcd, 0x5f, 0x30, 0x0b, 0x63, 0xa8, 0xf9, 0x02,
	0xfd, 0x7c, 0x81, 0x49, 0xac, 0x5e, 0x85, 0xf5, 0x92, 0x73, 0x79, 0x60, 0xd3, 0xb7, 0x8a, 0x1d,
	0xd8, 0xae, 0xf5, 0xaf, 0x12, 0x1d, 0x7c, 0x72, 0x06, 0xad, 0x8b, 0x6e, 0xee, 0x9e, 0x9a, 0xce,
	0xc0, 0x3c, 0x1c, 0x64, 0xaf, 0x62, 0x21, 0xec, 0xca, 0xda, 0x5b, 0x59, 0x42, 0xdb, 0x93, 0xf6,
	0xbb, 0xb2, 0x06, 0xfe, 0xe5, 0x91, 0x4f, 0xc0, 0x2e, 0xbc, 0x26, 0x9f, 0x00, 0xc3, 0x1b, 0xba,
	0x07, 0xa7, 0x99, 0xd7, 0x3f, 0xb5, 0xd0, 0x13, 0x42, 0x36, 0xc0, 0x90, 0xf4, 0x5e, 0xdb, 0x92,
	0x75, 0x1a, 0xff, 0xf6, 0x7f, 0xfa, 0x7a, 0x05, 0x2e, 0xcf, 0xf5, 0x35, 0x3b, 0x73, 0xdf, 0x84,
	0xfd, 0x87, 0xb0, 0x95, 0x1f, 0xc3, 0xc6, 0xdd, 0xd4, 0xda, 0x7b, 0x7b, 0xbd, 0x0f, 0xbb, 0xb6,
	0x13, 0xdd, 0x1a, 0x1e, 0xea, 0x16, 0x75, 0x0d, 0xe6, 0x52, 0xfe, 0xed, 0x65, 0xee, 0x07, 0xa6,
	0xc3, 0x76, 0xf2, 0xb9, 0xa5, 0xf7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0xb4, 0xc5, 0x43,
	0xa9, 0x13, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InfraProxyClient is the client API for InfraProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfraProxyClient interface {
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	GetServers(ctx context.Context, in *request.GetServers, opts ...grpc.CallOption) (*response.GetServers, error)
	GetServer(ctx context.Context, in *request.GetServer, opts ...grpc.CallOption) (*response.GetServer, error)
	GetServerByName(ctx context.Context, in *request.GetServerByName, opts ...grpc.CallOption) (*response.GetServer, error)
	CreateServer(ctx context.Context, in *request.CreateServer, opts ...grpc.CallOption) (*response.CreateServer, error)
	UpdateServer(ctx context.Context, in *request.UpdateServer, opts ...grpc.CallOption) (*response.UpdateServer, error)
	DeleteServer(ctx context.Context, in *request.DeleteServer, opts ...grpc.CallOption) (*response.DeleteServer, error)
	GetOrgs(ctx context.Context, in *request.GetOrgs, opts ...grpc.CallOption) (*response.GetOrgs, error)
	GetOrg(ctx context.Context, in *request.GetOrg, opts ...grpc.CallOption) (*response.GetOrg, error)
	GetOrgByName(ctx context.Context, in *request.GetOrgByName, opts ...grpc.CallOption) (*response.GetOrg, error)
	CreateOrg(ctx context.Context, in *request.CreateOrg, opts ...grpc.CallOption) (*response.CreateOrg, error)
	UpdateOrg(ctx context.Context, in *request.UpdateOrg, opts ...grpc.CallOption) (*response.UpdateOrg, error)
	DeleteOrg(ctx context.Context, in *request.DeleteOrg, opts ...grpc.CallOption) (*response.DeleteOrg, error)
	GetCookbooks(ctx context.Context, in *request.Cookbooks, opts ...grpc.CallOption) (*response.Cookbooks, error)
	GetCookbooksAvailableVersions(ctx context.Context, in *request.CookbooksAvailableVersions, opts ...grpc.CallOption) (*response.CookbooksAvailableVersions, error)
	GetCookbook(ctx context.Context, in *request.Cookbook, opts ...grpc.CallOption) (*response.Cookbook, error)
}

type infraProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewInfraProxyClient(cc grpc.ClientConnInterface) InfraProxyClient {
	return &infraProxyClient{cc}
}

func (c *infraProxyClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServers(ctx context.Context, in *request.GetServers, opts ...grpc.CallOption) (*response.GetServers, error) {
	out := new(response.GetServers)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServer(ctx context.Context, in *request.GetServer, opts ...grpc.CallOption) (*response.GetServer, error) {
	out := new(response.GetServer)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServerByName(ctx context.Context, in *request.GetServerByName, opts ...grpc.CallOption) (*response.GetServer, error) {
	out := new(response.GetServer)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetServerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) CreateServer(ctx context.Context, in *request.CreateServer, opts ...grpc.CallOption) (*response.CreateServer, error) {
	out := new(response.CreateServer)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) UpdateServer(ctx context.Context, in *request.UpdateServer, opts ...grpc.CallOption) (*response.UpdateServer, error) {
	out := new(response.UpdateServer)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) DeleteServer(ctx context.Context, in *request.DeleteServer, opts ...grpc.CallOption) (*response.DeleteServer, error) {
	out := new(response.DeleteServer)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrgs(ctx context.Context, in *request.GetOrgs, opts ...grpc.CallOption) (*response.GetOrgs, error) {
	out := new(response.GetOrgs)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrg(ctx context.Context, in *request.GetOrg, opts ...grpc.CallOption) (*response.GetOrg, error) {
	out := new(response.GetOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrgByName(ctx context.Context, in *request.GetOrgByName, opts ...grpc.CallOption) (*response.GetOrg, error) {
	out := new(response.GetOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetOrgByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) CreateOrg(ctx context.Context, in *request.CreateOrg, opts ...grpc.CallOption) (*response.CreateOrg, error) {
	out := new(response.CreateOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) UpdateOrg(ctx context.Context, in *request.UpdateOrg, opts ...grpc.CallOption) (*response.UpdateOrg, error) {
	out := new(response.UpdateOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) DeleteOrg(ctx context.Context, in *request.DeleteOrg, opts ...grpc.CallOption) (*response.DeleteOrg, error) {
	out := new(response.DeleteOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbooks(ctx context.Context, in *request.Cookbooks, opts ...grpc.CallOption) (*response.Cookbooks, error) {
	out := new(response.Cookbooks)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbooksAvailableVersions(ctx context.Context, in *request.CookbooksAvailableVersions, opts ...grpc.CallOption) (*response.CookbooksAvailableVersions, error) {
	out := new(response.CookbooksAvailableVersions)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetCookbooksAvailableVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbook(ctx context.Context, in *request.Cookbook, opts ...grpc.CallOption) (*response.Cookbook, error) {
	out := new(response.Cookbook)
	err := c.cc.Invoke(ctx, "/chef.automate.api.infra_proxy.InfraProxy/GetCookbook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraProxyServer is the server API for InfraProxy service.
type InfraProxyServer interface {
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	GetServers(context.Context, *request.GetServers) (*response.GetServers, error)
	GetServer(context.Context, *request.GetServer) (*response.GetServer, error)
	GetServerByName(context.Context, *request.GetServerByName) (*response.GetServer, error)
	CreateServer(context.Context, *request.CreateServer) (*response.CreateServer, error)
	UpdateServer(context.Context, *request.UpdateServer) (*response.UpdateServer, error)
	DeleteServer(context.Context, *request.DeleteServer) (*response.DeleteServer, error)
	GetOrgs(context.Context, *request.GetOrgs) (*response.GetOrgs, error)
	GetOrg(context.Context, *request.GetOrg) (*response.GetOrg, error)
	GetOrgByName(context.Context, *request.GetOrgByName) (*response.GetOrg, error)
	CreateOrg(context.Context, *request.CreateOrg) (*response.CreateOrg, error)
	UpdateOrg(context.Context, *request.UpdateOrg) (*response.UpdateOrg, error)
	DeleteOrg(context.Context, *request.DeleteOrg) (*response.DeleteOrg, error)
	GetCookbooks(context.Context, *request.Cookbooks) (*response.Cookbooks, error)
	GetCookbooksAvailableVersions(context.Context, *request.CookbooksAvailableVersions) (*response.CookbooksAvailableVersions, error)
	GetCookbook(context.Context, *request.Cookbook) (*response.Cookbook, error)
}

// UnimplementedInfraProxyServer can be embedded to have forward compatible implementations.
type UnimplementedInfraProxyServer struct {
}

func (*UnimplementedInfraProxyServer) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedInfraProxyServer) GetServers(ctx context.Context, req *request.GetServers) (*response.GetServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (*UnimplementedInfraProxyServer) GetServer(ctx context.Context, req *request.GetServer) (*response.GetServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (*UnimplementedInfraProxyServer) GetServerByName(ctx context.Context, req *request.GetServerByName) (*response.GetServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerByName not implemented")
}
func (*UnimplementedInfraProxyServer) CreateServer(ctx context.Context, req *request.CreateServer) (*response.CreateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (*UnimplementedInfraProxyServer) UpdateServer(ctx context.Context, req *request.UpdateServer) (*response.UpdateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (*UnimplementedInfraProxyServer) DeleteServer(ctx context.Context, req *request.DeleteServer) (*response.DeleteServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrgs(ctx context.Context, req *request.GetOrgs) (*response.GetOrgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgs not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrg(ctx context.Context, req *request.GetOrg) (*response.GetOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrgByName(ctx context.Context, req *request.GetOrgByName) (*response.GetOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgByName not implemented")
}
func (*UnimplementedInfraProxyServer) CreateOrg(ctx context.Context, req *request.CreateOrg) (*response.CreateOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (*UnimplementedInfraProxyServer) UpdateOrg(ctx context.Context, req *request.UpdateOrg) (*response.UpdateOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrg not implemented")
}
func (*UnimplementedInfraProxyServer) DeleteOrg(ctx context.Context, req *request.DeleteOrg) (*response.DeleteOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrg not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbooks(ctx context.Context, req *request.Cookbooks) (*response.Cookbooks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbooks not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbooksAvailableVersions(ctx context.Context, req *request.CookbooksAvailableVersions) (*response.CookbooksAvailableVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbooksAvailableVersions not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbook(ctx context.Context, req *request.Cookbook) (*response.Cookbook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbook not implemented")
}

func RegisterInfraProxyServer(s *grpc.Server, srv InfraProxyServer) {
	s.RegisterService(&_InfraProxy_serviceDesc, srv)
}

func _InfraProxy_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServers(ctx, req.(*request.GetServers))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServer(ctx, req.(*request.GetServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServerByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetServerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServerByName(ctx, req.(*request.GetServerByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).CreateServer(ctx, req.(*request.CreateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/UpdateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).UpdateServer(ctx, req.(*request.UpdateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/DeleteServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).DeleteServer(ctx, req.(*request.DeleteServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetOrgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrgs(ctx, req.(*request.GetOrgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrg(ctx, req.(*request.GetOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrgByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrgByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrgByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetOrgByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrgByName(ctx, req.(*request.GetOrgByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/CreateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).CreateOrg(ctx, req.(*request.CreateOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_UpdateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).UpdateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).UpdateOrg(ctx, req.(*request.UpdateOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_DeleteOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).DeleteOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).DeleteOrg(ctx, req.(*request.DeleteOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Cookbooks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbooks(ctx, req.(*request.Cookbooks))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbooksAvailableVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CookbooksAvailableVersions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbooksAvailableVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetCookbooksAvailableVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbooksAvailableVersions(ctx, req.(*request.CookbooksAvailableVersions))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Cookbook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.infra_proxy.InfraProxy/GetCookbook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbook(ctx, req.(*request.Cookbook))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfraProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.infra_proxy.InfraProxy",
	HandlerType: (*InfraProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _InfraProxy_GetVersion_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _InfraProxy_GetServers_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _InfraProxy_GetServer_Handler,
		},
		{
			MethodName: "GetServerByName",
			Handler:    _InfraProxy_GetServerByName_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _InfraProxy_CreateServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _InfraProxy_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _InfraProxy_DeleteServer_Handler,
		},
		{
			MethodName: "GetOrgs",
			Handler:    _InfraProxy_GetOrgs_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _InfraProxy_GetOrg_Handler,
		},
		{
			MethodName: "GetOrgByName",
			Handler:    _InfraProxy_GetOrgByName_Handler,
		},
		{
			MethodName: "CreateOrg",
			Handler:    _InfraProxy_CreateOrg_Handler,
		},
		{
			MethodName: "UpdateOrg",
			Handler:    _InfraProxy_UpdateOrg_Handler,
		},
		{
			MethodName: "DeleteOrg",
			Handler:    _InfraProxy_DeleteOrg_Handler,
		},
		{
			MethodName: "GetCookbooks",
			Handler:    _InfraProxy_GetCookbooks_Handler,
		},
		{
			MethodName: "GetCookbooksAvailableVersions",
			Handler:    _InfraProxy_GetCookbooksAvailableVersions_Handler,
		},
		{
			MethodName: "GetCookbook",
			Handler:    _InfraProxy_GetCookbook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/external/infra_proxy/infra_proxy.proto",
}
