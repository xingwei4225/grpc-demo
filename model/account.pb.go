// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package model

import (
	context "context"
	fmt "fmt"
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

type AccountId struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountId.Unmarshal(m, b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
}
func (m *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(m, src)
}
func (m *AccountId) XXX_Size() int {
	return xxx_messageInfo_AccountId.Size(m)
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ResponseMsg struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMsg) Reset()         { *m = ResponseMsg{} }
func (m *ResponseMsg) String() string { return proto.CompactTextString(m) }
func (*ResponseMsg) ProtoMessage()    {}
func (*ResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *ResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMsg.Unmarshal(m, b)
}
func (m *ResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMsg.Marshal(b, m, deterministic)
}
func (m *ResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMsg.Merge(m, src)
}
func (m *ResponseMsg) XXX_Size() int {
	return xxx_messageInfo_ResponseMsg.Size(m)
}
func (m *ResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMsg proto.InternalMessageInfo

func (m *ResponseMsg) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseMsg) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqMsg struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMsg) Reset()         { *m = ReqMsg{} }
func (m *ReqMsg) String() string { return proto.CompactTextString(m) }
func (*ReqMsg) ProtoMessage()    {}
func (*ReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *ReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMsg.Unmarshal(m, b)
}
func (m *ReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMsg.Marshal(b, m, deterministic)
}
func (m *ReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMsg.Merge(m, src)
}
func (m *ReqMsg) XXX_Size() int {
	return xxx_messageInfo_ReqMsg.Size(m)
}
func (m *ReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMsg proto.InternalMessageInfo

func (m *ReqMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AccountInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *AccountInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountId)(nil), "model.AccountId")
	proto.RegisterType((*ResponseMsg)(nil), "model.ResponseMsg")
	proto.RegisterType((*ReqMsg)(nil), "model.ReqMsg")
	proto.RegisterType((*AccountInfo)(nil), "model.AccountInfo")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x3b, 0xe9, 0xc7, 0x9f, 0xde, 0xd2, 0xfe, 0xc3, 0xa0, 0x10, 0x5a, 0x91, 0x32, 0xab,
	0xe2, 0xa2, 0xf1, 0x63, 0xa5, 0xe2, 0xa2, 0x2b, 0x11, 0x74, 0x93, 0xae, 0x5d, 0x8c, 0x99, 0xdb,
	0x18, 0x6c, 0xe6, 0xc6, 0xce, 0x58, 0x5c, 0xfb, 0x0a, 0x3e, 0x9a, 0x0f, 0xe0, 0xc6, 0x07, 0x91,
	0x4c, 0xd2, 0x12, 0x4b, 0x77, 0xf7, 0x0c, 0xe7, 0x9c, 0x1f, 0x73, 0x2f, 0xf4, 0x65, 0x1c, 0xd3,
	0x9b, 0xb6, 0xd3, 0x7c, 0x45, 0x96, 0x78, 0x3b, 0x23, 0x85, 0xcb, 0xe1, 0x51, 0x42, 0x94, 0x2c,
	0x31, 0x94, 0x79, 0x1a, 0x4a, 0xad, 0xc9, 0x4a, 0x9b, 0x92, 0x36, 0xa5, 0x49, 0x8c, 0xa0, 0x3b,
	0x2b, 0x53, 0x77, 0x8a, 0x0f, 0xc0, 0x4b, 0x55, 0xc0, 0xc6, 0x6c, 0xd2, 0x8f, 0xbc, 0x54, 0x89,
	0x6b, 0xe8, 0x45, 0x68, 0x72, 0xd2, 0x06, 0x1f, 0x4c, 0xc2, 0x39, 0xb4, 0x62, 0x52, 0x58, 0x19,
	0xdc, 0xcc, 0x03, 0xf8, 0x97, 0xa1, 0x31, 0x32, 0xc1, 0xc0, 0x1b, 0xb3, 0x49, 0x37, 0xda, 0x48,
	0x71, 0x0c, 0x9d, 0x08, 0x5f, 0x8b, 0xdc, 0x01, 0xb4, 0x2d, 0xbd, 0xa0, 0x76, 0xc1, 0x6e, 0x54,
	0x0a, 0xf1, 0x08, 0xbd, 0x0d, 0x59, 0x2f, 0x68, 0x97, 0x5d, 0xc0, 0xb4, 0xcc, 0x36, 0xad, 0x6e,
	0xe6, 0x3e, 0x34, 0x0b, 0x50, 0xd3, 0x99, 0x8a, 0xb1, 0xc0, 0x4b, 0xa5, 0x56, 0x68, 0x4c, 0xd0,
	0x2a, 0xf1, 0x95, 0x3c, 0xff, 0x66, 0x30, 0xa8, 0xfa, 0xe7, 0xb8, 0x5a, 0xa7, 0x31, 0xf2, 0x39,
	0x0c, 0x6e, 0xd1, 0xd6, 0xa1, 0xfe, 0xd4, 0xed, 0x68, 0xba, 0x5d, 0xc1, 0x90, 0xef, 0xbc, 0xe8,
	0x05, 0x89, 0xd1, 0xc7, 0xd7, 0xcf, 0xa7, 0x77, 0x28, 0xfc, 0x70, 0x7d, 0x16, 0xe2, 0xbb, 0xcc,
	0xf2, 0x25, 0x86, 0x18, 0x3f, 0xd3, 0x15, 0x3b, 0xe1, 0x97, 0xe0, 0xdf, 0xa7, 0xa6, 0xde, 0x6a,
	0x78, 0xbf, 0x2a, 0x29, 0xff, 0xbf, 0xb7, 0xb3, 0x71, 0xca, 0xf8, 0x0d, 0xfc, 0x9f, 0x29, 0xf5,
	0x27, 0xb9, 0xc7, 0xba, 0x8d, 0xd7, 0x4e, 0x21, 0x1a, 0x13, 0xf6, 0xd4, 0x71, 0x17, 0xbc, 0xf8,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x02, 0x25, 0x4b, 0xc5, 0xf7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	//根据标识获取账户详细信息
	GetAccountInfo(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*AccountInfo, error)
	//获取所有账户信息
	ListAccountInfos(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (AccountService_ListAccountInfosClient, error)
	//批量添加账户
	AddAccountInfos(ctx context.Context, opts ...grpc.CallOption) (AccountService_AddAccountInfosClient, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccountInfo(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/model.AccountService/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListAccountInfos(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (AccountService_ListAccountInfosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountService_serviceDesc.Streams[0], "/model.AccountService/ListAccountInfos", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountServiceListAccountInfosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountService_ListAccountInfosClient interface {
	Recv() (*AccountInfo, error)
	grpc.ClientStream
}

type accountServiceListAccountInfosClient struct {
	grpc.ClientStream
}

func (x *accountServiceListAccountInfosClient) Recv() (*AccountInfo, error) {
	m := new(AccountInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountServiceClient) AddAccountInfos(ctx context.Context, opts ...grpc.CallOption) (AccountService_AddAccountInfosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountService_serviceDesc.Streams[1], "/model.AccountService/AddAccountInfos", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountServiceAddAccountInfosClient{stream}
	return x, nil
}

type AccountService_AddAccountInfosClient interface {
	Send(*AccountInfo) error
	CloseAndRecv() (*ResponseMsg, error)
	grpc.ClientStream
}

type accountServiceAddAccountInfosClient struct {
	grpc.ClientStream
}

func (x *accountServiceAddAccountInfosClient) Send(m *AccountInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accountServiceAddAccountInfosClient) CloseAndRecv() (*ResponseMsg, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	//根据标识获取账户详细信息
	GetAccountInfo(context.Context, *AccountId) (*AccountInfo, error)
	//获取所有账户信息
	ListAccountInfos(*ReqMsg, AccountService_ListAccountInfosServer) error
	//批量添加账户
	AddAccountInfos(AccountService_AddAccountInfosServer) error
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) GetAccountInfo(ctx context.Context, req *AccountId) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (*UnimplementedAccountServiceServer) ListAccountInfos(req *ReqMsg, srv AccountService_ListAccountInfosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccountInfos not implemented")
}
func (*UnimplementedAccountServiceServer) AddAccountInfos(srv AccountService_AddAccountInfosServer) error {
	return status.Errorf(codes.Unimplemented, "method AddAccountInfos not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AccountService/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountInfo(ctx, req.(*AccountId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListAccountInfos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountServiceServer).ListAccountInfos(m, &accountServiceListAccountInfosServer{stream})
}

type AccountService_ListAccountInfosServer interface {
	Send(*AccountInfo) error
	grpc.ServerStream
}

type accountServiceListAccountInfosServer struct {
	grpc.ServerStream
}

func (x *accountServiceListAccountInfosServer) Send(m *AccountInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _AccountService_AddAccountInfos_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccountServiceServer).AddAccountInfos(&accountServiceAddAccountInfosServer{stream})
}

type AccountService_AddAccountInfosServer interface {
	SendAndClose(*ResponseMsg) error
	Recv() (*AccountInfo, error)
	grpc.ServerStream
}

type accountServiceAddAccountInfosServer struct {
	grpc.ServerStream
}

func (x *accountServiceAddAccountInfosServer) SendAndClose(m *ResponseMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accountServiceAddAccountInfosServer) Recv() (*AccountInfo, error) {
	m := new(AccountInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountInfo",
			Handler:    _AccountService_GetAccountInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccountInfos",
			Handler:       _AccountService_ListAccountInfos_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddAccountInfos",
			Handler:       _AccountService_AddAccountInfos_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "account.proto",
}
