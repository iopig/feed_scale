// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fs_gw.proto

package fsapi

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

type PigstyInfo struct {
	//猪圈ID
	PigstyId string `protobuf:"bytes,1,opt,name=PigstyId,proto3" json:"PigstyId,omitempty"`
	//猪数量
	PigNum string `protobuf:"bytes,2,opt,name=PigNum,proto3" json:"PigNum,omitempty"`
	//猪平均重量
	AverageWeight string `protobuf:"bytes,3,opt,name=AverageWeight,proto3" json:"AverageWeight,omitempty"`
	//推荐值
	AdviseWeight uint32 `protobuf:"varint,4,opt,name=AdviseWeight,proto3" json:"AdviseWeight,omitempty"`
	//上次喂料时间戳 unix_timestamp
	LastFed uint32 `protobuf:"varint,5,opt,name=LastFed,proto3" json:"LastFed,omitempty"`
	//猪的id列表
	PigId                []string `protobuf:"bytes,6,rep,name=PigId,proto3" json:"PigId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PigstyInfo) Reset()         { *m = PigstyInfo{} }
func (m *PigstyInfo) String() string { return proto.CompactTextString(m) }
func (*PigstyInfo) ProtoMessage()    {}
func (*PigstyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{0}
}

func (m *PigstyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigstyInfo.Unmarshal(m, b)
}
func (m *PigstyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigstyInfo.Marshal(b, m, deterministic)
}
func (m *PigstyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigstyInfo.Merge(m, src)
}
func (m *PigstyInfo) XXX_Size() int {
	return xxx_messageInfo_PigstyInfo.Size(m)
}
func (m *PigstyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PigstyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PigstyInfo proto.InternalMessageInfo

func (m *PigstyInfo) GetPigstyId() string {
	if m != nil {
		return m.PigstyId
	}
	return ""
}

func (m *PigstyInfo) GetPigNum() string {
	if m != nil {
		return m.PigNum
	}
	return ""
}

func (m *PigstyInfo) GetAverageWeight() string {
	if m != nil {
		return m.AverageWeight
	}
	return ""
}

func (m *PigstyInfo) GetAdviseWeight() uint32 {
	if m != nil {
		return m.AdviseWeight
	}
	return 0
}

func (m *PigstyInfo) GetLastFed() uint32 {
	if m != nil {
		return m.LastFed
	}
	return 0
}

func (m *PigstyInfo) GetPigId() []string {
	if m != nil {
		return m.PigId
	}
	return nil
}

type PigHouseInfo struct {
	//猪舍ID
	HouseId string `protobuf:"bytes,1,opt,name=HouseId,proto3" json:"HouseId,omitempty"`
	//猪圈信息
	PigstyInfo           []*PigstyInfo `protobuf:"bytes,2,rep,name=pigstyInfo,proto3" json:"pigstyInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PigHouseInfo) Reset()         { *m = PigHouseInfo{} }
func (m *PigHouseInfo) String() string { return proto.CompactTextString(m) }
func (*PigHouseInfo) ProtoMessage()    {}
func (*PigHouseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{1}
}

func (m *PigHouseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigHouseInfo.Unmarshal(m, b)
}
func (m *PigHouseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigHouseInfo.Marshal(b, m, deterministic)
}
func (m *PigHouseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigHouseInfo.Merge(m, src)
}
func (m *PigHouseInfo) XXX_Size() int {
	return xxx_messageInfo_PigHouseInfo.Size(m)
}
func (m *PigHouseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PigHouseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PigHouseInfo proto.InternalMessageInfo

func (m *PigHouseInfo) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func (m *PigHouseInfo) GetPigstyInfo() []*PigstyInfo {
	if m != nil {
		return m.PigstyInfo
	}
	return nil
}

type DevInfoReq struct {
	//Version ,Devid ,Timestamp
	ReqHeader            *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DevInfoReq) Reset()         { *m = DevInfoReq{} }
func (m *DevInfoReq) String() string { return proto.CompactTextString(m) }
func (*DevInfoReq) ProtoMessage()    {}
func (*DevInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{2}
}

func (m *DevInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevInfoReq.Unmarshal(m, b)
}
func (m *DevInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevInfoReq.Marshal(b, m, deterministic)
}
func (m *DevInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevInfoReq.Merge(m, src)
}
func (m *DevInfoReq) XXX_Size() int {
	return xxx_messageInfo_DevInfoReq.Size(m)
}
func (m *DevInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DevInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DevInfoReq proto.InternalMessageInfo

func (m *DevInfoReq) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

//
type PigstyInfoRes struct {
	//Version ,Devid ,Timestamp
	ReqHeader *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	//猪场ID
	PigFarmId string `protobuf:"bytes,2,opt,name=PigFarmId,proto3" json:"PigFarmId,omitempty"`
	//猪场主人
	Farmer               string          `protobuf:"bytes,3,opt,name=farmer,proto3" json:"farmer,omitempty"`
	PigHouseInfo         []*PigHouseInfo `protobuf:"bytes,4,rep,name=pigHouseInfo,proto3" json:"pigHouseInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PigstyInfoRes) Reset()         { *m = PigstyInfoRes{} }
func (m *PigstyInfoRes) String() string { return proto.CompactTextString(m) }
func (*PigstyInfoRes) ProtoMessage()    {}
func (*PigstyInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{3}
}

func (m *PigstyInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigstyInfoRes.Unmarshal(m, b)
}
func (m *PigstyInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigstyInfoRes.Marshal(b, m, deterministic)
}
func (m *PigstyInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigstyInfoRes.Merge(m, src)
}
func (m *PigstyInfoRes) XXX_Size() int {
	return xxx_messageInfo_PigstyInfoRes.Size(m)
}
func (m *PigstyInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PigstyInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_PigstyInfoRes proto.InternalMessageInfo

func (m *PigstyInfoRes) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

func (m *PigstyInfoRes) GetPigFarmId() string {
	if m != nil {
		return m.PigFarmId
	}
	return ""
}

func (m *PigstyInfoRes) GetFarmer() string {
	if m != nil {
		return m.Farmer
	}
	return ""
}

func (m *PigstyInfoRes) GetPigHouseInfo() []*PigHouseInfo {
	if m != nil {
		return m.PigHouseInfo
	}
	return nil
}

//LoadReq
type LoadReq struct {
	//Version ,Devid ,Timestamp
	ReqHeader *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	//当前重量
	CurrentWeight        string   `protobuf:"bytes,2,opt,name=CurrentWeight,proto3" json:"CurrentWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadReq) Reset()         { *m = LoadReq{} }
func (m *LoadReq) String() string { return proto.CompactTextString(m) }
func (*LoadReq) ProtoMessage()    {}
func (*LoadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{4}
}

func (m *LoadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadReq.Unmarshal(m, b)
}
func (m *LoadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadReq.Marshal(b, m, deterministic)
}
func (m *LoadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadReq.Merge(m, src)
}
func (m *LoadReq) XXX_Size() int {
	return xxx_messageInfo_LoadReq.Size(m)
}
func (m *LoadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoadReq proto.InternalMessageInfo

func (m *LoadReq) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

func (m *LoadReq) GetCurrentWeight() string {
	if m != nil {
		return m.CurrentWeight
	}
	return ""
}

//ChoosePigstyReq
//称在一秒内6次上传当前重量,重量选取算法：
//如果一秒内的重量变化（小于0.5kg），那么上传当前重量。
//如果下一秒的数据无变化，那么不上传数据。
type ChoosePigstyReq struct {
	//Version ,Devid ,Timestamp
	ReqHeader *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	//当前猪圈
	PigstyId string `protobuf:"bytes,2,opt,name=PigstyId,proto3" json:"PigstyId,omitempty"`
	//当前重量
	CurrentWeight        string   `protobuf:"bytes,3,opt,name=CurrentWeight,proto3" json:"CurrentWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChoosePigstyReq) Reset()         { *m = ChoosePigstyReq{} }
func (m *ChoosePigstyReq) String() string { return proto.CompactTextString(m) }
func (*ChoosePigstyReq) ProtoMessage()    {}
func (*ChoosePigstyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{5}
}

func (m *ChoosePigstyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChoosePigstyReq.Unmarshal(m, b)
}
func (m *ChoosePigstyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChoosePigstyReq.Marshal(b, m, deterministic)
}
func (m *ChoosePigstyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChoosePigstyReq.Merge(m, src)
}
func (m *ChoosePigstyReq) XXX_Size() int {
	return xxx_messageInfo_ChoosePigstyReq.Size(m)
}
func (m *ChoosePigstyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChoosePigstyReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChoosePigstyReq proto.InternalMessageInfo

func (m *ChoosePigstyReq) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

func (m *ChoosePigstyReq) GetPigstyId() string {
	if m != nil {
		return m.PigstyId
	}
	return ""
}

func (m *ChoosePigstyReq) GetCurrentWeight() string {
	if m != nil {
		return m.CurrentWeight
	}
	return ""
}

//CurrentFedRes
type CurrentFedRes struct {
	//error number ,error detail
	ResHeader *ResHeader `protobuf:"bytes,1,opt,name=ResHeader,proto3" json:"ResHeader,omitempty"`
	//当前猪圈
	PigstyId string `protobuf:"bytes,2,opt,name=PigstyId,proto3" json:"PigstyId,omitempty"`
	//已经喂料重量
	FedWeight            string   `protobuf:"bytes,3,opt,name=FedWeight,proto3" json:"FedWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentFedRes) Reset()         { *m = CurrentFedRes{} }
func (m *CurrentFedRes) String() string { return proto.CompactTextString(m) }
func (*CurrentFedRes) ProtoMessage()    {}
func (*CurrentFedRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{6}
}

func (m *CurrentFedRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentFedRes.Unmarshal(m, b)
}
func (m *CurrentFedRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentFedRes.Marshal(b, m, deterministic)
}
func (m *CurrentFedRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentFedRes.Merge(m, src)
}
func (m *CurrentFedRes) XXX_Size() int {
	return xxx_messageInfo_CurrentFedRes.Size(m)
}
func (m *CurrentFedRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentFedRes.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentFedRes proto.InternalMessageInfo

func (m *CurrentFedRes) GetResHeader() *ResHeader {
	if m != nil {
		return m.ResHeader
	}
	return nil
}

func (m *CurrentFedRes) GetPigstyId() string {
	if m != nil {
		return m.PigstyId
	}
	return ""
}

func (m *CurrentFedRes) GetFedWeight() string {
	if m != nil {
		return m.FedWeight
	}
	return ""
}

func init() {
	proto.RegisterType((*PigstyInfo)(nil), "fsapi.PigstyInfo")
	proto.RegisterType((*PigHouseInfo)(nil), "fsapi.PigHouseInfo")
	proto.RegisterType((*DevInfoReq)(nil), "fsapi.DevInfoReq")
	proto.RegisterType((*PigstyInfoRes)(nil), "fsapi.PigstyInfoRes")
	proto.RegisterType((*LoadReq)(nil), "fsapi.LoadReq")
	proto.RegisterType((*ChoosePigstyReq)(nil), "fsapi.ChoosePigstyReq")
	proto.RegisterType((*CurrentFedRes)(nil), "fsapi.CurrentFedRes")
}

func init() { proto.RegisterFile("fs_gw.proto", fileDescriptor_d1a5b6d85b3d9f79) }

var fileDescriptor_d1a5b6d85b3d9f79 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x8e, 0x21, 0x40, 0x18, 0xa0, 0x3f, 0xdb, 0x28, 0xb2, 0x50, 0x0e, 0xc8, 0xea, 0x81, 0x4b,
	0x6c, 0x95, 0xaa, 0xca, 0xa5, 0x6a, 0x45, 0xa9, 0x50, 0x90, 0x50, 0x65, 0xb9, 0xaa, 0x2a, 0xb5,
	0x07, 0xb4, 0xb0, 0xe3, 0x65, 0x25, 0x60, 0x1d, 0xaf, 0x21, 0xca, 0xb9, 0xef, 0xd3, 0x07, 0xe8,
	0x43, 0xf5, 0x19, 0x2a, 0xd6, 0x8b, 0x8d, 0x43, 0x2b, 0x95, 0x1c, 0xbf, 0xd9, 0xcf, 0x33, 0xdf,
	0x7c, 0x33, 0x1e, 0x68, 0x84, 0x6a, 0xc2, 0xef, 0xdc, 0x28, 0x96, 0x89, 0x24, 0x95, 0x50, 0xd1,
	0x48, 0xb4, 0xfb, 0x5c, 0x24, 0xf3, 0xf5, 0xd4, 0x9d, 0xc9, 0xa5, 0x27, 0x64, 0x24, 0xb8, 0x17,
	0x22, 0xb2, 0x2b, 0x35, 0xa3, 0x0b, 0xf4, 0xc4, 0x2a, 0xc1, 0x38, 0xa4, 0x33, 0xf4, 0x78, 0x1c,
	0xcd, 0x3c, 0xfd, 0xa1, 0xc7, 0x17, 0x72, 0x4a, 0x17, 0x13, 0x86, 0x61, 0x9a, 0xc9, 0xf9, 0x65,
	0x01, 0xf8, 0x82, 0xab, 0xe4, 0x7e, 0xb4, 0x0a, 0x25, 0x69, 0xc3, 0x99, 0x41, 0xcc, 0xb6, 0x3a,
	0x56, 0xb7, 0x1e, 0x64, 0x98, 0x5c, 0x40, 0xd5, 0x17, 0xfc, 0xd3, 0x7a, 0x69, 0x97, 0xf4, 0x8b,
	0x41, 0xe4, 0x25, 0xb4, 0xfa, 0x1b, 0x8c, 0x29, 0xc7, 0xaf, 0x28, 0xf8, 0x3c, 0xb1, 0xcb, 0xfa,
	0xb9, 0x18, 0x24, 0x0e, 0x34, 0xfb, 0x6c, 0x23, 0xd4, 0x8e, 0x74, 0xda, 0xb1, 0xba, 0xad, 0xa0,
	0x10, 0x23, 0x36, 0xd4, 0xc6, 0x54, 0x25, 0x43, 0x64, 0x76, 0x45, 0x3f, 0xef, 0x20, 0x39, 0x87,
	0x8a, 0x2f, 0xf8, 0x88, 0xd9, 0xd5, 0x4e, 0xb9, 0x5b, 0x0f, 0x52, 0xe0, 0x7c, 0x87, 0xa6, 0x2f,
	0xf8, 0x8d, 0x5c, 0x2b, 0xd4, 0xea, 0x6d, 0xa8, 0xa5, 0x60, 0x27, 0x7e, 0x07, 0xc9, 0x2b, 0x80,
	0x28, 0xeb, 0xd2, 0x2e, 0x75, 0xca, 0xdd, 0x46, 0xef, 0xb9, 0xab, 0x5d, 0x74, 0xf3, 0xf6, 0x83,
	0x3d, 0x92, 0xf3, 0x16, 0xe0, 0x23, 0x6e, 0x74, 0x18, 0x6f, 0x89, 0x0b, 0xf5, 0x00, 0x6f, 0x6f,
	0x90, 0x32, 0x8c, 0x75, 0xf2, 0x46, 0xef, 0x99, 0xf9, 0x3e, 0x8b, 0x07, 0x39, 0xc5, 0xf9, 0x69,
	0x41, 0x6b, 0x2f, 0x31, 0xaa, 0x63, 0x33, 0x90, 0x4b, 0xa8, 0xfb, 0x82, 0x0f, 0x69, 0xbc, 0x1c,
	0x31, 0xe3, 0x78, 0x1e, 0xd8, 0x0e, 0x23, 0xa4, 0xf1, 0x12, 0x63, 0xe3, 0xb6, 0x41, 0xe4, 0x1a,
	0x9a, 0xd1, 0x9e, 0x25, 0xf6, 0xa9, 0x6e, 0xf5, 0x45, 0xde, 0x6a, 0xf6, 0x14, 0x14, 0x88, 0xce,
	0x04, 0x6a, 0x63, 0x49, 0xd9, 0x23, 0x7a, 0xdd, 0x2e, 0xc0, 0x60, 0x1d, 0xc7, 0xb8, 0x4a, 0xcc,
	0x6c, 0x53, 0xb5, 0xc5, 0xa0, 0xf3, 0xc3, 0x82, 0xa7, 0x83, 0xb9, 0x94, 0x0a, 0x53, 0x5f, 0x1e,
	0x53, 0x69, 0x7f, 0x3d, 0x4b, 0x0f, 0xd6, 0xf3, 0x40, 0x45, 0xf9, 0x6f, 0x2a, 0xee, 0x33, 0xd6,
	0x10, 0x59, 0x36, 0x16, 0xf5, 0x0f, 0x09, 0x2a, 0x97, 0xa0, 0xfe, 0x43, 0xc2, 0x25, 0xd4, 0x87,
	0xc8, 0x0a, 0xe5, 0xf3, 0x40, 0xef, 0xb7, 0x05, 0x95, 0xa1, 0xf2, 0x29, 0x23, 0x6f, 0xe0, 0xcc,
	0xa7, 0x6c, 0x2c, 0xb9, 0x58, 0x91, 0xdd, 0x16, 0xe6, 0xbb, 0xd6, 0x3e, 0x3f, 0x5c, 0x4c, 0x54,
	0xce, 0x09, 0xb9, 0x4a, 0x47, 0x34, 0x58, 0x32, 0xf2, 0xc4, 0x50, 0xcc, 0xc8, 0xda, 0x07, 0x92,
	0x9d, 0x13, 0xf2, 0x0e, 0x9a, 0xfb, 0x7e, 0x93, 0x0b, 0xc3, 0x79, 0x30, 0x84, 0xac, 0x5c, 0xc1,
	0x17, 0xe7, 0x84, 0xbc, 0x87, 0xd6, 0x97, 0x68, 0xb1, 0x2d, 0x40, 0xef, 0xf4, 0xef, 0x75, 0x64,
	0x82, 0x0f, 0x9f, 0xc1, 0x9e, 0xad, 0x5c, 0x7d, 0x98, 0xdc, 0xed, 0x61, 0x72, 0xf5, 0x61, 0x72,
	0xb7, 0xe7, 0xe8, 0xdb, 0xf5, 0x11, 0xa7, 0x8b, 0xcb, 0x89, 0x5c, 0x27, 0x9e, 0x2e, 0x32, 0xad,
	0xea, 0xbb, 0xf5, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x99, 0x81, 0x67, 0x10, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FsPadClient is the client API for FsPad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FsPadClient interface {
	//PAD 登录
	PadLogin(ctx context.Context, in *DevInfoReq, opts ...grpc.CallOption) (*PigstyInfoRes, error)
	//上料命令
	LoadCmd(ctx context.Context, in *LoadReq, opts ...grpc.CallOption) (*ResHeader, error)
	//选择猪圈
	ChoosePigsty(ctx context.Context, in *ChoosePigstyReq, opts ...grpc.CallOption) (*CurrentFedRes, error)
	//上传当前重量
	UploadRawInfo(ctx context.Context, in *ChoosePigstyReq, opts ...grpc.CallOption) (*CurrentFedRes, error)
}

type fsPadClient struct {
	cc *grpc.ClientConn
}

func NewFsPadClient(cc *grpc.ClientConn) FsPadClient {
	return &fsPadClient{cc}
}

func (c *fsPadClient) PadLogin(ctx context.Context, in *DevInfoReq, opts ...grpc.CallOption) (*PigstyInfoRes, error) {
	out := new(PigstyInfoRes)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/PadLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsPadClient) LoadCmd(ctx context.Context, in *LoadReq, opts ...grpc.CallOption) (*ResHeader, error) {
	out := new(ResHeader)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/LoadCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsPadClient) ChoosePigsty(ctx context.Context, in *ChoosePigstyReq, opts ...grpc.CallOption) (*CurrentFedRes, error) {
	out := new(CurrentFedRes)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/ChoosePigsty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsPadClient) UploadRawInfo(ctx context.Context, in *ChoosePigstyReq, opts ...grpc.CallOption) (*CurrentFedRes, error) {
	out := new(CurrentFedRes)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/UploadRawInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FsPadServer is the server API for FsPad service.
type FsPadServer interface {
	//PAD 登录
	PadLogin(context.Context, *DevInfoReq) (*PigstyInfoRes, error)
	//上料命令
	LoadCmd(context.Context, *LoadReq) (*ResHeader, error)
	//选择猪圈
	ChoosePigsty(context.Context, *ChoosePigstyReq) (*CurrentFedRes, error)
	//上传当前重量
	UploadRawInfo(context.Context, *ChoosePigstyReq) (*CurrentFedRes, error)
}

// UnimplementedFsPadServer can be embedded to have forward compatible implementations.
type UnimplementedFsPadServer struct {
}

func (*UnimplementedFsPadServer) PadLogin(ctx context.Context, req *DevInfoReq) (*PigstyInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PadLogin not implemented")
}
func (*UnimplementedFsPadServer) LoadCmd(ctx context.Context, req *LoadReq) (*ResHeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCmd not implemented")
}
func (*UnimplementedFsPadServer) ChoosePigsty(ctx context.Context, req *ChoosePigstyReq) (*CurrentFedRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChoosePigsty not implemented")
}
func (*UnimplementedFsPadServer) UploadRawInfo(ctx context.Context, req *ChoosePigstyReq) (*CurrentFedRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadRawInfo not implemented")
}

func RegisterFsPadServer(s *grpc.Server, srv FsPadServer) {
	s.RegisterService(&_FsPad_serviceDesc, srv)
}

func _FsPad_PadLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).PadLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/PadLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).PadLogin(ctx, req.(*DevInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsPad_LoadCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).LoadCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/LoadCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).LoadCmd(ctx, req.(*LoadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsPad_ChoosePigsty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChoosePigstyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).ChoosePigsty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/ChoosePigsty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).ChoosePigsty(ctx, req.(*ChoosePigstyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsPad_UploadRawInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChoosePigstyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).UploadRawInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/UploadRawInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).UploadRawInfo(ctx, req.(*ChoosePigstyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FsPad_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fsapi.FsPad",
	HandlerType: (*FsPadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PadLogin",
			Handler:    _FsPad_PadLogin_Handler,
		},
		{
			MethodName: "LoadCmd",
			Handler:    _FsPad_LoadCmd_Handler,
		},
		{
			MethodName: "ChoosePigsty",
			Handler:    _FsPad_ChoosePigsty_Handler,
		},
		{
			MethodName: "UploadRawInfo",
			Handler:    _FsPad_UploadRawInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fs_gw.proto",
}