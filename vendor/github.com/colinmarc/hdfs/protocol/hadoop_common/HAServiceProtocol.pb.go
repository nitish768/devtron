// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HAServiceProtocol.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HAServiceStateProto int32

const (
	HAServiceStateProto_INITIALIZING HAServiceStateProto = 0
	HAServiceStateProto_ACTIVE       HAServiceStateProto = 1
	HAServiceStateProto_STANDBY      HAServiceStateProto = 2
)

var HAServiceStateProto_name = map[int32]string{
	0: "INITIALIZING",
	1: "ACTIVE",
	2: "STANDBY",
}
var HAServiceStateProto_value = map[string]int32{
	"INITIALIZING": 0,
	"ACTIVE":       1,
	"STANDBY":      2,
}

func (x HAServiceStateProto) Enum() *HAServiceStateProto {
	p := new(HAServiceStateProto)
	*p = x
	return p
}
func (x HAServiceStateProto) String() string {
	return proto.EnumName(HAServiceStateProto_name, int32(x))
}
func (x *HAServiceStateProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HAServiceStateProto_value, data, "HAServiceStateProto")
	if err != nil {
		return err
	}
	*x = HAServiceStateProto(value)
	return nil
}
func (HAServiceStateProto) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type HARequestSource int32

const (
	HARequestSource_REQUEST_BY_USER        HARequestSource = 0
	HARequestSource_REQUEST_BY_USER_FORCED HARequestSource = 1
	HARequestSource_REQUEST_BY_ZKFC        HARequestSource = 2
)

var HARequestSource_name = map[int32]string{
	0: "REQUEST_BY_USER",
	1: "REQUEST_BY_USER_FORCED",
	2: "REQUEST_BY_ZKFC",
}
var HARequestSource_value = map[string]int32{
	"REQUEST_BY_USER":        0,
	"REQUEST_BY_USER_FORCED": 1,
	"REQUEST_BY_ZKFC":        2,
}

func (x HARequestSource) Enum() *HARequestSource {
	p := new(HARequestSource)
	*p = x
	return p
}
func (x HARequestSource) String() string {
	return proto.EnumName(HARequestSource_name, int32(x))
}
func (x *HARequestSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HARequestSource_value, data, "HARequestSource")
	if err != nil {
		return err
	}
	*x = HARequestSource(value)
	return nil
}
func (HARequestSource) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

type HAStateChangeRequestInfoProto struct {
	ReqSource        *HARequestSource `protobuf:"varint,1,req,name=reqSource,enum=hadoop.common.HARequestSource" json:"reqSource,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HAStateChangeRequestInfoProto) Reset()                    { *m = HAStateChangeRequestInfoProto{} }
func (m *HAStateChangeRequestInfoProto) String() string            { return proto.CompactTextString(m) }
func (*HAStateChangeRequestInfoProto) ProtoMessage()               {}
func (*HAStateChangeRequestInfoProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *HAStateChangeRequestInfoProto) GetReqSource() HARequestSource {
	if m != nil && m.ReqSource != nil {
		return *m.ReqSource
	}
	return HARequestSource_REQUEST_BY_USER
}

// *
// void request
type MonitorHealthRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MonitorHealthRequestProto) Reset()                    { *m = MonitorHealthRequestProto{} }
func (m *MonitorHealthRequestProto) String() string            { return proto.CompactTextString(m) }
func (*MonitorHealthRequestProto) ProtoMessage()               {}
func (*MonitorHealthRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

// *
// void response
type MonitorHealthResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MonitorHealthResponseProto) Reset()                    { *m = MonitorHealthResponseProto{} }
func (m *MonitorHealthResponseProto) String() string            { return proto.CompactTextString(m) }
func (*MonitorHealthResponseProto) ProtoMessage()               {}
func (*MonitorHealthResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

// *
// void request
type TransitionToActiveRequestProto struct {
	ReqInfo          *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *TransitionToActiveRequestProto) Reset()                    { *m = TransitionToActiveRequestProto{} }
func (m *TransitionToActiveRequestProto) String() string            { return proto.CompactTextString(m) }
func (*TransitionToActiveRequestProto) ProtoMessage()               {}
func (*TransitionToActiveRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *TransitionToActiveRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

// *
// void response
type TransitionToActiveResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TransitionToActiveResponseProto) Reset()         { *m = TransitionToActiveResponseProto{} }
func (m *TransitionToActiveResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToActiveResponseProto) ProtoMessage()    {}
func (*TransitionToActiveResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{4}
}

// *
// void request
type TransitionToStandbyRequestProto struct {
	ReqInfo          *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *TransitionToStandbyRequestProto) Reset()         { *m = TransitionToStandbyRequestProto{} }
func (m *TransitionToStandbyRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyRequestProto) ProtoMessage()    {}
func (*TransitionToStandbyRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{5}
}

func (m *TransitionToStandbyRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

// *
// void response
type TransitionToStandbyResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TransitionToStandbyResponseProto) Reset()         { *m = TransitionToStandbyResponseProto{} }
func (m *TransitionToStandbyResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyResponseProto) ProtoMessage()    {}
func (*TransitionToStandbyResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{6}
}

// *
// void request
type GetServiceStatusRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetServiceStatusRequestProto) Reset()                    { *m = GetServiceStatusRequestProto{} }
func (m *GetServiceStatusRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetServiceStatusRequestProto) ProtoMessage()               {}
func (*GetServiceStatusRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

// *
// Returns the state of the service
type GetServiceStatusResponseProto struct {
	State *HAServiceStateProto `protobuf:"varint,1,req,name=state,enum=hadoop.common.HAServiceStateProto" json:"state,omitempty"`
	// If state is STANDBY, indicate whether it is
	// ready to become active.
	ReadyToBecomeActive *bool `protobuf:"varint,2,opt,name=readyToBecomeActive" json:"readyToBecomeActive,omitempty"`
	// If not ready to become active, a textual explanation of why not
	NotReadyReason   *string `protobuf:"bytes,3,opt,name=notReadyReason" json:"notReadyReason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetServiceStatusResponseProto) Reset()                    { *m = GetServiceStatusResponseProto{} }
func (m *GetServiceStatusResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetServiceStatusResponseProto) ProtoMessage()               {}
func (*GetServiceStatusResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *GetServiceStatusResponseProto) GetState() HAServiceStateProto {
	if m != nil && m.State != nil {
		return *m.State
	}
	return HAServiceStateProto_INITIALIZING
}

func (m *GetServiceStatusResponseProto) GetReadyToBecomeActive() bool {
	if m != nil && m.ReadyToBecomeActive != nil {
		return *m.ReadyToBecomeActive
	}
	return false
}

func (m *GetServiceStatusResponseProto) GetNotReadyReason() string {
	if m != nil && m.NotReadyReason != nil {
		return *m.NotReadyReason
	}
	return ""
}

func init() {
	proto.RegisterType((*HAStateChangeRequestInfoProto)(nil), "hadoop.common.HAStateChangeRequestInfoProto")
	proto.RegisterType((*MonitorHealthRequestProto)(nil), "hadoop.common.MonitorHealthRequestProto")
	proto.RegisterType((*MonitorHealthResponseProto)(nil), "hadoop.common.MonitorHealthResponseProto")
	proto.RegisterType((*TransitionToActiveRequestProto)(nil), "hadoop.common.TransitionToActiveRequestProto")
	proto.RegisterType((*TransitionToActiveResponseProto)(nil), "hadoop.common.TransitionToActiveResponseProto")
	proto.RegisterType((*TransitionToStandbyRequestProto)(nil), "hadoop.common.TransitionToStandbyRequestProto")
	proto.RegisterType((*TransitionToStandbyResponseProto)(nil), "hadoop.common.TransitionToStandbyResponseProto")
	proto.RegisterType((*GetServiceStatusRequestProto)(nil), "hadoop.common.GetServiceStatusRequestProto")
	proto.RegisterType((*GetServiceStatusResponseProto)(nil), "hadoop.common.GetServiceStatusResponseProto")
	proto.RegisterEnum("hadoop.common.HAServiceStateProto", HAServiceStateProto_name, HAServiceStateProto_value)
	proto.RegisterEnum("hadoop.common.HARequestSource", HARequestSource_name, HARequestSource_value)
}

func init() { proto.RegisterFile("HAServiceProtocol.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x9e, 0x3b, 0xc1, 0xd8, 0x19, 0xdb, 0x22, 0x57, 0x82, 0x12, 0xb6, 0x52, 0x72, 0x81, 0xca,
	0x18, 0x01, 0xf5, 0x0a, 0x09, 0x2e, 0x48, 0xbb, 0x74, 0x8d, 0x80, 0x02, 0x4e, 0x06, 0xda, 0x24,
	0x54, 0x99, 0xd4, 0x34, 0x91, 0x16, 0xbb, 0x4d, 0xdc, 0x4a, 0x7b, 0x03, 0x1e, 0x83, 0x77, 0xe0,
	0x75, 0x78, 0x18, 0x94, 0x26, 0xd3, 0xd2, 0x24, 0x84, 0xdd, 0xec, 0x2a, 0x89, 0xfd, 0xfd, 0xc5,
	0xe7, 0x1c, 0xc3, 0xfd, 0x81, 0x61, 0xb3, 0x70, 0xe1, 0xbb, 0xec, 0x53, 0x28, 0xa4, 0x70, 0xc5,
	0xb9, 0x3e, 0x8d, 0x5f, 0xf0, 0xb6, 0x47, 0xc7, 0x42, 0x4c, 0x75, 0x57, 0x04, 0x81, 0xe0, 0xda,
	0x37, 0xd8, 0x1f, 0x18, 0xb6, 0xa4, 0x92, 0xf5, 0x3c, 0xca, 0x27, 0x8c, 0xb0, 0xd9, 0x9c, 0x45,
	0xd2, 0xe2, 0x3f, 0xc4, 0x92, 0x88, 0xdf, 0xc0, 0x66, 0xc8, 0x66, 0xb6, 0x98, 0x87, 0x2e, 0x6b,
	0xa0, 0x56, 0xad, 0xbd, 0xd3, 0x69, 0xea, 0x2b, 0x1a, 0xfa, 0xc0, 0x48, 0x59, 0x09, 0x8a, 0x5c,
	0x11, 0xb4, 0x87, 0xf0, 0xe0, 0x83, 0xe0, 0xbe, 0x14, 0xe1, 0x80, 0xd1, 0x73, 0xe9, 0xa5, 0xc0,
	0xa5, 0xb4, 0xb6, 0x07, 0x6a, 0x6e, 0x33, 0x9a, 0x0a, 0x1e, 0x25, 0x89, 0x35, 0x0f, 0x9a, 0x4e,
	0x48, 0x79, 0xe4, 0x4b, 0x5f, 0x70, 0x47, 0x18, 0xae, 0xf4, 0x17, 0x2c, 0xcb, 0xc7, 0x7d, 0xd8,
	0x08, 0xd9, 0x2c, 0x8e, 0xba, 0x0c, 0xb6, 0xd5, 0x39, 0x2c, 0x04, 0xab, 0xf8, 0x33, 0x72, 0x49,
	0xd6, 0x1e, 0xc3, 0xa3, 0x32, 0xa7, 0x6c, 0x18, 0x7f, 0x15, 0x62, 0x4b, 0xca, 0xc7, 0xdf, 0x2f,
	0x6e, 0x24, 0x8d, 0x06, 0xad, 0x52, 0xab, 0x6c, 0x9c, 0x26, 0xec, 0x1d, 0x33, 0x99, 0x16, 0x38,
	0x56, 0x9d, 0x47, 0x2b, 0x27, 0xfb, 0x1b, 0xc1, 0x7e, 0x11, 0x90, 0x51, 0xc0, 0xaf, 0xe0, 0x56,
	0x14, 0xa7, 0x49, 0x4b, 0xaa, 0x15, 0xb3, 0x5e, 0x71, 0x13, 0x0a, 0x49, 0x08, 0xf8, 0x25, 0xd4,
	0x43, 0x46, 0xc7, 0x17, 0x8e, 0xe8, 0x32, 0x57, 0x04, 0x2c, 0x39, 0xae, 0x46, 0xad, 0x85, 0xda,
	0x77, 0x48, 0xd9, 0x16, 0x7e, 0x02, 0x3b, 0x5c, 0x48, 0x12, 0xef, 0x10, 0x46, 0x23, 0xc1, 0x1b,
	0xeb, 0x2d, 0xd4, 0xde, 0x24, 0xb9, 0xd5, 0x83, 0xb7, 0x50, 0x2f, 0xf1, 0xc5, 0x0a, 0xdc, 0xb5,
	0x86, 0x96, 0x63, 0x19, 0xef, 0xad, 0x33, 0x6b, 0x78, 0xac, 0xac, 0x61, 0x80, 0xdb, 0x46, 0xcf,
	0xb1, 0xbe, 0x98, 0x0a, 0xc2, 0x5b, 0xb0, 0x61, 0x3b, 0xc6, 0xf0, 0xa8, 0x7b, 0xaa, 0xd4, 0x0e,
	0xbe, 0xc2, 0x6e, 0xae, 0x19, 0x71, 0x1d, 0x76, 0x89, 0xf9, 0xf9, 0xc4, 0xb4, 0x9d, 0x51, 0xf7,
	0x74, 0x74, 0x62, 0x9b, 0x44, 0x59, 0xc3, 0x2a, 0xdc, 0xcb, 0x2d, 0x8e, 0xfa, 0x1f, 0x49, 0xcf,
	0x3c, 0x52, 0x50, 0x8e, 0x70, 0xf6, 0xae, 0xdf, 0x53, 0x6a, 0x9d, 0x3f, 0xeb, 0xd0, 0x28, 0x4c,
	0x54, 0xfa, 0x89, 0xc7, 0xb0, 0x1d, 0x64, 0xfb, 0x18, 0xb7, 0x73, 0xa7, 0xf9, 0xcf, 0x11, 0x50,
	0x9f, 0x56, 0x23, 0xb3, 0x15, 0x8b, 0x00, 0xcb, 0x42, 0x97, 0xe2, 0xe7, 0x39, 0x81, 0xea, 0x91,
	0x51, 0xf5, 0x6b, 0xc0, 0xb3, 0xa6, 0x0b, 0xa8, 0xcb, 0x62, 0x33, 0xe2, 0x2a, 0x99, 0x92, 0xd9,
	0x50, 0x5f, 0x5c, 0x07, 0x9f, 0xf5, 0x0d, 0x40, 0x99, 0xe4, 0xfa, 0x17, 0x3f, 0xcb, 0x89, 0x54,
	0x4d, 0x80, 0x7a, 0xf8, 0x5f, 0x70, 0xc6, 0xae, 0xfb, 0x1a, 0x54, 0x11, 0x4e, 0x74, 0x3a, 0xa5,
	0xae, 0xc7, 0x2e, 0x99, 0x1e, 0x4d, 0xae, 0xcc, 0x6e, 0xf1, 0x2e, 0x5d, 0x3e, 0xa3, 0x9f, 0x08,
	0xfd, 0x42, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd8, 0x57, 0xbf, 0x6b, 0x05, 0x00,
	0x00,
}