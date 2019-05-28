// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flock.proto

package flock

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

type FlockRequest struct {
	// Types that are valid to be assigned to Value:
	//	*FlockRequest_Start
	//	*FlockRequest_Ping
	//	*FlockRequest_Batch
	//	*FlockRequest_End
	Value                isFlockRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlockRequest) Reset()         { *m = FlockRequest{} }
func (m *FlockRequest) String() string { return proto.CompactTextString(m) }
func (*FlockRequest) ProtoMessage()    {}
func (*FlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{0}
}

func (m *FlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlockRequest.Unmarshal(m, b)
}
func (m *FlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlockRequest.Marshal(b, m, deterministic)
}
func (m *FlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlockRequest.Merge(m, src)
}
func (m *FlockRequest) XXX_Size() int {
	return xxx_messageInfo_FlockRequest.Size(m)
}
func (m *FlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlockRequest proto.InternalMessageInfo

type isFlockRequest_Value interface {
	isFlockRequest_Value()
}

type FlockRequest_Start struct {
	Start *Start `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type FlockRequest_Ping struct {
	Ping *Ping `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

type FlockRequest_Batch struct {
	Batch *Batch `protobuf:"bytes,3,opt,name=batch,proto3,oneof"`
}

type FlockRequest_End struct {
	End *EndStream `protobuf:"bytes,4,opt,name=end,proto3,oneof"`
}

func (*FlockRequest_Start) isFlockRequest_Value() {}

func (*FlockRequest_Ping) isFlockRequest_Value() {}

func (*FlockRequest_Batch) isFlockRequest_Value() {}

func (*FlockRequest_End) isFlockRequest_Value() {}

func (m *FlockRequest) GetValue() isFlockRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FlockRequest) GetStart() *Start {
	if x, ok := m.GetValue().(*FlockRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (m *FlockRequest) GetPing() *Ping {
	if x, ok := m.GetValue().(*FlockRequest_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *FlockRequest) GetBatch() *Batch {
	if x, ok := m.GetValue().(*FlockRequest_Batch); ok {
		return x.Batch
	}
	return nil
}

func (m *FlockRequest) GetEnd() *EndStream {
	if x, ok := m.GetValue().(*FlockRequest_End); ok {
		return x.End
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FlockRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FlockRequest_Start)(nil),
		(*FlockRequest_Ping)(nil),
		(*FlockRequest_Batch)(nil),
		(*FlockRequest_End)(nil),
	}
}

type FlockResponse struct {
	// Types that are valid to be assigned to Value:
	//	*FlockResponse_Pong
	//	*FlockResponse_Batch
	Value                isFlockResponse_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FlockResponse) Reset()         { *m = FlockResponse{} }
func (m *FlockResponse) String() string { return proto.CompactTextString(m) }
func (*FlockResponse) ProtoMessage()    {}
func (*FlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{1}
}

func (m *FlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlockResponse.Unmarshal(m, b)
}
func (m *FlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlockResponse.Marshal(b, m, deterministic)
}
func (m *FlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlockResponse.Merge(m, src)
}
func (m *FlockResponse) XXX_Size() int {
	return xxx_messageInfo_FlockResponse.Size(m)
}
func (m *FlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlockResponse proto.InternalMessageInfo

type isFlockResponse_Value interface {
	isFlockResponse_Value()
}

type FlockResponse_Pong struct {
	Pong *Pong `protobuf:"bytes,1,opt,name=pong,proto3,oneof"`
}

type FlockResponse_Batch struct {
	Batch *BatchInsertResponse `protobuf:"bytes,2,opt,name=batch,proto3,oneof"`
}

func (*FlockResponse_Pong) isFlockResponse_Value() {}

func (*FlockResponse_Batch) isFlockResponse_Value() {}

func (m *FlockResponse) GetValue() isFlockResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FlockResponse) GetPong() *Pong {
	if x, ok := m.GetValue().(*FlockResponse_Pong); ok {
		return x.Pong
	}
	return nil
}

func (m *FlockResponse) GetBatch() *BatchInsertResponse {
	if x, ok := m.GetValue().(*FlockResponse_Batch); ok {
		return x.Batch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FlockResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FlockResponse_Pong)(nil),
		(*FlockResponse_Batch)(nil),
	}
}

type Start struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Dollar               bool     `protobuf:"varint,3,opt,name=dollar,proto3" json:"dollar,omitempty"`
	Schema               []byte   `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Plugin               []byte   `protobuf:"bytes,5,opt,name=plugin,proto3" json:"plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Start) Reset()         { *m = Start{} }
func (m *Start) String() string { return proto.CompactTextString(m) }
func (*Start) ProtoMessage()    {}
func (*Start) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{2}
}

func (m *Start) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Start.Unmarshal(m, b)
}
func (m *Start) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Start.Marshal(b, m, deterministic)
}
func (m *Start) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Start.Merge(m, src)
}
func (m *Start) XXX_Size() int {
	return xxx_messageInfo_Start.Size(m)
}
func (m *Start) XXX_DiscardUnknown() {
	xxx_messageInfo_Start.DiscardUnknown(m)
}

var xxx_messageInfo_Start proto.InternalMessageInfo

func (m *Start) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Start) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Start) GetDollar() bool {
	if m != nil {
		return m.Dollar
	}
	return false
}

func (m *Start) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *Start) GetPlugin() []byte {
	if m != nil {
		return m.Plugin
	}
	return nil
}

type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{3}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

type Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{4}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

type DBPing struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBPing) Reset()         { *m = DBPing{} }
func (m *DBPing) String() string { return proto.CompactTextString(m) }
func (*DBPing) ProtoMessage()    {}
func (*DBPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{5}
}

func (m *DBPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBPing.Unmarshal(m, b)
}
func (m *DBPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBPing.Marshal(b, m, deterministic)
}
func (m *DBPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBPing.Merge(m, src)
}
func (m *DBPing) XXX_Size() int {
	return xxx_messageInfo_DBPing.Size(m)
}
func (m *DBPing) XXX_DiscardUnknown() {
	xxx_messageInfo_DBPing.DiscardUnknown(m)
}

var xxx_messageInfo_DBPing proto.InternalMessageInfo

func (m *DBPing) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DBPing) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type Batch struct {
	BatchId string `protobuf:"bytes,1,opt,name=BatchId,proto3" json:"BatchId,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Batch_Head
	//	*Batch_Chunk
	//	*Batch_Tail
	Value                isBatch_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{6}
}

func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

type isBatch_Value interface {
	isBatch_Value()
}

type Batch_Head struct {
	Head *BatchInsertHead `protobuf:"bytes,2,opt,name=head,proto3,oneof"`
}

type Batch_Chunk struct {
	Chunk *DataStream `protobuf:"bytes,3,opt,name=chunk,proto3,oneof"`
}

type Batch_Tail struct {
	Tail *BatchInsertTail `protobuf:"bytes,4,opt,name=tail,proto3,oneof"`
}

func (*Batch_Head) isBatch_Value() {}

func (*Batch_Chunk) isBatch_Value() {}

func (*Batch_Tail) isBatch_Value() {}

func (m *Batch) GetValue() isBatch_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Batch) GetHead() *BatchInsertHead {
	if x, ok := m.GetValue().(*Batch_Head); ok {
		return x.Head
	}
	return nil
}

func (m *Batch) GetChunk() *DataStream {
	if x, ok := m.GetValue().(*Batch_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (m *Batch) GetTail() *BatchInsertTail {
	if x, ok := m.GetValue().(*Batch_Tail); ok {
		return x.Tail
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Batch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Batch_Head)(nil),
		(*Batch_Chunk)(nil),
		(*Batch_Tail)(nil),
	}
}

type EndStream struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndStream) Reset()         { *m = EndStream{} }
func (m *EndStream) String() string { return proto.CompactTextString(m) }
func (*EndStream) ProtoMessage()    {}
func (*EndStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{7}
}

func (m *EndStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndStream.Unmarshal(m, b)
}
func (m *EndStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndStream.Marshal(b, m, deterministic)
}
func (m *EndStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndStream.Merge(m, src)
}
func (m *EndStream) XXX_Size() int {
	return xxx_messageInfo_EndStream.Size(m)
}
func (m *EndStream) XXX_DiscardUnknown() {
	xxx_messageInfo_EndStream.DiscardUnknown(m)
}

var xxx_messageInfo_EndStream proto.InternalMessageInfo

type BatchInsertHead struct {
	TableName            string   `protobuf:"bytes,2,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Chunks               int64    `protobuf:"varint,3,opt,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchInsertHead) Reset()         { *m = BatchInsertHead{} }
func (m *BatchInsertHead) String() string { return proto.CompactTextString(m) }
func (*BatchInsertHead) ProtoMessage()    {}
func (*BatchInsertHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{8}
}

func (m *BatchInsertHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchInsertHead.Unmarshal(m, b)
}
func (m *BatchInsertHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchInsertHead.Marshal(b, m, deterministic)
}
func (m *BatchInsertHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchInsertHead.Merge(m, src)
}
func (m *BatchInsertHead) XXX_Size() int {
	return xxx_messageInfo_BatchInsertHead.Size(m)
}
func (m *BatchInsertHead) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchInsertHead.DiscardUnknown(m)
}

var xxx_messageInfo_BatchInsertHead proto.InternalMessageInfo

func (m *BatchInsertHead) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *BatchInsertHead) GetChunks() int64 {
	if m != nil {
		return m.Chunks
	}
	return 0
}

type DataStream struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataStream) Reset()         { *m = DataStream{} }
func (m *DataStream) String() string { return proto.CompactTextString(m) }
func (*DataStream) ProtoMessage()    {}
func (*DataStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{9}
}

func (m *DataStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStream.Unmarshal(m, b)
}
func (m *DataStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStream.Marshal(b, m, deterministic)
}
func (m *DataStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStream.Merge(m, src)
}
func (m *DataStream) XXX_Size() int {
	return xxx_messageInfo_DataStream.Size(m)
}
func (m *DataStream) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStream.DiscardUnknown(m)
}

var xxx_messageInfo_DataStream proto.InternalMessageInfo

func (m *DataStream) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *DataStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BatchInsertTail struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchInsertTail) Reset()         { *m = BatchInsertTail{} }
func (m *BatchInsertTail) String() string { return proto.CompactTextString(m) }
func (*BatchInsertTail) ProtoMessage()    {}
func (*BatchInsertTail) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{10}
}

func (m *BatchInsertTail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchInsertTail.Unmarshal(m, b)
}
func (m *BatchInsertTail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchInsertTail.Marshal(b, m, deterministic)
}
func (m *BatchInsertTail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchInsertTail.Merge(m, src)
}
func (m *BatchInsertTail) XXX_Size() int {
	return xxx_messageInfo_BatchInsertTail.Size(m)
}
func (m *BatchInsertTail) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchInsertTail.DiscardUnknown(m)
}

var xxx_messageInfo_BatchInsertTail proto.InternalMessageInfo

type BatchInsertResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchInsertResponse) Reset()         { *m = BatchInsertResponse{} }
func (m *BatchInsertResponse) String() string { return proto.CompactTextString(m) }
func (*BatchInsertResponse) ProtoMessage()    {}
func (*BatchInsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87515fe46cf1e69, []int{11}
}

func (m *BatchInsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchInsertResponse.Unmarshal(m, b)
}
func (m *BatchInsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchInsertResponse.Marshal(b, m, deterministic)
}
func (m *BatchInsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchInsertResponse.Merge(m, src)
}
func (m *BatchInsertResponse) XXX_Size() int {
	return xxx_messageInfo_BatchInsertResponse.Size(m)
}
func (m *BatchInsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchInsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchInsertResponse proto.InternalMessageInfo

func (m *BatchInsertResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*FlockRequest)(nil), "flock.FlockRequest")
	proto.RegisterType((*FlockResponse)(nil), "flock.FlockResponse")
	proto.RegisterType((*Start)(nil), "flock.Start")
	proto.RegisterType((*Ping)(nil), "flock.Ping")
	proto.RegisterType((*Pong)(nil), "flock.Pong")
	proto.RegisterType((*DBPing)(nil), "flock.DBPing")
	proto.RegisterType((*Batch)(nil), "flock.Batch")
	proto.RegisterType((*EndStream)(nil), "flock.EndStream")
	proto.RegisterType((*BatchInsertHead)(nil), "flock.BatchInsertHead")
	proto.RegisterType((*DataStream)(nil), "flock.DataStream")
	proto.RegisterType((*BatchInsertTail)(nil), "flock.BatchInsertTail")
	proto.RegisterType((*BatchInsertResponse)(nil), "flock.BatchInsertResponse")
}

func init() { proto.RegisterFile("flock.proto", fileDescriptor_a87515fe46cf1e69) }

var fileDescriptor_a87515fe46cf1e69 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x68, 0xd3, 0xb5, 0xb7, 0x1d, 0x6c, 0xde, 0x34, 0x45, 0x15, 0x0f, 0x60, 0xed, 0x61,
	0x48, 0xd3, 0x40, 0x45, 0xea, 0x07, 0x54, 0x05, 0xca, 0x0b, 0x42, 0x1e, 0x3f, 0xe0, 0x26, 0x26,
	0x89, 0xe6, 0xda, 0x25, 0x76, 0x10, 0x2f, 0x7c, 0x04, 0xff, 0xc1, 0x0b, 0x7f, 0x88, 0xee, 0x8d,
	0x93, 0xb6, 0x6c, 0x2f, 0x3c, 0x35, 0xe7, 0xf8, 0xe4, 0xf8, 0xdc, 0x63, 0xa7, 0x30, 0xfe, 0xaa,
	0x6d, 0x7a, 0x7f, 0xbb, 0xad, 0xac, 0xb7, 0x2c, 0x26, 0xc0, 0x7f, 0x47, 0x30, 0x79, 0x8f, 0x4f,
	0x42, 0x7d, 0xab, 0x95, 0xf3, 0xec, 0x0a, 0x62, 0xe7, 0x65, 0xe5, 0x93, 0xe8, 0x45, 0x74, 0x3d,
	0x9e, 0x4d, 0x6e, 0x9b, 0x97, 0xee, 0x90, 0x5b, 0x1d, 0x89, 0x66, 0x91, 0xbd, 0x84, 0xfe, 0xb6,
	0x34, 0x79, 0xf2, 0x84, 0x44, 0xe3, 0x20, 0xfa, 0x5c, 0x9a, 0x7c, 0x75, 0x24, 0x68, 0x09, 0x8d,
	0xd6, 0xd2, 0xa7, 0x45, 0xd2, 0x3b, 0x30, 0x5a, 0x20, 0x87, 0x46, 0xb4, 0xc8, 0xae, 0xa0, 0xa7,
	0x4c, 0x96, 0xf4, 0x49, 0x73, 0x1a, 0x34, 0xef, 0x4c, 0x76, 0xe7, 0x2b, 0x25, 0x37, 0xab, 0x23,
	0x81, 0xcb, 0x8b, 0x63, 0x88, 0xbf, 0x4b, 0x5d, 0x2b, 0x6e, 0xe1, 0x24, 0xa4, 0x75, 0x5b, 0x6b,
	0x9c, 0xa2, 0x20, 0xd6, 0xe4, 0x21, 0x6d, 0x17, 0xc4, 0x86, 0x20, 0xd6, 0xe4, 0x6c, 0xd6, 0x06,
	0x69, 0xc2, 0x4e, 0xf7, 0x83, 0x7c, 0x34, 0x4e, 0x55, 0xbe, 0x75, 0xeb, 0x62, 0xed, 0x36, 0xfc,
	0x09, 0x31, 0x8d, 0xce, 0x4e, 0xa1, 0x57, 0x57, 0x9a, 0xf6, 0x19, 0x09, 0x7c, 0x64, 0x53, 0x18,
	0x66, 0xd2, 0xcb, 0xb5, 0x74, 0x8a, 0xac, 0x47, 0xa2, 0xc3, 0xec, 0x12, 0x06, 0x99, 0xd5, 0x5a,
	0x56, 0x34, 0xfd, 0x50, 0x04, 0x84, 0xbc, 0x4b, 0x0b, 0xb5, 0x91, 0x34, 0xf1, 0x44, 0x04, 0x84,
	0xfc, 0x56, 0xd7, 0x79, 0x69, 0x92, 0xb8, 0xe1, 0x1b, 0xc4, 0x07, 0xd0, 0xc7, 0x52, 0xe9, 0xd7,
	0x9a, 0x9c, 0xcf, 0x61, 0xb0, 0x5c, 0x20, 0xf3, 0x7f, 0x79, 0xf8, 0x9f, 0x08, 0x62, 0x1a, 0x98,
	0x25, 0x70, 0xdc, 0x4c, 0x9e, 0x85, 0x77, 0x5b, 0xc8, 0x6e, 0xa0, 0x5f, 0x28, 0x99, 0x85, 0x9a,
	0x2e, 0x1f, 0xd6, 0xb4, 0x52, 0x32, 0xc3, 0x56, 0x51, 0xc5, 0x5e, 0x41, 0x9c, 0x16, 0xb5, 0xb9,
	0x0f, 0xc7, 0x7b, 0x16, 0xe4, 0x4b, 0xe9, 0x65, 0x77, 0x76, 0x8d, 0x02, 0x8d, 0xbd, 0x2c, 0x75,
	0x38, 0xe4, 0x47, 0x8c, 0xbf, 0xc8, 0x52, 0xa3, 0x31, 0xaa, 0x76, 0xd5, 0x8f, 0x61, 0xd4, 0x5d,
	0x04, 0xfe, 0x01, 0x9e, 0xfd, 0x93, 0x84, 0x3d, 0x87, 0x91, 0x97, 0x6b, 0xad, 0x3e, 0xc9, 0x4d,
	0x3b, 0xf0, 0x8e, 0xc0, 0x46, 0x69, 0x77, 0x47, 0x01, 0x7b, 0x22, 0x20, 0x3e, 0x07, 0xd8, 0x65,
	0x64, 0x17, 0x10, 0x97, 0x26, 0x53, 0x3f, 0xa8, 0x8b, 0x9e, 0x68, 0x00, 0x63, 0xd0, 0xc7, 0xe6,
	0xc8, 0x74, 0x22, 0xe8, 0x99, 0x9f, 0x1d, 0x04, 0xc0, 0xc4, 0xfc, 0x35, 0x9c, 0x3f, 0x72, 0x89,
	0xb0, 0x61, 0x57, 0xa7, 0xa9, 0x72, 0x8e, 0x5c, 0x87, 0xa2, 0x85, 0xb3, 0x5f, 0x11, 0xc4, 0x74,
	0x7d, 0x19, 0x87, 0xc1, 0x4a, 0x49, 0xed, 0x0b, 0xb6, 0xff, 0xed, 0x4c, 0xf7, 0xef, 0x2f, 0xbb,
	0x81, 0xa7, 0xcb, 0x70, 0x7e, 0x41, 0x7b, 0xd2, 0x96, 0xbc, 0x78, 0xa8, 0x9e, 0xb7, 0xd6, 0xe7,
	0x81, 0xdd, 0xff, 0xaa, 0xa7, 0x17, 0x87, 0x64, 0x93, 0xf4, 0x3a, 0x7a, 0x13, 0xad, 0x07, 0xf4,
	0x77, 0xf0, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x36, 0x90, 0xa8, 0x1d, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlockClient is the client API for Flock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlockClient interface {
	Health(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	DatabaseHealth(ctx context.Context, in *DBPing, opts ...grpc.CallOption) (*Pong, error)
	Flock(ctx context.Context, opts ...grpc.CallOption) (Flock_FlockClient, error)
}

type flockClient struct {
	cc *grpc.ClientConn
}

func NewFlockClient(cc *grpc.ClientConn) FlockClient {
	return &flockClient{cc}
}

func (c *flockClient) Health(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/flock.Flock/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flockClient) DatabaseHealth(ctx context.Context, in *DBPing, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/flock.Flock/DatabaseHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flockClient) Flock(ctx context.Context, opts ...grpc.CallOption) (Flock_FlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Flock_serviceDesc.Streams[0], "/flock.Flock/Flock", opts...)
	if err != nil {
		return nil, err
	}
	x := &flockFlockClient{stream}
	return x, nil
}

type Flock_FlockClient interface {
	Send(*FlockRequest) error
	Recv() (*FlockResponse, error)
	grpc.ClientStream
}

type flockFlockClient struct {
	grpc.ClientStream
}

func (x *flockFlockClient) Send(m *FlockRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flockFlockClient) Recv() (*FlockResponse, error) {
	m := new(FlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlockServer is the server API for Flock service.
type FlockServer interface {
	Health(context.Context, *Ping) (*Pong, error)
	DatabaseHealth(context.Context, *DBPing) (*Pong, error)
	Flock(Flock_FlockServer) error
}

// UnimplementedFlockServer can be embedded to have forward compatible implementations.
type UnimplementedFlockServer struct {
}

func (*UnimplementedFlockServer) Health(ctx context.Context, req *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedFlockServer) DatabaseHealth(ctx context.Context, req *DBPing) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseHealth not implemented")
}
func (*UnimplementedFlockServer) Flock(srv Flock_FlockServer) error {
	return status.Errorf(codes.Unimplemented, "method Flock not implemented")
}

func RegisterFlockServer(s *grpc.Server, srv FlockServer) {
	s.RegisterService(&_Flock_serviceDesc, srv)
}

func _Flock_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlockServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flock.Flock/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlockServer).Health(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flock_DatabaseHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlockServer).DatabaseHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flock.Flock/DatabaseHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlockServer).DatabaseHealth(ctx, req.(*DBPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flock_Flock_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlockServer).Flock(&flockFlockServer{stream})
}

type Flock_FlockServer interface {
	Send(*FlockResponse) error
	Recv() (*FlockRequest, error)
	grpc.ServerStream
}

type flockFlockServer struct {
	grpc.ServerStream
}

func (x *flockFlockServer) Send(m *FlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flockFlockServer) Recv() (*FlockRequest, error) {
	m := new(FlockRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Flock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flock.Flock",
	HandlerType: (*FlockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Flock_Health_Handler,
		},
		{
			MethodName: "DatabaseHealth",
			Handler:    _Flock_DatabaseHealth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Flock",
			Handler:       _Flock_Flock_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "flock.proto",
}
