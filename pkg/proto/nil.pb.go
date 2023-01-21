// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/proto/nil.proto

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Y1   int64  `protobuf:"varint,2,opt,name=y1,proto3" json:"y1,omitempty"`
	Y2   int64  `protobuf:"varint,3,opt,name=y2,proto3" json:"y2,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RegisterRequest) GetY1() int64 {
	if x != nil {
		return x.Y1
	}
	return 0
}

func (x *RegisterRequest) GetY2() int64 {
	if x != nil {
		return x.Y2
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{1}
}

type AuthenticationChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	R1   int64  `protobuf:"varint,2,opt,name=r1,proto3" json:"r1,omitempty"`
	R2   int64  `protobuf:"varint,3,opt,name=r2,proto3" json:"r2,omitempty"`
}

func (x *AuthenticationChallengeRequest) Reset() {
	*x = AuthenticationChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationChallengeRequest) ProtoMessage() {}

func (x *AuthenticationChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationChallengeRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationChallengeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationChallengeRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthenticationChallengeRequest) GetR1() int64 {
	if x != nil {
		return x.R1
	}
	return 0
}

func (x *AuthenticationChallengeRequest) GetR2() int64 {
	if x != nil {
		return x.R2
	}
	return 0
}

type AuthenticationChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	C      int64  `protobuf:"varint,2,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *AuthenticationChallengeResponse) Reset() {
	*x = AuthenticationChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationChallengeResponse) ProtoMessage() {}

func (x *AuthenticationChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationChallengeResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationChallengeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationChallengeResponse) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *AuthenticationChallengeResponse) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

type AuthenticationAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId string `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	S      int64  `protobuf:"varint,2,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *AuthenticationAnswerRequest) Reset() {
	*x = AuthenticationAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationAnswerRequest) ProtoMessage() {}

func (x *AuthenticationAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationAnswerRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationAnswerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{4}
}

func (x *AuthenticationAnswerRequest) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *AuthenticationAnswerRequest) GetS() int64 {
	if x != nil {
		return x.S
	}
	return 0
}

type AuthenticationAnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *AuthenticationAnswerResponse) Reset() {
	*x = AuthenticationAnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nil_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationAnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationAnswerResponse) ProtoMessage() {}

func (x *AuthenticationAnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nil_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationAnswerResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationAnswerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nil_proto_rawDescGZIP(), []int{5}
}

func (x *AuthenticationAnswerResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_pkg_proto_nil_proto protoreflect.FileDescriptor

var file_pkg_proto_nil_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x79, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x31, 0x12, 0x0e, 0x0a, 0x02,
	0x79, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x32, 0x22, 0x12, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x54, 0x0a, 0x1e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x72, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x32, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x72, 0x32, 0x22, 0x48, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x63,
	0x22, 0x44, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x73, 0x22, 0x3d, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xf6, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x64,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12,
	0x1f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x77, 0x6f,
	0x6c, 0x73, 0x6b, 0x69, 0x32, 0x2f, 0x6e, 0x69, 0x6c, 0x2d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nil_proto_rawDescOnce sync.Once
	file_pkg_proto_nil_proto_rawDescData = file_pkg_proto_nil_proto_rawDesc
)

func file_pkg_proto_nil_proto_rawDescGZIP() []byte {
	file_pkg_proto_nil_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nil_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nil_proto_rawDescData)
	})
	return file_pkg_proto_nil_proto_rawDescData
}

var file_pkg_proto_nil_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_nil_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),                 // 0: RegisterRequest
	(*RegisterResponse)(nil),                // 1: RegisterResponse
	(*AuthenticationChallengeRequest)(nil),  // 2: AuthenticationChallengeRequest
	(*AuthenticationChallengeResponse)(nil), // 3: AuthenticationChallengeResponse
	(*AuthenticationAnswerRequest)(nil),     // 4: AuthenticationAnswerRequest
	(*AuthenticationAnswerResponse)(nil),    // 5: AuthenticationAnswerResponse
}
var file_pkg_proto_nil_proto_depIdxs = []int32{
	2, // 0: Auth.CreateAuthenticationChallenge:input_type -> AuthenticationChallengeRequest
	0, // 1: Auth.Register:input_type -> RegisterRequest
	4, // 2: Auth.VerifyAuthentication:input_type -> AuthenticationAnswerRequest
	3, // 3: Auth.CreateAuthenticationChallenge:output_type -> AuthenticationChallengeResponse
	1, // 4: Auth.Register:output_type -> RegisterResponse
	5, // 5: Auth.VerifyAuthentication:output_type -> AuthenticationAnswerResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_nil_proto_init() }
func file_pkg_proto_nil_proto_init() {
	if File_pkg_proto_nil_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nil_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_pkg_proto_nil_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_pkg_proto_nil_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationChallengeRequest); i {
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
		file_pkg_proto_nil_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationChallengeResponse); i {
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
		file_pkg_proto_nil_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationAnswerRequest); i {
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
		file_pkg_proto_nil_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationAnswerResponse); i {
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
			RawDescriptor: file_pkg_proto_nil_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_nil_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nil_proto_depIdxs,
		MessageInfos:      file_pkg_proto_nil_proto_msgTypes,
	}.Build()
	File_pkg_proto_nil_proto = out.File
	file_pkg_proto_nil_proto_rawDesc = nil
	file_pkg_proto_nil_proto_goTypes = nil
	file_pkg_proto_nil_proto_depIdxs = nil
}
