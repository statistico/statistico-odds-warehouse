// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/grpc/proto/market.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RunnerFilter_Operator int32

const (
	RunnerFilter_GTE RunnerFilter_Operator = 0
	RunnerFilter_LTE RunnerFilter_Operator = 1
)

// Enum value maps for RunnerFilter_Operator.
var (
	RunnerFilter_Operator_name = map[int32]string{
		0: "GTE",
		1: "LTE",
	}
	RunnerFilter_Operator_value = map[string]int32{
		"GTE": 0,
		"LTE": 1,
	}
)

func (x RunnerFilter_Operator) Enum() *RunnerFilter_Operator {
	p := new(RunnerFilter_Operator)
	*p = x
	return p
}

func (x RunnerFilter_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunnerFilter_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_grpc_proto_market_proto_enumTypes[0].Descriptor()
}

func (RunnerFilter_Operator) Type() protoreflect.EnumType {
	return &file_internal_grpc_proto_market_proto_enumTypes[0]
}

func (x RunnerFilter_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunnerFilter_Operator.Descriptor instead.
func (RunnerFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{1, 0}
}

type RunnerFilter_Line int32

const (
	RunnerFilter_CLOSING RunnerFilter_Line = 0
	RunnerFilter_MAX     RunnerFilter_Line = 1
)

// Enum value maps for RunnerFilter_Line.
var (
	RunnerFilter_Line_name = map[int32]string{
		0: "CLOSING",
		1: "MAX",
	}
	RunnerFilter_Line_value = map[string]int32{
		"CLOSING": 0,
		"MAX":     1,
	}
)

func (x RunnerFilter_Line) Enum() *RunnerFilter_Line {
	p := new(RunnerFilter_Line)
	*p = x
	return p
}

func (x RunnerFilter_Line) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunnerFilter_Line) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_grpc_proto_market_proto_enumTypes[1].Descriptor()
}

func (RunnerFilter_Line) Type() protoreflect.EnumType {
	return &file_internal_grpc_proto_market_proto_enumTypes[1]
}

func (x RunnerFilter_Line) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunnerFilter_Line.Descriptor instead.
func (RunnerFilter_Line) EnumDescriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{1, 1}
}

type MarketSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Filters        []*RunnerFilter       `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	CompetitionIds []uint64              `protobuf:"varint,3,rep,packed,name=competition_ids,json=competitionIds,proto3" json:"competition_ids,omitempty"`
	SeasonIds      []uint64              `protobuf:"varint,4,rep,packed,name=season_ids,json=seasonIds,proto3" json:"season_ids,omitempty"`
	DateFrom       *wrappers.StringValue `protobuf:"bytes,5,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo         *wrappers.StringValue `protobuf:"bytes,6,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *MarketSearchRequest) Reset() {
	*x = MarketSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketSearchRequest) ProtoMessage() {}

func (x *MarketSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketSearchRequest.ProtoReflect.Descriptor instead.
func (*MarketSearchRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{0}
}

func (x *MarketSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MarketSearchRequest) GetFilters() []*RunnerFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *MarketSearchRequest) GetCompetitionIds() []uint64 {
	if x != nil {
		return x.CompetitionIds
	}
	return nil
}

func (x *MarketSearchRequest) GetSeasonIds() []uint64 {
	if x != nil {
		return x.SeasonIds
	}
	return nil
}

func (x *MarketSearchRequest) GetDateFrom() *wrappers.StringValue {
	if x != nil {
		return x.DateFrom
	}
	return nil
}

func (x *MarketSearchRequest) GetDateTo() *wrappers.StringValue {
	if x != nil {
		return x.DateTo
	}
	return nil
}

type RunnerFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Operator RunnerFilter_Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=proto.RunnerFilter_Operator" json:"operator,omitempty"`
	Value    float32               `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
	Line     RunnerFilter_Line     `protobuf:"varint,5,opt,name=line,proto3,enum=proto.RunnerFilter_Line" json:"line,omitempty"`
}

func (x *RunnerFilter) Reset() {
	*x = RunnerFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerFilter) ProtoMessage() {}

func (x *RunnerFilter) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerFilter.ProtoReflect.Descriptor instead.
func (*RunnerFilter) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{1}
}

func (x *RunnerFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunnerFilter) GetOperator() RunnerFilter_Operator {
	if x != nil {
		return x.Operator
	}
	return RunnerFilter_GTE
}

func (x *RunnerFilter) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *RunnerFilter) GetLine() RunnerFilter_Line {
	if x != nil {
		return x.Line
	}
	return RunnerFilter_CLOSING
}

type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_       uint64    `protobuf:"varint,1,opt,name=string,proto3" json:"string,omitempty"`
	Name          string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EventId       uint64    `protobuf:"varint,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	CompetitionId uint64    `protobuf:"varint,4,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	SeasonId      uint64    `protobuf:"varint,5,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	EventDate     string    `protobuf:"bytes,6,opt,name=event_date,json=eventDate,proto3" json:"event_date,omitempty"`
	Side          string    `protobuf:"bytes,7,opt,name=side,proto3" json:"side,omitempty"`
	Exchange      string    `protobuf:"bytes,8,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Runners       []*Runner `protobuf:"bytes,9,rep,name=runners,proto3" json:"runners,omitempty"`
	Timestamp     int64     `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

func (x *Market) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{2}
}

func (x *Market) GetString_() uint64 {
	if x != nil {
		return x.String_
	}
	return 0
}

func (x *Market) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Market) GetEventId() uint64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *Market) GetCompetitionId() uint64 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

func (x *Market) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *Market) GetEventDate() string {
	if x != nil {
		return x.EventDate
	}
	return ""
}

func (x *Market) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *Market) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *Market) GetRunners() []*Runner {
	if x != nil {
		return x.Runners
	}
	return nil
}

func (x *Market) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Runner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Size  float32 `protobuf:"fixed32,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Runner) Reset() {
	*x = Runner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runner) ProtoMessage() {}

func (x *Runner) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runner.ProtoReflect.Descriptor instead.
func (*Runner) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_market_proto_rawDescGZIP(), []int{3}
}

func (x *Runner) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Runner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Runner) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Runner) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_internal_grpc_proto_market_proto protoreflect.FileDescriptor

var file_internal_grpc_proto_market_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x13, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0xdc, 0x01, 0x0a,
	0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0x1c, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x54, 0x45, 0x10, 0x01, 0x22, 0x1c, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10, 0x01, 0x22, 0xa9, 0x02, 0x0a, 0x06,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x27, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x56, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32,
	0x4e, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x15, 0x5a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_proto_market_proto_rawDescOnce sync.Once
	file_internal_grpc_proto_market_proto_rawDescData = file_internal_grpc_proto_market_proto_rawDesc
)

func file_internal_grpc_proto_market_proto_rawDescGZIP() []byte {
	file_internal_grpc_proto_market_proto_rawDescOnce.Do(func() {
		file_internal_grpc_proto_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_proto_market_proto_rawDescData)
	})
	return file_internal_grpc_proto_market_proto_rawDescData
}

var file_internal_grpc_proto_market_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_grpc_proto_market_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_grpc_proto_market_proto_goTypes = []interface{}{
	(RunnerFilter_Operator)(0),   // 0: proto.RunnerFilter.Operator
	(RunnerFilter_Line)(0),       // 1: proto.RunnerFilter.Line
	(*MarketSearchRequest)(nil),  // 2: proto.MarketSearchRequest
	(*RunnerFilter)(nil),         // 3: proto.RunnerFilter
	(*Market)(nil),               // 4: proto.Market
	(*Runner)(nil),               // 5: proto.Runner
	(*wrappers.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_internal_grpc_proto_market_proto_depIdxs = []int32{
	3, // 0: proto.MarketSearchRequest.filters:type_name -> proto.RunnerFilter
	6, // 1: proto.MarketSearchRequest.dateFrom:type_name -> google.protobuf.StringValue
	6, // 2: proto.MarketSearchRequest.dateTo:type_name -> google.protobuf.StringValue
	0, // 3: proto.RunnerFilter.operator:type_name -> proto.RunnerFilter.Operator
	1, // 4: proto.RunnerFilter.line:type_name -> proto.RunnerFilter.Line
	5, // 5: proto.Market.runners:type_name -> proto.Runner
	2, // 6: proto.MarketService.MarketSearch:input_type -> proto.MarketSearchRequest
	4, // 7: proto.MarketService.MarketSearch:output_type -> proto.Market
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_grpc_proto_market_proto_init() }
func file_internal_grpc_proto_market_proto_init() {
	if File_internal_grpc_proto_market_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_proto_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketSearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_proto_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnerFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_proto_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_proto_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_grpc_proto_market_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_proto_market_proto_goTypes,
		DependencyIndexes: file_internal_grpc_proto_market_proto_depIdxs,
		EnumInfos:         file_internal_grpc_proto_market_proto_enumTypes,
		MessageInfos:      file_internal_grpc_proto_market_proto_msgTypes,
	}.Build()
	File_internal_grpc_proto_market_proto = out.File
	file_internal_grpc_proto_market_proto_rawDesc = nil
	file_internal_grpc_proto_market_proto_goTypes = nil
	file_internal_grpc_proto_market_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MarketServiceClient is the client API for MarketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketServiceClient interface {
	MarketSearch(ctx context.Context, in *MarketSearchRequest, opts ...grpc.CallOption) (MarketService_MarketSearchClient, error)
}

type marketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketServiceClient(cc grpc.ClientConnInterface) MarketServiceClient {
	return &marketServiceClient{cc}
}

func (c *marketServiceClient) MarketSearch(ctx context.Context, in *MarketSearchRequest, opts ...grpc.CallOption) (MarketService_MarketSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MarketService_serviceDesc.Streams[0], "/proto.MarketService/MarketSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketServiceMarketSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketService_MarketSearchClient interface {
	Recv() (*Market, error)
	grpc.ClientStream
}

type marketServiceMarketSearchClient struct {
	grpc.ClientStream
}

func (x *marketServiceMarketSearchClient) Recv() (*Market, error) {
	m := new(Market)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketServiceServer is the server API for MarketService service.
type MarketServiceServer interface {
	MarketSearch(*MarketSearchRequest, MarketService_MarketSearchServer) error
}

// UnimplementedMarketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMarketServiceServer struct {
}

func (*UnimplementedMarketServiceServer) MarketSearch(*MarketSearchRequest, MarketService_MarketSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method MarketSearch not implemented")
}

func RegisterMarketServiceServer(s *grpc.Server, srv MarketServiceServer) {
	s.RegisterService(&_MarketService_serviceDesc, srv)
}

func _MarketService_MarketSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MarketSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketServiceServer).MarketSearch(m, &marketServiceMarketSearchServer{stream})
}

type MarketService_MarketSearchServer interface {
	Send(*Market) error
	grpc.ServerStream
}

type marketServiceMarketSearchServer struct {
	grpc.ServerStream
}

func (x *marketServiceMarketSearchServer) Send(m *Market) error {
	return x.ServerStream.SendMsg(m)
}

var _MarketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MarketService",
	HandlerType: (*MarketServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MarketSearch",
			Handler:       _MarketService_MarketSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/grpc/proto/market.proto",
}