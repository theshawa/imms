// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: inventory.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PurchaseInventoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int64                  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Quantity      int64                  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Reference     string                 `protobuf:"bytes,3,opt,name=Reference,proto3" json:"Reference,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PurchaseInventoryRequest) Reset() {
	*x = PurchaseInventoryRequest{}
	mi := &file_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PurchaseInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseInventoryRequest) ProtoMessage() {}

func (x *PurchaseInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseInventoryRequest.ProtoReflect.Descriptor instead.
func (*PurchaseInventoryRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseInventoryRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *PurchaseInventoryRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PurchaseInventoryRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

type StockMovement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProductId     int64                  `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Change        int64                  `protobuf:"varint,3,opt,name=Change,proto3" json:"Change,omitempty"`
	Type          string                 `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Reference     string                 `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Note          string                 `protobuf:"bytes,6,opt,name=Note,proto3" json:"Note,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StockMovement) Reset() {
	*x = StockMovement{}
	mi := &file_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockMovement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockMovement) ProtoMessage() {}

func (x *StockMovement) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockMovement.ProtoReflect.Descriptor instead.
func (*StockMovement) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *StockMovement) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StockMovement) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *StockMovement) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *StockMovement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StockMovement) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *StockMovement) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *StockMovement) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ManageInventoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int64                  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Quantity      int64                  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Note          string                 `protobuf:"bytes,3,opt,name=Note,proto3" json:"Note,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManageInventoryRequest) Reset() {
	*x = ManageInventoryRequest{}
	mi := &file_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManageInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageInventoryRequest) ProtoMessage() {}

func (x *ManageInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageInventoryRequest.ProtoReflect.Descriptor instead.
func (*ManageInventoryRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *ManageInventoryRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ManageInventoryRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ManageInventoryRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type ListStockMovementsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int64                  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStockMovementsRequest) Reset() {
	*x = ListStockMovementsRequest{}
	mi := &file_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStockMovementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStockMovementsRequest) ProtoMessage() {}

func (x *ListStockMovementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStockMovementsRequest.ProtoReflect.Descriptor instead.
func (*ListStockMovementsRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *ListStockMovementsRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type ListStockMovementsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Records       []*StockMovement       `protobuf:"bytes,1,rep,name=Records,proto3" json:"Records,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStockMovementsResponse) Reset() {
	*x = ListStockMovementsResponse{}
	mi := &file_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStockMovementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStockMovementsResponse) ProtoMessage() {}

func (x *ListStockMovementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStockMovementsResponse.ProtoReflect.Descriptor instead.
func (*ListStockMovementsResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *ListStockMovementsResponse) GetRecords() []*StockMovement {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_inventory_proto protoreflect.FileDescriptor

const file_inventory_proto_rawDesc = "" +
	"\n" +
	"\x0finventory.proto\"r\n" +
	"\x18PurchaseInventoryRequest\x12\x1c\n" +
	"\tProductId\x18\x01 \x01(\x03R\tProductId\x12\x1a\n" +
	"\bQuantity\x18\x02 \x01(\x03R\bQuantity\x12\x1c\n" +
	"\tReference\x18\x03 \x01(\tR\tReference\"\xb9\x01\n" +
	"\rStockMovement\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\x03R\x02Id\x12\x1c\n" +
	"\tProductId\x18\x02 \x01(\x03R\tProductId\x12\x16\n" +
	"\x06Change\x18\x03 \x01(\x03R\x06Change\x12\x12\n" +
	"\x04Type\x18\x04 \x01(\tR\x04Type\x12\x1c\n" +
	"\tReference\x18\x05 \x01(\tR\tReference\x12\x12\n" +
	"\x04Note\x18\x06 \x01(\tR\x04Note\x12\x1c\n" +
	"\tCreatedAt\x18\a \x01(\tR\tCreatedAt\"f\n" +
	"\x16ManageInventoryRequest\x12\x1c\n" +
	"\tProductId\x18\x01 \x01(\x03R\tProductId\x12\x1a\n" +
	"\bQuantity\x18\x02 \x01(\x03R\bQuantity\x12\x12\n" +
	"\x04Note\x18\x03 \x01(\tR\x04Note\"9\n" +
	"\x19ListStockMovementsRequest\x12\x1c\n" +
	"\tProductId\x18\x01 \x01(\x03R\tProductId\"F\n" +
	"\x1aListStockMovementsResponse\x12(\n" +
	"\aRecords\x18\x01 \x03(\v2\x0e.StockMovementR\aRecords2\xad\x02\n" +
	"\x10InventoryService\x12E\n" +
	"\x18PurchaseInventoryProduct\x12\x19.PurchaseInventoryRequest\x1a\x0e.StockMovement\x12A\n" +
	"\x16SupplyInventoryProduct\x12\x17.ManageInventoryRequest\x1a\x0e.StockMovement\x12@\n" +
	"\x15CorrectInventoryStock\x12\x17.ManageInventoryRequest\x1a\x0e.StockMovement\x12M\n" +
	"\x12ListStockMovements\x12\x1a.ListStockMovementsRequest\x1a\x1b.ListStockMovementsResponseB3Z1github.com/logan2k02/ims/shared/protobuf;protobufb\x06proto3"

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData []byte
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_inventory_proto_rawDesc), len(file_inventory_proto_rawDesc)))
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_inventory_proto_goTypes = []any{
	(*PurchaseInventoryRequest)(nil),   // 0: PurchaseInventoryRequest
	(*StockMovement)(nil),              // 1: StockMovement
	(*ManageInventoryRequest)(nil),     // 2: ManageInventoryRequest
	(*ListStockMovementsRequest)(nil),  // 3: ListStockMovementsRequest
	(*ListStockMovementsResponse)(nil), // 4: ListStockMovementsResponse
}
var file_inventory_proto_depIdxs = []int32{
	1, // 0: ListStockMovementsResponse.Records:type_name -> StockMovement
	0, // 1: InventoryService.PurchaseInventoryProduct:input_type -> PurchaseInventoryRequest
	2, // 2: InventoryService.SupplyInventoryProduct:input_type -> ManageInventoryRequest
	2, // 3: InventoryService.CorrectInventoryStock:input_type -> ManageInventoryRequest
	3, // 4: InventoryService.ListStockMovements:input_type -> ListStockMovementsRequest
	1, // 5: InventoryService.PurchaseInventoryProduct:output_type -> StockMovement
	1, // 6: InventoryService.SupplyInventoryProduct:output_type -> StockMovement
	1, // 7: InventoryService.CorrectInventoryStock:output_type -> StockMovement
	4, // 8: InventoryService.ListStockMovements:output_type -> ListStockMovementsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_inventory_proto_rawDesc), len(file_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}
