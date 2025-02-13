// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fence_module.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Fence struct {
	// service ports enable lazyload
	WormholePort []string `protobuf:"bytes,1,rep,name=wormholePort,proto3" json:"wormholePort,omitempty"`
	// whether enable ServiceFence auto generating
	// default value is false
	AutoFence bool `protobuf:"varint,2,opt,name=autoFence,proto3" json:"autoFence,omitempty"`
	// the namespace list which enable lazyload
	Namespace []string `protobuf:"bytes,3,rep,name=namespace,proto3" json:"namespace,omitempty"`
	// custom outside dispatch traffic rules
	Dispatches []*Dispatch `protobuf:"bytes,4,rep,name=dispatches,proto3" json:"dispatches,omitempty"`
	// can convert to one or many domain alias rules
	DomainAliases []*DomainAlias `protobuf:"bytes,5,rep,name=domainAliases,proto3" json:"domainAliases,omitempty"`
	// default behavior of create fence or not when autoFence is true
	// default value is false
	DefaultFence bool `protobuf:"varint,6,opt,name=defaultFence,proto3" json:"defaultFence,omitempty"`
	// whether enable http service port auto management
	// default value is false
	AutoPort bool `protobuf:"varint,7,opt,name=autoPort,proto3" json:"autoPort,omitempty"`
	// specify the ns of global-siecar, same as slimeNamespace by default
	ClusterGsNamespace   string   `protobuf:"bytes,8,opt,name=clusterGsNamespace,proto3" json:"clusterGsNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fence) Reset()         { *m = Fence{} }
func (m *Fence) String() string { return proto.CompactTextString(m) }
func (*Fence) ProtoMessage()    {}
func (*Fence) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eebc4b237a55c9b, []int{0}
}
func (m *Fence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fence.Unmarshal(m, b)
}
func (m *Fence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fence.Marshal(b, m, deterministic)
}
func (m *Fence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fence.Merge(m, src)
}
func (m *Fence) XXX_Size() int {
	return xxx_messageInfo_Fence.Size(m)
}
func (m *Fence) XXX_DiscardUnknown() {
	xxx_messageInfo_Fence.DiscardUnknown(m)
}

var xxx_messageInfo_Fence proto.InternalMessageInfo

func (m *Fence) GetWormholePort() []string {
	if m != nil {
		return m.WormholePort
	}
	return nil
}

func (m *Fence) GetAutoFence() bool {
	if m != nil {
		return m.AutoFence
	}
	return false
}

func (m *Fence) GetNamespace() []string {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Fence) GetDispatches() []*Dispatch {
	if m != nil {
		return m.Dispatches
	}
	return nil
}

func (m *Fence) GetDomainAliases() []*DomainAlias {
	if m != nil {
		return m.DomainAliases
	}
	return nil
}

func (m *Fence) GetDefaultFence() bool {
	if m != nil {
		return m.DefaultFence
	}
	return false
}

func (m *Fence) GetAutoPort() bool {
	if m != nil {
		return m.AutoPort
	}
	return false
}

func (m *Fence) GetClusterGsNamespace() string {
	if m != nil {
		return m.ClusterGsNamespace
	}
	return ""
}

// The general idea is to assign different default traffic to different targets
// for correct processing by means of domain matching.
type Dispatch struct {
	// dispatch rule name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// domain matching contents
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// target cluster
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dispatch) Reset()         { *m = Dispatch{} }
func (m *Dispatch) String() string { return proto.CompactTextString(m) }
func (*Dispatch) ProtoMessage()    {}
func (*Dispatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eebc4b237a55c9b, []int{1}
}
func (m *Dispatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dispatch.Unmarshal(m, b)
}
func (m *Dispatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dispatch.Marshal(b, m, deterministic)
}
func (m *Dispatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dispatch.Merge(m, src)
}
func (m *Dispatch) XXX_Size() int {
	return xxx_messageInfo_Dispatch.Size(m)
}
func (m *Dispatch) XXX_DiscardUnknown() {
	xxx_messageInfo_Dispatch.DiscardUnknown(m)
}

var xxx_messageInfo_Dispatch proto.InternalMessageInfo

func (m *Dispatch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dispatch) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *Dispatch) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

// DomainAlias regexp expression, which is alias for target domain
// default value is empty
// example:
// domainAliases:
//   - pattern: (?P<service>[^\.]+)\.(?P<namespace>[^\.]+)\.svc\.cluster\.local$
//     template:
//       - $namespace.$service.service.mailsaas
type DomainAlias struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Templates            []string `protobuf:"bytes,2,rep,name=templates,proto3" json:"templates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainAlias) Reset()         { *m = DomainAlias{} }
func (m *DomainAlias) String() string { return proto.CompactTextString(m) }
func (*DomainAlias) ProtoMessage()    {}
func (*DomainAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eebc4b237a55c9b, []int{2}
}
func (m *DomainAlias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainAlias.Unmarshal(m, b)
}
func (m *DomainAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainAlias.Marshal(b, m, deterministic)
}
func (m *DomainAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainAlias.Merge(m, src)
}
func (m *DomainAlias) XXX_Size() int {
	return xxx_messageInfo_DomainAlias.Size(m)
}
func (m *DomainAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainAlias.DiscardUnknown(m)
}

var xxx_messageInfo_DomainAlias proto.InternalMessageInfo

func (m *DomainAlias) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *DomainAlias) GetTemplates() []string {
	if m != nil {
		return m.Templates
	}
	return nil
}

func init() {
	proto.RegisterType((*Fence)(nil), "slime.microservice.lazyload.config.Fence")
	proto.RegisterType((*Dispatch)(nil), "slime.microservice.lazyload.config.Dispatch")
	proto.RegisterType((*DomainAlias)(nil), "slime.microservice.lazyload.config.DomainAlias")
}

func init() { proto.RegisterFile("fence_module.proto", fileDescriptor_8eebc4b237a55c9b) }

var fileDescriptor_8eebc4b237a55c9b = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc4, 0x30,
	0x14, 0xa4, 0xdb, 0xfd, 0xe8, 0x66, 0xf5, 0x92, 0x53, 0x10, 0x0f, 0xa5, 0xa7, 0x22, 0x4b, 0x0a,
	0xfa, 0x0b, 0x14, 0x3f, 0x2e, 0x22, 0x52, 0xf0, 0xe2, 0x45, 0x62, 0xfa, 0xd6, 0x0d, 0x24, 0x4d,
	0x49, 0x52, 0x45, 0x7f, 0xbb, 0x07, 0x49, 0xba, 0xdd, 0x6e, 0x41, 0xd0, 0x5b, 0xde, 0x4c, 0x66,
	0xc8, 0x4c, 0x1e, 0xc2, 0x1b, 0xa8, 0x39, 0xbc, 0x28, 0x5d, 0xb5, 0x12, 0x68, 0x63, 0xb4, 0xd3,
	0x38, 0xb3, 0x52, 0x28, 0xa0, 0x4a, 0x70, 0xa3, 0x2d, 0x98, 0x77, 0xc1, 0x81, 0x4a, 0xf6, 0xf5,
	0x29, 0x35, 0xab, 0x28, 0xd7, 0xf5, 0x46, 0xbc, 0x65, 0xdf, 0x13, 0x34, 0xbb, 0xf5, 0x52, 0x9c,
	0xa1, 0xa3, 0x0f, 0x6d, 0xd4, 0x56, 0x4b, 0x78, 0xd4, 0xc6, 0x91, 0x28, 0x8d, 0xf3, 0x65, 0x39,
	0xc2, 0xf0, 0x29, 0x5a, 0xb2, 0xd6, 0xe9, 0x20, 0x20, 0x93, 0x34, 0xca, 0x93, 0x72, 0x00, 0x3c,
	0x5b, 0x33, 0x05, 0xb6, 0x61, 0x1c, 0x48, 0x1c, 0xe4, 0x03, 0x80, 0xef, 0x11, 0xaa, 0x84, 0x6d,
	0x98, 0xe3, 0x5b, 0xb0, 0x64, 0x9a, 0xc6, 0xf9, 0xea, 0x7c, 0x4d, 0xff, 0x7e, 0x22, 0xbd, 0xde,
	0xa9, 0xca, 0x03, 0x3d, 0x7e, 0x42, 0xc7, 0x95, 0x56, 0x4c, 0xd4, 0x97, 0x52, 0x30, 0x0b, 0x96,
	0xcc, 0x82, 0x61, 0xf1, 0x2f, 0xc3, 0x41, 0x58, 0x8e, 0x5d, 0x7c, 0x09, 0x15, 0x6c, 0x58, 0x2b,
	0x5d, 0x97, 0x71, 0x1e, 0x32, 0x8e, 0x30, 0x7c, 0x82, 0x12, 0x9f, 0x39, 0x94, 0xb4, 0x08, 0xfc,
	0x7e, 0xc6, 0x14, 0x61, 0x2e, 0x5b, 0xeb, 0xc0, 0xdc, 0xd9, 0x87, 0x7d, 0x17, 0x49, 0x1a, 0xe5,
	0xcb, 0xf2, 0x17, 0x26, 0x2b, 0x51, 0xd2, 0xc7, 0xc3, 0x18, 0x4d, 0x7d, 0x5b, 0x24, 0x0a, 0xb7,
	0xc3, 0x19, 0x13, 0xb4, 0xe8, 0x1e, 0x68, 0xc9, 0x24, 0x14, 0xda, 0x8f, 0x9e, 0xd9, 0xf9, 0x91,
	0x38, 0x08, 0xfa, 0x31, 0xbb, 0x41, 0xab, 0x83, 0x84, 0xfe, 0x62, 0xc3, 0x9c, 0x03, 0x53, 0xef,
	0x9c, 0xfb, 0xd1, 0xff, 0x97, 0x03, 0xd5, 0x48, 0xe6, 0xa0, 0xb7, 0x1f, 0x80, 0xab, 0xf5, 0xf3,
	0x59, 0xd7, 0xa5, 0xd0, 0x45, 0x38, 0x14, 0xdd, 0x72, 0xd9, 0xa2, 0xef, 0xb3, 0x60, 0x8d, 0x28,
	0xba, 0x4e, 0x5f, 0xe7, 0x61, 0xe5, 0x2e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0xb4, 0x0c,
	0x4e, 0x88, 0x02, 0x00, 0x00,
}
