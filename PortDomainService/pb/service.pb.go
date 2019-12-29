// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package ports

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

type SearchRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Port struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Latitude             float64  `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Province             string   `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	LocationName         string   `protobuf:"bytes,9,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Unlocs               []string `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string   `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
	Key                  string   `protobuf:"bytes,12,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Port) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Port) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type UpsertSummary struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	ElapsedTime          int64    `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertSummary) Reset()         { *m = UpsertSummary{} }
func (m *UpsertSummary) String() string { return proto.CompactTextString(m) }
func (*UpsertSummary) ProtoMessage()    {}
func (*UpsertSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *UpsertSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertSummary.Unmarshal(m, b)
}
func (m *UpsertSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertSummary.Marshal(b, m, deterministic)
}
func (m *UpsertSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertSummary.Merge(m, src)
}
func (m *UpsertSummary) XXX_Size() int {
	return xxx_messageInfo_UpsertSummary.Size(m)
}
func (m *UpsertSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertSummary.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertSummary proto.InternalMessageInfo

func (m *UpsertSummary) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UpsertSummary) GetElapsedTime() int64 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

// we may not have a port in DB, this is a poor's man Option<T> :)
type MaybePort struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Port                 *Port    `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaybePort) Reset()         { *m = MaybePort{} }
func (m *MaybePort) String() string { return proto.CompactTextString(m) }
func (*MaybePort) ProtoMessage()    {}
func (*MaybePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *MaybePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaybePort.Unmarshal(m, b)
}
func (m *MaybePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaybePort.Marshal(b, m, deterministic)
}
func (m *MaybePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaybePort.Merge(m, src)
}
func (m *MaybePort) XXX_Size() int {
	return xxx_messageInfo_MaybePort.Size(m)
}
func (m *MaybePort) XXX_DiscardUnknown() {
	xxx_messageInfo_MaybePort.DiscardUnknown(m)
}

var xxx_messageInfo_MaybePort proto.InternalMessageInfo

func (m *MaybePort) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *MaybePort) GetPort() *Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "ports.SearchRequest")
	proto.RegisterType((*Port)(nil), "ports.Port")
	proto.RegisterType((*UpsertSummary)(nil), "ports.UpsertSummary")
	proto.RegisterType((*MaybePort)(nil), "ports.MaybePort")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0x9b, 0xe6, 0xa3, 0x9b, 0xc9, 0x46, 0xaa, 0xac, 0x15, 0xb2, 0x56, 0x48, 0x2c, 0xe9,
	0x25, 0xa7, 0xad, 0x54, 0xde, 0x80, 0x0b, 0x5c, 0x40, 0xc8, 0x85, 0x73, 0xe5, 0x26, 0x43, 0xb1,
	0x48, 0xec, 0x60, 0x3b, 0x2b, 0xe5, 0xb1, 0x78, 0x43, 0xe4, 0x49, 0xb2, 0xb0, 0x12, 0xb7, 0xf9,
	0xff, 0xe6, 0x23, 0xfe, 0xcf, 0x04, 0x4a, 0x87, 0xf6, 0xa4, 0x1a, 0x3c, 0x0e, 0xd6, 0x78, 0xc3,
	0xd2, 0xc1, 0x58, 0xef, 0xaa, 0x3b, 0x28, 0x1f, 0x51, 0xda, 0xe6, 0x87, 0xc0, 0x5f, 0x23, 0x3a,
	0xcf, 0x18, 0x24, 0x5a, 0xf6, 0xc8, 0xa3, 0x43, 0x54, 0xe7, 0x82, 0xe2, 0xea, 0xf7, 0x35, 0x24,
	0x5f, 0x8c, 0xfd, 0x6f, 0x32, 0xb0, 0x46, 0xf9, 0x89, 0x5f, 0xcf, 0x2c, 0xc4, 0x8c, 0xc3, 0x4d,
	0x63, 0x46, 0xed, 0xed, 0xc4, 0x63, 0xc2, 0xab, 0x64, 0x3b, 0x48, 0x65, 0xa7, 0xa4, 0xe3, 0xc9,
	0x21, 0xae, 0x73, 0x31, 0x8b, 0x50, 0x6f, 0xf1, 0x45, 0x19, 0xed, 0x78, 0x4a, 0x7c, 0x95, 0x6c,
	0x0f, 0x9b, 0x4e, 0x7a, 0xe5, 0xc7, 0x16, 0x79, 0x76, 0x88, 0xea, 0x48, 0x9c, 0x35, 0x7b, 0x0d,
	0x79, 0x67, 0xf4, 0xcb, 0x9c, 0xbc, 0xa1, 0xe4, 0x5f, 0x10, 0x3a, 0x07, 0x6b, 0x4e, 0x4a, 0x37,
	0xc8, 0x37, 0xf4, 0x88, 0xb3, 0x66, 0x77, 0x50, 0x76, 0xa6, 0x91, 0x5e, 0x19, 0xfd, 0x44, 0x86,
	0x72, 0x2a, 0xd8, 0xae, 0xf0, 0x73, 0x30, 0xf6, 0x0a, 0xb2, 0x51, 0x77, 0xa6, 0x71, 0x1c, 0xe8,
	0x4d, 0x8b, 0x22, 0xc3, 0xa6, 0x45, 0x5e, 0x2c, 0x86, 0x4d, 0x8b, 0xec, 0x16, 0xe2, 0x9f, 0x38,
	0xf1, 0x2d, 0xa1, 0x10, 0x56, 0x1f, 0xa1, 0xfc, 0x36, 0x38, 0xb4, 0xfe, 0x71, 0xec, 0x7b, 0x39,
	0x3b, 0xa7, 0x25, 0xd0, 0xf2, 0x62, 0x31, 0x0b, 0xf6, 0x16, 0xb6, 0xd8, 0xc9, 0xc1, 0x61, 0xfb,
	0xe4, 0x55, 0x8f, 0xb4, 0xc5, 0x58, 0x14, 0x0b, 0xfb, 0xaa, 0x7a, 0xac, 0xde, 0x43, 0xfe, 0x49,
	0x4e, 0xcf, 0x48, 0x17, 0xd8, 0x41, 0xfa, 0xdd, 0x8c, 0xba, 0xa5, 0x29, 0x1b, 0x31, 0x0b, 0xf6,
	0x06, 0x92, 0x70, 0x4e, 0xea, 0x2e, 0x1e, 0x8a, 0x23, 0xdd, 0xf6, 0x18, 0x1a, 0x04, 0x25, 0x1e,
	0x14, 0xa4, 0x41, 0x39, 0x76, 0x0f, 0xf1, 0x07, 0xf4, 0x6c, 0xb7, 0x94, 0x5c, 0xdc, 0x7e, 0x7f,
	0xbb, 0xd0, 0xf3, 0xe7, 0xaa, 0x2b, 0x76, 0x0f, 0xd9, 0xec, 0x83, 0xfd, 0x3b, 0x76, 0xbf, 0x0e,
	0xb8, 0xf0, 0x58, 0x5d, 0xd5, 0xd1, 0x73, 0x46, 0xff, 0xd7, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x14, 0x11, 0x33, 0x70, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortsClient is the client API for Ports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortsClient interface {
	Get(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*MaybePort, error)
	Upsert(ctx context.Context, opts ...grpc.CallOption) (Ports_UpsertClient, error)
}

type portsClient struct {
	cc *grpc.ClientConn
}

func NewPortsClient(cc *grpc.ClientConn) PortsClient {
	return &portsClient{cc}
}

func (c *portsClient) Get(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*MaybePort, error) {
	out := new(MaybePort)
	err := c.cc.Invoke(ctx, "/ports.Ports/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsClient) Upsert(ctx context.Context, opts ...grpc.CallOption) (Ports_UpsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ports_serviceDesc.Streams[0], "/ports.Ports/Upsert", opts...)
	if err != nil {
		return nil, err
	}
	x := &portsUpsertClient{stream}
	return x, nil
}

type Ports_UpsertClient interface {
	Send(*Port) error
	CloseAndRecv() (*UpsertSummary, error)
	grpc.ClientStream
}

type portsUpsertClient struct {
	grpc.ClientStream
}

func (x *portsUpsertClient) Send(m *Port) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portsUpsertClient) CloseAndRecv() (*UpsertSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpsertSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortsServer is the server API for Ports service.
type PortsServer interface {
	Get(context.Context, *SearchRequest) (*MaybePort, error)
	Upsert(Ports_UpsertServer) error
}

// UnimplementedPortsServer can be embedded to have forward compatible implementations.
type UnimplementedPortsServer struct {
}

func (*UnimplementedPortsServer) Get(ctx context.Context, req *SearchRequest) (*MaybePort, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedPortsServer) Upsert(srv Ports_UpsertServer) error {
	return status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}

func RegisterPortsServer(s *grpc.Server, srv PortsServer) {
	s.RegisterService(&_Ports_serviceDesc, srv)
}

func _Ports_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ports.Ports/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsServer).Get(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ports_Upsert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PortsServer).Upsert(&portsUpsertServer{stream})
}

type Ports_UpsertServer interface {
	SendAndClose(*UpsertSummary) error
	Recv() (*Port, error)
	grpc.ServerStream
}

type portsUpsertServer struct {
	grpc.ServerStream
}

func (x *portsUpsertServer) SendAndClose(m *UpsertSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portsUpsertServer) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ports_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ports.Ports",
	HandlerType: (*PortsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Ports_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upsert",
			Handler:       _Ports_Upsert_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
