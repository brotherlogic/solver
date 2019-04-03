// Code generated by protoc-gen-go. DO NOT EDIT.
// source: solver.proto

package solver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SolveRequest struct {
	KeyStart             int64    `protobuf:"varint,1,opt,name=key_start,json=keyStart,proto3" json:"key_start,omitempty"`
	KeyEnd               int64    `protobuf:"varint,2,opt,name=key_end,json=keyEnd,proto3" json:"key_end,omitempty"`
	Step                 int64    `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SolveRequest) Reset()         { *m = SolveRequest{} }
func (m *SolveRequest) String() string { return proto.CompactTextString(m) }
func (*SolveRequest) ProtoMessage()    {}
func (*SolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7779dcc6dfedf133, []int{0}
}

func (m *SolveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SolveRequest.Unmarshal(m, b)
}
func (m *SolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SolveRequest.Marshal(b, m, deterministic)
}
func (m *SolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SolveRequest.Merge(m, src)
}
func (m *SolveRequest) XXX_Size() int {
	return xxx_messageInfo_SolveRequest.Size(m)
}
func (m *SolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SolveRequest proto.InternalMessageInfo

func (m *SolveRequest) GetKeyStart() int64 {
	if m != nil {
		return m.KeyStart
	}
	return 0
}

func (m *SolveRequest) GetKeyEnd() int64 {
	if m != nil {
		return m.KeyEnd
	}
	return 0
}

func (m *SolveRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

type SolveResponse struct {
	Solution             int64    `protobuf:"varint,1,opt,name=solution,proto3" json:"solution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SolveResponse) Reset()         { *m = SolveResponse{} }
func (m *SolveResponse) String() string { return proto.CompactTextString(m) }
func (*SolveResponse) ProtoMessage()    {}
func (*SolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7779dcc6dfedf133, []int{1}
}

func (m *SolveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SolveResponse.Unmarshal(m, b)
}
func (m *SolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SolveResponse.Marshal(b, m, deterministic)
}
func (m *SolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SolveResponse.Merge(m, src)
}
func (m *SolveResponse) XXX_Size() int {
	return xxx_messageInfo_SolveResponse.Size(m)
}
func (m *SolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SolveResponse proto.InternalMessageInfo

func (m *SolveResponse) GetSolution() int64 {
	if m != nil {
		return m.Solution
	}
	return 0
}

func init() {
	proto.RegisterType((*SolveRequest)(nil), "solver.SolveRequest")
	proto.RegisterType((*SolveResponse)(nil), "solver.SolveResponse")
}

func init() { proto.RegisterFile("solver.proto", fileDescriptor_7779dcc6dfedf133) }

var fileDescriptor_7779dcc6dfedf133 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0xcf, 0x29,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x22, 0xb8, 0x78,
	0x82, 0x41, 0xac, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x69, 0x2e, 0xce, 0xec, 0xd4,
	0xca, 0xf8, 0xe2, 0x92, 0xc4, 0xa2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x8e, 0xec,
	0xd4, 0xca, 0x60, 0x10, 0x5f, 0x48, 0x9c, 0x8b, 0x1d, 0x24, 0x99, 0x9a, 0x97, 0x22, 0xc1, 0x04,
	0x96, 0x62, 0xcb, 0x4e, 0xad, 0x74, 0xcd, 0x4b, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x2e, 0x49, 0x2d,
	0x90, 0x60, 0x06, 0x8b, 0x82, 0xd9, 0x4a, 0xda, 0x5c, 0xbc, 0x50, 0x93, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x8a, 0xf3, 0x73, 0x4a, 0x4b, 0x32, 0xf3, 0xf3, 0x60, 0x26,
	0xc3, 0xf8, 0x46, 0xee, 0x50, 0xc5, 0x45, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x66,
	0x5c, 0xac, 0x60, 0x01, 0x21, 0x11, 0x3d, 0xa8, 0xbb, 0x91, 0x9d, 0x29, 0x25, 0x8a, 0x26, 0x0a,
	0xb1, 0x42, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x3d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0xd6, 0xc9, 0xf9, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SolverServiceClient is the client API for SolverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SolverServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type solverServiceClient struct {
	cc *grpc.ClientConn
}

func NewSolverServiceClient(cc *grpc.ClientConn) SolverServiceClient {
	return &solverServiceClient{cc}
}

func (c *solverServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, "/solver.SolverService/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolverServiceServer is the server API for SolverService service.
type SolverServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

func RegisterSolverServiceServer(s *grpc.Server, srv SolverServiceServer) {
	s.RegisterService(&_SolverService_serviceDesc, srv)
}

func _SolverService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolverServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solver.SolverService/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolverServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SolverService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "solver.SolverService",
	HandlerType: (*SolverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _SolverService_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solver.proto",
}
