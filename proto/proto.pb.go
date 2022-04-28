// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/proto.proto

package proto

import (
	context "context"
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

type ResultStatus int32

const (
	ResultStatus_RESULT_OK     ResultStatus = 0
	ResultStatus_RESULT_FAILED ResultStatus = 1
)

// Enum value maps for ResultStatus.
var (
	ResultStatus_name = map[int32]string{
		0: "RESULT_OK",
		1: "RESULT_FAILED",
	}
	ResultStatus_value = map[string]int32{
		"RESULT_OK":     0,
		"RESULT_FAILED": 1,
	}
)

func (x ResultStatus) Enum() *ResultStatus {
	p := new(ResultStatus)
	*p = x
	return p
}

func (x ResultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_proto_proto_enumTypes[0].Descriptor()
}

func (ResultStatus) Type() protoreflect.EnumType {
	return &file_proto_proto_proto_enumTypes[0]
}

func (x ResultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultStatus.Descriptor instead.
func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

type ResultInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResultInt64) Reset() {
	*x = ResultInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultInt64) ProtoMessage() {}

func (x *ResultInt64) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultInt64.ProtoReflect.Descriptor instead.
func (*ResultInt64) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *ResultInt64) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResultInt64) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ResultFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResultFloat) Reset() {
	*x = ResultFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultFloat) ProtoMessage() {}

func (x *ResultFloat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultFloat.ProtoReflect.Descriptor instead.
func (*ResultFloat) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ResultFloat) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResultFloat) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ResultRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*ResultRow_Floatresult
	//	*ResultRow_Intresult
	Result isResultRow_Result `protobuf_oneof:"result"`
}

func (x *ResultRow) Reset() {
	*x = ResultRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRow) ProtoMessage() {}

func (x *ResultRow) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRow.ProtoReflect.Descriptor instead.
func (*ResultRow) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{2}
}

func (m *ResultRow) GetResult() isResultRow_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ResultRow) GetFloatresult() *ResultFloat {
	if x, ok := x.GetResult().(*ResultRow_Floatresult); ok {
		return x.Floatresult
	}
	return nil
}

func (x *ResultRow) GetIntresult() *ResultInt64 {
	if x, ok := x.GetResult().(*ResultRow_Intresult); ok {
		return x.Intresult
	}
	return nil
}

type isResultRow_Result interface {
	isResultRow_Result()
}

type ResultRow_Floatresult struct {
	Floatresult *ResultFloat `protobuf:"bytes,1,opt,name=floatresult,proto3,oneof"`
}

type ResultRow_Intresult struct {
	Intresult *ResultInt64 `protobuf:"bytes,2,opt,name=intresult,proto3,oneof"`
}

func (*ResultRow_Floatresult) isResultRow_Result() {}

func (*ResultRow_Intresult) isResultRow_Result() {}

type Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source     string       `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Host       string       `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Probe      string       `protobuf:"bytes,3,opt,name=probe,proto3" json:"probe,omitempty"`
	Time       uint64       `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Resultrows []*ResultRow `protobuf:"bytes,5,rep,name=resultrows,proto3" json:"resultrows,omitempty"`
}

func (x *Results) Reset() {
	*x = Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Results) ProtoMessage() {}

func (x *Results) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Results.ProtoReflect.Descriptor instead.
func (*Results) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{3}
}

func (x *Results) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Results) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Results) GetProbe() string {
	if x != nil {
		return x.Probe
	}
	return ""
}

func (x *Results) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Results) GetResultrows() []*ResultRow {
	if x != nil {
		return x.Resultrows
	}
	return nil
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ResultStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ResultStatus" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{4}
}

func (x *ResultResponse) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_RESULT_OK
}

func (x *ResultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_proto_proto protoreflect.FileDescriptor

var file_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x35, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x36, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x48,
	0x00, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x91, 0x01, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x6f, 0x77, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x77, 0x73,
	0x22, 0x57, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x30, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0x47, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x61, 0x6e, 0x6e, 0x74, 0x74, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x42,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_proto_rawDescOnce sync.Once
	file_proto_proto_proto_rawDescData = file_proto_proto_proto_rawDesc
)

func file_proto_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_proto_rawDescData)
	})
	return file_proto_proto_proto_rawDescData
}

var file_proto_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_proto_proto_goTypes = []interface{}{
	(ResultStatus)(0),      // 0: proto.ResultStatus
	(*ResultInt64)(nil),    // 1: proto.ResultInt64
	(*ResultFloat)(nil),    // 2: proto.ResultFloat
	(*ResultRow)(nil),      // 3: proto.ResultRow
	(*Results)(nil),        // 4: proto.Results
	(*ResultResponse)(nil), // 5: proto.ResultResponse
}
var file_proto_proto_proto_depIdxs = []int32{
	2, // 0: proto.ResultRow.floatresult:type_name -> proto.ResultFloat
	1, // 1: proto.ResultRow.intresult:type_name -> proto.ResultInt64
	3, // 2: proto.Results.resultrows:type_name -> proto.ResultRow
	0, // 3: proto.ResultResponse.status:type_name -> proto.ResultStatus
	4, // 4: proto.ResultService.SendResults:input_type -> proto.Results
	5, // 5: proto.ResultService.SendResults:output_type -> proto.ResultResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_proto_proto_init() }
func file_proto_proto_proto_init() {
	if File_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultInt64); i {
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
		file_proto_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultFloat); i {
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
		file_proto_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRow); i {
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
		file_proto_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Results); i {
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
		file_proto_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
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
	file_proto_proto_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ResultRow_Floatresult)(nil),
		(*ResultRow_Intresult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_proto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_proto_depIdxs,
		EnumInfos:         file_proto_proto_proto_enumTypes,
		MessageInfos:      file_proto_proto_proto_msgTypes,
	}.Build()
	File_proto_proto_proto = out.File
	file_proto_proto_proto_rawDesc = nil
	file_proto_proto_proto_goTypes = nil
	file_proto_proto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResultServiceClient is the client API for ResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResultServiceClient interface {
	SendResults(ctx context.Context, in *Results, opts ...grpc.CallOption) (*ResultResponse, error)
}

type resultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResultServiceClient(cc grpc.ClientConnInterface) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) SendResults(ctx context.Context, in *Results, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/proto.ResultService/SendResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultServiceServer is the server API for ResultService service.
type ResultServiceServer interface {
	SendResults(context.Context, *Results) (*ResultResponse, error)
}

// UnimplementedResultServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResultServiceServer struct {
}

func (*UnimplementedResultServiceServer) SendResults(context.Context, *Results) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResults not implemented")
}

func RegisterResultServiceServer(s *grpc.Server, srv ResultServiceServer) {
	s.RegisterService(&_ResultService_serviceDesc, srv)
}

func _ResultService_SendResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Results)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).SendResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ResultService/SendResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).SendResults(ctx, req.(*Results))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendResults",
			Handler:    _ResultService_SendResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}