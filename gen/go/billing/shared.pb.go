// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: billing/shared.proto

package pmtstmv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type UUID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UUID) Reset() {
	*x = UUID{}
	mi := &file_billing_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UserSession struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSession) Reset() {
	*x = UserSession{}
	mi := &file_billing_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSession) ProtoMessage() {}

func (x *UserSession) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSession.ProtoReflect.Descriptor instead.
func (*UserSession) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{1}
}

func (x *UserSession) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserSession) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount        float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Description   *string                `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_billing_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Payment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Payment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_billing_shared_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *UUID                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_billing_shared_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{4}
}

func (x *Permission) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Relation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SourceId      *UUID                  `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargetId      *UUID                  `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	RelationType  string                 `protobuf:"bytes,3,opt,name=relation_type,json=relationType,proto3" json:"relation_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Relation) Reset() {
	*x = Relation{}
	mi := &file_billing_shared_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_billing_shared_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_billing_shared_proto_rawDescGZIP(), []int{5}
}

func (x *Relation) GetSourceId() *UUID {
	if x != nil {
		return x.SourceId
	}
	return nil
}

func (x *Relation) GetTargetId() *UUID {
	if x != nil {
		return x.TargetId
	}
	return nil
}

func (x *Relation) GetRelationType() string {
	if x != nil {
		return x.RelationType
	}
	return ""
}

var File_billing_shared_proto protoreflect.FileDescriptor

const file_billing_shared_proto_rawDesc = "" +
	"\n" +
	"\x14billing/shared.proto\x12\abilling\x1a\x1fgoogle/protobuf/timestamp.proto\"\x1c\n" +
	"\x04UUID\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\"U\n" +
	"\vUserSession\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\"\x92\x02\n" +
	"\aPayment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x02R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\x12%\n" +
	"\vdescription\x18\x05 \x01(\tH\x00R\vdescription\x88\x01\x01\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAtB\x0e\n" +
	"\f_description\"\xbd\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1b\n" +
	"\tfull_name\x18\x03 \x01(\tR\bfullName\x128\n" +
	"\tcreatedAt\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x128\n" +
	"\tupdatedAt\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"a\n" +
	"\n" +
	"Permission\x12\x1d\n" +
	"\x02id\x18\x01 \x01(\v2\r.billing.UUIDR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"\x87\x01\n" +
	"\bRelation\x12*\n" +
	"\tsource_id\x18\x01 \x01(\v2\r.billing.UUIDR\bsourceId\x12*\n" +
	"\ttarget_id\x18\x02 \x01(\v2\r.billing.UUIDR\btargetId\x12#\n" +
	"\rrelation_type\x18\x03 \x01(\tR\frelationTypeB\x14Z\x12pmtstm.v1;pmtstmv1b\x06proto3"

var (
	file_billing_shared_proto_rawDescOnce sync.Once
	file_billing_shared_proto_rawDescData []byte
)

func file_billing_shared_proto_rawDescGZIP() []byte {
	file_billing_shared_proto_rawDescOnce.Do(func() {
		file_billing_shared_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_billing_shared_proto_rawDesc), len(file_billing_shared_proto_rawDesc)))
	})
	return file_billing_shared_proto_rawDescData
}

var file_billing_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_billing_shared_proto_goTypes = []any{
	(*UUID)(nil),                  // 0: billing.UUID
	(*UserSession)(nil),           // 1: billing.UserSession
	(*Payment)(nil),               // 2: billing.Payment
	(*User)(nil),                  // 3: billing.User
	(*Permission)(nil),            // 4: billing.Permission
	(*Relation)(nil),              // 5: billing.Relation
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_billing_shared_proto_depIdxs = []int32{
	6, // 0: billing.Payment.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: billing.Payment.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: billing.User.createdAt:type_name -> google.protobuf.Timestamp
	6, // 3: billing.User.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 4: billing.Permission.id:type_name -> billing.UUID
	0, // 5: billing.Relation.source_id:type_name -> billing.UUID
	0, // 6: billing.Relation.target_id:type_name -> billing.UUID
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_billing_shared_proto_init() }
func file_billing_shared_proto_init() {
	if File_billing_shared_proto != nil {
		return
	}
	file_billing_shared_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_billing_shared_proto_rawDesc), len(file_billing_shared_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_billing_shared_proto_goTypes,
		DependencyIndexes: file_billing_shared_proto_depIdxs,
		MessageInfos:      file_billing_shared_proto_msgTypes,
	}.Build()
	File_billing_shared_proto = out.File
	file_billing_shared_proto_goTypes = nil
	file_billing_shared_proto_depIdxs = nil
}
