// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: bufbuild/example/asset/v1/asset.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mtyrrell/buf-example/pkg/opt/v1"
	v1 "github.com/mtyrrell/buf-example/pkg/semver/v1"
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

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version    *v1.Version          `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Data       []byte               `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetVersion() *v1.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Asset) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Asset) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type CreateAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *CreateAssetRequest) Reset() {
	*x = CreateAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetRequest) ProtoMessage() {}

func (x *CreateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetRequest.ProtoReflect.Descriptor instead.
func (*CreateAssetRequest) Descriptor() ([]byte, []int) {
	return file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAssetRequest) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

type CreateAssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateAssetResponse) Reset() {
	*x = CreateAssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetResponse) ProtoMessage() {}

func (x *CreateAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetResponse.ProtoReflect.Descriptor instead.
func (*CreateAssetResponse) Descriptor() ([]byte, []int) {
	return file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAssetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAssetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssetByIDRequest) Reset() {
	*x = GetAssetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetByIDRequest) ProtoMessage() {}

func (x *GetAssetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetAssetByIDRequest) Descriptor() ([]byte, []int) {
	return file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAssetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *GetAssetByIDResponse) Reset() {
	*x = GetAssetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetByIDResponse) ProtoMessage() {}

func (x *GetAssetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_example_asset_v1_asset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetByIDResponse.ProtoReflect.Descriptor instead.
func (*GetAssetByIDResponse) Descriptor() ([]byte, []int) {
	return file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssetByIDResponse) GetAsset() *Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

var File_bufbuild_example_asset_v1_asset_proto protoreflect.FileDescriptor

var file_bufbuild_example_asset_v1_asset_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc1, 0x01, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xd0, 0xe5, 0x1f, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x32,
	0xf1, 0x01, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x2d, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x74, 0x79, 0x72, 0x72, 0x65, 0x6c, 0x6c, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bufbuild_example_asset_v1_asset_proto_rawDescOnce sync.Once
	file_bufbuild_example_asset_v1_asset_proto_rawDescData = file_bufbuild_example_asset_v1_asset_proto_rawDesc
)

func file_bufbuild_example_asset_v1_asset_proto_rawDescGZIP() []byte {
	file_bufbuild_example_asset_v1_asset_proto_rawDescOnce.Do(func() {
		file_bufbuild_example_asset_v1_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_bufbuild_example_asset_v1_asset_proto_rawDescData)
	})
	return file_bufbuild_example_asset_v1_asset_proto_rawDescData
}

var file_bufbuild_example_asset_v1_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bufbuild_example_asset_v1_asset_proto_goTypes = []interface{}{
	(*Asset)(nil),                // 0: bufbuild.example.asset.v1.Asset
	(*CreateAssetRequest)(nil),   // 1: bufbuild.example.asset.v1.CreateAssetRequest
	(*CreateAssetResponse)(nil),  // 2: bufbuild.example.asset.v1.CreateAssetResponse
	(*GetAssetByIDRequest)(nil),  // 3: bufbuild.example.asset.v1.GetAssetByIDRequest
	(*GetAssetByIDResponse)(nil), // 4: bufbuild.example.asset.v1.GetAssetByIDResponse
	(*v1.Version)(nil),           // 5: bufbuild.example.semver.v1.Version
	(*timestamp.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_bufbuild_example_asset_v1_asset_proto_depIdxs = []int32{
	5, // 0: bufbuild.example.asset.v1.Asset.version:type_name -> bufbuild.example.semver.v1.Version
	6, // 1: bufbuild.example.asset.v1.Asset.create_time:type_name -> google.protobuf.Timestamp
	0, // 2: bufbuild.example.asset.v1.CreateAssetRequest.asset:type_name -> bufbuild.example.asset.v1.Asset
	0, // 3: bufbuild.example.asset.v1.GetAssetByIDResponse.asset:type_name -> bufbuild.example.asset.v1.Asset
	1, // 4: bufbuild.example.asset.v1.AssetService.CreateAsset:input_type -> bufbuild.example.asset.v1.CreateAssetRequest
	3, // 5: bufbuild.example.asset.v1.AssetService.GetAssetByID:input_type -> bufbuild.example.asset.v1.GetAssetByIDRequest
	2, // 6: bufbuild.example.asset.v1.AssetService.CreateAsset:output_type -> bufbuild.example.asset.v1.CreateAssetResponse
	4, // 7: bufbuild.example.asset.v1.AssetService.GetAssetByID:output_type -> bufbuild.example.asset.v1.GetAssetByIDResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bufbuild_example_asset_v1_asset_proto_init() }
func file_bufbuild_example_asset_v1_asset_proto_init() {
	if File_bufbuild_example_asset_v1_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bufbuild_example_asset_v1_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_bufbuild_example_asset_v1_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssetRequest); i {
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
		file_bufbuild_example_asset_v1_asset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssetResponse); i {
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
		file_bufbuild_example_asset_v1_asset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetByIDRequest); i {
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
		file_bufbuild_example_asset_v1_asset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetByIDResponse); i {
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
			RawDescriptor: file_bufbuild_example_asset_v1_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bufbuild_example_asset_v1_asset_proto_goTypes,
		DependencyIndexes: file_bufbuild_example_asset_v1_asset_proto_depIdxs,
		MessageInfos:      file_bufbuild_example_asset_v1_asset_proto_msgTypes,
	}.Build()
	File_bufbuild_example_asset_v1_asset_proto = out.File
	file_bufbuild_example_asset_v1_asset_proto_rawDesc = nil
	file_bufbuild_example_asset_v1_asset_proto_goTypes = nil
	file_bufbuild_example_asset_v1_asset_proto_depIdxs = nil
}
