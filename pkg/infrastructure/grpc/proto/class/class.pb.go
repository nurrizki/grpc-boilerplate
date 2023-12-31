// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: class.proto

package miniProject

import (
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

type ClassMP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId          int32  `protobuf:"varint,1,opt,name=ClassId,proto3" json:"ClassId,omitempty"`
	ClassName        string `protobuf:"bytes,2,opt,name=ClassName,proto3" json:"ClassName,omitempty"`
	ClassDescription string `protobuf:"bytes,3,opt,name=ClassDescription,proto3" json:"ClassDescription,omitempty"`
	Level            string `protobuf:"bytes,4,opt,name=Level,proto3" json:"Level,omitempty"`
	Status           string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ClassMP) Reset() {
	*x = ClassMP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassMP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassMP) ProtoMessage() {}

func (x *ClassMP) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassMP.ProtoReflect.Descriptor instead.
func (*ClassMP) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{0}
}

func (x *ClassMP) GetClassId() int32 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *ClassMP) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *ClassMP) GetClassDescription() string {
	if x != nil {
		return x.ClassDescription
	}
	return ""
}

func (x *ClassMP) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *ClassMP) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ClassDistinct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName string `protobuf:"bytes,1,opt,name=ClassName,proto3" json:"ClassName,omitempty"`
}

func (x *ClassDistinct) Reset() {
	*x = ClassDistinct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassDistinct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassDistinct) ProtoMessage() {}

func (x *ClassDistinct) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassDistinct.ProtoReflect.Descriptor instead.
func (*ClassDistinct) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{1}
}

func (x *ClassDistinct) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type GetClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ClassMP `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetClassResponse) Reset() {
	*x = GetClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassResponse) ProtoMessage() {}

func (x *GetClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassResponse.ProtoReflect.Descriptor instead.
func (*GetClassResponse) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{2}
}

func (x *GetClassResponse) GetData() []*ClassMP {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetClassDistinctResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ClassDistinct `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetClassDistinctResponse) Reset() {
	*x = GetClassDistinctResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassDistinctResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassDistinctResponse) ProtoMessage() {}

func (x *GetClassDistinctResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassDistinctResponse.ProtoReflect.Descriptor instead.
func (*GetClassDistinctResponse) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{3}
}

func (x *GetClassDistinctResponse) GetData() []*ClassDistinct {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetClassDistinctRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetClassDistinctRequest) Reset() {
	*x = GetClassDistinctRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassDistinctRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassDistinctRequest) ProtoMessage() {}

func (x *GetClassDistinctRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassDistinctRequest.ProtoReflect.Descriptor instead.
func (*GetClassDistinctRequest) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{4}
}

func (x *GetClassDistinctRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName string `protobuf:"bytes,1,opt,name=ClassName,proto3" json:"ClassName,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Level     string `protobuf:"bytes,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *GetClassRequest) Reset() {
	*x = GetClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassRequest) ProtoMessage() {}

func (x *GetClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassRequest.ProtoReflect.Descriptor instead.
func (*GetClassRequest) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{5}
}

func (x *GetClassRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *GetClassRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetClassRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

var File_class_proto protoreflect.FileDescriptor

var file_class_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d,
	0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x4d, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x50, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x32, 0xb8, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63,
	0x74, 0x12, 0x24, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x44, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_class_proto_rawDescOnce sync.Once
	file_class_proto_rawDescData = file_class_proto_rawDesc
)

func file_class_proto_rawDescGZIP() []byte {
	file_class_proto_rawDescOnce.Do(func() {
		file_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_class_proto_rawDescData)
	})
	return file_class_proto_rawDescData
}

var file_class_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_class_proto_goTypes = []interface{}{
	(*ClassMP)(nil),                  // 0: miniProject.ClassMP
	(*ClassDistinct)(nil),            // 1: miniProject.ClassDistinct
	(*GetClassResponse)(nil),         // 2: miniProject.GetClassResponse
	(*GetClassDistinctResponse)(nil), // 3: miniProject.GetClassDistinctResponse
	(*GetClassDistinctRequest)(nil),  // 4: miniProject.GetClassDistinctRequest
	(*GetClassRequest)(nil),          // 5: miniProject.GetClassRequest
}
var file_class_proto_depIdxs = []int32{
	0, // 0: miniProject.GetClassResponse.Data:type_name -> miniProject.ClassMP
	1, // 1: miniProject.GetClassDistinctResponse.Data:type_name -> miniProject.ClassDistinct
	5, // 2: miniProject.ClassService.GetClass:input_type -> miniProject.GetClassRequest
	4, // 3: miniProject.ClassService.GetClassDistinct:input_type -> miniProject.GetClassDistinctRequest
	2, // 4: miniProject.ClassService.GetClass:output_type -> miniProject.GetClassResponse
	3, // 5: miniProject.ClassService.GetClassDistinct:output_type -> miniProject.GetClassDistinctResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_class_proto_init() }
func file_class_proto_init() {
	if File_class_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_class_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassMP); i {
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
		file_class_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassDistinct); i {
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
		file_class_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassResponse); i {
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
		file_class_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassDistinctResponse); i {
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
		file_class_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassDistinctRequest); i {
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
		file_class_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassRequest); i {
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
			RawDescriptor: file_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_class_proto_goTypes,
		DependencyIndexes: file_class_proto_depIdxs,
		MessageInfos:      file_class_proto_msgTypes,
	}.Build()
	File_class_proto = out.File
	file_class_proto_rawDesc = nil
	file_class_proto_goTypes = nil
	file_class_proto_depIdxs = nil
}
