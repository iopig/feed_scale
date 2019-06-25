// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global_def.proto

package fsapi

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

type ErrCode int32

const (
	ErrCode_ERR_SUCCESS              ErrCode = 0
	ErrCode_ERR_FAILED               ErrCode = 1
	ErrCode_ERR_DEVID_NOT_EXIST      ErrCode = 2
	ErrCode_ERR_CONN_TO_DA_NOT_READY ErrCode = 3
)

var ErrCode_name = map[int32]string{
	0: "ERR_SUCCESS",
	1: "ERR_FAILED",
	2: "ERR_DEVID_NOT_EXIST",
	3: "ERR_CONN_TO_DA_NOT_READY",
}

var ErrCode_value = map[string]int32{
	"ERR_SUCCESS":              0,
	"ERR_FAILED":               1,
	"ERR_DEVID_NOT_EXIST":      2,
	"ERR_CONN_TO_DA_NOT_READY": 3,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{0}
}

type FeedType int32

const (
	//上料信息
	FeedType_LOAD FeedType = 0
	//选择猪栏
	FeedType_CHOOSE_PIGSTY FeedType = 1
	//喂料信息
	FeedType_WEIGHT_INFO FeedType = 2
)

var FeedType_name = map[int32]string{
	0: "LOAD",
	1: "CHOOSE_PIGSTY",
	2: "WEIGHT_INFO",
}

var FeedType_value = map[string]int32{
	"LOAD":          0,
	"CHOOSE_PIGSTY": 1,
	"WEIGHT_INFO":   2,
}

func (x FeedType) String() string {
	return proto.EnumName(FeedType_name, int32(x))
}

func (FeedType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{1}
}

type ReqHeader struct {
	// 现在版本默认为 1
	Version int32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	// 设备ID
	DevId string `protobuf:"bytes,2,opt,name=DevId,proto3" json:"DevId,omitempty"`
	//时间戳 unix_timestamp
	Ts                   uint64   `protobuf:"varint,3,opt,name=Ts,proto3" json:"Ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{0}
}

func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (m *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(m, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReqHeader) GetDevId() string {
	if m != nil {
		return m.DevId
	}
	return ""
}

func (m *ReqHeader) GetTs() uint64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type ResHeader struct {
	// 现在版本默认为 1
	Version              int32    `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	ErrCode              ErrCode  `protobuf:"varint,2,opt,name=ErrCode,proto3,enum=fsapi.ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,3,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	ReqId                string   `protobuf:"bytes,4,opt,name=ReqId,proto3" json:"ReqId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResHeader) Reset()         { *m = ResHeader{} }
func (m *ResHeader) String() string { return proto.CompactTextString(m) }
func (*ResHeader) ProtoMessage()    {}
func (*ResHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{1}
}

func (m *ResHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResHeader.Unmarshal(m, b)
}
func (m *ResHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResHeader.Marshal(b, m, deterministic)
}
func (m *ResHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResHeader.Merge(m, src)
}
func (m *ResHeader) XXX_Size() int {
	return xxx_messageInfo_ResHeader.Size(m)
}
func (m *ResHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResHeader proto.InternalMessageInfo

func (m *ResHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ResHeader) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_ERR_SUCCESS
}

func (m *ResHeader) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ResHeader) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func init() {
	proto.RegisterEnum("fsapi.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterEnum("fsapi.FeedType", FeedType_name, FeedType_value)
	proto.RegisterType((*ReqHeader)(nil), "fsapi.ReqHeader")
	proto.RegisterType((*ResHeader)(nil), "fsapi.ResHeader")
}

func init() { proto.RegisterFile("global_def.proto", fileDescriptor_24a8c9fe6e60032d) }

var fileDescriptor_24a8c9fe6e60032d = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6f, 0xd3, 0x30,
	0x18, 0xc5, 0x97, 0xb4, 0xdd, 0xd6, 0x0f, 0x51, 0x8c, 0x41, 0x90, 0x03, 0x87, 0x6a, 0xa7, 0x6a,
	0x12, 0x8e, 0x04, 0x07, 0xb8, 0x96, 0xc4, 0x5d, 0x2d, 0x4a, 0x33, 0x39, 0x66, 0x63, 0x70, 0xb0,
	0xd2, 0xe4, 0x6b, 0x88, 0x54, 0xea, 0xcc, 0xc9, 0x90, 0x38, 0xf0, 0xcf, 0xf0, 0x97, 0x22, 0x7b,
	0x2b, 0x57, 0x8e, 0xbf, 0xf7, 0x9e, 0x9e, 0x9f, 0xfc, 0x01, 0xa9, 0x77, 0x66, 0x53, 0xec, 0x74,
	0x85, 0x5b, 0xd6, 0x5a, 0xd3, 0x1b, 0x3a, 0xda, 0x76, 0x45, 0xdb, 0x9c, 0x7d, 0x84, 0xb1, 0xc4,
	0xdb, 0x25, 0x16, 0x15, 0x5a, 0x1a, 0xc1, 0xc9, 0x15, 0xda, 0xae, 0x31, 0xfb, 0x28, 0x98, 0x06,
	0xb3, 0x91, 0x3c, 0x20, 0x7d, 0x0e, 0xa3, 0x14, 0x7f, 0x8a, 0x2a, 0x0a, 0xa7, 0xc1, 0x6c, 0x2c,
	0xef, 0x81, 0x4e, 0x20, 0x54, 0x5d, 0x34, 0x98, 0x06, 0xb3, 0xa1, 0x0c, 0x55, 0x77, 0xf6, 0xdb,
	0x95, 0x75, 0xff, 0x2d, 0x9b, 0xc1, 0x09, 0xb7, 0x36, 0x31, 0x15, 0xfa, 0xba, 0xc9, 0x9b, 0x09,
	0xf3, 0x63, 0xd8, 0x83, 0x2a, 0x0f, 0x36, 0x7d, 0x01, 0xc7, 0xdc, 0xda, 0x4f, 0x5d, 0xed, 0x1f,
	0x19, 0xcb, 0x07, 0x72, 0x73, 0x24, 0xde, 0x8a, 0x2a, 0x1a, 0xde, 0xcf, 0xf1, 0x70, 0x5e, 0xfc,
	0xeb, 0xa5, 0x4f, 0xe0, 0x11, 0x97, 0x52, 0xe7, 0x9f, 0x93, 0x84, 0xe7, 0x39, 0x39, 0xa2, 0x13,
	0x00, 0x27, 0x2c, 0xe6, 0x62, 0xc5, 0x53, 0x12, 0xd0, 0x97, 0xf0, 0xcc, 0x71, 0xca, 0xaf, 0x44,
	0xaa, 0xd7, 0x99, 0xd2, 0xfc, 0x8b, 0xc8, 0x15, 0x09, 0xe9, 0x2b, 0x88, 0x9c, 0x91, 0x64, 0xeb,
	0xb5, 0x56, 0x99, 0x4e, 0xe7, 0xde, 0x95, 0x7c, 0x9e, 0xde, 0x90, 0xc1, 0xf9, 0x7b, 0x38, 0x5d,
	0x20, 0x56, 0xea, 0x57, 0x8b, 0xf4, 0x14, 0x86, 0xab, 0x6c, 0x9e, 0x92, 0x23, 0xfa, 0x14, 0x1e,
	0x27, 0xcb, 0x2c, 0xcb, 0xb9, 0xbe, 0x14, 0x17, 0xb9, 0xba, 0x21, 0x81, 0x1b, 0x70, 0xcd, 0xc5,
	0xc5, 0x52, 0x69, 0xb1, 0x5e, 0x64, 0x24, 0xfc, 0xf0, 0x0d, 0xa2, 0x72, 0xcf, 0x1a, 0xd3, 0x36,
	0x35, 0xdb, 0x22, 0x56, 0xac, 0x2b, 0x8b, 0x1d, 0xb2, 0xda, 0xb6, 0xe5, 0x65, 0xf0, 0xf5, 0x5d,
	0xdd, 0xf4, 0xdf, 0xef, 0x36, 0xac, 0x34, 0x3f, 0x62, 0x9f, 0x89, 0x5d, 0xe6, 0xb5, 0xcf, 0xc4,
	0xcd, 0xbe, 0x47, 0xbb, 0x2d, 0x4a, 0x8c, 0x5d, 0x3a, 0xae, 0x8d, 0x36, 0x77, 0x7d, 0xec, 0xbf,
	0xec, 0x4f, 0x38, 0x58, 0xae, 0xae, 0x37, 0xc7, 0xfe, 0xa6, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x00, 0xf6, 0x16, 0xe1, 0xe7, 0x01, 0x00, 0x00,
}
