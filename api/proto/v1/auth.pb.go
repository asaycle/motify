// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: v1/auth.proto

package routiq

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExchangeOAuthCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeOAuthCodeRequest) Reset() {
	*x = ExchangeOAuthCodeRequest{}
	mi := &file_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeOAuthCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeOAuthCodeRequest) ProtoMessage() {}

func (x *ExchangeOAuthCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeOAuthCodeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeOAuthCodeRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeOAuthCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ExchangeOAuthCodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RedirectUrl   string                 `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	AccessToken   string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Role          string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeOAuthCodeResponse) Reset() {
	*x = ExchangeOAuthCodeResponse{}
	mi := &file_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeOAuthCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeOAuthCodeResponse) ProtoMessage() {}

func (x *ExchangeOAuthCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeOAuthCodeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeOAuthCodeResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeOAuthCodeResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *ExchangeOAuthCodeResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ExchangeOAuthCodeResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *ExchangeOAuthCodeResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`    // アクセストークン
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // リフレッシュトークン
	ExpiresIn     int64                  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`         // 有効期限 (秒)
	TokenType     string                 `protobuf:"bytes,4,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`          // トークンのタイプ (例: Bearer)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *RefreshTokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

type VerifyTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	mi := &file_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type VerifyTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenResponse) Reset() {
	*x = VerifyTokenResponse{}
	mi := &file_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenResponse) ProtoMessage() {}

func (x *VerifyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VerifyTokenResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_v1_auth_proto protoreflect.FileDescriptor

var file_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x18,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9a, 0x01, 0x0a,
	0x19, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a,
	0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x32, 0x9a, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x12, 0x7e, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x73,
	0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x12, 0x79, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x61, 0x79,
	0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x61,
	0x79, 0x63, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_auth_proto_rawDescOnce sync.Once
	file_v1_auth_proto_rawDescData = file_v1_auth_proto_rawDesc
)

func file_v1_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_auth_proto_rawDescData)
	})
	return file_v1_auth_proto_rawDescData
}

var file_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_auth_proto_goTypes = []any{
	(*ExchangeOAuthCodeRequest)(nil),  // 0: asaycle.routiq.v1.ExchangeOAuthCodeRequest
	(*ExchangeOAuthCodeResponse)(nil), // 1: asaycle.routiq.v1.ExchangeOAuthCodeResponse
	(*RefreshTokenRequest)(nil),       // 2: asaycle.routiq.v1.RefreshTokenRequest
	(*RefreshTokenResponse)(nil),      // 3: asaycle.routiq.v1.RefreshTokenResponse
	(*VerifyTokenRequest)(nil),        // 4: asaycle.routiq.v1.VerifyTokenRequest
	(*VerifyTokenResponse)(nil),       // 5: asaycle.routiq.v1.VerifyTokenResponse
}
var file_v1_auth_proto_depIdxs = []int32{
	0, // 0: asaycle.routiq.v1.AuthService.ExchangeOAuthCode:input_type -> asaycle.routiq.v1.ExchangeOAuthCodeRequest
	2, // 1: asaycle.routiq.v1.AuthService.RefreshToken:input_type -> asaycle.routiq.v1.RefreshTokenRequest
	4, // 2: asaycle.routiq.v1.AuthService.VerifyToken:input_type -> asaycle.routiq.v1.VerifyTokenRequest
	1, // 3: asaycle.routiq.v1.AuthService.ExchangeOAuthCode:output_type -> asaycle.routiq.v1.ExchangeOAuthCodeResponse
	3, // 4: asaycle.routiq.v1.AuthService.RefreshToken:output_type -> asaycle.routiq.v1.RefreshTokenResponse
	5, // 5: asaycle.routiq.v1.AuthService.VerifyToken:output_type -> asaycle.routiq.v1.VerifyTokenResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_auth_proto_init() }
func file_v1_auth_proto_init() {
	if File_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_proto_depIdxs,
		MessageInfos:      file_v1_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_proto = out.File
	file_v1_auth_proto_rawDesc = nil
	file_v1_auth_proto_goTypes = nil
	file_v1_auth_proto_depIdxs = nil
}
