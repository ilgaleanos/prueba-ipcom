// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: compra.proto

package src

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

//*
//
// protoc -I=. --go_out=. compra.proto
type Compra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int64   `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Phone    string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Nombre   string  `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Compro   bool    `protobuf:"varint,4,opt,name=compro,proto3" json:"compro,omitempty"`
	Tdc      string  `protobuf:"bytes,5,opt,name=tdc,proto3" json:"tdc,omitempty"`
	Monto    float64 `protobuf:"fixed64,6,opt,name=monto,proto3" json:"monto,omitempty"`
	Date     string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Compra) Reset() {
	*x = Compra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compra) ProtoMessage() {}

func (x *Compra) ProtoReflect() protoreflect.Message {
	mi := &file_compra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compra.ProtoReflect.Descriptor instead.
func (*Compra) Descriptor() ([]byte, []int) {
	return file_compra_proto_rawDescGZIP(), []int{0}
}

func (x *Compra) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Compra) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Compra) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Compra) GetCompro() bool {
	if x != nil {
		return x.Compro
	}
	return false
}

func (x *Compra) GetTdc() string {
	if x != nil {
		return x.Tdc
	}
	return ""
}

func (x *Compra) GetMonto() float64 {
	if x != nil {
		return x.Monto
	}
	return 0
}

func (x *Compra) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_compra_proto protoreflect.FileDescriptor

var file_compra_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x64, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x64, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2e, 0x2f,
	0x73, 0x72, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compra_proto_rawDescOnce sync.Once
	file_compra_proto_rawDescData = file_compra_proto_rawDesc
)

func file_compra_proto_rawDescGZIP() []byte {
	file_compra_proto_rawDescOnce.Do(func() {
		file_compra_proto_rawDescData = protoimpl.X.CompressGZIP(file_compra_proto_rawDescData)
	})
	return file_compra_proto_rawDescData
}

var file_compra_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compra_proto_goTypes = []interface{}{
	(*Compra)(nil), // 0: Compra
}
var file_compra_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_compra_proto_init() }
func file_compra_proto_init() {
	if File_compra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compra); i {
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
			RawDescriptor: file_compra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compra_proto_goTypes,
		DependencyIndexes: file_compra_proto_depIdxs,
		MessageInfos:      file_compra_proto_msgTypes,
	}.Build()
	File_compra_proto = out.File
	file_compra_proto_rawDesc = nil
	file_compra_proto_goTypes = nil
	file_compra_proto_depIdxs = nil
}
