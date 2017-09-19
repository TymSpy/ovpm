// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetworkCreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cidr string `protobuf:"bytes,2,opt,name=cidr" json:"cidr,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Via  string `protobuf:"bytes,4,opt,name=via" json:"via,omitempty"`
}

func (m *NetworkCreateRequest) Reset()                    { *m = NetworkCreateRequest{} }
func (m *NetworkCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkCreateRequest) ProtoMessage()               {}
func (*NetworkCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NetworkCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkCreateRequest) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *NetworkCreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkCreateRequest) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

type NetworkListRequest struct {
}

func (m *NetworkListRequest) Reset()                    { *m = NetworkListRequest{} }
func (m *NetworkListRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkListRequest) ProtoMessage()               {}
func (*NetworkListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type NetworkDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NetworkDeleteRequest) Reset()                    { *m = NetworkDeleteRequest{} }
func (m *NetworkDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkDeleteRequest) ProtoMessage()               {}
func (*NetworkDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *NetworkDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetworkGetAllTypesRequest struct {
}

func (m *NetworkGetAllTypesRequest) Reset()                    { *m = NetworkGetAllTypesRequest{} }
func (m *NetworkGetAllTypesRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkGetAllTypesRequest) ProtoMessage()               {}
func (*NetworkGetAllTypesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type NetworkAssociateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
}

func (m *NetworkAssociateRequest) Reset()                    { *m = NetworkAssociateRequest{} }
func (m *NetworkAssociateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkAssociateRequest) ProtoMessage()               {}
func (*NetworkAssociateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *NetworkAssociateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkAssociateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type NetworkDissociateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
}

func (m *NetworkDissociateRequest) Reset()                    { *m = NetworkDissociateRequest{} }
func (m *NetworkDissociateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkDissociateRequest) ProtoMessage()               {}
func (*NetworkDissociateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *NetworkDissociateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkDissociateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type NetworkGetAssociatedUsersRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NetworkGetAssociatedUsersRequest) Reset()         { *m = NetworkGetAssociatedUsersRequest{} }
func (m *NetworkGetAssociatedUsersRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkGetAssociatedUsersRequest) ProtoMessage()    {}
func (*NetworkGetAssociatedUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{6}
}

func (m *NetworkGetAssociatedUsersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Network struct {
	Name                string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cidr                string   `protobuf:"bytes,2,opt,name=cidr" json:"cidr,omitempty"`
	Type                string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	CreatedAt           string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	AssociatedUsernames []string `protobuf:"bytes,5,rep,name=associated_usernames,json=associatedUsernames" json:"associated_usernames,omitempty"`
	Via                 string   `protobuf:"bytes,6,opt,name=via" json:"via,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Network) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Network) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Network) GetAssociatedUsernames() []string {
	if m != nil {
		return m.AssociatedUsernames
	}
	return nil
}

func (m *Network) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

type NetworkType struct {
	Type        string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *NetworkType) Reset()                    { *m = NetworkType{} }
func (m *NetworkType) String() string            { return proto.CompactTextString(m) }
func (*NetworkType) ProtoMessage()               {}
func (*NetworkType) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *NetworkType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NetworkCreateResponse struct {
	Network *Network `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
}

func (m *NetworkCreateResponse) Reset()                    { *m = NetworkCreateResponse{} }
func (m *NetworkCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkCreateResponse) ProtoMessage()               {}
func (*NetworkCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *NetworkCreateResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type NetworkListResponse struct {
	Networks []*Network `protobuf:"bytes,1,rep,name=networks" json:"networks,omitempty"`
}

func (m *NetworkListResponse) Reset()                    { *m = NetworkListResponse{} }
func (m *NetworkListResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkListResponse) ProtoMessage()               {}
func (*NetworkListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *NetworkListResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

type NetworkDeleteResponse struct {
	Network *Network `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
}

func (m *NetworkDeleteResponse) Reset()                    { *m = NetworkDeleteResponse{} }
func (m *NetworkDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkDeleteResponse) ProtoMessage()               {}
func (*NetworkDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *NetworkDeleteResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type NetworkGetAllTypesResponse struct {
	Types []*NetworkType `protobuf:"bytes,1,rep,name=types" json:"types,omitempty"`
}

func (m *NetworkGetAllTypesResponse) Reset()                    { *m = NetworkGetAllTypesResponse{} }
func (m *NetworkGetAllTypesResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkGetAllTypesResponse) ProtoMessage()               {}
func (*NetworkGetAllTypesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *NetworkGetAllTypesResponse) GetTypes() []*NetworkType {
	if m != nil {
		return m.Types
	}
	return nil
}

type NetworkAssociateResponse struct {
}

func (m *NetworkAssociateResponse) Reset()                    { *m = NetworkAssociateResponse{} }
func (m *NetworkAssociateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkAssociateResponse) ProtoMessage()               {}
func (*NetworkAssociateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

type NetworkDissociateResponse struct {
}

func (m *NetworkDissociateResponse) Reset()                    { *m = NetworkDissociateResponse{} }
func (m *NetworkDissociateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkDissociateResponse) ProtoMessage()               {}
func (*NetworkDissociateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

type NetworkGetAssociatedUsersResponse struct {
	Usernames []string `protobuf:"bytes,1,rep,name=usernames" json:"usernames,omitempty"`
}

func (m *NetworkGetAssociatedUsersResponse) Reset()         { *m = NetworkGetAssociatedUsersResponse{} }
func (m *NetworkGetAssociatedUsersResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkGetAssociatedUsersResponse) ProtoMessage()    {}
func (*NetworkGetAssociatedUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{15}
}

func (m *NetworkGetAssociatedUsersResponse) GetUsernames() []string {
	if m != nil {
		return m.Usernames
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkCreateRequest)(nil), "pb.NetworkCreateRequest")
	proto.RegisterType((*NetworkListRequest)(nil), "pb.NetworkListRequest")
	proto.RegisterType((*NetworkDeleteRequest)(nil), "pb.NetworkDeleteRequest")
	proto.RegisterType((*NetworkGetAllTypesRequest)(nil), "pb.NetworkGetAllTypesRequest")
	proto.RegisterType((*NetworkAssociateRequest)(nil), "pb.NetworkAssociateRequest")
	proto.RegisterType((*NetworkDissociateRequest)(nil), "pb.NetworkDissociateRequest")
	proto.RegisterType((*NetworkGetAssociatedUsersRequest)(nil), "pb.NetworkGetAssociatedUsersRequest")
	proto.RegisterType((*Network)(nil), "pb.Network")
	proto.RegisterType((*NetworkType)(nil), "pb.NetworkType")
	proto.RegisterType((*NetworkCreateResponse)(nil), "pb.NetworkCreateResponse")
	proto.RegisterType((*NetworkListResponse)(nil), "pb.NetworkListResponse")
	proto.RegisterType((*NetworkDeleteResponse)(nil), "pb.NetworkDeleteResponse")
	proto.RegisterType((*NetworkGetAllTypesResponse)(nil), "pb.NetworkGetAllTypesResponse")
	proto.RegisterType((*NetworkAssociateResponse)(nil), "pb.NetworkAssociateResponse")
	proto.RegisterType((*NetworkDissociateResponse)(nil), "pb.NetworkDissociateResponse")
	proto.RegisterType((*NetworkGetAssociatedUsersResponse)(nil), "pb.NetworkGetAssociatedUsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkService service

type NetworkServiceClient interface {
	Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error)
	List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error)
	Delete(ctx context.Context, in *NetworkDeleteRequest, opts ...grpc.CallOption) (*NetworkDeleteResponse, error)
	GetAllTypes(ctx context.Context, in *NetworkGetAllTypesRequest, opts ...grpc.CallOption) (*NetworkGetAllTypesResponse, error)
	GetAssociatedUsers(ctx context.Context, in *NetworkGetAssociatedUsersRequest, opts ...grpc.CallOption) (*NetworkGetAssociatedUsersResponse, error)
	Associate(ctx context.Context, in *NetworkAssociateRequest, opts ...grpc.CallOption) (*NetworkAssociateResponse, error)
	Dissociate(ctx context.Context, in *NetworkDissociateRequest, opts ...grpc.CallOption) (*NetworkDissociateResponse, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error) {
	out := new(NetworkCreateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error) {
	out := new(NetworkListResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Delete(ctx context.Context, in *NetworkDeleteRequest, opts ...grpc.CallOption) (*NetworkDeleteResponse, error) {
	out := new(NetworkDeleteResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetAllTypes(ctx context.Context, in *NetworkGetAllTypesRequest, opts ...grpc.CallOption) (*NetworkGetAllTypesResponse, error) {
	out := new(NetworkGetAllTypesResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/GetAllTypes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetAssociatedUsers(ctx context.Context, in *NetworkGetAssociatedUsersRequest, opts ...grpc.CallOption) (*NetworkGetAssociatedUsersResponse, error) {
	out := new(NetworkGetAssociatedUsersResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/GetAssociatedUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Associate(ctx context.Context, in *NetworkAssociateRequest, opts ...grpc.CallOption) (*NetworkAssociateResponse, error) {
	out := new(NetworkAssociateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Associate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Dissociate(ctx context.Context, in *NetworkDissociateRequest, opts ...grpc.CallOption) (*NetworkDissociateResponse, error) {
	out := new(NetworkDissociateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Dissociate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkService service

type NetworkServiceServer interface {
	Create(context.Context, *NetworkCreateRequest) (*NetworkCreateResponse, error)
	List(context.Context, *NetworkListRequest) (*NetworkListResponse, error)
	Delete(context.Context, *NetworkDeleteRequest) (*NetworkDeleteResponse, error)
	GetAllTypes(context.Context, *NetworkGetAllTypesRequest) (*NetworkGetAllTypesResponse, error)
	GetAssociatedUsers(context.Context, *NetworkGetAssociatedUsersRequest) (*NetworkGetAssociatedUsersResponse, error)
	Associate(context.Context, *NetworkAssociateRequest) (*NetworkAssociateResponse, error)
	Dissociate(context.Context, *NetworkDissociateRequest) (*NetworkDissociateResponse, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Create(ctx, req.(*NetworkCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).List(ctx, req.(*NetworkListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Delete(ctx, req.(*NetworkDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetAllTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkGetAllTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetAllTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/GetAllTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetAllTypes(ctx, req.(*NetworkGetAllTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetAssociatedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkGetAssociatedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetAssociatedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/GetAssociatedUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetAssociatedUsers(ctx, req.(*NetworkGetAssociatedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Associate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkAssociateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Associate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Associate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Associate(ctx, req.(*NetworkAssociateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Dissociate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkDissociateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Dissociate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Dissociate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Dissociate(ctx, req.(*NetworkDissociateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkService_Delete_Handler,
		},
		{
			MethodName: "GetAllTypes",
			Handler:    _NetworkService_GetAllTypes_Handler,
		},
		{
			MethodName: "GetAssociatedUsers",
			Handler:    _NetworkService_GetAssociatedUsers_Handler,
		},
		{
			MethodName: "Associate",
			Handler:    _NetworkService_Associate_Handler,
		},
		{
			MethodName: "Dissociate",
			Handler:    _NetworkService_Dissociate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}

func init() { proto.RegisterFile("network.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0x81, 0x90, 0x30, 0xa8, 0x1f, 0xda, 0x50, 0x62, 0x0c, 0xa4, 0x64, 0x0b, 0x6a, 0xc4,
	0x01, 0x94, 0x54, 0xea, 0xa1, 0x87, 0x48, 0x88, 0x4a, 0x55, 0xab, 0xaa, 0x07, 0xda, 0x9c, 0x13,
	0x83, 0x57, 0xc8, 0xaa, 0x6b, 0x3b, 0xde, 0x85, 0x2a, 0xd7, 0xde, 0xdb, 0x4b, 0x7f, 0x45, 0x7f,
	0x4f, 0xff, 0x42, 0x7f, 0x48, 0xb5, 0xeb, 0xb1, 0x59, 0x9b, 0x8f, 0x28, 0xca, 0x6d, 0x99, 0x99,
	0x7d, 0xef, 0x79, 0xdf, 0xd3, 0x00, 0x8f, 0x7c, 0x26, 0xbe, 0x07, 0xd1, 0xd7, 0x41, 0x18, 0x05,
	0x22, 0x20, 0x85, 0x70, 0x6a, 0xb5, 0xe6, 0x41, 0x30, 0xf7, 0xd8, 0xd0, 0x0e, 0xdd, 0xa1, 0xed,
	0xfb, 0x81, 0xb0, 0x85, 0x1b, 0xf8, 0x3c, 0x9e, 0xa0, 0x0e, 0xd4, 0x3e, 0xc5, 0x57, 0xc6, 0x11,
	0xb3, 0x05, 0x9b, 0xb0, 0x9b, 0x05, 0xe3, 0x82, 0x10, 0x28, 0xf9, 0xf6, 0x37, 0x66, 0x1a, 0x1d,
	0xe3, 0xb4, 0x32, 0x51, 0x67, 0x59, 0x9b, 0xb9, 0x4e, 0x64, 0x16, 0xe2, 0x9a, 0x3c, 0xcb, 0x9a,
	0xb8, 0x0d, 0x99, 0x59, 0x8c, 0x6b, 0xf2, 0x4c, 0x9e, 0x42, 0x71, 0xe9, 0xda, 0x66, 0x49, 0x95,
	0xe4, 0x91, 0xd6, 0x80, 0x20, 0xcb, 0x47, 0x97, 0x0b, 0xe4, 0xa0, 0xfd, 0x94, 0xfb, 0x2d, 0xf3,
	0xd8, 0x4e, 0x6e, 0xda, 0x84, 0x06, 0xce, 0xbe, 0x63, 0x62, 0xe4, 0x79, 0x5f, 0x6e, 0x43, 0xc6,
	0x13, 0xa0, 0xf7, 0x70, 0x84, 0xcd, 0x11, 0xe7, 0xc1, 0xcc, 0xbd, 0xe3, 0x3b, 0x2c, 0x38, 0x58,
	0x70, 0x16, 0xa9, 0x7a, 0xfc, 0x2d, 0xe9, 0x6f, 0xfa, 0x01, 0xcc, 0x44, 0x93, 0xfb, 0x50, 0xac,
	0xd7, 0xd0, 0xd1, 0x34, 0x27, 0x68, 0xce, 0x25, 0x67, 0x11, 0xdf, 0xf5, 0xad, 0x7f, 0x0c, 0xd8,
	0xc7, 0x8b, 0x0f, 0xf2, 0xa1, 0x0d, 0x30, 0x53, 0xa6, 0x3a, 0x57, 0xb6, 0x40, 0x3b, 0x2a, 0x58,
	0x19, 0x09, 0x72, 0x06, 0x35, 0x3b, 0x15, 0x75, 0x95, 0xa8, 0xe6, 0xe6, 0x5e, 0xa7, 0x78, 0x5a,
	0x99, 0x1c, 0xda, 0x19, 0xc1, 0xaa, 0x95, 0x38, 0x5b, 0x5e, 0x39, 0x3b, 0x86, 0x2a, 0x4a, 0x95,
	0x8e, 0xa4, 0x32, 0x0c, 0x4d, 0x46, 0x07, 0xaa, 0x0e, 0xe3, 0xb3, 0xc8, 0x0d, 0x65, 0xf0, 0x50,
	0xb5, 0x5e, 0xa2, 0x17, 0xf0, 0x2c, 0x17, 0x42, 0x1e, 0x06, 0x3e, 0x67, 0xa4, 0x07, 0xfb, 0x18,
	0x68, 0x85, 0x58, 0x3d, 0xaf, 0x0e, 0xc2, 0xe9, 0x00, 0x67, 0x27, 0x49, 0x8f, 0x5e, 0xc0, 0x61,
	0x26, 0x5e, 0x78, 0xfb, 0x25, 0x1c, 0xe0, 0x04, 0x37, 0x8d, 0x4e, 0x31, 0x7f, 0x3d, 0x6d, 0x6a,
	0xfc, 0x49, 0x10, 0xef, 0xc7, 0x3f, 0x06, 0x6b, 0x53, 0x38, 0x53, 0x90, 0x3d, 0xf9, 0x0e, 0x89,
	0x86, 0x27, 0x1a, 0x84, 0x1c, 0x9c, 0xc4, 0x5d, 0x6a, 0xa5, 0xc9, 0xd3, 0x42, 0x1c, 0x43, 0x68,
	0xe9, 0xd7, 0x53, 0x89, 0xcd, 0x11, 0x9c, 0xec, 0x88, 0x19, 0x8a, 0x68, 0x41, 0x65, 0xe5, 0xb0,
	0xa1, 0x1c, 0x5e, 0x15, 0xce, 0x7f, 0x95, 0xe1, 0x31, 0x62, 0x7c, 0x66, 0xd1, 0xd2, 0x9d, 0x31,
	0x72, 0x0d, 0xe5, 0xd8, 0x0c, 0x62, 0x6a, 0x82, 0x33, 0x4b, 0xc2, 0x6a, 0x6c, 0xe8, 0xa0, 0xa8,
	0x93, 0x1f, 0x7f, 0xff, 0xfd, 0x2e, 0x34, 0x69, 0x5d, 0xed, 0x9d, 0xe5, 0xd9, 0x10, 0xdf, 0x6a,
	0x18, 0xe7, 0xef, 0x8d, 0xd1, 0x27, 0x97, 0x50, 0x92, 0x76, 0x91, 0xba, 0x86, 0xa2, 0xad, 0x07,
	0xeb, 0x68, 0xad, 0x8e, 0xd8, 0x2d, 0x85, 0x5d, 0x27, 0xb5, 0x3c, 0xb6, 0x27, 0xe1, 0xae, 0xa1,
	0x1c, 0xbb, 0x98, 0x11, 0x9e, 0xd9, 0x30, 0x19, 0xe1, 0x59, 0xcb, 0xb7, 0x0b, 0x77, 0xd4, 0x9c,
	0x14, 0x7e, 0x03, 0x55, 0xcd, 0x67, 0xd2, 0xd6, 0xc0, 0xd6, 0x97, 0x93, 0x75, 0xbc, 0xad, 0x8d,
	0x84, 0x2f, 0x14, 0x61, 0x9b, 0x34, 0xf3, 0x84, 0x73, 0x26, 0x6c, 0xcf, 0x53, 0xe1, 0x20, 0x3f,
	0x0d, 0x20, 0xeb, 0xee, 0x92, 0x6e, 0x0e, 0x7b, 0xe3, 0x8e, 0xb1, 0x7a, 0x77, 0x4c, 0xa1, 0x90,
	0xbe, 0x12, 0xd2, 0x25, 0x74, 0x93, 0x90, 0xf4, 0xce, 0x42, 0x11, 0x7b, 0x50, 0x49, 0x61, 0x48,
	0x53, 0xc3, 0xcf, 0x2f, 0x60, 0xab, 0xb5, 0xb9, 0x89, 0x9c, 0x5d, 0xc5, 0x79, 0x4c, 0x1b, 0x79,
	0xce, 0x94, 0x50, 0x3e, 0x78, 0x08, 0xb0, 0xca, 0x3d, 0xd1, 0x11, 0xd7, 0x96, 0xb4, 0xd5, 0xde,
	0xd2, 0x45, 0xc2, 0x9e, 0x22, 0x7c, 0x4e, 0xad, 0x35, 0x7b, 0x5d, 0x8d, 0x71, 0x5a, 0x56, 0xff,
	0x8e, 0xaf, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xab, 0x14, 0xfe, 0x39, 0x50, 0x07, 0x00, 0x00,
}
