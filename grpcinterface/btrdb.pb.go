// Code generated by protoc-gen-go.
// source: btrdb.proto
// DO NOT EDIT!

/*
Package grpcinterface is a generated protocol buffer package.

Version 4.0

It is generated from these files:
	btrdb.proto

It has these top-level messages:
	RawValuesParams
	RawValuesResponse
	AlignedWindowsParams
	AlignedWindowsResponse
	WindowsParams
	WindowsResponse
	VersionParams
	VersionResponse
	NearestParams
	NearestResponse
	ChangesParams
	ChangesResponse
	InsertParams
	InsertResponse
	DeleteParams
	DeleteResponse
	InfoParams
	InfoResponse
	RawPoint
	StatPoint
	ChangedRange
	Status
	Mash
*/
package grpcinterface

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RawValuesParams struct {
	Uuid         []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Start        int64  `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End          int64  `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
	VersionMajor uint64 `protobuf:"varint,4,opt,name=versionMajor" json:"versionMajor,omitempty"`
}

func (m *RawValuesParams) Reset()                    { *m = RawValuesParams{} }
func (m *RawValuesParams) String() string            { return proto.CompactTextString(m) }
func (*RawValuesParams) ProtoMessage()               {}
func (*RawValuesParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RawValuesResponse struct {
	Stat         *Status     `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64      `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64      `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
	Values       []*RawPoint `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *RawValuesResponse) Reset()                    { *m = RawValuesResponse{} }
func (m *RawValuesResponse) String() string            { return proto.CompactTextString(m) }
func (*RawValuesResponse) ProtoMessage()               {}
func (*RawValuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RawValuesResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *RawValuesResponse) GetValues() []*RawPoint {
	if m != nil {
		return m.Values
	}
	return nil
}

type AlignedWindowsParams struct {
	Uuid         []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Start        int64  `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End          int64  `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
	VersionMajor uint64 `protobuf:"varint,4,opt,name=versionMajor" json:"versionMajor,omitempty"`
	PointWidth   uint32 `protobuf:"varint,5,opt,name=pointWidth" json:"pointWidth,omitempty"`
}

func (m *AlignedWindowsParams) Reset()                    { *m = AlignedWindowsParams{} }
func (m *AlignedWindowsParams) String() string            { return proto.CompactTextString(m) }
func (*AlignedWindowsParams) ProtoMessage()               {}
func (*AlignedWindowsParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AlignedWindowsResponse struct {
	Stat         *Status      `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64       `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64       `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
	Values       []*StatPoint `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *AlignedWindowsResponse) Reset()                    { *m = AlignedWindowsResponse{} }
func (m *AlignedWindowsResponse) String() string            { return proto.CompactTextString(m) }
func (*AlignedWindowsResponse) ProtoMessage()               {}
func (*AlignedWindowsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AlignedWindowsResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *AlignedWindowsResponse) GetValues() []*StatPoint {
	if m != nil {
		return m.Values
	}
	return nil
}

type WindowsParams struct {
	Uuid         []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Start        int64  `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End          int64  `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
	VersionMajor uint64 `protobuf:"varint,4,opt,name=versionMajor" json:"versionMajor,omitempty"`
	Width        uint64 `protobuf:"varint,5,opt,name=width" json:"width,omitempty"`
	Depth        uint32 `protobuf:"varint,6,opt,name=depth" json:"depth,omitempty"`
}

func (m *WindowsParams) Reset()                    { *m = WindowsParams{} }
func (m *WindowsParams) String() string            { return proto.CompactTextString(m) }
func (*WindowsParams) ProtoMessage()               {}
func (*WindowsParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type WindowsResponse struct {
	Stat         *Status      `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64       `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64       `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
	Values       []*StatPoint `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *WindowsResponse) Reset()                    { *m = WindowsResponse{} }
func (m *WindowsResponse) String() string            { return proto.CompactTextString(m) }
func (*WindowsResponse) ProtoMessage()               {}
func (*WindowsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WindowsResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *WindowsResponse) GetValues() []*StatPoint {
	if m != nil {
		return m.Values
	}
	return nil
}

type VersionParams struct {
	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *VersionParams) Reset()                    { *m = VersionParams{} }
func (m *VersionParams) String() string            { return proto.CompactTextString(m) }
func (*VersionParams) ProtoMessage()               {}
func (*VersionParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type VersionResponse struct {
	Stat         *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64  `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64  `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *VersionResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type NearestParams struct {
	Uuid         []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Time         int64  `protobuf:"fixed64,2,opt,name=time" json:"time,omitempty"`
	VersionMajor uint64 `protobuf:"varint,3,opt,name=versionMajor" json:"versionMajor,omitempty"`
	Backward     bool   `protobuf:"varint,4,opt,name=backward" json:"backward,omitempty"`
}

func (m *NearestParams) Reset()                    { *m = NearestParams{} }
func (m *NearestParams) String() string            { return proto.CompactTextString(m) }
func (*NearestParams) ProtoMessage()               {}
func (*NearestParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type NearestResponse struct {
	Stat         *Status   `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64    `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64    `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
	Value        *RawPoint `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *NearestResponse) Reset()                    { *m = NearestResponse{} }
func (m *NearestResponse) String() string            { return proto.CompactTextString(m) }
func (*NearestResponse) ProtoMessage()               {}
func (*NearestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *NearestResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *NearestResponse) GetValue() *RawPoint {
	if m != nil {
		return m.Value
	}
	return nil
}

type ChangesParams struct {
	Uuid       []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FromMajor  uint64 `protobuf:"varint,2,opt,name=fromMajor" json:"fromMajor,omitempty"`
	ToMajor    uint64 `protobuf:"varint,3,opt,name=toMajor" json:"toMajor,omitempty"`
	Resolution uint32 `protobuf:"varint,4,opt,name=resolution" json:"resolution,omitempty"`
}

func (m *ChangesParams) Reset()                    { *m = ChangesParams{} }
func (m *ChangesParams) String() string            { return proto.CompactTextString(m) }
func (*ChangesParams) ProtoMessage()               {}
func (*ChangesParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ChangesResponse struct {
	Stat         *Status         `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64          `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64          `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
	Ranges       []*ChangedRange `protobuf:"bytes,4,rep,name=ranges" json:"ranges,omitempty"`
}

func (m *ChangesResponse) Reset()                    { *m = ChangesResponse{} }
func (m *ChangesResponse) String() string            { return proto.CompactTextString(m) }
func (*ChangesResponse) ProtoMessage()               {}
func (*ChangesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ChangesResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ChangesResponse) GetRanges() []*ChangedRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type InsertParams struct {
	Uuid   []byte      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Sync   bool        `protobuf:"varint,2,opt,name=sync" json:"sync,omitempty"`
	Values []*RawPoint `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *InsertParams) Reset()                    { *m = InsertParams{} }
func (m *InsertParams) String() string            { return proto.CompactTextString(m) }
func (*InsertParams) ProtoMessage()               {}
func (*InsertParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *InsertParams) GetValues() []*RawPoint {
	if m != nil {
		return m.Values
	}
	return nil
}

type InsertResponse struct {
	Stat         *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64  `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64  `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
}

func (m *InsertResponse) Reset()                    { *m = InsertResponse{} }
func (m *InsertResponse) String() string            { return proto.CompactTextString(m) }
func (*InsertResponse) ProtoMessage()               {}
func (*InsertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *InsertResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type DeleteParams struct {
	Uuid  []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Start int64  `protobuf:"fixed64,2,opt,name=start" json:"start,omitempty"`
	End   int64  `protobuf:"fixed64,3,opt,name=end" json:"end,omitempty"`
}

func (m *DeleteParams) Reset()                    { *m = DeleteParams{} }
func (m *DeleteParams) String() string            { return proto.CompactTextString(m) }
func (*DeleteParams) ProtoMessage()               {}
func (*DeleteParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type DeleteResponse struct {
	Stat         *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	VersionMajor uint64  `protobuf:"varint,2,opt,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor uint64  `protobuf:"varint,3,opt,name=versionMinor" json:"versionMinor,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type InfoParams struct {
}

func (m *InfoParams) Reset()                    { *m = InfoParams{} }
func (m *InfoParams) String() string            { return proto.CompactTextString(m) }
func (*InfoParams) ProtoMessage()               {}
func (*InfoParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type InfoResponse struct {
	Stat *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Mash *Mash   `protobuf:"bytes,2,opt,name=mash" json:"mash,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *InfoResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *InfoResponse) GetMash() *Mash {
	if m != nil {
		return m.Mash
	}
	return nil
}

type RawPoint struct {
	Time  int64   `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *RawPoint) Reset()                    { *m = RawPoint{} }
func (m *RawPoint) String() string            { return proto.CompactTextString(m) }
func (*RawPoint) ProtoMessage()               {}
func (*RawPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type StatPoint struct {
	Time  int64   `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Min   float64 `protobuf:"fixed64,2,opt,name=min" json:"min,omitempty"`
	Mean  float64 `protobuf:"fixed64,3,opt,name=mean" json:"mean,omitempty"`
	Max   float64 `protobuf:"fixed64,4,opt,name=max" json:"max,omitempty"`
	Count uint64  `protobuf:"fixed64,5,opt,name=count" json:"count,omitempty"`
}

func (m *StatPoint) Reset()                    { *m = StatPoint{} }
func (m *StatPoint) String() string            { return proto.CompactTextString(m) }
func (*StatPoint) ProtoMessage()               {}
func (*StatPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

type ChangedRange struct {
	Start int64 `protobuf:"fixed64,1,opt,name=start" json:"start,omitempty"`
	End   int64 `protobuf:"fixed64,2,opt,name=end" json:"end,omitempty"`
}

func (m *ChangedRange) Reset()                    { *m = ChangedRange{} }
func (m *ChangedRange) String() string            { return proto.CompactTextString(m) }
func (*ChangedRange) ProtoMessage()               {}
func (*ChangedRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type Status struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Mash *Mash  `protobuf:"bytes,3,opt,name=mash" json:"mash,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *Status) GetMash() *Mash {
	if m != nil {
		return m.Mash
	}
	return nil
}

type Mash struct {
}

func (m *Mash) Reset()                    { *m = Mash{} }
func (m *Mash) String() string            { return proto.CompactTextString(m) }
func (*Mash) ProtoMessage()               {}
func (*Mash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func init() {
	proto.RegisterType((*RawValuesParams)(nil), "grpcinterface.RawValuesParams")
	proto.RegisterType((*RawValuesResponse)(nil), "grpcinterface.RawValuesResponse")
	proto.RegisterType((*AlignedWindowsParams)(nil), "grpcinterface.AlignedWindowsParams")
	proto.RegisterType((*AlignedWindowsResponse)(nil), "grpcinterface.AlignedWindowsResponse")
	proto.RegisterType((*WindowsParams)(nil), "grpcinterface.WindowsParams")
	proto.RegisterType((*WindowsResponse)(nil), "grpcinterface.WindowsResponse")
	proto.RegisterType((*VersionParams)(nil), "grpcinterface.VersionParams")
	proto.RegisterType((*VersionResponse)(nil), "grpcinterface.VersionResponse")
	proto.RegisterType((*NearestParams)(nil), "grpcinterface.NearestParams")
	proto.RegisterType((*NearestResponse)(nil), "grpcinterface.NearestResponse")
	proto.RegisterType((*ChangesParams)(nil), "grpcinterface.ChangesParams")
	proto.RegisterType((*ChangesResponse)(nil), "grpcinterface.ChangesResponse")
	proto.RegisterType((*InsertParams)(nil), "grpcinterface.InsertParams")
	proto.RegisterType((*InsertResponse)(nil), "grpcinterface.InsertResponse")
	proto.RegisterType((*DeleteParams)(nil), "grpcinterface.DeleteParams")
	proto.RegisterType((*DeleteResponse)(nil), "grpcinterface.DeleteResponse")
	proto.RegisterType((*InfoParams)(nil), "grpcinterface.InfoParams")
	proto.RegisterType((*InfoResponse)(nil), "grpcinterface.InfoResponse")
	proto.RegisterType((*RawPoint)(nil), "grpcinterface.RawPoint")
	proto.RegisterType((*StatPoint)(nil), "grpcinterface.StatPoint")
	proto.RegisterType((*ChangedRange)(nil), "grpcinterface.ChangedRange")
	proto.RegisterType((*Status)(nil), "grpcinterface.Status")
	proto.RegisterType((*Mash)(nil), "grpcinterface.Mash")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BTrDB service

type BTrDBClient interface {
	RawValues(ctx context.Context, in *RawValuesParams, opts ...grpc.CallOption) (BTrDB_RawValuesClient, error)
	AlignedWindows(ctx context.Context, in *AlignedWindowsParams, opts ...grpc.CallOption) (BTrDB_AlignedWindowsClient, error)
	Windows(ctx context.Context, in *WindowsParams, opts ...grpc.CallOption) (BTrDB_WindowsClient, error)
	Version(ctx context.Context, in *VersionParams, opts ...grpc.CallOption) (*VersionResponse, error)
	Nearest(ctx context.Context, in *NearestParams, opts ...grpc.CallOption) (*NearestResponse, error)
	Changes(ctx context.Context, in *ChangesParams, opts ...grpc.CallOption) (*ChangesResponse, error)
	Insert(ctx context.Context, in *InsertParams, opts ...grpc.CallOption) (*InsertResponse, error)
	Delete(ctx context.Context, in *DeleteParams, opts ...grpc.CallOption) (*DeleteResponse, error)
	Info(ctx context.Context, in *InfoParams, opts ...grpc.CallOption) (*InfoResponse, error)
}

type bTrDBClient struct {
	cc *grpc.ClientConn
}

func NewBTrDBClient(cc *grpc.ClientConn) BTrDBClient {
	return &bTrDBClient{cc}
}

func (c *bTrDBClient) RawValues(ctx context.Context, in *RawValuesParams, opts ...grpc.CallOption) (BTrDB_RawValuesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BTrDB_serviceDesc.Streams[0], c.cc, "/grpcinterface.BTrDB/RawValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &bTrDBRawValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BTrDB_RawValuesClient interface {
	Recv() (*RawValuesResponse, error)
	grpc.ClientStream
}

type bTrDBRawValuesClient struct {
	grpc.ClientStream
}

func (x *bTrDBRawValuesClient) Recv() (*RawValuesResponse, error) {
	m := new(RawValuesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bTrDBClient) AlignedWindows(ctx context.Context, in *AlignedWindowsParams, opts ...grpc.CallOption) (BTrDB_AlignedWindowsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BTrDB_serviceDesc.Streams[1], c.cc, "/grpcinterface.BTrDB/AlignedWindows", opts...)
	if err != nil {
		return nil, err
	}
	x := &bTrDBAlignedWindowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BTrDB_AlignedWindowsClient interface {
	Recv() (*AlignedWindowsResponse, error)
	grpc.ClientStream
}

type bTrDBAlignedWindowsClient struct {
	grpc.ClientStream
}

func (x *bTrDBAlignedWindowsClient) Recv() (*AlignedWindowsResponse, error) {
	m := new(AlignedWindowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bTrDBClient) Windows(ctx context.Context, in *WindowsParams, opts ...grpc.CallOption) (BTrDB_WindowsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BTrDB_serviceDesc.Streams[2], c.cc, "/grpcinterface.BTrDB/Windows", opts...)
	if err != nil {
		return nil, err
	}
	x := &bTrDBWindowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BTrDB_WindowsClient interface {
	Recv() (*WindowsResponse, error)
	grpc.ClientStream
}

type bTrDBWindowsClient struct {
	grpc.ClientStream
}

func (x *bTrDBWindowsClient) Recv() (*WindowsResponse, error) {
	m := new(WindowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bTrDBClient) Version(ctx context.Context, in *VersionParams, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBClient) Nearest(ctx context.Context, in *NearestParams, opts ...grpc.CallOption) (*NearestResponse, error) {
	out := new(NearestResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Nearest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBClient) Changes(ctx context.Context, in *ChangesParams, opts ...grpc.CallOption) (*ChangesResponse, error) {
	out := new(ChangesResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Changes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBClient) Insert(ctx context.Context, in *InsertParams, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBClient) Delete(ctx context.Context, in *DeleteParams, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBClient) Info(ctx context.Context, in *InfoParams, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/grpcinterface.BTrDB/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BTrDB service

type BTrDBServer interface {
	RawValues(*RawValuesParams, BTrDB_RawValuesServer) error
	AlignedWindows(*AlignedWindowsParams, BTrDB_AlignedWindowsServer) error
	Windows(*WindowsParams, BTrDB_WindowsServer) error
	Version(context.Context, *VersionParams) (*VersionResponse, error)
	Nearest(context.Context, *NearestParams) (*NearestResponse, error)
	Changes(context.Context, *ChangesParams) (*ChangesResponse, error)
	Insert(context.Context, *InsertParams) (*InsertResponse, error)
	Delete(context.Context, *DeleteParams) (*DeleteResponse, error)
	Info(context.Context, *InfoParams) (*InfoResponse, error)
}

func RegisterBTrDBServer(s *grpc.Server, srv BTrDBServer) {
	s.RegisterService(&_BTrDB_serviceDesc, srv)
}

func _BTrDB_RawValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RawValuesParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BTrDBServer).RawValues(m, &bTrDBRawValuesServer{stream})
}

type BTrDB_RawValuesServer interface {
	Send(*RawValuesResponse) error
	grpc.ServerStream
}

type bTrDBRawValuesServer struct {
	grpc.ServerStream
}

func (x *bTrDBRawValuesServer) Send(m *RawValuesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BTrDB_AlignedWindows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AlignedWindowsParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BTrDBServer).AlignedWindows(m, &bTrDBAlignedWindowsServer{stream})
}

type BTrDB_AlignedWindowsServer interface {
	Send(*AlignedWindowsResponse) error
	grpc.ServerStream
}

type bTrDBAlignedWindowsServer struct {
	grpc.ServerStream
}

func (x *bTrDBAlignedWindowsServer) Send(m *AlignedWindowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BTrDB_Windows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WindowsParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BTrDBServer).Windows(m, &bTrDBWindowsServer{stream})
}

type BTrDB_WindowsServer interface {
	Send(*WindowsResponse) error
	grpc.ServerStream
}

type bTrDBWindowsServer struct {
	grpc.ServerStream
}

func (x *bTrDBWindowsServer) Send(m *WindowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BTrDB_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Version(ctx, req.(*VersionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDB_Nearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NearestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Nearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Nearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Nearest(ctx, req.(*NearestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDB_Changes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Changes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Changes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Changes(ctx, req.(*ChangesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDB_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Insert(ctx, req.(*InsertParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Delete(ctx, req.(*DeleteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDB_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinterface.BTrDB/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBServer).Info(ctx, req.(*InfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BTrDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcinterface.BTrDB",
	HandlerType: (*BTrDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _BTrDB_Version_Handler,
		},
		{
			MethodName: "Nearest",
			Handler:    _BTrDB_Nearest_Handler,
		},
		{
			MethodName: "Changes",
			Handler:    _BTrDB_Changes_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _BTrDB_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BTrDB_Delete_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _BTrDB_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RawValues",
			Handler:       _BTrDB_RawValues_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AlignedWindows",
			Handler:       _BTrDB_AlignedWindows_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Windows",
			Handler:       _BTrDB_Windows_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("btrdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0x96, 0x9b, 0xc9, 0xa4, 0x39, 0xcd, 0x34, 0xc5, 0x04, 0x18, 0xb2, 0xa5, 0x8a, 0xbc, 0x42,
	0x04, 0x24, 0x9a, 0xaa, 0xbb, 0xe2, 0x82, 0x0b, 0x24, 0x96, 0x4a, 0x55, 0x91, 0x16, 0xad, 0x0c,
	0xda, 0xde, 0x21, 0x39, 0x19, 0x37, 0x19, 0x48, 0xec, 0x60, 0x3b, 0xcd, 0x22, 0xee, 0x80, 0x37,
	0x40, 0xe2, 0x86, 0xc7, 0x40, 0x02, 0x71, 0x81, 0xc4, 0x3b, 0xf0, 0x0a, 0x3c, 0x08, 0xb2, 0x67,
	0x26, 0x9d, 0xbf, 0xed, 0x22, 0x90, 0x1a, 0x71, 0x13, 0x1d, 0x7b, 0x8e, 0xbf, 0xf3, 0xf9, 0x1c,
	0xfb, 0x3b, 0x0e, 0xec, 0x8d, 0x8d, 0x8a, 0xc6, 0xc7, 0x4b, 0x25, 0x8d, 0xc4, 0xc1, 0x54, 0x2d,
	0x27, 0xb1, 0x30, 0x5c, 0x5d, 0xb1, 0x09, 0xef, 0x1f, 0x4e, 0xa5, 0x9c, 0xce, 0xf9, 0x88, 0x2d,
	0xe3, 0x11, 0x13, 0x42, 0x1a, 0x66, 0x62, 0x29, 0x74, 0xe2, 0x4c, 0xbe, 0x82, 0x2e, 0x65, 0xeb,
	0xa7, 0x6c, 0xbe, 0xe2, 0xfa, 0x09, 0x53, 0x6c, 0xa1, 0x31, 0x06, 0x6f, 0xb5, 0x8a, 0xa3, 0x10,
	0x0d, 0xd0, 0xb0, 0x43, 0x9d, 0x8d, 0x7b, 0xd0, 0xd4, 0x86, 0x29, 0x13, 0xee, 0x0c, 0xd0, 0xf0,
	0x80, 0x26, 0x03, 0x7c, 0x00, 0x0d, 0x2e, 0xa2, 0xb0, 0xe1, 0xe6, 0xac, 0x89, 0x09, 0x74, 0xae,
	0xb9, 0xd2, 0xb1, 0x14, 0x8f, 0xd9, 0x17, 0x52, 0x85, 0xde, 0x00, 0x0d, 0x3d, 0x5a, 0x98, 0x23,
	0xbf, 0x22, 0x78, 0x69, 0x13, 0x93, 0x72, 0xbd, 0x94, 0x42, 0x73, 0xfc, 0x36, 0x78, 0xda, 0x30,
	0xe3, 0xa2, 0xee, 0x9d, 0xbe, 0x72, 0x5c, 0xd8, 0xc4, 0xf1, 0xa7, 0x86, 0x99, 0x95, 0xa6, 0xce,
	0xa5, 0x12, 0x64, 0xa7, 0x1a, 0x24, 0xef, 0x13, 0x0b, 0xa9, 0x1c, 0xc7, 0x9c, 0x8f, 0x9d, 0xc3,
	0x23, 0xf0, 0xaf, 0x1d, 0x89, 0xd0, 0x1b, 0x34, 0x86, 0x7b, 0xa7, 0xaf, 0x95, 0x82, 0x52, 0xb6,
	0x7e, 0x22, 0x63, 0x61, 0x68, 0xea, 0x46, 0x7e, 0x44, 0xd0, 0xfb, 0x70, 0x1e, 0x4f, 0x05, 0x8f,
	0x2e, 0x63, 0x11, 0xc9, 0xf5, 0x1d, 0xa5, 0x0c, 0x1f, 0x01, 0x2c, 0x2d, 0x93, 0xcb, 0x38, 0x32,
	0xb3, 0xb0, 0x39, 0x40, 0xc3, 0x80, 0xe6, 0x66, 0xc8, 0xef, 0x08, 0x5e, 0x2d, 0x12, 0xdb, 0x66,
	0x5e, 0x4f, 0x4a, 0x79, 0x0d, 0x6b, 0x82, 0x16, 0x13, 0xfb, 0x13, 0x82, 0xe0, 0x6e, 0x33, 0xda,
	0x83, 0xe6, 0x7a, 0x93, 0x4c, 0x8f, 0x26, 0x03, 0x3b, 0x1b, 0xf1, 0xa5, 0x99, 0x85, 0xbe, 0x4b,
	0x71, 0x32, 0x20, 0xbf, 0x20, 0xe8, 0xfe, 0x2f, 0xd3, 0x7a, 0x1f, 0x82, 0xa7, 0x09, 0xc2, 0xf3,
	0xb3, 0x4a, 0xbe, 0x47, 0xd0, 0x4d, 0xbd, 0xb6, 0xb8, 0x3b, 0xb2, 0x86, 0xe0, 0x13, 0xce, 0x14,
	0xd7, 0xe6, 0x96, 0x13, 0x80, 0xc1, 0x33, 0xf1, 0x82, 0xa7, 0x07, 0xc0, 0xd9, 0x15, 0x02, 0x8d,
	0x1a, 0x02, 0x7d, 0xd8, 0x1d, 0xb3, 0xc9, 0x97, 0x6b, 0xa6, 0x22, 0x77, 0x1a, 0x76, 0xe9, 0x66,
	0x4c, 0x7e, 0x46, 0xd0, 0x4d, 0x23, 0x6f, 0xb3, 0xba, 0xef, 0x42, 0xd3, 0x55, 0xcd, 0xf1, 0xbb,
	0x45, 0x8b, 0x12, 0x2f, 0xf2, 0x0d, 0x04, 0x1f, 0xcd, 0x98, 0x98, 0xde, 0xaa, 0xda, 0x87, 0xd0,
	0xbe, 0x52, 0x72, 0x91, 0x27, 0x76, 0x33, 0x81, 0x43, 0x68, 0x19, 0x99, 0xcf, 0x59, 0x36, 0xb4,
	0x72, 0xa3, 0xb8, 0x96, 0xf3, 0x95, 0xed, 0x14, 0x8e, 0x50, 0x40, 0x73, 0x33, 0xe4, 0x37, 0x04,
	0xdd, 0x34, 0xfa, 0x36, 0x53, 0xf6, 0x00, 0x7c, 0xe5, 0x48, 0xa4, 0x17, 0xe2, 0x5e, 0x29, 0x68,
	0x42, 0x31, 0xa2, 0xf6, 0x97, 0xa6, 0xae, 0x64, 0x0a, 0x9d, 0x0b, 0xa1, 0xb9, 0x7a, 0xc1, 0x31,
	0xd3, 0x5f, 0x8b, 0x89, 0x23, 0xb6, 0x4b, 0x9d, 0x9d, 0x6b, 0x16, 0x8d, 0x7f, 0xd6, 0x2c, 0xbe,
	0x43, 0xb0, 0x9f, 0x44, 0xda, 0xe6, 0xb5, 0xfa, 0x18, 0x3a, 0x67, 0x7c, 0xce, 0x0d, 0xff, 0xef,
	0xba, 0xea, 0x76, 0x94, 0x80, 0x6d, 0x73, 0x47, 0x1d, 0x80, 0x0b, 0x71, 0x25, 0x93, 0xfd, 0x90,
	0xb1, 0x2d, 0xe7, 0x95, 0xfc, 0x37, 0x84, 0xde, 0x02, 0x6f, 0xc1, 0xf4, 0xcc, 0x11, 0xd9, 0x3b,
	0x7d, 0xb9, 0xe4, 0xfa, 0x98, 0xe9, 0x19, 0x75, 0x0e, 0xe4, 0x21, 0xec, 0x66, 0xd5, 0xdd, 0x28,
	0x10, 0xca, 0x29, 0x50, 0x2f, 0xbb, 0xba, 0x16, 0x09, 0x65, 0x37, 0x74, 0x01, 0xed, 0x8d, 0x22,
	0xd7, 0x2e, 0x3b, 0x80, 0xc6, 0x22, 0x16, 0xe9, 0x22, 0x6b, 0x5a, 0xaf, 0x05, 0x67, 0xc2, 0x6d,
	0x1b, 0x51, 0x67, 0x3b, 0x2f, 0xf6, 0xcc, 0x5d, 0x42, 0xeb, 0xc5, 0x9e, 0xd9, 0x70, 0x13, 0xb9,
	0x12, 0xc6, 0xb5, 0x2e, 0x9f, 0x26, 0x03, 0xf2, 0x1e, 0x74, 0xf2, 0xe7, 0xfd, 0xa6, 0xa8, 0xa8,
	0xa6, 0xa8, 0x3b, 0x37, 0x45, 0xbd, 0x04, 0x3f, 0xc9, 0x8a, 0x8d, 0x3e, 0x91, 0x51, 0xc2, 0x31,
	0xa0, 0xce, 0x76, 0xd1, 0xf5, 0xd4, 0xf9, 0xb7, 0xa9, 0x35, 0x37, 0x59, 0x6b, 0xbc, 0x28, 0x6b,
	0x3e, 0x78, 0x76, 0x74, 0xfa, 0x47, 0x13, 0x9a, 0x8f, 0x3e, 0x53, 0x67, 0x8f, 0x30, 0x87, 0xf6,
	0xe6, 0xdd, 0x87, 0x8f, 0xaa, 0xf7, 0x27, 0xff, 0x0a, 0xed, 0x0f, 0x9e, 0xf7, 0x3d, 0x2b, 0x35,
	0xe9, 0x7d, 0xfb, 0xe7, 0x5f, 0x3f, 0xec, 0xec, 0x93, 0xf6, 0xe8, 0xfa, 0xe1, 0xf1, 0xc9, 0x48,
	0xb1, 0xf5, 0xfb, 0xe8, 0x9d, 0x13, 0x84, 0x3f, 0x87, 0xfd, 0xe2, 0x5b, 0x08, 0xdf, 0x2f, 0x61,
	0xd5, 0xbd, 0xe1, 0xfa, 0x6f, 0xde, 0xea, 0x94, 0x45, 0x3d, 0x41, 0xf8, 0x02, 0x5a, 0x19, 0xf0,
	0x61, 0x69, 0x4d, 0x11, 0xf1, 0xa8, 0xfe, 0x6b, 0x0e, 0xea, 0x1c, 0x5a, 0x69, 0xeb, 0xad, 0x40,
	0x15, 0x1a, 0x77, 0x05, 0xaa, 0xdc, 0xb0, 0xcf, 0xa1, 0x95, 0xf6, 0xb0, 0x0a, 0x50, 0xa1, 0xab,
	0x56, 0x80, 0xca, 0x9d, 0xef, 0x1c, 0x5a, 0xa9, 0xb2, 0x57, 0x80, 0x0a, 0xfd, 0xa6, 0x02, 0x54,
	0xee, 0x07, 0x67, 0xe0, 0x27, 0xea, 0x87, 0xcb, 0xb2, 0x9c, 0x97, 0xdf, 0xfe, 0x1b, 0xb5, 0x1f,
	0xf3, 0x28, 0x89, 0xe2, 0x54, 0x50, 0xf2, 0xaa, 0x56, 0x41, 0x29, 0xa9, 0xd4, 0x07, 0xe0, 0x59,
	0x91, 0xc0, 0xaf, 0x57, 0x82, 0x65, 0x3a, 0xd2, 0xbf, 0x57, 0xf3, 0x29, 0x5b, 0x3f, 0xf6, 0xdd,
	0x7f, 0xa5, 0x07, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf1, 0xf5, 0x6d, 0x67, 0x0d, 0x00,
	0x00,
}
