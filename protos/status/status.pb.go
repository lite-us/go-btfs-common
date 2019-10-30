// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/status/status.proto

package status

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	shared "github.com/tron-us/go-btfs-common/protos/shared"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SignedMetrics struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedMetrics) Reset()         { *m = SignedMetrics{} }
func (m *SignedMetrics) String() string { return proto.CompactTextString(m) }
func (*SignedMetrics) ProtoMessage()    {}
func (*SignedMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9255cc60c5ca429, []int{0}
}
func (m *SignedMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedMetrics.Unmarshal(m, b)
}
func (m *SignedMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedMetrics.Marshal(b, m, deterministic)
}
func (m *SignedMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMetrics.Merge(m, src)
}
func (m *SignedMetrics) XXX_Size() int {
	return xxx_messageInfo_SignedMetrics.Size(m)
}
func (m *SignedMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMetrics proto.InternalMessageInfo

func (m *SignedMetrics) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignedMetrics) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMetrics) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type NodeHealth struct {
	NodeId               string     `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	BtfsVersion          string     `protobuf:"bytes,2,opt,name=btfs_version,json=btfsVersion,proto3" json:"btfs_version,omitempty"`
	FailurePoint         string     `protobuf:"bytes,3,opt,name=failure_point,json=failurePoint,proto3" json:"failure_point,omitempty"`
	TimeCreated          *time.Time `protobuf:"bytes,4,opt,name=time_created,json=timeCreated,proto3,stdtime" json:"time_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NodeHealth) Reset()         { *m = NodeHealth{} }
func (m *NodeHealth) String() string { return proto.CompactTextString(m) }
func (*NodeHealth) ProtoMessage()    {}
func (*NodeHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9255cc60c5ca429, []int{1}
}
func (m *NodeHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeHealth.Unmarshal(m, b)
}
func (m *NodeHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeHealth.Marshal(b, m, deterministic)
}
func (m *NodeHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeHealth.Merge(m, src)
}
func (m *NodeHealth) XXX_Size() int {
	return xxx_messageInfo_NodeHealth.Size(m)
}
func (m *NodeHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeHealth.DiscardUnknown(m)
}

var xxx_messageInfo_NodeHealth proto.InternalMessageInfo

func (m *NodeHealth) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeHealth) GetBtfsVersion() string {
	if m != nil {
		return m.BtfsVersion
	}
	return ""
}

func (m *NodeHealth) GetFailurePoint() string {
	if m != nil {
		return m.FailurePoint
	}
	return ""
}

func (m *NodeHealth) GetTimeCreated() *time.Time {
	if m != nil {
		return m.TimeCreated
	}
	return nil
}

type NodeError struct {
	HVal                 string     `protobuf:"bytes,1,opt,name=h_val,json=hVal,proto3" json:"h_val,omitempty"`
	PeerId               string     `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	ErrorStatus          string     `protobuf:"bytes,3,opt,name=error_status,json=errorStatus,proto3" json:"error_status,omitempty"`
	TimeCreated          *time.Time `protobuf:"bytes,4,opt,name=time_created,json=timeCreated,proto3,stdtime" json:"time_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NodeError) Reset()         { *m = NodeError{} }
func (m *NodeError) String() string { return proto.CompactTextString(m) }
func (*NodeError) ProtoMessage()    {}
func (*NodeError) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9255cc60c5ca429, []int{2}
}
func (m *NodeError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeError.Unmarshal(m, b)
}
func (m *NodeError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeError.Marshal(b, m, deterministic)
}
func (m *NodeError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeError.Merge(m, src)
}
func (m *NodeError) XXX_Size() int {
	return xxx_messageInfo_NodeError.Size(m)
}
func (m *NodeError) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeError.DiscardUnknown(m)
}

var xxx_messageInfo_NodeError proto.InternalMessageInfo

func (m *NodeError) GetHVal() string {
	if m != nil {
		return m.HVal
	}
	return ""
}

func (m *NodeError) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *NodeError) GetErrorStatus() string {
	if m != nil {
		return m.ErrorStatus
	}
	return ""
}

func (m *NodeError) GetTimeCreated() *time.Time {
	if m != nil {
		return m.TimeCreated
	}
	return nil
}

type Ip2Location struct {
	Ipv4                 string   `protobuf:"bytes,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	CountryShort         string   `protobuf:"bytes,2,opt,name=country_short,json=countryShort,proto3" json:"country_short,omitempty"`
	CountryLong          string   `protobuf:"bytes,3,opt,name=country_long,json=countryLong,proto3" json:"country_long,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	City                 string   `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Latitude             float32  `protobuf:"fixed32,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Zipcode              string   `protobuf:"bytes,8,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ip2Location) Reset()         { *m = Ip2Location{} }
func (m *Ip2Location) String() string { return proto.CompactTextString(m) }
func (*Ip2Location) ProtoMessage()    {}
func (*Ip2Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9255cc60c5ca429, []int{3}
}
func (m *Ip2Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ip2Location.Unmarshal(m, b)
}
func (m *Ip2Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ip2Location.Marshal(b, m, deterministic)
}
func (m *Ip2Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ip2Location.Merge(m, src)
}
func (m *Ip2Location) XXX_Size() int {
	return xxx_messageInfo_Ip2Location.Size(m)
}
func (m *Ip2Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Ip2Location.DiscardUnknown(m)
}

var xxx_messageInfo_Ip2Location proto.InternalMessageInfo

func (m *Ip2Location) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Ip2Location) GetCountryShort() string {
	if m != nil {
		return m.CountryShort
	}
	return ""
}

func (m *Ip2Location) GetCountryLong() string {
	if m != nil {
		return m.CountryLong
	}
	return ""
}

func (m *Ip2Location) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Ip2Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Ip2Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Ip2Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Ip2Location) GetZipcode() string {
	if m != nil {
		return m.Zipcode
	}
	return ""
}

func init() {
	proto.RegisterType((*SignedMetrics)(nil), "status.SignedMetrics")
	proto.RegisterType((*NodeHealth)(nil), "status.NodeHealth")
	proto.RegisterType((*NodeError)(nil), "status.NodeError")
	proto.RegisterType((*Ip2Location)(nil), "status.Ip2Location")
}

func init() { proto.RegisterFile("protos/status/status.proto", fileDescriptor_e9255cc60c5ca429) }

var fileDescriptor_e9255cc60c5ca429 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x55, 0xfa, 0x4d, 0xd3, 0xce, 0x9d, 0xcc, 0x27, 0x61, 0x44, 0x09, 0x01, 0xd4, 0x52, 0x16,
	0x74, 0x43, 0x2a, 0x15, 0x36, 0x08, 0x89, 0x45, 0x47, 0x95, 0xa8, 0x28, 0xa8, 0x4a, 0xa1, 0x0b,
	0x36, 0x91, 0x27, 0xb9, 0x93, 0x58, 0xcd, 0xc4, 0xc1, 0x71, 0x2a, 0x85, 0x17, 0x60, 0xcb, 0x2b,
	0xf0, 0x0a, 0xbc, 0x10, 0x6b, 0xde, 0x02, 0x5d, 0xdb, 0xe9, 0xf0, 0xd7, 0x1d, 0xab, 0xf8, 0x9e,
	0xeb, 0x33, 0xe7, 0xf8, 0xf8, 0x7a, 0x20, 0x6a, 0x94, 0xd4, 0xb2, 0xdd, 0x6f, 0x35, 0xd7, 0xdd,
	0xf0, 0x89, 0x0d, 0xc8, 0x7c, 0x5b, 0x45, 0x8f, 0x0b, 0xa1, 0xcb, 0x6e, 0x1e, 0x67, 0x72, 0xb9,
	0x5f, 0xc8, 0x42, 0xee, 0x9b, 0xf6, 0xbc, 0x5b, 0x98, 0xca, 0x14, 0x66, 0x65, 0x69, 0xd1, 0xdd,
	0x42, 0xca, 0xa2, 0xc2, 0xd5, 0x2e, 0x5c, 0x36, 0xba, 0x77, 0xcd, 0xed, 0xdf, 0x9b, 0x5a, 0x2c,
	0xb1, 0xd5, 0x7c, 0xd9, 0xb8, 0x0d, 0x57, 0x86, 0x4a, 0xae, 0x30, 0x77, 0x1f, 0xdb, 0xdb, 0x5d,
	0xc0, 0xf4, 0x4c, 0x14, 0x35, 0xe6, 0xaf, 0x51, 0x2b, 0x91, 0xb5, 0xec, 0x3e, 0x40, 0xd3, 0xcd,
	0x2b, 0x91, 0xa5, 0x17, 0xd8, 0x87, 0xde, 0x8e, 0xb7, 0x17, 0x24, 0x63, 0x8b, 0xbc, 0xc2, 0x9e,
	0xdd, 0x83, 0x71, 0x2b, 0x8a, 0x9a, 0xeb, 0x4e, 0x61, 0xb8, 0x66, 0xbb, 0x57, 0x00, 0x0b, 0x61,
	0xa3, 0xe1, 0x7d, 0x25, 0x79, 0x1e, 0xfe, 0x67, 0x7a, 0x43, 0xb9, 0xfb, 0xd5, 0x03, 0x78, 0x23,
	0x73, 0x7c, 0x89, 0xbc, 0xd2, 0x25, 0xbb, 0x0d, 0x1b, 0xb5, 0xcc, 0x31, 0x15, 0xb9, 0x91, 0x18,
	0x27, 0x3e, 0x95, 0xc7, 0x39, 0x7b, 0x00, 0xc1, 0x5c, 0x2f, 0xda, 0xf4, 0x12, 0x55, 0x2b, 0x64,
	0x6d, 0x24, 0xc6, 0xc9, 0x84, 0xb0, 0x73, 0x0b, 0xb1, 0x87, 0x30, 0x5d, 0x70, 0x51, 0x75, 0x0a,
	0xd3, 0x46, 0x8a, 0x5a, 0x1b, 0xa9, 0x71, 0x12, 0x38, 0xf0, 0x94, 0x30, 0x36, 0x83, 0x80, 0x62,
	0x48, 0x33, 0x85, 0x5c, 0x63, 0x1e, 0x8e, 0x76, 0xbc, 0xbd, 0xc9, 0x41, 0x14, 0xdb, 0xac, 0xe2,
	0x21, 0xab, 0xf8, 0xed, 0x90, 0xd5, 0xe1, 0xe8, 0xf3, 0xb7, 0x6d, 0x2f, 0x99, 0x10, 0x6b, 0x66,
	0x49, 0xbb, 0x5f, 0x3c, 0x18, 0x93, 0xe9, 0x23, 0xa5, 0xa4, 0x62, 0x37, 0x61, 0xbd, 0x4c, 0x2f,
	0x79, 0xe5, 0x1c, 0x8f, 0xca, 0x73, 0x5e, 0xd1, 0x41, 0x1a, 0x44, 0x45, 0x07, 0xb1, 0x56, 0x7d,
	0x2a, 0xed, 0x41, 0x90, 0x68, 0xa9, 0xbd, 0x71, 0x67, 0x72, 0x62, 0xb0, 0x33, 0x03, 0xfd, 0x1b,
	0x8f, 0xdf, 0x3d, 0x98, 0x1c, 0x37, 0x07, 0x27, 0x32, 0xe3, 0x9a, 0xd2, 0x61, 0x30, 0x12, 0xcd,
	0xe5, 0xd3, 0xc1, 0x24, 0xad, 0x29, 0xb1, 0x4c, 0x76, 0xb5, 0x56, 0x7d, 0xda, 0x96, 0x52, 0x69,
	0x67, 0x35, 0x70, 0xe0, 0x19, 0x61, 0x64, 0x78, 0xd8, 0x54, 0xc9, 0xba, 0x18, 0x0c, 0x3b, 0xec,
	0x44, 0xd6, 0x05, 0xdb, 0x02, 0x5f, 0x61, 0x41, 0xd7, 0x32, 0xb2, 0x67, 0xb5, 0x15, 0x69, 0x66,
	0x42, 0xf7, 0xe1, 0xba, 0xd5, 0xa4, 0x35, 0x8b, 0x60, 0xb3, 0xe2, 0x5a, 0xe8, 0x2e, 0xc7, 0xd0,
	0xdf, 0xf1, 0xf6, 0xd6, 0x92, 0xab, 0x9a, 0x86, 0x88, 0x24, 0x6c, 0x73, 0xc3, 0x34, 0x57, 0x00,
	0x0d, 0xd1, 0x47, 0xd1, 0x64, 0x32, 0xc7, 0x70, 0xd3, 0xfc, 0xe0, 0x50, 0x1e, 0x7c, 0x5a, 0x03,
	0xdf, 0x65, 0xf7, 0x02, 0xa6, 0xef, 0x9a, 0x9c, 0x6b, 0x1c, 0xe6, 0xf6, 0x56, 0xec, 0x1e, 0xda,
	0x2f, 0xe3, 0x1c, 0x6d, 0xfd, 0x91, 0xe6, 0x11, 0x3d, 0x1d, 0xf6, 0x1c, 0xa6, 0x33, 0x59, 0x55,
	0x98, 0x69, 0x37, 0x91, 0x6c, 0xe0, 0xaf, 0xa6, 0xf4, 0x5a, 0xf2, 0x33, 0x08, 0x1c, 0xd9, 0x4e,
	0xc6, 0x8d, 0x9f, 0xb9, 0x06, 0xba, 0x96, 0x7a, 0x04, 0xc1, 0xac, 0xc4, 0xec, 0x22, 0xe9, 0x6a,
	0xba, 0x45, 0x16, 0xc5, 0xee, 0x39, 0x3a, 0xe0, 0xb8, 0x5e, 0xc8, 0x04, 0x3f, 0x74, 0xd8, 0xea,
	0xe8, 0xce, 0x5f, 0x7b, 0x8d, 0x54, 0xfa, 0xf0, 0x11, 0xfc, 0x2f, 0x64, 0x4c, 0xaf, 0xc2, 0x49,
	0x1f, 0x4e, 0x6c, 0x30, 0xa7, 0x24, 0x77, 0xea, 0xbd, 0x77, 0x7f, 0x34, 0x73, 0xdf, 0xe8, 0x3f,
	0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0xef, 0x7b, 0x05, 0x22, 0x95, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatusClient is the client API for Status service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusClient interface {
	UpdateMetrics(ctx context.Context, in *SignedMetrics, opts ...grpc.CallOption) (*types.Empty, error)
	CollectHealth(ctx context.Context, in *NodeHealth, opts ...grpc.CallOption) (*types.Empty, error)
	CollectError(ctx context.Context, in *NodeError, opts ...grpc.CallOption) (*types.Empty, error)
	CheckRuntime(ctx context.Context, in *shared.RuntimeInfoRequest, opts ...grpc.CallOption) (*shared.RuntimeInfoReport, error)
}

type statusClient struct {
	cc *grpc.ClientConn
}

func NewStatusClient(cc *grpc.ClientConn) StatusClient {
	return &statusClient{cc}
}

func (c *statusClient) UpdateMetrics(ctx context.Context, in *SignedMetrics, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/status.Status/UpdateMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusClient) CollectHealth(ctx context.Context, in *NodeHealth, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/status.Status/CollectHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusClient) CollectError(ctx context.Context, in *NodeError, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/status.Status/CollectError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusClient) CheckRuntime(ctx context.Context, in *shared.RuntimeInfoRequest, opts ...grpc.CallOption) (*shared.RuntimeInfoReport, error) {
	out := new(shared.RuntimeInfoReport)
	err := c.cc.Invoke(ctx, "/status.Status/CheckRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServer is the server API for Status service.
type StatusServer interface {
	UpdateMetrics(context.Context, *SignedMetrics) (*types.Empty, error)
	CollectHealth(context.Context, *NodeHealth) (*types.Empty, error)
	CollectError(context.Context, *NodeError) (*types.Empty, error)
	CheckRuntime(context.Context, *shared.RuntimeInfoRequest) (*shared.RuntimeInfoReport, error)
}

// UnimplementedStatusServer can be embedded to have forward compatible implementations.
type UnimplementedStatusServer struct {
}

func (*UnimplementedStatusServer) UpdateMetrics(ctx context.Context, req *SignedMetrics) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetrics not implemented")
}
func (*UnimplementedStatusServer) CollectHealth(ctx context.Context, req *NodeHealth) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectHealth not implemented")
}
func (*UnimplementedStatusServer) CollectError(ctx context.Context, req *NodeError) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectError not implemented")
}
func (*UnimplementedStatusServer) CheckRuntime(ctx context.Context, req *shared.RuntimeInfoRequest) (*shared.RuntimeInfoReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRuntime not implemented")
}

func RegisterStatusServer(s *grpc.Server, srv StatusServer) {
	s.RegisterService(&_Status_serviceDesc, srv)
}

func _Status_UpdateMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).UpdateMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.Status/UpdateMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).UpdateMetrics(ctx, req.(*SignedMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _Status_CollectHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeHealth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).CollectHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.Status/CollectHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).CollectHealth(ctx, req.(*NodeHealth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Status_CollectError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeError)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).CollectError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.Status/CollectError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).CollectError(ctx, req.(*NodeError))
	}
	return interceptor(ctx, in, info, handler)
}

func _Status_CheckRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.RuntimeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).CheckRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.Status/CheckRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).CheckRuntime(ctx, req.(*shared.RuntimeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Status_serviceDesc = grpc.ServiceDesc{
	ServiceName: "status.Status",
	HandlerType: (*StatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMetrics",
			Handler:    _Status_UpdateMetrics_Handler,
		},
		{
			MethodName: "CollectHealth",
			Handler:    _Status_CollectHealth_Handler,
		},
		{
			MethodName: "CollectError",
			Handler:    _Status_CollectError_Handler,
		},
		{
			MethodName: "CheckRuntime",
			Handler:    _Status_CheckRuntime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/status/status.proto",
}