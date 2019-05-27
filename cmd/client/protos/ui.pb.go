// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ui.proto

package UIproto

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

type ReportRequest struct {
	Server               *Server           `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	ServerDB             *ServerDB         `protobuf:"bytes,2,opt,name=serverDB,proto3" json:"serverDB,omitempty"`
	ClientDB             *ClientDB         `protobuf:"bytes,3,opt,name=clientDB,proto3" json:"clientDB,omitempty"`
	Flock                []byte            `protobuf:"bytes,4,opt,name=flock,proto3" json:"flock,omitempty"`
	Plugin               []byte            `protobuf:"bytes,5,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Params               map[string]string `protobuf:"bytes,6,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{0}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *ReportRequest) GetServerDB() *ServerDB {
	if m != nil {
		return m.ServerDB
	}
	return nil
}

func (m *ReportRequest) GetClientDB() *ClientDB {
	if m != nil {
		return m.ClientDB
	}
	return nil
}

func (m *ReportRequest) GetFlock() []byte {
	if m != nil {
		return m.Flock
	}
	return nil
}

func (m *ReportRequest) GetPlugin() []byte {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (m *ReportRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type ReportResponse struct {
	Chunks               int64    `protobuf:"varint,1,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Tables               int64    `protobuf:"varint,2,opt,name=tables,proto3" json:"tables,omitempty"`
	Percentage           int64    `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{1}
}

func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (m *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(m, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func (m *ReportResponse) GetChunks() int64 {
	if m != nil {
		return m.Chunks
	}
	return 0
}

func (m *ReportResponse) GetTables() int64 {
	if m != nil {
		return m.Tables
	}
	return 0
}

func (m *ReportResponse) GetPercentage() int64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type PingRequest struct {
	// Types that are valid to be assigned to Value:
	//	*PingRequest_Server
	//	*PingRequest_ClientDB
	//	*PingRequest_ServerDB
	Value                isPingRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{2}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type isPingRequest_Value interface {
	isPingRequest_Value()
}

type PingRequest_Server struct {
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3,oneof"`
}

type PingRequest_ClientDB struct {
	ClientDB *ClientDB `protobuf:"bytes,2,opt,name=clientDB,proto3,oneof"`
}

type PingRequest_ServerDB struct {
	ServerDB *ServerDB `protobuf:"bytes,3,opt,name=serverDB,proto3,oneof"`
}

func (*PingRequest_Server) isPingRequest_Value() {}

func (*PingRequest_ClientDB) isPingRequest_Value() {}

func (*PingRequest_ServerDB) isPingRequest_Value() {}

func (m *PingRequest) GetValue() isPingRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PingRequest) GetServer() *Server {
	if x, ok := m.GetValue().(*PingRequest_Server); ok {
		return x.Server
	}
	return nil
}

func (m *PingRequest) GetClientDB() *ClientDB {
	if x, ok := m.GetValue().(*PingRequest_ClientDB); ok {
		return x.ClientDB
	}
	return nil
}

func (m *PingRequest) GetServerDB() *ServerDB {
	if x, ok := m.GetValue().(*PingRequest_ServerDB); ok {
		return x.ServerDB
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PingRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PingRequest_Server)(nil),
		(*PingRequest_ClientDB)(nil),
		(*PingRequest_ServerDB)(nil),
	}
}

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{3}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

type SchemaFile struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaFile) Reset()         { *m = SchemaFile{} }
func (m *SchemaFile) String() string { return proto.CompactTextString(m) }
func (*SchemaFile) ProtoMessage()    {}
func (*SchemaFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{4}
}

func (m *SchemaFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaFile.Unmarshal(m, b)
}
func (m *SchemaFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaFile.Marshal(b, m, deterministic)
}
func (m *SchemaFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaFile.Merge(m, src)
}
func (m *SchemaFile) XXX_Size() int {
	return xxx_messageInfo_SchemaFile.Size(m)
}
func (m *SchemaFile) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaFile.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaFile proto.InternalMessageInfo

func (m *SchemaFile) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type SchemaResponse struct {
	Params               []string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaResponse) Reset()         { *m = SchemaResponse{} }
func (m *SchemaResponse) String() string { return proto.CompactTextString(m) }
func (*SchemaResponse) ProtoMessage()    {}
func (*SchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{5}
}

func (m *SchemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaResponse.Unmarshal(m, b)
}
func (m *SchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaResponse.Marshal(b, m, deterministic)
}
func (m *SchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaResponse.Merge(m, src)
}
func (m *SchemaResponse) XXX_Size() int {
	return xxx_messageInfo_SchemaResponse.Size(m)
}
func (m *SchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaResponse proto.InternalMessageInfo

func (m *SchemaResponse) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Server struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{6}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ClientDB struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientDB) Reset()         { *m = ClientDB{} }
func (m *ClientDB) String() string { return proto.CompactTextString(m) }
func (*ClientDB) ProtoMessage()    {}
func (*ClientDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{7}
}

func (m *ClientDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientDB.Unmarshal(m, b)
}
func (m *ClientDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientDB.Marshal(b, m, deterministic)
}
func (m *ClientDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientDB.Merge(m, src)
}
func (m *ClientDB) XXX_Size() int {
	return xxx_messageInfo_ClientDB.Size(m)
}
func (m *ClientDB) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientDB.DiscardUnknown(m)
}

var xxx_messageInfo_ClientDB proto.InternalMessageInfo

func (m *ClientDB) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ClientDB) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type ServerDB struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Database             string   `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerDB) Reset()         { *m = ServerDB{} }
func (m *ServerDB) String() string { return proto.CompactTextString(m) }
func (*ServerDB) ProtoMessage()    {}
func (*ServerDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_63867a62624c1283, []int{8}
}

func (m *ServerDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerDB.Unmarshal(m, b)
}
func (m *ServerDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerDB.Marshal(b, m, deterministic)
}
func (m *ServerDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerDB.Merge(m, src)
}
func (m *ServerDB) XXX_Size() int {
	return xxx_messageInfo_ServerDB.Size(m)
}
func (m *ServerDB) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerDB.DiscardUnknown(m)
}

var xxx_messageInfo_ServerDB proto.InternalMessageInfo

func (m *ServerDB) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *ServerDB) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ServerDB) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func init() {
	proto.RegisterType((*ReportRequest)(nil), "UIproto.ReportRequest")
	proto.RegisterMapType((map[string]string)(nil), "UIproto.ReportRequest.ParamsEntry")
	proto.RegisterType((*ReportResponse)(nil), "UIproto.ReportResponse")
	proto.RegisterType((*PingRequest)(nil), "UIproto.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "UIproto.PingResponse")
	proto.RegisterType((*SchemaFile)(nil), "UIproto.SchemaFile")
	proto.RegisterType((*SchemaResponse)(nil), "UIproto.SchemaResponse")
	proto.RegisterType((*Server)(nil), "UIproto.Server")
	proto.RegisterType((*ClientDB)(nil), "UIproto.ClientDB")
	proto.RegisterType((*ServerDB)(nil), "UIproto.ServerDB")
}

func init() { proto.RegisterFile("ui.proto", fileDescriptor_63867a62624c1283) }

var fileDescriptor_63867a62624c1283 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0xec, 0x2e, 0x4b, 0x5f, 0x4b, 0x01, 0x33, 0x46, 0xd4, 0x03, 0x8a, 0x7c, 0x21, 0x1c,
	0xe8, 0x50, 0x77, 0x19, 0xe5, 0x56, 0x0a, 0xda, 0x6e, 0x93, 0xc7, 0xee, 0xb8, 0xc1, 0xeb, 0xa2,
	0x7a, 0x49, 0x70, 0x92, 0x49, 0xfb, 0x7b, 0xf8, 0x0f, 0xb8, 0xf2, 0xcf, 0x21, 0x3b, 0x76, 0x48,
	0xa3, 0x4d, 0x70, 0x7b, 0xef, 0xcb, 0xf7, 0x7e, 0x7d, 0x9f, 0x03, 0x41, 0x9d, 0xce, 0x0a, 0x95,
	0x57, 0x39, 0x39, 0xb8, 0x3a, 0x37, 0x01, 0xfd, 0x8d, 0xe0, 0x09, 0x13, 0x45, 0xae, 0x2a, 0x26,
	0x7e, 0xd4, 0xa2, 0xac, 0xc8, 0x1b, 0xf0, 0x4b, 0xa1, 0xee, 0x84, 0x0a, 0xbd, 0xc8, 0x8b, 0x47,
	0xf3, 0xa7, 0x33, 0xcb, 0x9d, 0x5d, 0x1a, 0x98, 0xd9, 0xcf, 0xe4, 0x1d, 0x04, 0x4d, 0xb4, 0x5a,
	0x86, 0xc8, 0x50, 0x9f, 0xf7, 0xa8, 0xab, 0x25, 0x6b, 0x29, 0x9a, 0x9e, 0xc8, 0x54, 0x64, 0xd5,
	0x6a, 0x19, 0xe2, 0x1e, 0xfd, 0x93, 0xfd, 0xc0, 0x5a, 0x0a, 0x39, 0x84, 0xfd, 0x6b, 0x99, 0x27,
	0xdb, 0x70, 0x10, 0x79, 0xf1, 0x98, 0x35, 0x09, 0x39, 0x02, 0xbf, 0x90, 0xf5, 0x26, 0xcd, 0xc2,
	0x7d, 0x03, 0xdb, 0x8c, 0x2c, 0xc0, 0x2f, 0xb8, 0xe2, 0xb7, 0x65, 0xe8, 0x47, 0x38, 0x1e, 0xcd,
	0x69, 0xdb, 0x7a, 0xe7, 0xb8, 0xd9, 0x85, 0x21, 0x7d, 0xce, 0x2a, 0x75, 0xcf, 0x6c, 0xc5, 0xf4,
	0x03, 0x8c, 0x3a, 0x30, 0x79, 0x06, 0x78, 0x2b, 0xee, 0xcd, 0xf1, 0x43, 0xa6, 0x43, 0xbd, 0xca,
	0x1d, 0x97, 0xb5, 0x30, 0x57, 0x0e, 0x59, 0x93, 0x2c, 0xd0, 0xa9, 0x47, 0xbf, 0xc1, 0xc4, 0xf5,
	0x2f, 0x8b, 0x3c, 0x2b, 0x85, 0x5e, 0x30, 0xb9, 0xa9, 0xb3, 0x6d, 0x69, 0x1a, 0x60, 0x66, 0x33,
	0x8d, 0x57, 0x7c, 0x2d, 0x45, 0x69, 0x9a, 0x60, 0x66, 0x33, 0xf2, 0x1a, 0xa0, 0x10, 0x2a, 0x11,
	0x59, 0xc5, 0x37, 0xc2, 0xe8, 0x82, 0x59, 0x07, 0xa1, 0x3f, 0x3d, 0x18, 0x5d, 0xa4, 0xd9, 0xc6,
	0xb9, 0xf3, 0xf6, 0x1f, 0xee, 0x9c, 0xed, 0xb5, 0xfe, 0x1c, 0x77, 0x04, 0x47, 0x8f, 0x08, 0x7e,
	0xb6, 0xd7, 0x91, 0xfc, 0xb8, 0x63, 0x28, 0x7e, 0xc4, 0x50, 0x5d, 0xe0, 0x48, 0xcb, 0x03, 0x2b,
	0x0c, 0x9d, 0xc0, 0xb8, 0x59, 0xb2, 0x51, 0x81, 0x46, 0x00, 0x97, 0xc9, 0x8d, 0xb8, 0xe5, 0x5f,
	0x52, 0x29, 0x08, 0x81, 0xc1, 0x75, 0x2a, 0x85, 0xd9, 0x78, 0xcc, 0x4c, 0x4c, 0x63, 0x98, 0x34,
	0x8c, 0xae, 0x72, 0xd6, 0x42, 0x14, 0xe1, 0x78, 0xe8, 0xec, 0xa1, 0x21, 0xf8, 0xcd, 0x70, 0x32,
	0x01, 0x94, 0x16, 0xd6, 0x18, 0x94, 0x16, 0xf4, 0x14, 0x02, 0x77, 0x87, 0x76, 0xad, 0x56, 0xd2,
	0xb9, 0x56, 0x2b, 0x49, 0xa6, 0x10, 0x7c, 0xe7, 0x15, 0x5f, 0xf3, 0xd2, 0x19, 0xd7, 0xe6, 0x94,
	0x43, 0xe0, 0x0e, 0xfa, 0xff, 0xf7, 0x6e, 0x47, 0xa0, 0x87, 0x47, 0xe0, 0xdd, 0x11, 0xf3, 0x5f,
	0x1e, 0xa0, 0xab, 0x73, 0x72, 0x02, 0x03, 0xad, 0x0c, 0x39, 0x6c, 0xbb, 0x76, 0xdc, 0x9c, 0xbe,
	0xec, 0xa1, 0x56, 0x8a, 0x85, 0x93, 0xef, 0xab, 0xb6, 0xfc, 0xc5, 0xdf, 0x85, 0x5a, 0x4d, 0xa7,
	0xaf, 0x7a, 0x60, 0x5b, 0xfb, 0x11, 0xfc, 0xe6, 0x49, 0x92, 0xa3, 0x87, 0xff, 0x81, 0x4e, 0xe9,
	0xee, 0xdb, 0x7d, 0xef, 0xad, 0x7d, 0x83, 0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xeb,
	0x41, 0x7a, 0x29, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UIClient is the client API for UI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SchemaTest(ctx context.Context, in *SchemaFile, opts ...grpc.CallOption) (*SchemaResponse, error)
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (UI_ReportClient, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/UIproto.UI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) SchemaTest(ctx context.Context, in *SchemaFile, opts ...grpc.CallOption) (*SchemaResponse, error) {
	out := new(SchemaResponse)
	err := c.cc.Invoke(ctx, "/UIproto.UI/SchemaTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (UI_ReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UI_serviceDesc.Streams[0], "/UIproto.UI/Report", opts...)
	if err != nil {
		return nil, err
	}
	x := &uIReportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UI_ReportClient interface {
	Recv() (*ReportResponse, error)
	grpc.ClientStream
}

type uIReportClient struct {
	grpc.ClientStream
}

func (x *uIReportClient) Recv() (*ReportResponse, error) {
	m := new(ReportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UIServer is the server API for UI service.
type UIServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SchemaTest(context.Context, *SchemaFile) (*SchemaResponse, error)
	Report(*ReportRequest, UI_ReportServer) error
}

// UnimplementedUIServer can be embedded to have forward compatible implementations.
type UnimplementedUIServer struct {
}

func (*UnimplementedUIServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedUIServer) SchemaTest(ctx context.Context, req *SchemaFile) (*SchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaTest not implemented")
}
func (*UnimplementedUIServer) Report(req *ReportRequest, srv UI_ReportServer) error {
	return status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UIproto.UI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_SchemaTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).SchemaTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UIproto.UI/SchemaTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).SchemaTest(ctx, req.(*SchemaFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Report_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UIServer).Report(m, &uIReportServer{stream})
}

type UI_ReportServer interface {
	Send(*ReportResponse) error
	grpc.ServerStream
}

type uIReportServer struct {
	grpc.ServerStream
}

func (x *uIReportServer) Send(m *ReportResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UIproto.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "SchemaTest",
			Handler:    _UI_SchemaTest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Report",
			Handler:       _UI_Report_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ui.proto",
}
