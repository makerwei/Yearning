// Code generated by protoc-gen-go. DO NOT EDIT.
// source: detection.proto

package proto

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

type LibraAuditOrder struct {
	SQL                  string   `protobuf:"bytes,1,opt,name=SQL,proto3" json:"SQL"`
	DataBase             string   `protobuf:"bytes,2,opt,name=DataBase,proto3" json:"DataBase"`
	Table                string   `protobuf:"bytes,3,opt,name=Table,proto3" json:"Table"`
	Execute              bool     `protobuf:"varint,4,opt,name=Execute,proto3" json:"Execute"`
	Check                bool     `protobuf:"varint,5,opt,name=Check,proto3" json:"Check"`
	IsDML                bool     `protobuf:"varint,6,opt,name=IsDML,proto3" json:"IsDML"`
	Backup               bool     `protobuf:"varint,7,opt,name=Backup,proto3" json:"Backup"`
	Source               *Source  `protobuf:"bytes,8,opt,name=source,proto3" json:"source"`
	WorkId               string   `protobuf:"bytes,9,opt,name=WorkId,proto3" json:"WorkId"`
	IsAutoTask           bool     `protobuf:"varint,10,opt,name=IsAutoTask,proto3" json:"IsAutoTask"`
	Name                 string   `protobuf:"bytes,11,opt,name=Name,proto3" json:"Name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LibraAuditOrder) Reset()         { *m = LibraAuditOrder{} }
func (m *LibraAuditOrder) String() string { return proto.CompactTextString(m) }
func (*LibraAuditOrder) ProtoMessage()    {}
func (*LibraAuditOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{0}
}

func (m *LibraAuditOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LibraAuditOrder.Unmarshal(m, b)
}
func (m *LibraAuditOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LibraAuditOrder.Marshal(b, m, deterministic)
}
func (m *LibraAuditOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LibraAuditOrder.Merge(m, src)
}
func (m *LibraAuditOrder) XXX_Size() int {
	return xxx_messageInfo_LibraAuditOrder.Size(m)
}
func (m *LibraAuditOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_LibraAuditOrder.DiscardUnknown(m)
}

var xxx_messageInfo_LibraAuditOrder proto.InternalMessageInfo

func (m *LibraAuditOrder) GetSQL() string {
	if m != nil {
		return m.SQL
	}
	return ""
}

func (m *LibraAuditOrder) GetDataBase() string {
	if m != nil {
		return m.DataBase
	}
	return ""
}

func (m *LibraAuditOrder) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *LibraAuditOrder) GetExecute() bool {
	if m != nil {
		return m.Execute
	}
	return false
}

func (m *LibraAuditOrder) GetCheck() bool {
	if m != nil {
		return m.Check
	}
	return false
}

func (m *LibraAuditOrder) GetIsDML() bool {
	if m != nil {
		return m.IsDML
	}
	return false
}

func (m *LibraAuditOrder) GetBackup() bool {
	if m != nil {
		return m.Backup
	}
	return false
}

func (m *LibraAuditOrder) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *LibraAuditOrder) GetWorkId() string {
	if m != nil {
		return m.WorkId
	}
	return ""
}

func (m *LibraAuditOrder) GetIsAutoTask() bool {
	if m != nil {
		return m.IsAutoTask
	}
	return false
}

func (m *LibraAuditOrder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Source struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr"`
	User                 string   `protobuf:"bytes,2,opt,name=User,proto3" json:"User"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password"`
	Port                 int32    `protobuf:"varint,4,opt,name=Port,proto3" json:"Port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{1}
}

func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Source) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Source) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Source) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Record struct {
	SQL                  string   `protobuf:"bytes,1,opt,name=SQL,proto3" json:"SQL"`
	AffectRows           int32    `protobuf:"varint,2,opt,name=AffectRows,proto3" json:"AffectRows"`
	Status               string   `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status"`
	Error                string   `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error"`
	Level                int32    `protobuf:"varint,6,opt,name=Level,proto3" json:"Level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{2}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetSQL() string {
	if m != nil {
		return m.SQL
	}
	return ""
}

func (m *Record) GetAffectRows() int32 {
	if m != nil {
		return m.AffectRows
	}
	return 0
}

func (m *Record) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Record) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Record) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type RecordSet struct {
	Record               []*Record `protobuf:"bytes,1,rep,name=record,proto3" json:"record"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RecordSet) Reset()         { *m = RecordSet{} }
func (m *RecordSet) String() string { return proto.CompactTextString(m) }
func (*RecordSet) ProtoMessage()    {}
func (*RecordSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{3}
}

func (m *RecordSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordSet.Unmarshal(m, b)
}
func (m *RecordSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordSet.Marshal(b, m, deterministic)
}
func (m *RecordSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordSet.Merge(m, src)
}
func (m *RecordSet) XXX_Size() int {
	return xxx_messageInfo_RecordSet.Size(m)
}
func (m *RecordSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordSet.DiscardUnknown(m)
}

var xxx_messageInfo_RecordSet proto.InternalMessageInfo

func (m *RecordSet) GetRecord() []*Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ExecOrder struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecOrder) Reset()         { *m = ExecOrder{} }
func (m *ExecOrder) String() string { return proto.CompactTextString(m) }
func (*ExecOrder) ProtoMessage()    {}
func (*ExecOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{4}
}

func (m *ExecOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecOrder.Unmarshal(m, b)
}
func (m *ExecOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecOrder.Marshal(b, m, deterministic)
}
func (m *ExecOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecOrder.Merge(m, src)
}
func (m *ExecOrder) XXX_Size() int {
	return xxx_messageInfo_ExecOrder.Size(m)
}
func (m *ExecOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecOrder.DiscardUnknown(m)
}

var xxx_messageInfo_ExecOrder proto.InternalMessageInfo

func (m *ExecOrder) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Isok struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Isok) Reset()         { *m = Isok{} }
func (m *Isok) String() string { return proto.CompactTextString(m) }
func (*Isok) ProtoMessage()    {}
func (*Isok) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{5}
}

func (m *Isok) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Isok.Unmarshal(m, b)
}
func (m *Isok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Isok.Marshal(b, m, deterministic)
}
func (m *Isok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Isok.Merge(m, src)
}
func (m *Isok) XXX_Size() int {
	return xxx_messageInfo_Isok.Size(m)
}
func (m *Isok) XXX_DiscardUnknown() {
	xxx_messageInfo_Isok.DiscardUnknown(m)
}

var xxx_messageInfo_Isok proto.InternalMessageInfo

func (m *Isok) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type InsulateWordList struct {
	InsulateWordList     []string `protobuf:"bytes,1,rep,name=InsulateWordList,proto3" json:"InsulateWordList"`
	SQL                  string   `protobuf:"bytes,2,opt,name=SQL,proto3" json:"SQL"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsulateWordList) Reset()         { *m = InsulateWordList{} }
func (m *InsulateWordList) String() string { return proto.CompactTextString(m) }
func (*InsulateWordList) ProtoMessage()    {}
func (*InsulateWordList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0932ade6a231544d, []int{6}
}

func (m *InsulateWordList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsulateWordList.Unmarshal(m, b)
}
func (m *InsulateWordList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsulateWordList.Marshal(b, m, deterministic)
}
func (m *InsulateWordList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsulateWordList.Merge(m, src)
}
func (m *InsulateWordList) XXX_Size() int {
	return xxx_messageInfo_InsulateWordList.Size(m)
}
func (m *InsulateWordList) XXX_DiscardUnknown() {
	xxx_messageInfo_InsulateWordList.DiscardUnknown(m)
}

var xxx_messageInfo_InsulateWordList proto.InternalMessageInfo

func (m *InsulateWordList) GetInsulateWordList() []string {
	if m != nil {
		return m.InsulateWordList
	}
	return nil
}

func (m *InsulateWordList) GetSQL() string {
	if m != nil {
		return m.SQL
	}
	return ""
}

func init() {
	proto.RegisterType((*LibraAuditOrder)(nil), "proto.LibraAuditOrder")
	proto.RegisterType((*Source)(nil), "proto.Source")
	proto.RegisterType((*Record)(nil), "proto.Record")
	proto.RegisterType((*RecordSet)(nil), "proto.RecordSet")
	proto.RegisterType((*ExecOrder)(nil), "proto.ExecOrder")
	proto.RegisterType((*Isok)(nil), "proto.Isok")
	proto.RegisterType((*InsulateWordList)(nil), "proto.InsulateWordList")
}

func init() { proto.RegisterFile("detection.proto", fileDescriptor_0932ade6a231544d) }

var fileDescriptor_0932ade6a231544d = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x6d, 0xdc, 0xd8, 0x89, 0x27, 0xdf, 0x47, 0xa3, 0x15, 0x0a, 0xab, 0x1e, 0xaa, 0xc8, 0x52,
	0xa5, 0x88, 0x43, 0x05, 0xe1, 0x80, 0xc4, 0x2d, 0x25, 0x3d, 0x04, 0x1c, 0x92, 0x6e, 0x8a, 0x7a,
	0x65, 0x63, 0x4f, 0x21, 0xb2, 0xc9, 0x56, 0xbb, 0x6b, 0x0a, 0x17, 0x7e, 0x0e, 0xbf, 0x89, 0x9f,
	0x83, 0x76, 0xbc, 0x89, 0x42, 0xa3, 0x0a, 0x38, 0x79, 0xde, 0x9b, 0x79, 0x9e, 0x99, 0xb7, 0x03,
	0x47, 0x39, 0x5a, 0xcc, 0xec, 0x4a, 0xad, 0xcf, 0x6e, 0xb5, 0xb2, 0x8a, 0x85, 0xf4, 0x49, 0x7e,
	0x04, 0x70, 0x94, 0xae, 0x96, 0x5a, 0x8e, 0xaa, 0x7c, 0x65, 0x67, 0x3a, 0x47, 0xcd, 0xba, 0x70,
	0xb8, 0xb8, 0x4c, 0x79, 0xa3, 0xdf, 0x18, 0xc4, 0xc2, 0x85, 0xec, 0x18, 0xda, 0x63, 0x69, 0xe5,
	0xb9, 0x34, 0xc8, 0x03, 0xa2, 0xb7, 0x98, 0x3d, 0x86, 0xf0, 0x4a, 0x2e, 0x4b, 0xe4, 0x87, 0x94,
	0xa8, 0x01, 0xe3, 0xd0, 0xba, 0xf8, 0x8a, 0x59, 0x65, 0x91, 0x37, 0xfb, 0x8d, 0x41, 0x5b, 0x6c,
	0xa0, 0xab, 0x7f, 0xfd, 0x09, 0xb3, 0x82, 0x87, 0xc4, 0xd7, 0xc0, 0xb1, 0x13, 0x33, 0x9e, 0xa6,
	0x3c, 0xaa, 0x59, 0x02, 0xac, 0x07, 0xd1, 0xb9, 0xcc, 0x8a, 0xea, 0x96, 0xb7, 0x88, 0xf6, 0x88,
	0x9d, 0x42, 0x64, 0x54, 0xa5, 0x33, 0xe4, 0xed, 0x7e, 0x63, 0xd0, 0x19, 0xfe, 0x5f, 0x2f, 0x75,
	0xb6, 0x20, 0x52, 0xf8, 0xa4, 0x93, 0x5f, 0x2b, 0x5d, 0x4c, 0x72, 0x1e, 0xd3, 0x6c, 0x1e, 0xb1,
	0x13, 0x80, 0x89, 0x19, 0x55, 0x56, 0x5d, 0x49, 0x53, 0x70, 0xa0, 0x5f, 0xef, 0x30, 0x8c, 0x41,
	0xf3, 0x9d, 0xfc, 0x8c, 0xbc, 0x43, 0x2a, 0x8a, 0x93, 0x0f, 0x10, 0xd5, 0x7f, 0x77, 0xd9, 0x51,
	0x9e, 0x6b, 0xef, 0x0f, 0xc5, 0x8e, 0x7b, 0x6f, 0x50, 0x7b, 0x73, 0x28, 0x76, 0xa6, 0xcd, 0xa5,
	0x31, 0x77, 0x4a, 0xe7, 0xde, 0x9b, 0x2d, 0x76, 0xf5, 0x73, 0xa5, 0x2d, 0x79, 0x13, 0x0a, 0x8a,
	0x93, 0xef, 0x10, 0x09, 0xcc, 0x5c, 0x76, 0xff, 0x01, 0x4e, 0x00, 0x46, 0x37, 0x37, 0x98, 0x59,
	0xa1, 0xee, 0x0c, 0x75, 0x09, 0xc5, 0x0e, 0xe3, 0x36, 0x5d, 0x58, 0x69, 0x2b, 0xe3, 0x3b, 0x79,
	0xe4, 0x6c, 0xbd, 0xd0, 0x5a, 0x69, 0x6a, 0x14, 0x8b, 0x1a, 0x38, 0x36, 0xc5, 0x2f, 0x58, 0x92,
	0xd9, 0xa1, 0xa8, 0x41, 0x32, 0x84, 0xb8, 0xee, 0xbf, 0x40, 0xeb, 0x1c, 0xd6, 0x04, 0x78, 0xa3,
	0x7f, 0xb8, 0xe3, 0x70, 0x5d, 0x21, 0x7c, 0x32, 0x39, 0x85, 0xd8, 0xbd, 0x6b, 0x7d, 0x37, 0x1c,
	0x5a, 0x53, 0x34, 0x46, 0x7e, 0x44, 0x3f, 0xfa, 0x06, 0x26, 0x3d, 0x68, 0x4e, 0x8c, 0x2a, 0xd8,
	0x23, 0x08, 0x66, 0x05, 0x25, 0xdb, 0x22, 0x98, 0x15, 0xc9, 0x1c, 0xba, 0x93, 0xb5, 0xa9, 0x4a,
	0x69, 0xf1, 0x5a, 0xe9, 0x3c, 0x5d, 0x19, 0xcb, 0x9e, 0xee, 0x73, 0x34, 0x43, 0x2c, 0xf6, 0x6b,
	0xbd, 0x51, 0xc1, 0xd6, 0xa8, 0xe1, 0xcf, 0x00, 0x9a, 0x6f, 0xaa, 0xb5, 0x62, 0x2f, 0x21, 0xa6,
	0xa9, 0xc6, 0x28, 0x4b, 0xd6, 0xf3, 0xd3, 0xdf, 0xbb, 0xf4, 0xe3, 0xee, 0x6f, 0x5b, 0x2d, 0xd0,
	0x26, 0x07, 0xec, 0x15, 0xfc, 0x57, 0x0b, 0xa7, 0xa9, 0x5b, 0xed, 0x8f, 0xda, 0xed, 0xfe, 0xbb,
	0xda, 0xf1, 0xbf, 0x6b, 0x9f, 0x43, 0x7b, 0x7b, 0x80, 0x0f, 0xe9, 0x3a, 0x9e, 0x77, 0x66, 0x52,
	0xbb, 0xf0, 0xb2, 0x42, 0xfd, 0xed, 0xc1, 0xfa, 0x27, 0x9b, 0xfa, 0x7b, 0xc6, 0x25, 0x07, 0xec,
	0x19, 0xb4, 0xde, 0xae, 0xca, 0x72, 0x66, 0xb2, 0xbf, 0xec, 0xb6, 0x8c, 0x08, 0xbd, 0xf8, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x28, 0xa9, 0x6c, 0x4b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JunoClient is the client API for Juno service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JunoClient interface {
	OrderDeal(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*RecordSet, error)
	OrderDMLExec(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*ExecOrder, error)
	OrderDDLExec(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*ExecOrder, error)
	AutoTask(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*Isok, error)
	Query(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*InsulateWordList, error)
	KillOsc(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*Isok, error)
}

type junoClient struct {
	cc *grpc.ClientConn
}

func NewJunoClient(cc *grpc.ClientConn) JunoClient {
	return &junoClient{cc}
}

func (c *junoClient) OrderDeal(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*RecordSet, error) {
	out := new(RecordSet)
	err := c.cc.Invoke(ctx, "/proto.Juno/OrderDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *junoClient) OrderDMLExec(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*ExecOrder, error) {
	out := new(ExecOrder)
	err := c.cc.Invoke(ctx, "/proto.Juno/OrderDMLExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *junoClient) OrderDDLExec(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*ExecOrder, error) {
	out := new(ExecOrder)
	err := c.cc.Invoke(ctx, "/proto.Juno/OrderDDLExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *junoClient) AutoTask(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*Isok, error) {
	out := new(Isok)
	err := c.cc.Invoke(ctx, "/proto.Juno/AutoTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *junoClient) Query(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*InsulateWordList, error) {
	out := new(InsulateWordList)
	err := c.cc.Invoke(ctx, "/proto.Juno/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *junoClient) KillOsc(ctx context.Context, in *LibraAuditOrder, opts ...grpc.CallOption) (*Isok, error) {
	out := new(Isok)
	err := c.cc.Invoke(ctx, "/proto.Juno/KillOsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JunoServer is the server API for Juno service.
type JunoServer interface {
	OrderDeal(context.Context, *LibraAuditOrder) (*RecordSet, error)
	OrderDMLExec(context.Context, *LibraAuditOrder) (*ExecOrder, error)
	OrderDDLExec(context.Context, *LibraAuditOrder) (*ExecOrder, error)
	AutoTask(context.Context, *LibraAuditOrder) (*Isok, error)
	Query(context.Context, *LibraAuditOrder) (*InsulateWordList, error)
	KillOsc(context.Context, *LibraAuditOrder) (*Isok, error)
}

// UnimplementedJunoServer can be embedded to have forward compatible implementations.
type UnimplementedJunoServer struct {
}

func (*UnimplementedJunoServer) OrderDeal(ctx context.Context, req *LibraAuditOrder) (*RecordSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDeal not implemented")
}
func (*UnimplementedJunoServer) OrderDMLExec(ctx context.Context, req *LibraAuditOrder) (*ExecOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDMLExec not implemented")
}
func (*UnimplementedJunoServer) OrderDDLExec(ctx context.Context, req *LibraAuditOrder) (*ExecOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDDLExec not implemented")
}
func (*UnimplementedJunoServer) AutoTask(ctx context.Context, req *LibraAuditOrder) (*Isok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutoTask not implemented")
}
func (*UnimplementedJunoServer) Query(ctx context.Context, req *LibraAuditOrder) (*InsulateWordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedJunoServer) KillOsc(ctx context.Context, req *LibraAuditOrder) (*Isok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillOsc not implemented")
}

func RegisterJunoServer(s *grpc.Server, srv JunoServer) {
	s.RegisterService(&_Juno_serviceDesc, srv)
}

func _Juno_OrderDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).OrderDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/OrderDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).OrderDeal(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Juno_OrderDMLExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).OrderDMLExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/OrderDMLExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).OrderDMLExec(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Juno_OrderDDLExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).OrderDDLExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/OrderDDLExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).OrderDDLExec(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Juno_AutoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).AutoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/AutoTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).AutoTask(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Juno_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).Query(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Juno_KillOsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibraAuditOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunoServer).KillOsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juno/KillOsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunoServer).KillOsc(ctx, req.(*LibraAuditOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Juno_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Juno",
	HandlerType: (*JunoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderDeal",
			Handler:    _Juno_OrderDeal_Handler,
		},
		{
			MethodName: "OrderDMLExec",
			Handler:    _Juno_OrderDMLExec_Handler,
		},
		{
			MethodName: "OrderDDLExec",
			Handler:    _Juno_OrderDDLExec_Handler,
		},
		{
			MethodName: "AutoTask",
			Handler:    _Juno_AutoTask_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Juno_Query_Handler,
		},
		{
			MethodName: "KillOsc",
			Handler:    _Juno_KillOsc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "detection.proto",
}
