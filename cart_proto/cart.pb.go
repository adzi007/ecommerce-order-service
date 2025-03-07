// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: cart.proto

package proto

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

type CartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CartRequest) Reset() {
	*x = CartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartRequest) ProtoMessage() {}

func (x *CartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartRequest.ProtoReflect.Descriptor instead.
func (*CartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*CartItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartResponse) GetData() []*CartItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId uint64           `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Name      string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Slug      string           `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Price     uint64           `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Qty       uint64           `protobuf:"varint,6,opt,name=qty,proto3" json:"qty,omitempty"`
	Category  *ProductCategory `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartItem) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartItem) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CartItem) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetQty() uint64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *CartItem) GetCategory() *ProductCategory {
	if x != nil {
		return x.Category
	}
	return nil
}

type ProductCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *ProductCategory) Reset() {
	*x = ProductCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategory) ProtoMessage() {}

func (x *ProductCategory) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategory.ProtoReflect.Descriptor instead.
func (*ProductCategory) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ProductCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCategory) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type DeleteCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCartResponse) Reset() {
	*x = DeleteCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartResponse) ProtoMessage() {}

func (x *DeleteCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteCartResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x25, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbb, 0x01, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x39, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x86, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x7a,
	0x69, 0x30, 0x30, 0x37, 0x2f, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2d, 0x63,
	0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cart_proto_goTypes = []interface{}{
	(*CartRequest)(nil),        // 0: cart.CartRequest
	(*UserRequest)(nil),        // 1: cart.UserRequest
	(*CartResponse)(nil),       // 2: cart.CartResponse
	(*CartItem)(nil),           // 3: cart.CartItem
	(*ProductCategory)(nil),    // 4: cart.ProductCategory
	(*DeleteCartResponse)(nil), // 5: cart.DeleteCartResponse
}
var file_cart_proto_depIdxs = []int32{
	3, // 0: cart.CartResponse.data:type_name -> cart.CartItem
	4, // 1: cart.CartItem.category:type_name -> cart.ProductCategory
	0, // 2: cart.CartService.GetCartUser:input_type -> cart.CartRequest
	1, // 3: cart.CartService.DeleteCartUser:input_type -> cart.UserRequest
	2, // 4: cart.CartService.GetCartUser:output_type -> cart.CartResponse
	5, // 5: cart.CartService.DeleteCartUser:output_type -> cart.DeleteCartResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartRequest); i {
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
		file_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartResponse); i {
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
		file_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategory); i {
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
		file_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartResponse); i {
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
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
