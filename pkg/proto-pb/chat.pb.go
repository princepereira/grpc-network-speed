// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package student

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

type Request struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	SlNo                 int32    `protobuf:"varint,2,opt,name=sl_no,json=slNo,proto3" json:"sl_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Request) GetSlNo() int32 {
	if m != nil {
		return m.SlNo
	}
	return 0
}

type Response struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	SlNo                 int32    `protobuf:"varint,2,opt,name=sl_no,json=slNo,proto3" json:"sl_no,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Response) GetSlNo() int32 {
	if m != nil {
		return m.SlNo
	}
	return 0
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "student.Request")
	proto.RegisterType((*Response)(nil), "student.Response")
}

func init() {
	proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54)
}

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e, 0x29, 0x4d, 0x49, 0xcd, 0x2b, 0x51,
	0x32, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe5, 0x62, 0xcb, 0xca,
	0x4f, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcd, 0xca, 0x4f, 0xf2,
	0x4c, 0x11, 0x12, 0xe6, 0x62, 0x2d, 0xce, 0x89, 0xcf, 0xcb, 0x97, 0x60, 0x52, 0x60, 0xd4, 0x60,
	0x0d, 0x62, 0x29, 0xce, 0xf1, 0xcb, 0x57, 0xf2, 0xe2, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x25, 0x45, 0x9f, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x33, 0x58,
	0x25, 0x98, 0x6d, 0xe4, 0xc0, 0xc5, 0x17, 0x0c, 0x71, 0x4d, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x90, 0x1e, 0x17, 0xbb, 0x7b, 0x6a, 0x89, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x80, 0x1e, 0xd4,
	0xa5, 0x7a, 0x50, 0x67, 0x4a, 0x09, 0x22, 0x89, 0x40, 0x5c, 0x90, 0xc4, 0x06, 0xf6, 0x94, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x37, 0xa2, 0x5f, 0x7f, 0xe2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	GetName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) GetName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	GetName(context.Context, *Request) (*Response, error)
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) GetName(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetName(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _StudentService_GetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
