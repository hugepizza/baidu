// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: body.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 人体检测参数
type TrackOption struct {
	CaseId   int64  `protobuf:"varint,1,opt,name=caseId,proto3" json:"case_id"`
	CaseInit bool   `protobuf:"varint,2,opt,name=caseInit,proto3" json:"case_init"`
	Image    string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Dynamic  bool   `protobuf:"varint,4,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
}

func (m *TrackOption) Reset()         { *m = TrackOption{} }
func (m *TrackOption) String() string { return proto.CompactTextString(m) }
func (*TrackOption) ProtoMessage()    {}
func (*TrackOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_body_6f9116684fe79234, []int{0}
}
func (m *TrackOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackOption.Unmarshal(m, b)
}
func (m *TrackOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackOption.Marshal(b, m, deterministic)
}
func (dst *TrackOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackOption.Merge(dst, src)
}
func (m *TrackOption) XXX_Size() int {
	return xxx_messageInfo_TrackOption.Size(m)
}
func (m *TrackOption) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackOption.DiscardUnknown(m)
}

var xxx_messageInfo_TrackOption proto.InternalMessageInfo

// 人体位置
type Location struct {
	Left   int32 `protobuf:"varint,1,opt,name=left,proto3" json:"left,omitempty"`
	Top    int32 `protobuf:"varint,2,opt,name=top,proto3" json:"top,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_body_6f9116684fe79234, []int{1}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

// 人体信息
type Info struct {
	Location *Location `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Id       int32     `protobuf:"varint,2,opt,name=id,proto3" json:"ID"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_body_6f9116684fe79234, []int{2}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

// 人体数量
type Count struct {
	In  int32 `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
	Out int32 `protobuf:"varint,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_body_6f9116684fe79234, []int{3}
}
func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (dst *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(dst, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

// 人体检测结果
type TrackResult struct {
	PersonNum   int32   `protobuf:"varint,1,opt,name=personNum,proto3" json:"person_num"`
	PersonInfo  []*Info `protobuf:"bytes,2,rep,name=personInfo" json:"person_info"`
	PersonCount *Count  `protobuf:"bytes,3,opt,name=personCount" json:"person_count"`
	Image       string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *TrackResult) Reset()         { *m = TrackResult{} }
func (m *TrackResult) String() string { return proto.CompactTextString(m) }
func (*TrackResult) ProtoMessage()    {}
func (*TrackResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_body_6f9116684fe79234, []int{4}
}
func (m *TrackResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackResult.Unmarshal(m, b)
}
func (m *TrackResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackResult.Marshal(b, m, deterministic)
}
func (dst *TrackResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackResult.Merge(dst, src)
}
func (m *TrackResult) XXX_Size() int {
	return xxx_messageInfo_TrackResult.Size(m)
}
func (m *TrackResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackResult.DiscardUnknown(m)
}

var xxx_messageInfo_TrackResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TrackOption)(nil), "funxdata.baidu.TrackOption")
	proto.RegisterType((*Location)(nil), "funxdata.baidu.Location")
	proto.RegisterType((*Info)(nil), "funxdata.baidu.Info")
	proto.RegisterType((*Count)(nil), "funxdata.baidu.Count")
	proto.RegisterType((*TrackResult)(nil), "funxdata.baidu.TrackResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BodyTrackClient is the client API for BodyTrack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BodyTrackClient interface {
	Track(ctx context.Context, in *TrackOption, opts ...grpc.CallOption) (*TrackResult, error)
}

type bodyTrackClient struct {
	cc *grpc.ClientConn
}

func NewBodyTrackClient(cc *grpc.ClientConn) BodyTrackClient {
	return &bodyTrackClient{cc}
}

func (c *bodyTrackClient) Track(ctx context.Context, in *TrackOption, opts ...grpc.CallOption) (*TrackResult, error) {
	out := new(TrackResult)
	err := c.cc.Invoke(ctx, "/funxdata.baidu.BodyTrack/Track", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BodyTrackServer is the server API for BodyTrack service.
type BodyTrackServer interface {
	Track(context.Context, *TrackOption) (*TrackResult, error)
}

func RegisterBodyTrackServer(s *grpc.Server, srv BodyTrackServer) {
	s.RegisterService(&_BodyTrack_serviceDesc, srv)
}

func _BodyTrack_Track_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyTrackServer).Track(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funxdata.baidu.BodyTrack/Track",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyTrackServer).Track(ctx, req.(*TrackOption))
	}
	return interceptor(ctx, in, info, handler)
}

var _BodyTrack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "funxdata.baidu.BodyTrack",
	HandlerType: (*BodyTrackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Track",
			Handler:    _BodyTrack_Track_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "body.proto",
}

func init() { proto.RegisterFile("body.proto", fileDescriptor_body_6f9116684fe79234) }

var fileDescriptor_body_6f9116684fe79234 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x9d, 0x38, 0x4d, 0xc6, 0x10, 0xaa, 0x55, 0xa9, 0xac, 0x22, 0xd9, 0x91, 0xe1, 0x90,
	0x4a, 0xe0, 0x4a, 0x81, 0x2f, 0x70, 0xe1, 0x10, 0x09, 0x01, 0x5a, 0xf5, 0xc4, 0x81, 0xca, 0xf6,
	0xda, 0xce, 0x8a, 0x78, 0xd7, 0x4a, 0x76, 0x05, 0xf9, 0x08, 0x24, 0x3e, 0x2b, 0xc7, 0x1e, 0x39,
	0x59, 0x10, 0xdf, 0xfc, 0x15, 0xc8, 0x6b, 0xa7, 0x0d, 0x11, 0xb7, 0x79, 0xfb, 0xe6, 0xcd, 0x7b,
	0x3b, 0xbb, 0x00, 0x11, 0x27, 0x1b, 0xbf, 0x58, 0x71, 0xc1, 0xd1, 0x38, 0x95, 0xec, 0x3b, 0x09,
	0x45, 0xe8, 0x47, 0x21, 0x25, 0xf2, 0xe2, 0x55, 0x46, 0xc5, 0x42, 0x46, 0x7e, 0xcc, 0xf3, 0xab,
	0x8c, 0x67, 0xfc, 0x4a, 0xb5, 0x45, 0x32, 0x55, 0x48, 0x01, 0x55, 0xb5, 0x72, 0xef, 0x87, 0x0e,
	0xd6, 0xcd, 0x2a, 0x8c, 0xbf, 0x7e, 0x2c, 0x04, 0xe5, 0x0c, 0x3d, 0x87, 0x41, 0x1c, 0xae, 0x93,
	0x39, 0xb1, 0xf5, 0x89, 0x3e, 0xed, 0x05, 0x56, 0x5d, 0xba, 0x27, 0xcd, 0xc9, 0x2d, 0x25, 0xb8,
	0xa3, 0xd0, 0x25, 0x0c, 0x55, 0xc5, 0xa8, 0xb0, 0x8d, 0x89, 0x3e, 0x1d, 0x06, 0x8f, 0xeb, 0xd2,
	0x1d, 0xb5, 0x6d, 0x8c, 0x0a, 0x7c, 0x4f, 0xa3, 0x33, 0x30, 0x69, 0x1e, 0x66, 0x89, 0xdd, 0x9b,
	0xe8, 0xd3, 0x11, 0x6e, 0x01, 0xb2, 0xe1, 0x84, 0x6c, 0x58, 0x98, 0xd3, 0xd8, 0xee, 0x37, 0x7a,
	0xbc, 0x87, 0xde, 0x17, 0x18, 0xbe, 0xe7, 0x71, 0xa8, 0xb2, 0x20, 0xe8, 0x2f, 0x93, 0x54, 0xa8,
	0x24, 0x26, 0x56, 0x35, 0x3a, 0x85, 0x9e, 0xe0, 0x85, 0x72, 0x35, 0x71, 0x53, 0x36, 0x0e, 0xdf,
	0x28, 0x11, 0x0b, 0xe5, 0x60, 0xe2, 0x16, 0xa0, 0x73, 0x18, 0x2c, 0x12, 0x9a, 0x2d, 0x84, 0x32,
	0x30, 0x71, 0x87, 0xbc, 0x1b, 0xe8, 0xcf, 0x59, 0xca, 0xd1, 0x1b, 0x18, 0x2e, 0x3b, 0x1f, 0x35,
	0xdf, 0x9a, 0xd9, 0xfe, 0xbf, 0x9b, 0xf4, 0xf7, 0x39, 0xf0, 0x7d, 0x27, 0x3a, 0x07, 0x83, 0x92,
	0xd6, 0x3c, 0x18, 0xd4, 0xa5, 0x6b, 0xcc, 0xdf, 0x62, 0x83, 0x12, 0xef, 0x12, 0xcc, 0x6b, 0x2e,
	0x99, 0x40, 0x63, 0x30, 0x28, 0xeb, 0x02, 0x1b, 0x94, 0x35, 0x71, 0xb9, 0x14, 0xfb, 0xb8, 0x5c,
	0x0a, 0xaf, 0xda, 0x2f, 0x1c, 0x27, 0x6b, 0xb9, 0x14, 0xe8, 0x25, 0x8c, 0x8a, 0x64, 0xb5, 0xe6,
	0xec, 0x83, 0xcc, 0x5b, 0x61, 0x30, 0xae, 0x4b, 0x17, 0xda, 0xc3, 0x5b, 0x26, 0x73, 0xfc, 0xd0,
	0x80, 0xde, 0x41, 0x47, 0x34, 0x97, 0xb0, 0x8d, 0x49, 0x6f, 0x6a, 0xcd, 0xce, 0x8e, 0x83, 0x37,
	0x5c, 0xf0, 0xa4, 0x2e, 0x5d, 0xab, 0x1b, 0x42, 0x59, 0xca, 0xf1, 0x81, 0x10, 0xcd, 0xa1, 0xa3,
	0x54, 0x6a, 0xb5, 0x39, 0x6b, 0xf6, 0xf4, 0x78, 0x8e, 0x22, 0x83, 0xd3, 0xba, 0x74, 0x1f, 0x75,
	0x83, 0xe2, 0xe6, 0x04, 0x1f, 0x6a, 0x1f, 0x1e, 0xb8, 0x7f, 0xf0, 0xc0, 0xb3, 0x4f, 0x30, 0x0a,
	0x38, 0xd9, 0xa8, 0x8b, 0xa2, 0x6b, 0x30, 0xdb, 0xe2, 0xd9, 0xb1, 0xc3, 0xc1, 0xcf, 0xbb, 0xf8,
	0x3f, 0xd9, 0x6e, 0xc9, 0xd3, 0x82, 0x17, 0xdb, 0x3f, 0x8e, 0xb6, 0xdd, 0x39, 0xda, 0xdd, 0xce,
	0xd1, 0x7e, 0xef, 0x1c, 0xed, 0x67, 0xe5, 0x68, 0xdb, 0xca, 0xd1, 0xee, 0x2a, 0x47, 0xfb, 0x55,
	0x39, 0xda, 0x67, 0xa3, 0x88, 0xa2, 0x81, 0xfa, 0xd5, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xf4, 0x2e, 0xb0, 0x22, 0x03, 0x00, 0x00,
}