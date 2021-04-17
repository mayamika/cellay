// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: cellay/v1/cellay.proto

package cellayv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GamesServiceGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GamesServiceGetResponse) Reset() {
	*x = GamesServiceGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cellay_v1_cellay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GamesServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GamesServiceGetResponse) ProtoMessage() {}

func (x *GamesServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cellay_v1_cellay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GamesServiceGetResponse.ProtoReflect.Descriptor instead.
func (*GamesServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_cellay_v1_cellay_proto_rawDescGZIP(), []int{0}
}

func (x *GamesServiceGetResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GamesServiceGetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GamesServiceGetResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GamesServiceGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GamesServiceGetAllRequest) Reset() {
	*x = GamesServiceGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cellay_v1_cellay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GamesServiceGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GamesServiceGetAllRequest) ProtoMessage() {}

func (x *GamesServiceGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cellay_v1_cellay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GamesServiceGetAllRequest.ProtoReflect.Descriptor instead.
func (*GamesServiceGetAllRequest) Descriptor() ([]byte, []int) {
	return file_cellay_v1_cellay_proto_rawDescGZIP(), []int{1}
}

type GamesServiceGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*GamesServiceGetResponse `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	Total int32                      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GamesServiceGetAllResponse) Reset() {
	*x = GamesServiceGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cellay_v1_cellay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GamesServiceGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GamesServiceGetAllResponse) ProtoMessage() {}

func (x *GamesServiceGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cellay_v1_cellay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GamesServiceGetAllResponse.ProtoReflect.Descriptor instead.
func (*GamesServiceGetAllResponse) Descriptor() ([]byte, []int) {
	return file_cellay_v1_cellay_proto_rawDescGZIP(), []int{2}
}

func (x *GamesServiceGetAllResponse) GetGames() []*GamesServiceGetResponse {
	if x != nil {
		return x.Games
	}
	return nil
}

func (x *GamesServiceGetAllResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type MatchesServiceStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId int32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *MatchesServiceStartRequest) Reset() {
	*x = MatchesServiceStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cellay_v1_cellay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchesServiceStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchesServiceStartRequest) ProtoMessage() {}

func (x *MatchesServiceStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cellay_v1_cellay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchesServiceStartRequest.ProtoReflect.Descriptor instead.
func (*MatchesServiceStartRequest) Descriptor() ([]byte, []int) {
	return file_cellay_v1_cellay_proto_rawDescGZIP(), []int{3}
}

func (x *MatchesServiceStartRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type MatchesServiceStartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MatchesServiceStartResponse) Reset() {
	*x = MatchesServiceStartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cellay_v1_cellay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchesServiceStartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchesServiceStartResponse) ProtoMessage() {}

func (x *MatchesServiceStartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cellay_v1_cellay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchesServiceStartResponse.ProtoReflect.Descriptor instead.
func (*MatchesServiceStartResponse) Descriptor() ([]byte, []int) {
	return file_cellay_v1_cellay_proto_rawDescGZIP(), []int{4}
}

func (x *MatchesServiceStartResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_cellay_v1_cellay_proto protoreflect.FileDescriptor

var file_cellay_v1_cellay_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x17, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6c, 0x0a, 0x1a, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x35, 0x0a,
	0x1a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x1b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x8f, 0x01, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x24,
	0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x92, 0x41, 0x17,
	0x1a, 0x15, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x98, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x92, 0x41, 0x11, 0x1a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x3a, 0x01, 0x2a,
	0x32, 0x0f, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x89, 0x02, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x61, 0x79, 0x61, 0x6d, 0x69, 0x6b, 0x61, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x76, 0x31, 0x92, 0x41, 0xd0, 0x01, 0x12, 0x0f, 0x0a,
	0x06, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x79, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x22, 0x07,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50,
	0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cellay_v1_cellay_proto_rawDescOnce sync.Once
	file_cellay_v1_cellay_proto_rawDescData = file_cellay_v1_cellay_proto_rawDesc
)

func file_cellay_v1_cellay_proto_rawDescGZIP() []byte {
	file_cellay_v1_cellay_proto_rawDescOnce.Do(func() {
		file_cellay_v1_cellay_proto_rawDescData = protoimpl.X.CompressGZIP(file_cellay_v1_cellay_proto_rawDescData)
	})
	return file_cellay_v1_cellay_proto_rawDescData
}

var file_cellay_v1_cellay_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cellay_v1_cellay_proto_goTypes = []interface{}{
	(*GamesServiceGetResponse)(nil),     // 0: cellay.v1.GamesServiceGetResponse
	(*GamesServiceGetAllRequest)(nil),   // 1: cellay.v1.GamesServiceGetAllRequest
	(*GamesServiceGetAllResponse)(nil),  // 2: cellay.v1.GamesServiceGetAllResponse
	(*MatchesServiceStartRequest)(nil),  // 3: cellay.v1.MatchesServiceStartRequest
	(*MatchesServiceStartResponse)(nil), // 4: cellay.v1.MatchesServiceStartResponse
}
var file_cellay_v1_cellay_proto_depIdxs = []int32{
	0, // 0: cellay.v1.GamesServiceGetAllResponse.games:type_name -> cellay.v1.GamesServiceGetResponse
	1, // 1: cellay.v1.GamesService.GetAll:input_type -> cellay.v1.GamesServiceGetAllRequest
	3, // 2: cellay.v1.MatchesService.Start:input_type -> cellay.v1.MatchesServiceStartRequest
	2, // 3: cellay.v1.GamesService.GetAll:output_type -> cellay.v1.GamesServiceGetAllResponse
	4, // 4: cellay.v1.MatchesService.Start:output_type -> cellay.v1.MatchesServiceStartResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cellay_v1_cellay_proto_init() }
func file_cellay_v1_cellay_proto_init() {
	if File_cellay_v1_cellay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cellay_v1_cellay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GamesServiceGetResponse); i {
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
		file_cellay_v1_cellay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GamesServiceGetAllRequest); i {
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
		file_cellay_v1_cellay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GamesServiceGetAllResponse); i {
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
		file_cellay_v1_cellay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchesServiceStartRequest); i {
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
		file_cellay_v1_cellay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchesServiceStartResponse); i {
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
			RawDescriptor: file_cellay_v1_cellay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cellay_v1_cellay_proto_goTypes,
		DependencyIndexes: file_cellay_v1_cellay_proto_depIdxs,
		MessageInfos:      file_cellay_v1_cellay_proto_msgTypes,
	}.Build()
	File_cellay_v1_cellay_proto = out.File
	file_cellay_v1_cellay_proto_rawDesc = nil
	file_cellay_v1_cellay_proto_goTypes = nil
	file_cellay_v1_cellay_proto_depIdxs = nil
}
