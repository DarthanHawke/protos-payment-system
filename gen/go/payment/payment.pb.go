// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: payment/payment.proto

package pmtstmv1

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

type CreatePaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Amount        float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	mi := &file_payment_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreatePaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreatePaymentRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	mi := &file_payment_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaymentResponse) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type GetPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentRequest) Reset() {
	*x = GetPaymentRequest{}
	mi := &file_payment_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentRequest) ProtoMessage() {}

func (x *GetPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GetPaymentRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type GetPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payment       *Payment               `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentResponse) Reset() {
	*x = GetPaymentResponse{}
	mi := &file_payment_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentResponse) ProtoMessage() {}

func (x *GetPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type UpdateStatusPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStatusPaymentRequest) Reset() {
	*x = UpdateStatusPaymentRequest{}
	mi := &file_payment_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusPaymentRequest) ProtoMessage() {}

func (x *UpdateStatusPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusPaymentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStatusPaymentRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *UpdateStatusPaymentRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateStatusPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStatusPaymentResponse) Reset() {
	*x = UpdateStatusPaymentResponse{}
	mi := &file_payment_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusPaymentResponse) ProtoMessage() {}

func (x *UpdateStatusPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusPaymentResponse.ProtoReflect.Descriptor instead.
func (*UpdateStatusPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{5}
}

type CancelPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelPaymentRequest) Reset() {
	*x = CancelPaymentRequest{}
	mi := &file_payment_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentRequest) ProtoMessage() {}

func (x *CancelPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentRequest.ProtoReflect.Descriptor instead.
func (*CancelPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{6}
}

func (x *CancelPaymentRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type CancelPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelPaymentResponse) Reset() {
	*x = CancelPaymentResponse{}
	mi := &file_payment_payment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentResponse) ProtoMessage() {}

func (x *CancelPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentResponse.ProtoReflect.Descriptor instead.
func (*CancelPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{7}
}

var File_payment_payment_proto protoreflect.FileDescriptor

const file_payment_payment_proto_rawDesc = "" +
	"\n" +
	"\x15payment/payment.proto\x12\apayment\x1a\x14payment/shared.proto\"\x81\x01\n" +
	"\x14CreatePaymentRequest\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x02R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12%\n" +
	"\vdescription\x18\x04 \x01(\tH\x00R\vdescription\x88\x01\x01B\x0e\n" +
	"\f_description\"6\n" +
	"\x15CreatePaymentResponse\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\tR\tpaymentId\"2\n" +
	"\x11GetPaymentRequest\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\tR\tpaymentId\"@\n" +
	"\x12GetPaymentResponse\x12*\n" +
	"\apayment\x18\x01 \x01(\v2\x10.payment.PaymentR\apayment\"S\n" +
	"\x1aUpdateStatusPaymentRequest\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\tR\tpaymentId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"\x1d\n" +
	"\x1bUpdateStatusPaymentResponse\"5\n" +
	"\x14CancelPaymentRequest\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\tR\tpaymentId\"\x17\n" +
	"\x15CancelPaymentResponse2\xd9\x02\n" +
	"\x0ePaymentService\x12N\n" +
	"\rCreatePayment\x12\x1d.payment.CreatePaymentRequest\x1a\x1e.payment.CreatePaymentResponse\x12E\n" +
	"\n" +
	"GetPayment\x12\x1a.payment.GetPaymentRequest\x1a\x1b.payment.GetPaymentResponse\x12`\n" +
	"\x13UpdateStatusPayment\x12#.payment.UpdateStatusPaymentRequest\x1a$.payment.UpdateStatusPaymentResponse\x12N\n" +
	"\rCancelPayment\x12\x1d.payment.CancelPaymentRequest\x1a\x1e.payment.CancelPaymentResponseB\x14Z\x12pmtstm.v1;pmtstmv1b\x06proto3"

var (
	file_payment_payment_proto_rawDescOnce sync.Once
	file_payment_payment_proto_rawDescData []byte
)

func file_payment_payment_proto_rawDescGZIP() []byte {
	file_payment_payment_proto_rawDescOnce.Do(func() {
		file_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payment_payment_proto_rawDesc), len(file_payment_payment_proto_rawDesc)))
	})
	return file_payment_payment_proto_rawDescData
}

var file_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_payment_payment_proto_goTypes = []any{
	(*CreatePaymentRequest)(nil),        // 0: payment.CreatePaymentRequest
	(*CreatePaymentResponse)(nil),       // 1: payment.CreatePaymentResponse
	(*GetPaymentRequest)(nil),           // 2: payment.GetPaymentRequest
	(*GetPaymentResponse)(nil),          // 3: payment.GetPaymentResponse
	(*UpdateStatusPaymentRequest)(nil),  // 4: payment.UpdateStatusPaymentRequest
	(*UpdateStatusPaymentResponse)(nil), // 5: payment.UpdateStatusPaymentResponse
	(*CancelPaymentRequest)(nil),        // 6: payment.CancelPaymentRequest
	(*CancelPaymentResponse)(nil),       // 7: payment.CancelPaymentResponse
	(*Payment)(nil),                     // 8: payment.Payment
}
var file_payment_payment_proto_depIdxs = []int32{
	8, // 0: payment.GetPaymentResponse.payment:type_name -> payment.Payment
	0, // 1: payment.PaymentService.CreatePayment:input_type -> payment.CreatePaymentRequest
	2, // 2: payment.PaymentService.GetPayment:input_type -> payment.GetPaymentRequest
	4, // 3: payment.PaymentService.UpdateStatusPayment:input_type -> payment.UpdateStatusPaymentRequest
	6, // 4: payment.PaymentService.CancelPayment:input_type -> payment.CancelPaymentRequest
	1, // 5: payment.PaymentService.CreatePayment:output_type -> payment.CreatePaymentResponse
	3, // 6: payment.PaymentService.GetPayment:output_type -> payment.GetPaymentResponse
	5, // 7: payment.PaymentService.UpdateStatusPayment:output_type -> payment.UpdateStatusPaymentResponse
	7, // 8: payment.PaymentService.CancelPayment:output_type -> payment.CancelPaymentResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_payment_proto_init() }
func file_payment_payment_proto_init() {
	if File_payment_payment_proto != nil {
		return
	}
	file_payment_shared_proto_init()
	file_payment_payment_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payment_payment_proto_rawDesc), len(file_payment_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_payment_proto_goTypes,
		DependencyIndexes: file_payment_payment_proto_depIdxs,
		MessageInfos:      file_payment_payment_proto_msgTypes,
	}.Build()
	File_payment_payment_proto = out.File
	file_payment_payment_proto_goTypes = nil
	file_payment_payment_proto_depIdxs = nil
}
