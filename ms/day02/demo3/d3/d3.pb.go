// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: d3.proto

package d3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// wrapper
	Price  *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Price2 *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=price2,proto3" json:"price2,omitempty"`
	Foo    *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=foo,proto3" json:"foo,omitempty"`
	Info   *Book_Info              `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_d3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_d3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_d3_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Book) GetPrice2() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Price2
	}
	return nil
}

func (x *Book) GetFoo() *wrapperspb.StringValue {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *Book) GetInfo() *Book_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

// update book request message
type UpdateBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// operator
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// object of book
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	// 要更新的字段
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBook) Reset() {
	*x = UpdateBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_d3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBook) ProtoMessage() {}

func (x *UpdateBook) ProtoReflect() protoreflect.Message {
	mi := &file_d3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBook.ProtoReflect.Descriptor instead.
func (*UpdateBook) Descriptor() ([]byte, []int) {
	return file_d3_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBook) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdateBook) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBook) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Book_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Time string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_d3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_d3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_d3_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Book_Info) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Book_Info) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_d3_proto protoreflect.FileDescriptor

var file_d3_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x12, 0x2e, 0x0a, 0x03,
	0x66, 0x6f, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x1e, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x2e, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x64, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_d3_proto_rawDescOnce sync.Once
	file_d3_proto_rawDescData = file_d3_proto_rawDesc
)

func file_d3_proto_rawDescGZIP() []byte {
	file_d3_proto_rawDescOnce.Do(func() {
		file_d3_proto_rawDescData = protoimpl.X.CompressGZIP(file_d3_proto_rawDescData)
	})
	return file_d3_proto_rawDescData
}

var file_d3_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_d3_proto_goTypes = []any{
	(*Book)(nil),                   // 0: Book
	(*UpdateBook)(nil),             // 1: UpdateBook
	(*Book_Info)(nil),              // 2: Book.Info
	(*wrapperspb.Int64Value)(nil),  // 3: google.protobuf.Int64Value
	(*wrapperspb.DoubleValue)(nil), // 4: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*fieldmaskpb.FieldMask)(nil),  // 6: google.protobuf.FieldMask
}
var file_d3_proto_depIdxs = []int32{
	3, // 0: Book.price:type_name -> google.protobuf.Int64Value
	4, // 1: Book.price2:type_name -> google.protobuf.DoubleValue
	5, // 2: Book.foo:type_name -> google.protobuf.StringValue
	2, // 3: Book.info:type_name -> Book.Info
	0, // 4: UpdateBook.book:type_name -> Book
	6, // 5: UpdateBook.update_mask:type_name -> google.protobuf.FieldMask
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_d3_proto_init() }
func file_d3_proto_init() {
	if File_d3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_d3_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Book); i {
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
		file_d3_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBook); i {
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
		file_d3_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Book_Info); i {
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
			RawDescriptor: file_d3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_d3_proto_goTypes,
		DependencyIndexes: file_d3_proto_depIdxs,
		MessageInfos:      file_d3_proto_msgTypes,
	}.Build()
	File_d3_proto = out.File
	file_d3_proto_rawDesc = nil
	file_d3_proto_goTypes = nil
	file_d3_proto_depIdxs = nil
}
