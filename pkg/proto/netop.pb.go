// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netop.proto

package netops

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// TcType represents Traffic Control Type.
type TcType int32

const (
	TcType_BANDWIDTH_CONTROL TcType = 0
)

var TcType_name = map[int32]string{
	0: "BANDWIDTH_CONTROL",
}

var TcType_value = map[string]int32{
	"BANDWIDTH_CONTROL": 0,
}

func (x TcType) String() string {
	return proto.EnumName(TcType_name, int32(x))
}

func (TcType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{0}
}

type ClassType int32

const (
	ClassType_HTB ClassType = 0
	ClassType_TBF ClassType = 1
)

var ClassType_name = map[int32]string{
	0: "HTB",
	1: "TBF",
}

var ClassType_value = map[string]int32{
	"HTB": 0,
	"TBF": 1,
}

func (x ClassType) String() string {
	return proto.EnumName(ClassType_name, int32(x))
}

func (ClassType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{1}
}

// Generic event types enum
type EventType int32

const (
	EventType_EV_CREATE EventType = 0
	EventType_EV_UPDATE EventType = 1
	EventType_EV_DELETE EventType = 2
)

var EventType_name = map[int32]string{
	0: "EV_CREATE",
	1: "EV_UPDATE",
	2: "EV_DELETE",
}

var EventType_value = map[string]int32{
	"EV_CREATE": 0,
	"EV_UPDATE": 1,
	"EV_DELETE": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{2}
}

// slice gateway-host-type
type SliceGwHostType int32

const (
	SliceGwHostType_SLICE_GW_SERVER SliceGwHostType = 0
	SliceGwHostType_SLICE_GW_CLIENT SliceGwHostType = 1
)

var SliceGwHostType_name = map[int32]string{
	0: "SLICE_GW_SERVER",
	1: "SLICE_GW_CLIENT",
}

var SliceGwHostType_value = map[string]int32{
	"SLICE_GW_SERVER": 0,
	"SLICE_GW_CLIENT": 1,
}

func (x SliceGwHostType) String() string {
	return proto.EnumName(SliceGwHostType_name, int32(x))
}

func (SliceGwHostType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{3}
}

// Response represents the netops response format.
type Response struct {
	StatusMsg            string   `protobuf:"bytes,1,opt,name=statusMsg,proto3" json:"statusMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{0}
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

func (m *Response) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

// Slice QoS Profile
type SliceQosProfile struct {
	// Name of the slice
	SliceName string `protobuf:"bytes,1,opt,name=sliceName,proto3" json:"sliceName,omitempty"`
	// Slice Identifier //TODO
	SliceId string `protobuf:"bytes,2,opt,name=sliceId,proto3" json:"sliceId,omitempty"`
	// Name of the QoS profile attached to the slice
	QosProfileName string `protobuf:"bytes,3,opt,name=qosProfileName,proto3" json:"qosProfileName,omitempty"`
	// TC type -  Bandwidth control
	TcType TcType `protobuf:"varint,4,opt,name=tcType,proto3,enum=netops.TcType" json:"tcType,omitempty"`
	// ClassType - HTB   ( HTB)
	ClassType ClassType `protobuf:"varint,5,opt,name=ClassType,proto3,enum=netops.ClassType" json:"ClassType,omitempty"`
	// Bandwidth Ceiling in Mbps  - 5 Mbps (100k - 100 Mbps)
	BwCeiling uint32 `protobuf:"varint,6,opt,name=bwCeiling,proto3" json:"bwCeiling,omitempty"`
	// Bandwidth Guaranteed -  1 Mbps ( 100k- 100 Mbps)
	BwGuaranteed uint32 `protobuf:"varint,7,opt,name=bwGuaranteed,proto3" json:"bwGuaranteed,omitempty"`
	// Priority - 2 (Number 0-3)
	Priority uint32 `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	// Dscp class to mark inter cluster traffic
	DscpClass            string   `protobuf:"bytes,9,opt,name=dscpClass,proto3" json:"dscpClass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SliceQosProfile) Reset()         { *m = SliceQosProfile{} }
func (m *SliceQosProfile) String() string { return proto.CompactTextString(m) }
func (*SliceQosProfile) ProtoMessage()    {}
func (*SliceQosProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{1}
}

func (m *SliceQosProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceQosProfile.Unmarshal(m, b)
}
func (m *SliceQosProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceQosProfile.Marshal(b, m, deterministic)
}
func (m *SliceQosProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceQosProfile.Merge(m, src)
}
func (m *SliceQosProfile) XXX_Size() int {
	return xxx_messageInfo_SliceQosProfile.Size(m)
}
func (m *SliceQosProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceQosProfile.DiscardUnknown(m)
}

var xxx_messageInfo_SliceQosProfile proto.InternalMessageInfo

func (m *SliceQosProfile) GetSliceName() string {
	if m != nil {
		return m.SliceName
	}
	return ""
}

func (m *SliceQosProfile) GetSliceId() string {
	if m != nil {
		return m.SliceId
	}
	return ""
}

func (m *SliceQosProfile) GetQosProfileName() string {
	if m != nil {
		return m.QosProfileName
	}
	return ""
}

func (m *SliceQosProfile) GetTcType() TcType {
	if m != nil {
		return m.TcType
	}
	return TcType_BANDWIDTH_CONTROL
}

func (m *SliceQosProfile) GetClassType() ClassType {
	if m != nil {
		return m.ClassType
	}
	return ClassType_HTB
}

func (m *SliceQosProfile) GetBwCeiling() uint32 {
	if m != nil {
		return m.BwCeiling
	}
	return 0
}

func (m *SliceQosProfile) GetBwGuaranteed() uint32 {
	if m != nil {
		return m.BwGuaranteed
	}
	return 0
}

func (m *SliceQosProfile) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SliceQosProfile) GetDscpClass() string {
	if m != nil {
		return m.DscpClass
	}
	return ""
}

// Slice event message
type SliceLifeCycleEvent struct {
	// Name of the slice
	SliceName string `protobuf:"bytes,1,opt,name=sliceName,proto3" json:"sliceName,omitempty"`
	// Event type
	Event                EventType `protobuf:"varint,2,opt,name=event,proto3,enum=netops.EventType" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SliceLifeCycleEvent) Reset()         { *m = SliceLifeCycleEvent{} }
func (m *SliceLifeCycleEvent) String() string { return proto.CompactTextString(m) }
func (*SliceLifeCycleEvent) ProtoMessage()    {}
func (*SliceLifeCycleEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{2}
}

func (m *SliceLifeCycleEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceLifeCycleEvent.Unmarshal(m, b)
}
func (m *SliceLifeCycleEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceLifeCycleEvent.Marshal(b, m, deterministic)
}
func (m *SliceLifeCycleEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceLifeCycleEvent.Merge(m, src)
}
func (m *SliceLifeCycleEvent) XXX_Size() int {
	return xxx_messageInfo_SliceLifeCycleEvent.Size(m)
}
func (m *SliceLifeCycleEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceLifeCycleEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SliceLifeCycleEvent proto.InternalMessageInfo

func (m *SliceLifeCycleEvent) GetSliceName() string {
	if m != nil {
		return m.SliceName
	}
	return ""
}

func (m *SliceLifeCycleEvent) GetEvent() EventType {
	if m != nil {
		return m.Event
	}
	return EventType_EV_CREATE
}

// NetOpConnectionContext - NetOp Connection Context.
type NetOpConnectionContext struct {
	// Slice-Id
	SliceId string `protobuf:"bytes,1,opt,name=sliceId,proto3" json:"sliceId,omitempty"`
	// Local slice gateway ID
	LocalSliceGwId string `protobuf:"bytes,2,opt,name=localSliceGwId,proto3" json:"localSliceGwId,omitempty"`
	// Local slice gateway VPN IP
	LocalSliceGwVpnIP string `protobuf:"bytes,3,opt,name=localSliceGwVpnIP,proto3" json:"localSliceGwVpnIP,omitempty"`
	// Local slice gateway-host-type  -  client/server
	LocalSliceGwHostType SliceGwHostType `protobuf:"varint,4,opt,name=localSliceGwHostType,proto3,enum=netops.SliceGwHostType" json:"localSliceGwHostType,omitempty"`
	// Local slice gateway NSM Subnet
	LocalSliceGwNsmSubnet string `protobuf:"bytes,5,opt,name=localSliceGwNsmSubnet,proto3" json:"localSliceGwNsmSubnet,omitempty"`
	// Local slice gateway Node IP
	LocalSliceGwNodeIP string `protobuf:"bytes,6,opt,name=localSliceGwNodeIP,proto3" json:"localSliceGwNodeIP,omitempty"`
	// Local slice gateway Node Port
	LocalSliceGwNodePort string `protobuf:"bytes,7,opt,name=localSliceGwNodePort,proto3" json:"localSliceGwNodePort,omitempty"`
	// Remote slice gateway ID
	RemoteSliceGwId string `protobuf:"bytes,8,opt,name=remoteSliceGwId,proto3" json:"remoteSliceGwId,omitempty"`
	// Remote slice gateway VPN IP
	RemoteSliceGwVpnIP string `protobuf:"bytes,9,opt,name=remoteSliceGwVpnIP,proto3" json:"remoteSliceGwVpnIP,omitempty"`
	// Remote-slice gateway-host-type client or server
	RemoteSliceGwHostType SliceGwHostType `protobuf:"varint,10,opt,name=remoteSliceGwHostType,proto3,enum=netops.SliceGwHostType" json:"remoteSliceGwHostType,omitempty"`
	// Remote slice gateway NSM subnet
	RemoteSliceGwNsmSubnet string `protobuf:"bytes,11,opt,name=remoteSliceGwNsmSubnet,proto3" json:"remoteSliceGwNsmSubnet,omitempty"`
	// Remote slice gateway Node IP
	RemoteSliceGwNodeIP string `protobuf:"bytes,12,opt,name=remoteSliceGwNodeIP,proto3" json:"remoteSliceGwNodeIP,omitempty"`
	// Remote slice gateway Node Port
	RemoteSliceGwNodePort string   `protobuf:"bytes,13,opt,name=remoteSliceGwNodePort,proto3" json:"remoteSliceGwNodePort,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *NetOpConnectionContext) Reset()         { *m = NetOpConnectionContext{} }
func (m *NetOpConnectionContext) String() string { return proto.CompactTextString(m) }
func (*NetOpConnectionContext) ProtoMessage()    {}
func (*NetOpConnectionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0dbd33d19c0b5c, []int{3}
}

func (m *NetOpConnectionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetOpConnectionContext.Unmarshal(m, b)
}
func (m *NetOpConnectionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetOpConnectionContext.Marshal(b, m, deterministic)
}
func (m *NetOpConnectionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetOpConnectionContext.Merge(m, src)
}
func (m *NetOpConnectionContext) XXX_Size() int {
	return xxx_messageInfo_NetOpConnectionContext.Size(m)
}
func (m *NetOpConnectionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_NetOpConnectionContext.DiscardUnknown(m)
}

var xxx_messageInfo_NetOpConnectionContext proto.InternalMessageInfo

func (m *NetOpConnectionContext) GetSliceId() string {
	if m != nil {
		return m.SliceId
	}
	return ""
}

func (m *NetOpConnectionContext) GetLocalSliceGwId() string {
	if m != nil {
		return m.LocalSliceGwId
	}
	return ""
}

func (m *NetOpConnectionContext) GetLocalSliceGwVpnIP() string {
	if m != nil {
		return m.LocalSliceGwVpnIP
	}
	return ""
}

func (m *NetOpConnectionContext) GetLocalSliceGwHostType() SliceGwHostType {
	if m != nil {
		return m.LocalSliceGwHostType
	}
	return SliceGwHostType_SLICE_GW_SERVER
}

func (m *NetOpConnectionContext) GetLocalSliceGwNsmSubnet() string {
	if m != nil {
		return m.LocalSliceGwNsmSubnet
	}
	return ""
}

func (m *NetOpConnectionContext) GetLocalSliceGwNodeIP() string {
	if m != nil {
		return m.LocalSliceGwNodeIP
	}
	return ""
}

func (m *NetOpConnectionContext) GetLocalSliceGwNodePort() string {
	if m != nil {
		return m.LocalSliceGwNodePort
	}
	return ""
}

func (m *NetOpConnectionContext) GetRemoteSliceGwId() string {
	if m != nil {
		return m.RemoteSliceGwId
	}
	return ""
}

func (m *NetOpConnectionContext) GetRemoteSliceGwVpnIP() string {
	if m != nil {
		return m.RemoteSliceGwVpnIP
	}
	return ""
}

func (m *NetOpConnectionContext) GetRemoteSliceGwHostType() SliceGwHostType {
	if m != nil {
		return m.RemoteSliceGwHostType
	}
	return SliceGwHostType_SLICE_GW_SERVER
}

func (m *NetOpConnectionContext) GetRemoteSliceGwNsmSubnet() string {
	if m != nil {
		return m.RemoteSliceGwNsmSubnet
	}
	return ""
}

func (m *NetOpConnectionContext) GetRemoteSliceGwNodeIP() string {
	if m != nil {
		return m.RemoteSliceGwNodeIP
	}
	return ""
}

func (m *NetOpConnectionContext) GetRemoteSliceGwNodePort() string {
	if m != nil {
		return m.RemoteSliceGwNodePort
	}
	return ""
}

func init() {
	proto.RegisterEnum("netops.TcType", TcType_name, TcType_value)
	proto.RegisterEnum("netops.ClassType", ClassType_name, ClassType_value)
	proto.RegisterEnum("netops.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("netops.SliceGwHostType", SliceGwHostType_name, SliceGwHostType_value)
	proto.RegisterType((*Response)(nil), "netops.Response")
	proto.RegisterType((*SliceQosProfile)(nil), "netops.SliceQosProfile")
	proto.RegisterType((*SliceLifeCycleEvent)(nil), "netops.SliceLifeCycleEvent")
	proto.RegisterType((*NetOpConnectionContext)(nil), "netops.NetOpConnectionContext")
}

func init() {
	proto.RegisterFile("netop.proto", fileDescriptor_de0dbd33d19c0b5c)
}

var fileDescriptor_de0dbd33d19c0b5c = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0x8d, 0xe1, 0x12, 0xe2, 0x81, 0x80, 0x59, 0x6e, 0xc0, 0x97, 0x7b, 0x6f, 0x8b, 0xf2, 0x40,
	0xa3, 0xa8, 0x0a, 0x15, 0xad, 0xaa, 0x4a, 0x3c, 0x11, 0xc7, 0x85, 0xa8, 0x21, 0xa4, 0x1b, 0x03,
	0x52, 0x55, 0x29, 0x72, 0x9c, 0x05, 0x59, 0x32, 0x5e, 0xd7, 0xbb, 0x40, 0xf9, 0xb3, 0x7e, 0x41,
	0x3f, 0xa8, 0x5f, 0x50, 0x79, 0x1c, 0x3b, 0x76, 0xe2, 0xb6, 0x6f, 0x99, 0x73, 0x66, 0x26, 0xb3,
	0xe7, 0xec, 0x8e, 0x61, 0xcd, 0x67, 0x92, 0x07, 0xad, 0x20, 0xe4, 0x92, 0x93, 0x32, 0x06, 0xa2,
	0xde, 0x80, 0x0a, 0x65, 0x22, 0xe0, 0xbe, 0x60, 0xe4, 0x3f, 0x50, 0x85, 0xb4, 0xe5, 0xbd, 0x38,
	0x17, 0xb7, 0xba, 0xb2, 0xaf, 0x34, 0x54, 0x3a, 0x03, 0xea, 0xdf, 0x97, 0x60, 0x73, 0xe8, 0xb9,
	0x0e, 0xfb, 0xc8, 0xc5, 0x20, 0xe4, 0x37, 0xae, 0x17, 0x57, 0x44, 0x50, 0xdf, 0xbe, 0x63, 0x69,
	0x45, 0x02, 0x10, 0x1d, 0x56, 0x31, 0xe8, 0x4e, 0xf4, 0x25, 0xe4, 0x92, 0x90, 0x1c, 0xc0, 0xc6,
	0x97, 0xb4, 0x0b, 0x16, 0x2f, 0x63, 0xc2, 0x1c, 0x4a, 0x0e, 0xa0, 0x2c, 0x1d, 0xeb, 0x29, 0x60,
	0xfa, 0x5f, 0xfb, 0x4a, 0x63, 0xe3, 0x68, 0xa3, 0x15, 0x8f, 0xdd, 0xb2, 0x10, 0xa5, 0x53, 0x96,
	0x1c, 0x82, 0x6a, 0x78, 0xb6, 0x10, 0x98, 0xba, 0x82, 0xa9, 0x5b, 0x49, 0x6a, 0x4a, 0xd0, 0x59,
	0x4e, 0x34, 0xf8, 0xf8, 0xd1, 0x60, 0xae, 0xe7, 0xfa, 0xb7, 0x7a, 0x79, 0x5f, 0x69, 0x54, 0xe9,
	0x0c, 0x20, 0x75, 0x58, 0x1f, 0x3f, 0x9e, 0xde, 0xdb, 0xa1, 0xed, 0x4b, 0xc6, 0x26, 0xfa, 0x2a,
	0x26, 0xe4, 0x30, 0xb2, 0x07, 0x95, 0x20, 0x74, 0x79, 0xe8, 0xca, 0x27, 0xbd, 0x82, 0x7c, 0x1a,
	0x47, 0xdd, 0x27, 0xc2, 0x09, 0xf0, 0xef, 0x74, 0x35, 0x96, 0x25, 0x05, 0xea, 0x9f, 0x61, 0x1b,
	0x75, 0xec, 0xb9, 0x37, 0xcc, 0x78, 0x72, 0x3c, 0x66, 0x3e, 0x30, 0x5f, 0xfe, 0x41, 0xcb, 0x17,
	0xb0, 0xc2, 0xa2, 0x34, 0x54, 0x32, 0x73, 0x3a, 0xac, 0xc5, 0xd3, 0xc5, 0x7c, 0xfd, 0xdb, 0x0a,
	0xec, 0xf4, 0x99, 0xbc, 0x08, 0x0c, 0xee, 0xfb, 0xcc, 0x91, 0x2e, 0xf7, 0x0d, 0xee, 0x4b, 0xf6,
	0x55, 0x66, 0xfd, 0x50, 0x16, 0xfc, 0xf0, 0xb8, 0x63, 0x7b, 0x38, 0xd7, 0xe9, 0x63, 0x6a, 0xd8,
	0x1c, 0x4a, 0x5e, 0xc2, 0x56, 0x16, 0xb9, 0x0a, 0xfc, 0xee, 0x60, 0x6a, 0xdd, 0x22, 0x41, 0x3e,
	0xc0, 0xdf, 0x59, 0xf0, 0x8c, 0x0b, 0x99, 0xf1, 0x72, 0x37, 0x39, 0xc2, 0x1c, 0x4d, 0x0b, 0x8b,
	0xc8, 0x1b, 0xa8, 0x65, 0xf1, 0xbe, 0xb8, 0x1b, 0xde, 0x8f, 0x7d, 0x26, 0xd1, 0x6e, 0x95, 0x16,
	0x93, 0xa4, 0x05, 0x24, 0x47, 0xf0, 0x09, 0xeb, 0x0e, 0xd0, 0x70, 0x95, 0x16, 0x30, 0xe4, 0x28,
	0x3f, 0x72, 0x84, 0x0e, 0x78, 0x28, 0xf1, 0x06, 0xa8, 0xb4, 0x90, 0x23, 0x0d, 0xd8, 0x0c, 0xd9,
	0x1d, 0x97, 0x6c, 0xa6, 0x5e, 0x05, 0xd3, 0xe7, 0xe1, 0x68, 0x9a, 0x1c, 0x14, 0xeb, 0x17, 0x5f,
	0x90, 0x02, 0x86, 0x9c, 0x43, 0x2d, 0x87, 0xa6, 0x0a, 0xc2, 0xef, 0x15, 0x2c, 0xae, 0x22, 0x6f,
	0x61, 0x27, 0x47, 0xcc, 0x34, 0x5c, 0xc3, 0x11, 0x7e, 0xc1, 0x92, 0x57, 0xb0, 0x9d, 0x67, 0x62,
	0x15, 0xd7, 0xb1, 0xa8, 0x88, 0x8a, 0xcc, 0x5a, 0x80, 0x51, 0xc7, 0x6a, 0x6c, 0x56, 0x21, 0xd9,
	0x7c, 0x0e, 0xe5, 0xf8, 0x5d, 0x93, 0x1a, 0x6c, 0xb5, 0x4f, 0xfa, 0x9d, 0xeb, 0x6e, 0xc7, 0x3a,
	0x1b, 0x19, 0x17, 0x7d, 0x8b, 0x5e, 0xf4, 0xb4, 0x52, 0xf3, 0xff, 0xcc, 0x33, 0x27, 0xab, 0xb0,
	0x7c, 0x66, 0xb5, 0xb5, 0x52, 0xf4, 0xc3, 0x6a, 0xbf, 0xd7, 0x94, 0xe6, 0x3b, 0x50, 0xd3, 0xe7,
	0x40, 0xaa, 0xa0, 0x9a, 0x57, 0x23, 0x83, 0x9a, 0x27, 0x96, 0xa9, 0x95, 0xa6, 0xe1, 0xe5, 0xa0,
	0x13, 0x85, 0xca, 0x34, 0xec, 0x98, 0x3d, 0xd3, 0x32, 0xb5, 0xa5, 0xe6, 0xf1, 0x74, 0xb5, 0x65,
	0xc4, 0xda, 0x86, 0xcd, 0x61, 0xaf, 0x6b, 0x98, 0xa3, 0xd3, 0xeb, 0xd1, 0xd0, 0xa4, 0x57, 0x26,
	0xd5, 0x4a, 0x39, 0xd0, 0xe8, 0x75, 0xcd, 0xbe, 0xa5, 0x29, 0x47, 0x3f, 0x14, 0xa8, 0xe2, 0x8b,
	0x13, 0x43, 0x16, 0x3e, 0xb8, 0x0e, 0x23, 0x1d, 0xa8, 0x5d, 0x06, 0x13, 0x7b, 0x7a, 0xc2, 0xcc,
	0xbe, 0xcc, 0x3b, 0x36, 0x23, 0xf6, 0xb4, 0x84, 0x48, 0x96, 0x71, 0xbd, 0x44, 0x7a, 0xf0, 0x4f,
	0xa6, 0xcb, 0xdc, 0xb6, 0xf8, 0x37, 0xd7, 0x29, 0x4f, 0x16, 0x76, 0x3b, 0x87, 0xdd, 0xb8, 0xdb,
	0xe2, 0x5e, 0x78, 0x96, 0xa4, 0x17, 0xef, 0x8d, 0xa2, 0x76, 0xed, 0xb5, 0x4f, 0x6a, 0xeb, 0xf0,
	0x38, 0xc6, 0xc7, 0x65, 0xfc, 0xa6, 0xbc, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x5e, 0xd1,
	0x4f, 0x62, 0x06, 0x00, 0x00,
}