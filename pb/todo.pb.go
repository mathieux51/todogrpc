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

func init() {
	proto.RegisterType((*TodoByIDRequest)(nil), "todopb.TodoByIDRequest")
	proto.RegisterType((*TodoByIDResponse)(nil), "todopb.TodoByIDResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4f, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0xb1, 0x0b, 0x92, 0x94, 0x14, 0xb9, 0xf8,
	0x43, 0xf2, 0x53, 0xf2, 0x9d, 0x2a, 0x3d, 0x5d, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x94,
	0x42, 0xb8, 0x04, 0x10, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xd1, 0xd5, 0x08, 0x09, 0x71,
	0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x32,
	0x5c, 0x9c, 0xc9, 0xf9, 0xb9, 0x05, 0x39, 0xa9, 0x25, 0xa9, 0x29, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0x1c, 0x41, 0x08, 0x01, 0x23, 0x2f, 0x2e, 0x16, 0x90, 0xa9, 0x42, 0x4e, 0x5c, 0xdc, 0xee, 0xa9,
	0x25, 0x30, 0x0b, 0x84, 0xc4, 0xf5, 0x20, 0x0e, 0xd3, 0x43, 0x73, 0x95, 0x94, 0x04, 0xa6, 0x04,
	0xc4, 0x2d, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x3f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86,
	0x47, 0x93, 0x80, 0xe1, 0x00, 0x00, 0x00,
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

// TodoServer is the server API for Todo service.
type TodoServer interface {
	GetTodoByID(context.Context, *TodoByIDRequest) (*TodoByIDResponse, error)
}

// UnimplementedTodoServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (*UnimplementedTodoServer) GetTodoByID(ctx context.Context, req *TodoByIDRequest) (*TodoByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByID not implemented")
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

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todopb.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodoByID",
			Handler:    _Todo_GetTodoByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
