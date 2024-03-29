// 定义文件级别的Json标签复制
// @Tag("gorm", "autoUpdateTime:nano")

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: api_demo.proto

package api

import (
	_ "github.com/go-home-admin/go-admin/generate/proto/common/http"
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

type ApiDemoHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @Tag("bson")
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" form:"id"`
}

func (x *ApiDemoHomeRequest) Reset() {
	*x = ApiDemoHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDemoHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDemoHomeRequest) ProtoMessage() {}

func (x *ApiDemoHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDemoHomeRequest.ProtoReflect.Descriptor instead.
func (*ApiDemoHomeRequest) Descriptor() ([]byte, []int) {
	return file_api_demo_proto_rawDescGZIP(), []int{0}
}

func (x *ApiDemoHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ApiDemoHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime string `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" form:"start_time"`
	Tip       string `protobuf:"bytes,2,opt,name=tip,proto3" json:"tip,omitempty" form:"tip"`
	// Types that are assignable to Data:
	//	*ApiDemoHomeResponse_Name
	Data isApiDemoHomeResponse_Data `protobuf_oneof:"data" form:"data"`
}

func (x *ApiDemoHomeResponse) Reset() {
	*x = ApiDemoHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDemoHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDemoHomeResponse) ProtoMessage() {}

func (x *ApiDemoHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDemoHomeResponse.ProtoReflect.Descriptor instead.
func (*ApiDemoHomeResponse) Descriptor() ([]byte, []int) {
	return file_api_demo_proto_rawDescGZIP(), []int{1}
}

func (x *ApiDemoHomeResponse) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ApiDemoHomeResponse) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (m *ApiDemoHomeResponse) GetData() isApiDemoHomeResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ApiDemoHomeResponse) GetName() string {
	if x, ok := x.GetData().(*ApiDemoHomeResponse_Name); ok {
		return x.Name
	}
	return ""
}

type isApiDemoHomeResponse_Data interface {
	isApiDemoHomeResponse_Data()
}

type ApiDemoHomeResponse_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof" form:"name"`
}

func (*ApiDemoHomeResponse_Name) isApiDemoHomeResponse_Data() {}

var File_api_demo_proto protoreflect.FileDescriptor

var file_api_demo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x44,
	0x65, 0x6d, 0x6f, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64,
	0x0a, 0x13, 0x41, 0x70, 0x69, 0x44, 0x65, 0x6d, 0x6f, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x5c, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x44, 0x65, 0x6d, 0x6f, 0x12,
	0x48, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70,
	0x69, 0x44, 0x65, 0x6d, 0x6f, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x6d, 0x6f, 0x48, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xb5, 0x18, 0x09,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x1a, 0x07, 0x82, 0xb5, 0x18, 0x03, 0x61,
	0x70, 0x69, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67,
	0x6f, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_demo_proto_rawDescOnce sync.Once
	file_api_demo_proto_rawDescData = file_api_demo_proto_rawDesc
)

func file_api_demo_proto_rawDescGZIP() []byte {
	file_api_demo_proto_rawDescOnce.Do(func() {
		file_api_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_demo_proto_rawDescData)
	})
	return file_api_demo_proto_rawDescData
}

var file_api_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_demo_proto_goTypes = []interface{}{
	(*ApiDemoHomeRequest)(nil),  // 0: api.ApiDemoHomeRequest
	(*ApiDemoHomeResponse)(nil), // 1: api.ApiDemoHomeResponse
}
var file_api_demo_proto_depIdxs = []int32{
	0, // 0: api.ApiDemo.Home:input_type -> api.ApiDemoHomeRequest
	1, // 1: api.ApiDemo.Home:output_type -> api.ApiDemoHomeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_demo_proto_init() }
func file_api_demo_proto_init() {
	if File_api_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDemoHomeRequest); i {
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
		file_api_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDemoHomeResponse); i {
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
	file_api_demo_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ApiDemoHomeResponse_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_demo_proto_goTypes,
		DependencyIndexes: file_api_demo_proto_depIdxs,
		MessageInfos:      file_api_demo_proto_msgTypes,
	}.Build()
	File_api_demo_proto = out.File
	file_api_demo_proto_rawDesc = nil
	file_api_demo_proto_goTypes = nil
	file_api_demo_proto_depIdxs = nil
}
