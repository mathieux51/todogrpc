// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TodoByIDRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoByIDRequest) Reset()         { *m = TodoByIDRequest{} }
func (m *TodoByIDRequest) String() string { return proto.CompactTextString(m) }
func (*TodoByIDRequest) ProtoMessage()    {}
func (*TodoByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *TodoByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoByIDRequest.Unmarshal(m, b)
}
func (m *TodoByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoByIDRequest.Marshal(b, m, deterministic)
}
func (m *TodoByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoByIDRequest.Merge(m, src)
}
func (m *TodoByIDRequest) XXX_Size() int {
	return xxx_messageInfo_TodoByIDRequest.Size(m)
}
func (m *TodoByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodoByIDRequest proto.InternalMessageInfo

func (m *TodoByIDRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TodoByIDResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoByIDResponse) Reset()         { *m = TodoByIDResponse{} }
func (m *TodoByIDResponse) String() string { return proto.CompactTextString(m) }
func (*TodoByIDResponse) ProtoMessage()    {}
func (*TodoByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *TodoByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoByIDResponse.Unmarshal(m, b)
}
func (m *TodoByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoByIDResponse.Marshal(b, m, deterministic)
}
func (m *TodoByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoByIDResponse.Merge(m, src)
}
func (m *TodoByIDResponse) XXX_Size() int {
	return xxx_messageInfo_TodoByIDResponse.Size(m)
}
func (m *TodoByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TodoByIDResponse proto.InternalMessageInfo

func (m *TodoByIDResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TodoByIDResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *TodoByIDResponse) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type UploadImageRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

func (m *UploadImageRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadImageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageResponse) Reset()         { *m = UploadImageResponse{} }
func (m *UploadImageResponse) String() string { return proto.CompactTextString(m) }
func (*UploadImageResponse) ProtoMessage()    {}
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *UploadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageResponse.Unmarshal(m, b)
}
func (m *UploadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageResponse.Marshal(b, m, deterministic)
}
func (m *UploadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageResponse.Merge(m, src)
}
func (m *UploadImageResponse) XXX_Size() int {
	return xxx_messageInfo_UploadImageResponse.Size(m)
}
func (m *UploadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TodoByIDRequest)(nil), "todopb.TodoByIDRequest")
	proto.RegisterType((*TodoByIDResponse)(nil), "todopb.TodoByIDResponse")
	proto.RegisterType((*UploadImageRequest)(nil), "todopb.UploadImageRequest")
	proto.RegisterType((*UploadImageResponse)(nil), "todopb.UploadImageResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xcd, 0xba, 0x16, 0x3b, 0x15, 0x95, 0x11, 0x31, 0xac, 0x1e, 0xd6, 0x9c, 0x72, 0xda,
	0x83, 0xbe, 0x41, 0x11, 0xa4, 0x1e, 0x43, 0x7d, 0x80, 0xd4, 0x0c, 0x52, 0x68, 0x9d, 0xe8, 0x8e,
	0xa0, 0x2f, 0xe2, 0xf3, 0x4a, 0xb3, 0x84, 0xd5, 0xdd, 0xde, 0x7e, 0x32, 0x1f, 0xff, 0x7c, 0x13,
	0x00, 0xe1, 0xc0, 0x4d, 0xfc, 0x60, 0x61, 0x9c, 0xec, 0x72, 0x5c, 0x99, 0x5b, 0x38, 0x5b, 0x72,
	0xe0, 0xf9, 0xf7, 0xe2, 0xc1, 0xd1, 0xfb, 0x27, 0xb5, 0x82, 0xa7, 0x50, 0xac, 0x83, 0x56, 0xb5,
	0xb2, 0x47, 0xae, 0x58, 0x07, 0xb3, 0x84, 0xf3, 0x1e, 0x69, 0x23, 0xbf, 0xb5, 0x34, 0x64, 0x10,
	0xa1, 0x14, 0xfa, 0x12, 0x5d, 0xd4, 0xca, 0x4e, 0x5d, 0xca, 0x78, 0x03, 0xd3, 0x17, 0xde, 0xc6,
	0x0d, 0x09, 0x05, 0x7d, 0x58, 0x2b, 0x7b, 0xec, 0xfa, 0x07, 0x63, 0x01, 0x9f, 0xe3, 0x86, 0x7d,
	0x58, 0x6c, 0xfd, 0x2b, 0xe5, 0xdd, 0x08, 0x65, 0xf0, 0xe2, 0x53, 0xf3, 0x89, 0x4b, 0xd9, 0x5c,
	0xc2, 0xc5, 0x3f, 0xb2, 0x53, 0xb8, 0xfb, 0x51, 0x50, 0xee, 0xbc, 0x70, 0x0e, 0xb3, 0x47, 0x92,
	0xac, 0x88, 0x57, 0x4d, 0x77, 0x5a, 0x33, 0xb8, 0xab, 0xd2, 0xe3, 0x41, 0x57, 0x65, 0x0e, 0xf0,
	0x09, 0x66, 0x7f, 0x76, 0x60, 0x95, 0xd1, 0xb1, 0x62, 0x75, 0xbd, 0x77, 0x96, 0x9b, 0xac, 0x5a,
	0x4d, 0xd2, 0x0f, 0xdf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x77, 0x5a, 0x1f, 0x6f, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	GetTodoByID(ctx context.Context, in *TodoByIDRequest, opts ...grpc.CallOption) (*TodoByIDResponse, error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (Todo_UploadImageClient, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetTodoByID(ctx context.Context, in *TodoByIDRequest, opts ...grpc.CallOption) (*TodoByIDResponse, error) {
	out := new(TodoByIDResponse)
	err := c.cc.Invoke(ctx, "/todopb.Todo/GetTodoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (Todo_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Todo_serviceDesc.Streams[0], "/todopb.Todo/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoUploadImageClient{stream}
	return x, nil
}

type Todo_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type todoUploadImageClient struct {
	grpc.ClientStream
}

func (x *todoUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoUploadImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	GetTodoByID(context.Context, *TodoByIDRequest) (*TodoByIDResponse, error)
	UploadImage(Todo_UploadImageServer) error
}

// UnimplementedTodoServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (*UnimplementedTodoServer) GetTodoByID(ctx context.Context, req *TodoByIDRequest) (*TodoByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByID not implemented")
}
func (*UnimplementedTodoServer) UploadImage(srv Todo_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_GetTodoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTodoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.Todo/GetTodoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTodoByID(ctx, req.(*TodoByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoServer).UploadImage(&todoUploadImageServer{stream})
}

type Todo_UploadImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type todoUploadImageServer struct {
	grpc.ServerStream
}

func (x *todoUploadImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todopb.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodoByID",
			Handler:    _Todo_GetTodoByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _Todo_UploadImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "todo.proto",
}
