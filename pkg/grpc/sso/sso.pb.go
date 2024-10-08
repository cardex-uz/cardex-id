// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: api/service/sso.proto

package sso

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Avatar    string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_service_sso_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type ValidatePhoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ValidatePhoneReq) Reset() {
	*x = ValidatePhoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePhoneReq) ProtoMessage() {}

func (x *ValidatePhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePhoneReq.ProtoReflect.Descriptor instead.
func (*ValidatePhoneReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{1}
}

func (x *ValidatePhoneReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ValidatePhoneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidatePhoneRes) Reset() {
	*x = ValidatePhoneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePhoneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePhoneRes) ProtoMessage() {}

func (x *ValidatePhoneRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePhoneRes.ProtoReflect.Descriptor instead.
func (*ValidatePhoneRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{2}
}

type ConfirmSubmitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ConfirmSubmitReq) Reset() {
	*x = ConfirmSubmitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmSubmitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSubmitReq) ProtoMessage() {}

func (x *ConfirmSubmitReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmSubmitReq.ProtoReflect.Descriptor instead.
func (*ConfirmSubmitReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{3}
}

func (x *ConfirmSubmitReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ConfirmSubmitRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfirmSubmitRes) Reset() {
	*x = ConfirmSubmitRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmSubmitRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSubmitRes) ProtoMessage() {}

func (x *ConfirmSubmitRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmSubmitRes.ProtoReflect.Descriptor instead.
func (*ConfirmSubmitRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{4}
}

type ConfirmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ConfirmReq) Reset() {
	*x = ConfirmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmReq) ProtoMessage() {}

func (x *ConfirmReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmReq.ProtoReflect.Descriptor instead.
func (*ConfirmReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ConfirmReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ConfirmRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	User    *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ConfirmRes) Reset() {
	*x = ConfirmRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRes) ProtoMessage() {}

func (x *ConfirmRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRes.ProtoReflect.Descriptor instead.
func (*ConfirmRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{6}
}

func (x *ConfirmRes) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *ConfirmRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FindAccountsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackId string `protobuf:"bytes,1,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *FindAccountsReq) Reset() {
	*x = FindAccountsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAccountsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAccountsReq) ProtoMessage() {}

func (x *FindAccountsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAccountsReq.ProtoReflect.Descriptor instead.
func (*FindAccountsReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{7}
}

func (x *FindAccountsReq) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type FindAccountsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAccountsRes) Reset() {
	*x = FindAccountsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAccountsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAccountsRes) ProtoMessage() {}

func (x *FindAccountsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAccountsRes.ProtoReflect.Descriptor instead.
func (*FindAccountsRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{8}
}

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{9}
}

func (x *SignInReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SignInRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignInRes) Reset() {
	*x = SignInRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRes) ProtoMessage() {}

func (x *SignInRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRes.ProtoReflect.Descriptor instead.
func (*SignInRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{10}
}

func (x *SignInRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SignInRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ValidateUsernameReq) Reset() {
	*x = ValidateUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUsernameReq) ProtoMessage() {}

func (x *ValidateUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUsernameReq.ProtoReflect.Descriptor instead.
func (*ValidateUsernameReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{11}
}

func (x *ValidateUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ValidateUsernameRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidateUsernameRes) Reset() {
	*x = ValidateUsernameRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUsernameRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUsernameRes) ProtoMessage() {}

func (x *ValidateUsernameRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUsernameRes.ProtoReflect.Descriptor instead.
func (*ValidateUsernameRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{12}
}

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{13}
}

func (x *SignUpReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SignUpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignUpRes) Reset() {
	*x = SignUpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRes) ProtoMessage() {}

func (x *SignUpRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRes.ProtoReflect.Descriptor instead.
func (*SignUpRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{14}
}

func (x *SignUpRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SignUpRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetMeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMeReq) Reset() {
	*x = GetMeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeReq) ProtoMessage() {}

func (x *GetMeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeReq.ProtoReflect.Descriptor instead.
func (*GetMeReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{15}
}

type GetMeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetMeRes) Reset() {
	*x = GetMeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeRes) ProtoMessage() {}

func (x *GetMeRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeRes.ProtoReflect.Descriptor instead.
func (*GetMeRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{16}
}

func (x *GetMeRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignOutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutReq) Reset() {
	*x = SignOutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutReq) ProtoMessage() {}

func (x *SignOutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutReq.ProtoReflect.Descriptor instead.
func (*SignOutReq) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{17}
}

type SignOutRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutRes) Reset() {
	*x = SignOutRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_sso_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRes) ProtoMessage() {}

func (x *SignOutRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_sso_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRes.ProtoReflect.Descriptor instead.
func (*SignOutRes) Descriptor() ([]byte, []int) {
	return file_api_service_sso_proto_rawDescGZIP(), []int{18}
}

var File_api_service_sso_proto protoreflect.FileDescriptor

var file_api_service_sso_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x73,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x73, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x28, 0x0a, 0x10, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x31, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65,
	0x71, 0x22, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0c, 0x0a, 0x0a,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x69,
	0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x32, 0xe0, 0x03, 0x0a, 0x03, 0x53, 0x53, 0x4f,
	0x12, 0x4a, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x15, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x0f, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x0e, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x12, 0x0d, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2b,
	0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x3b, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_service_sso_proto_rawDescOnce sync.Once
	file_api_service_sso_proto_rawDescData = file_api_service_sso_proto_rawDesc
)

func file_api_service_sso_proto_rawDescGZIP() []byte {
	file_api_service_sso_proto_rawDescOnce.Do(func() {
		file_api_service_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_sso_proto_rawDescData)
	})
	return file_api_service_sso_proto_rawDescData
}

var file_api_service_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_api_service_sso_proto_goTypes = []any{
	(*User)(nil),                // 0: sso.User
	(*ValidatePhoneReq)(nil),    // 1: sso.ValidatePhoneReq
	(*ValidatePhoneRes)(nil),    // 2: sso.ValidatePhoneRes
	(*ConfirmSubmitReq)(nil),    // 3: sso.ConfirmSubmitReq
	(*ConfirmSubmitRes)(nil),    // 4: sso.ConfirmSubmitRes
	(*ConfirmReq)(nil),          // 5: sso.ConfirmReq
	(*ConfirmRes)(nil),          // 6: sso.ConfirmRes
	(*FindAccountsReq)(nil),     // 7: sso.FindAccountsReq
	(*FindAccountsRes)(nil),     // 8: sso.FindAccountsRes
	(*SignInReq)(nil),           // 9: sso.SignInReq
	(*SignInRes)(nil),           // 10: sso.SignInRes
	(*ValidateUsernameReq)(nil), // 11: sso.ValidateUsernameReq
	(*ValidateUsernameRes)(nil), // 12: sso.ValidateUsernameRes
	(*SignUpReq)(nil),           // 13: sso.SignUpReq
	(*SignUpRes)(nil),           // 14: sso.SignUpRes
	(*GetMeReq)(nil),            // 15: sso.GetMeReq
	(*GetMeRes)(nil),            // 16: sso.GetMeRes
	(*SignOutReq)(nil),          // 17: sso.SignOutReq
	(*SignOutRes)(nil),          // 18: sso.SignOutRes
}
var file_api_service_sso_proto_depIdxs = []int32{
	0,  // 0: sso.ConfirmRes.user:type_name -> sso.User
	0,  // 1: sso.SignInRes.user:type_name -> sso.User
	0,  // 2: sso.SignUpRes.user:type_name -> sso.User
	0,  // 3: sso.GetMeRes.user:type_name -> sso.User
	11, // 4: sso.SSO.ValidateUsername:input_type -> sso.ValidateUsernameReq
	1,  // 5: sso.SSO.ValidatePhone:input_type -> sso.ValidatePhoneReq
	3,  // 6: sso.SSO.ConfirmSubmit:input_type -> sso.ConfirmSubmitReq
	5,  // 7: sso.SSO.Confirm:input_type -> sso.ConfirmReq
	7,  // 8: sso.SSO.FindAccounts:input_type -> sso.FindAccountsReq
	9,  // 9: sso.SSO.SignIn:input_type -> sso.SignInReq
	13, // 10: sso.SSO.SignUp:input_type -> sso.SignUpReq
	15, // 11: sso.SSO.GetMe:input_type -> sso.GetMeReq
	17, // 12: sso.SSO.SignOut:input_type -> sso.SignOutReq
	12, // 13: sso.SSO.ValidateUsername:output_type -> sso.ValidateUsernameRes
	2,  // 14: sso.SSO.ValidatePhone:output_type -> sso.ValidatePhoneRes
	4,  // 15: sso.SSO.ConfirmSubmit:output_type -> sso.ConfirmSubmitRes
	6,  // 16: sso.SSO.Confirm:output_type -> sso.ConfirmRes
	8,  // 17: sso.SSO.FindAccounts:output_type -> sso.FindAccountsRes
	10, // 18: sso.SSO.SignIn:output_type -> sso.SignInRes
	14, // 19: sso.SSO.SignUp:output_type -> sso.SignUpRes
	16, // 20: sso.SSO.GetMe:output_type -> sso.GetMeRes
	18, // 21: sso.SSO.SignOut:output_type -> sso.SignOutRes
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_service_sso_proto_init() }
func file_api_service_sso_proto_init() {
	if File_api_service_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_service_sso_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_api_service_sso_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ValidatePhoneReq); i {
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
		file_api_service_sso_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ValidatePhoneRes); i {
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
		file_api_service_sso_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmSubmitReq); i {
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
		file_api_service_sso_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmSubmitRes); i {
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
		file_api_service_sso_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmReq); i {
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
		file_api_service_sso_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmRes); i {
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
		file_api_service_sso_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FindAccountsReq); i {
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
		file_api_service_sso_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*FindAccountsRes); i {
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
		file_api_service_sso_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SignInReq); i {
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
		file_api_service_sso_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SignInRes); i {
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
		file_api_service_sso_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ValidateUsernameReq); i {
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
		file_api_service_sso_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ValidateUsernameRes); i {
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
		file_api_service_sso_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*SignUpReq); i {
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
		file_api_service_sso_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*SignUpRes); i {
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
		file_api_service_sso_proto_msgTypes[15].Exporter = func(v any, i int) any {
			switch v := v.(*GetMeReq); i {
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
		file_api_service_sso_proto_msgTypes[16].Exporter = func(v any, i int) any {
			switch v := v.(*GetMeRes); i {
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
		file_api_service_sso_proto_msgTypes[17].Exporter = func(v any, i int) any {
			switch v := v.(*SignOutReq); i {
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
		file_api_service_sso_proto_msgTypes[18].Exporter = func(v any, i int) any {
			switch v := v.(*SignOutRes); i {
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
			RawDescriptor: file_api_service_sso_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_sso_proto_goTypes,
		DependencyIndexes: file_api_service_sso_proto_depIdxs,
		MessageInfos:      file_api_service_sso_proto_msgTypes,
	}.Build()
	File_api_service_sso_proto = out.File
	file_api_service_sso_proto_rawDesc = nil
	file_api_service_sso_proto_goTypes = nil
	file_api_service_sso_proto_depIdxs = nil
}
