// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: punch.proto

package pb

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type PunchInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PunchInRequest) Reset() {
	*x = PunchInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunchInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunchInRequest) ProtoMessage() {}

func (x *PunchInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunchInRequest.ProtoReflect.Descriptor instead.
func (*PunchInRequest) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{0}
}

func (x *PunchInRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PunchOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PunchOutRequest) Reset() {
	*x = PunchOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunchOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunchOutRequest) ProtoMessage() {}

func (x *PunchOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunchOutRequest.ProtoReflect.Descriptor instead.
func (*PunchOutRequest) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{1}
}

func (x *PunchOutRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PunchInRequestOtp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PunchInRequestOtp) Reset() {
	*x = PunchInRequestOtp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunchInRequestOtp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunchInRequestOtp) ProtoMessage() {}

func (x *PunchInRequestOtp) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunchInRequestOtp.ProtoReflect.Descriptor instead.
func (*PunchInRequestOtp) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{2}
}

func (x *PunchInRequestOtp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PunchInRequestOtp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetDutyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDutyRequest) Reset() {
	*x = GetDutyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDutyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDutyRequest) ProtoMessage() {}

func (x *GetDutyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDutyRequest.ProtoReflect.Descriptor instead.
func (*GetDutyRequest) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{3}
}

func (x *GetDutyRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDutyResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DutyType string `protobuf:"bytes,1,opt,name=dutyType,proto3" json:"dutyType,omitempty"`
	Time     string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetDutyResponce) Reset() {
	*x = GetDutyResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDutyResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDutyResponce) ProtoMessage() {}

func (x *GetDutyResponce) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDutyResponce.ProtoReflect.Descriptor instead.
func (*GetDutyResponce) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{4}
}

func (x *GetDutyResponce) GetDutyType() string {
	if x != nil {
		return x.DutyType
	}
	return ""
}

func (x *GetDutyResponce) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuscode int32     `protobuf:"varint,1,opt,name=statuscode,proto3" json:"statuscode,omitempty"`
	Message    string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Errors     *any1.Any `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
	Data       *any1.Any `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_punch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_punch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_punch_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetStatuscode() int32 {
	if x != nil {
		return x.Statuscode
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetErrors() *any1.Any {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Response) GetData() *any1.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_punch_proto protoreflect.FileDescriptor

var file_punch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x75, 0x6e, 0x63, 0x68, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0e, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x21, 0x0a, 0x0f, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x74, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xf0, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x12, 0x15, 0x2e,
	0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x74, 0x70, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x75,
	0x6e, 0x63, 0x68, 0x2e, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4f, 0x74, 0x70, 0x1a, 0x0f, 0x2e, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x50, 0x75, 0x6e, 0x63,
	0x68, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x50, 0x75, 0x6e,
	0x63, 0x68, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x75, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x75, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x70, 0x75, 0x6e,
	0x63, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x75,
	0x6e, 0x63, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_punch_proto_rawDescOnce sync.Once
	file_punch_proto_rawDescData = file_punch_proto_rawDesc
)

func file_punch_proto_rawDescGZIP() []byte {
	file_punch_proto_rawDescOnce.Do(func() {
		file_punch_proto_rawDescData = protoimpl.X.CompressGZIP(file_punch_proto_rawDescData)
	})
	return file_punch_proto_rawDescData
}

var file_punch_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_punch_proto_goTypes = []interface{}{
	(*PunchInRequest)(nil),    // 0: punch.PunchInRequest
	(*PunchOutRequest)(nil),   // 1: punch.PunchOutRequest
	(*PunchInRequestOtp)(nil), // 2: punch.PunchInRequestOtp
	(*GetDutyRequest)(nil),    // 3: punch.GetDutyRequest
	(*GetDutyResponce)(nil),   // 4: punch.GetDutyResponce
	(*Response)(nil),          // 5: punch.Response
	(*any1.Any)(nil),          // 6: google.protobuf.Any
}
var file_punch_proto_depIdxs = []int32{
	6, // 0: punch.Response.errors:type_name -> google.protobuf.Any
	6, // 1: punch.Response.data:type_name -> google.protobuf.Any
	0, // 2: punch.PunchService.PunchIn:input_type -> punch.PunchInRequest
	2, // 3: punch.PunchService.VerifyOtpPunchin:input_type -> punch.PunchInRequestOtp
	1, // 4: punch.PunchService.PunchOut:input_type -> punch.PunchOutRequest
	3, // 5: punch.PunchService.GetDuty:input_type -> punch.GetDutyRequest
	5, // 6: punch.PunchService.PunchIn:output_type -> punch.Response
	5, // 7: punch.PunchService.VerifyOtpPunchin:output_type -> punch.Response
	5, // 8: punch.PunchService.PunchOut:output_type -> punch.Response
	5, // 9: punch.PunchService.GetDuty:output_type -> punch.Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_punch_proto_init() }
func file_punch_proto_init() {
	if File_punch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_punch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunchInRequest); i {
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
		file_punch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunchOutRequest); i {
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
		file_punch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunchInRequestOtp); i {
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
		file_punch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDutyRequest); i {
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
		file_punch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDutyResponce); i {
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
		file_punch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_punch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_punch_proto_goTypes,
		DependencyIndexes: file_punch_proto_depIdxs,
		MessageInfos:      file_punch_proto_msgTypes,
	}.Build()
	File_punch_proto = out.File
	file_punch_proto_rawDesc = nil
	file_punch_proto_goTypes = nil
	file_punch_proto_depIdxs = nil
}
