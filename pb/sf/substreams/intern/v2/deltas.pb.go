// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sf/substreams/intern/v2/deltas.proto

package pbssinternal

import (
	v1 "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ModuleOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// Types that are assignable to Data:
	//
	//	*ModuleOutput_MapOutput
	//	*ModuleOutput_StoreDeltas
	Data               isModuleOutput_Data `protobuf_oneof:"data"`
	Logs               []string            `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	DebugLogsTruncated bool                `protobuf:"varint,5,opt,name=debug_logs_truncated,json=debugLogsTruncated,proto3" json:"debug_logs_truncated,omitempty"`
	Cached             bool                `protobuf:"varint,6,opt,name=cached,proto3" json:"cached,omitempty"`
}

func (x *ModuleOutput) Reset() {
	*x = ModuleOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleOutput) ProtoMessage() {}

func (x *ModuleOutput) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleOutput.ProtoReflect.Descriptor instead.
func (*ModuleOutput) Descriptor() ([]byte, []int) {
	return file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP(), []int{0}
}

func (x *ModuleOutput) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (m *ModuleOutput) GetData() isModuleOutput_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ModuleOutput) GetMapOutput() *anypb.Any {
	if x, ok := x.GetData().(*ModuleOutput_MapOutput); ok {
		return x.MapOutput
	}
	return nil
}

func (x *ModuleOutput) GetStoreDeltas() *v1.StoreDeltas {
	if x, ok := x.GetData().(*ModuleOutput_StoreDeltas); ok {
		return x.StoreDeltas
	}
	return nil
}

func (x *ModuleOutput) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *ModuleOutput) GetDebugLogsTruncated() bool {
	if x != nil {
		return x.DebugLogsTruncated
	}
	return false
}

func (x *ModuleOutput) GetCached() bool {
	if x != nil {
		return x.Cached
	}
	return false
}

type isModuleOutput_Data interface {
	isModuleOutput_Data()
}

type ModuleOutput_MapOutput struct {
	MapOutput *anypb.Any `protobuf:"bytes,2,opt,name=map_output,json=mapOutput,proto3,oneof"`
}

type ModuleOutput_StoreDeltas struct {
	StoreDeltas *v1.StoreDeltas `protobuf:"bytes,3,opt,name=store_deltas,json=storeDeltas,proto3,oneof"`
}

func (*ModuleOutput_MapOutput) isModuleOutput_Data() {}

func (*ModuleOutput_StoreDeltas) isModuleOutput_Data() {}

var File_sf_substreams_intern_v2_deltas_proto protoreflect.FileDescriptor

var file_sf_substreams_intern_v2_deltas_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76,
	0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x66,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x6d, 0x61, 0x70, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x4d,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x32,
	0x3b, 0x70, 0x62, 0x73, 0x73, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_intern_v2_deltas_proto_rawDescOnce sync.Once
	file_sf_substreams_intern_v2_deltas_proto_rawDescData = file_sf_substreams_intern_v2_deltas_proto_rawDesc
)

func file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP() []byte {
	file_sf_substreams_intern_v2_deltas_proto_rawDescOnce.Do(func() {
		file_sf_substreams_intern_v2_deltas_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_intern_v2_deltas_proto_rawDescData)
	})
	return file_sf_substreams_intern_v2_deltas_proto_rawDescData
}

var file_sf_substreams_intern_v2_deltas_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sf_substreams_intern_v2_deltas_proto_goTypes = []interface{}{
	(*ModuleOutput)(nil),   // 0: sf.substreams.internal.v2.ModuleOutput
	(*anypb.Any)(nil),      // 1: google.protobuf.Any
	(*v1.StoreDeltas)(nil), // 2: sf.substreams.v1.StoreDeltas
}
var file_sf_substreams_intern_v2_deltas_proto_depIdxs = []int32{
	1, // 0: sf.substreams.internal.v2.ModuleOutput.map_output:type_name -> google.protobuf.Any
	2, // 1: sf.substreams.internal.v2.ModuleOutput.store_deltas:type_name -> sf.substreams.v1.StoreDeltas
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sf_substreams_intern_v2_deltas_proto_init() }
func file_sf_substreams_intern_v2_deltas_proto_init() {
	if File_sf_substreams_intern_v2_deltas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_intern_v2_deltas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleOutput); i {
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
	file_sf_substreams_intern_v2_deltas_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ModuleOutput_MapOutput)(nil),
		(*ModuleOutput_StoreDeltas)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_substreams_intern_v2_deltas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_intern_v2_deltas_proto_goTypes,
		DependencyIndexes: file_sf_substreams_intern_v2_deltas_proto_depIdxs,
		MessageInfos:      file_sf_substreams_intern_v2_deltas_proto_msgTypes,
	}.Build()
	File_sf_substreams_intern_v2_deltas_proto = out.File
	file_sf_substreams_intern_v2_deltas_proto_rawDesc = nil
	file_sf_substreams_intern_v2_deltas_proto_goTypes = nil
	file_sf_substreams_intern_v2_deltas_proto_depIdxs = nil
}
