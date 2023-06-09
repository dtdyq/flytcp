// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: test_res.proto

package pb

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

type GeneralQuality int32

const (
	GeneralQuality_level_1 GeneralQuality = 0
	GeneralQuality_level_2 GeneralQuality = 1
	GeneralQuality_level_3 GeneralQuality = 2
	GeneralQuality_level_4 GeneralQuality = 3
	GeneralQuality_level_5 GeneralQuality = 4
)

// Enum value maps for GeneralQuality.
var (
	GeneralQuality_name = map[int32]string{
		0: "level_1",
		1: "level_2",
		2: "level_3",
		3: "level_4",
		4: "level_5",
	}
	GeneralQuality_value = map[string]int32{
		"level_1": 0,
		"level_2": 1,
		"level_3": 2,
		"level_4": 3,
		"level_5": 4,
	}
)

func (x GeneralQuality) Enum() *GeneralQuality {
	p := new(GeneralQuality)
	*p = x
	return p
}

func (x GeneralQuality) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeneralQuality) Descriptor() protoreflect.EnumDescriptor {
	return file_test_res_proto_enumTypes[0].Descriptor()
}

func (GeneralQuality) Type() protoreflect.EnumType {
	return &file_test_res_proto_enumTypes[0]
}

func (x GeneralQuality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeneralQuality.Descriptor instead.
func (GeneralQuality) EnumDescriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{0}
}

type GeneralEffect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Val   float32          `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Param map[string]int32 `protobuf:"bytes,3,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GeneralEffect) Reset() {
	*x = GeneralEffect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralEffect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralEffect) ProtoMessage() {}

func (x *GeneralEffect) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralEffect.ProtoReflect.Descriptor instead.
func (*GeneralEffect) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralEffect) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GeneralEffect) GetVal() float32 {
	if x != nil {
		return x.Val
	}
	return 0
}

func (x *GeneralEffect) GetParam() map[string]int32 {
	if x != nil {
		return x.Param
	}
	return nil
}

type GeneralSkill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Effect  *GeneralEffect           `protobuf:"bytes,4,opt,name=effect,proto3" json:"effect,omitempty"`
	Effects []*GeneralEffect         `protobuf:"bytes,5,rep,name=effects,proto3" json:"effects,omitempty"`
	EffMap  map[int32]*GeneralEffect `protobuf:"bytes,6,rep,name=effMap,proto3" json:"effMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GeneralSkill) Reset() {
	*x = GeneralSkill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralSkill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralSkill) ProtoMessage() {}

func (x *GeneralSkill) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralSkill.ProtoReflect.Descriptor instead.
func (*GeneralSkill) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralSkill) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GeneralSkill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneralSkill) GetEffect() *GeneralEffect {
	if x != nil {
		return x.Effect
	}
	return nil
}

func (x *GeneralSkill) GetEffects() []*GeneralEffect {
	if x != nil {
		return x.Effects
	}
	return nil
}

func (x *GeneralSkill) GetEffMap() map[int32]*GeneralEffect {
	if x != nil {
		return x.EffMap
	}
	return nil
}

type GeneralConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quality  GeneralQuality          `protobuf:"varint,3,opt,name=quality,proto3,enum=model.GeneralQuality" json:"quality,omitempty"`
	FSkill   *GeneralSkill           `protobuf:"bytes,4,opt,name=fSkill,proto3" json:"fSkill,omitempty"`
	Favor    []string                `protobuf:"bytes,5,rep,name=favor,proto3" json:"favor,omitempty"`
	Skills   []*GeneralSkill         `protobuf:"bytes,6,rep,name=skills,proto3" json:"skills,omitempty"`
	Link     map[int32]string        `protobuf:"bytes,7,rep,name=link,proto3" json:"link,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SkillMap map[int32]*GeneralSkill `protobuf:"bytes,8,rep,name=skillMap,proto3" json:"skillMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GeneralConf) Reset() {
	*x = GeneralConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralConf) ProtoMessage() {}

func (x *GeneralConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralConf.ProtoReflect.Descriptor instead.
func (*GeneralConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{2}
}

func (x *GeneralConf) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GeneralConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneralConf) GetQuality() GeneralQuality {
	if x != nil {
		return x.Quality
	}
	return GeneralQuality_level_1
}

func (x *GeneralConf) GetFSkill() *GeneralSkill {
	if x != nil {
		return x.FSkill
	}
	return nil
}

func (x *GeneralConf) GetFavor() []string {
	if x != nil {
		return x.Favor
	}
	return nil
}

func (x *GeneralConf) GetSkills() []*GeneralSkill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *GeneralConf) GetLink() map[int32]string {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *GeneralConf) GetSkillMap() map[int32]*GeneralSkill {
	if x != nil {
		return x.SkillMap
	}
	return nil
}

type ResGeneralConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenConfMap map[int32]*GeneralConf `protobuf:"bytes,1,rep,name=genConfMap,proto3" json:"genConfMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResGeneralConf) Reset() {
	*x = ResGeneralConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGeneralConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGeneralConf) ProtoMessage() {}

func (x *ResGeneralConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGeneralConf.ProtoReflect.Descriptor instead.
func (*ResGeneralConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{3}
}

func (x *ResGeneralConf) GetGenConfMap() map[int32]*GeneralConf {
	if x != nil {
		return x.GenConfMap
	}
	return nil
}

type CommonConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StarMax    int32   `protobuf:"varint,2,opt,name=starMax,proto3" json:"starMax,omitempty"`
	SkillCount int32   `protobuf:"varint,3,opt,name=skillCount,proto3" json:"skillCount,omitempty"`
	HotPoint   float32 `protobuf:"fixed32,4,opt,name=hotPoint,proto3" json:"hotPoint,omitempty"`
	Template   string  `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *CommonConf) Reset() {
	*x = CommonConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonConf) ProtoMessage() {}

func (x *CommonConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonConf.ProtoReflect.Descriptor instead.
func (*CommonConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{4}
}

func (x *CommonConf) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommonConf) GetStarMax() int32 {
	if x != nil {
		return x.StarMax
	}
	return 0
}

func (x *CommonConf) GetSkillCount() int32 {
	if x != nil {
		return x.SkillCount
	}
	return 0
}

func (x *CommonConf) GetHotPoint() float32 {
	if x != nil {
		return x.HotPoint
	}
	return 0
}

func (x *CommonConf) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type ResCommonConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf *CommonConf `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *ResCommonConf) Reset() {
	*x = ResCommonConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResCommonConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResCommonConf) ProtoMessage() {}

func (x *ResCommonConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResCommonConf.ProtoReflect.Descriptor instead.
func (*ResCommonConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{5}
}

func (x *ResCommonConf) GetConf() *CommonConf {
	if x != nil {
		return x.Conf
	}
	return nil
}

type WhiteListConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid        string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	SupportSvr []string `protobuf:"bytes,3,rep,name=supportSvr,proto3" json:"supportSvr,omitempty"`
}

func (x *WhiteListConf) Reset() {
	*x = WhiteListConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhiteListConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhiteListConf) ProtoMessage() {}

func (x *WhiteListConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhiteListConf.ProtoReflect.Descriptor instead.
func (*WhiteListConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{6}
}

func (x *WhiteListConf) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WhiteListConf) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *WhiteListConf) GetSupportSvr() []string {
	if x != nil {
		return x.SupportSvr
	}
	return nil
}

type ResWhiteListConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf []*WhiteListConf `protobuf:"bytes,1,rep,name=conf,proto3" json:"conf,omitempty"`
}

func (x *ResWhiteListConf) Reset() {
	*x = ResWhiteListConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_res_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResWhiteListConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResWhiteListConf) ProtoMessage() {}

func (x *ResWhiteListConf) ProtoReflect() protoreflect.Message {
	mi := &file_test_res_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResWhiteListConf.ProtoReflect.Descriptor instead.
func (*ResWhiteListConf) Descriptor() ([]byte, []int) {
	return file_test_res_proto_rawDescGZIP(), []int{7}
}

func (x *ResWhiteListConf) GetConf() []*WhiteListConf {
	if x != nil {
		return x.Conf
	}
	return nil
}

var File_test_res_proto protoreflect.FileDescriptor

var file_test_res_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x02, 0x0a,
	0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12,
	0x2e, 0x0a, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x37, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x4d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x45, 0x66, 0x66, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x65, 0x66, 0x66, 0x4d, 0x61, 0x70, 0x1a, 0x4f, 0x0a, 0x0b, 0x45, 0x66, 0x66, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcd, 0x03, 0x0a, 0x0b, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a,
	0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2b,
	0x0a, 0x06, 0x66, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x52, 0x06, 0x66, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x30,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x3c, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x1a, 0x37,
	0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x45, 0x0a, 0x0a,
	0x67, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x4d, 0x61, 0x70, 0x1a, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x4d, 0x61, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x4d, 0x61, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x68, 0x6f, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22,
	0x51, 0x0a, 0x0d, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x76, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x76, 0x72, 0x22, 0x3c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x28, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66,
	0x2a, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x31, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x32, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x33, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x34, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x35, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_res_proto_rawDescOnce sync.Once
	file_test_res_proto_rawDescData = file_test_res_proto_rawDesc
)

func file_test_res_proto_rawDescGZIP() []byte {
	file_test_res_proto_rawDescOnce.Do(func() {
		file_test_res_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_res_proto_rawDescData)
	})
	return file_test_res_proto_rawDescData
}

var file_test_res_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_res_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_test_res_proto_goTypes = []interface{}{
	(GeneralQuality)(0),      // 0: model.GeneralQuality
	(*GeneralEffect)(nil),    // 1: model.GeneralEffect
	(*GeneralSkill)(nil),     // 2: model.GeneralSkill
	(*GeneralConf)(nil),      // 3: model.GeneralConf
	(*ResGeneralConf)(nil),   // 4: model.ResGeneralConf
	(*CommonConf)(nil),       // 5: model.CommonConf
	(*ResCommonConf)(nil),    // 6: model.ResCommonConf
	(*WhiteListConf)(nil),    // 7: model.WhiteListConf
	(*ResWhiteListConf)(nil), // 8: model.ResWhiteListConf
	nil,                      // 9: model.GeneralEffect.ParamEntry
	nil,                      // 10: model.GeneralSkill.EffMapEntry
	nil,                      // 11: model.GeneralConf.LinkEntry
	nil,                      // 12: model.GeneralConf.SkillMapEntry
	nil,                      // 13: model.ResGeneralConf.GenConfMapEntry
}
var file_test_res_proto_depIdxs = []int32{
	9,  // 0: model.GeneralEffect.param:type_name -> model.GeneralEffect.ParamEntry
	1,  // 1: model.GeneralSkill.effect:type_name -> model.GeneralEffect
	1,  // 2: model.GeneralSkill.effects:type_name -> model.GeneralEffect
	10, // 3: model.GeneralSkill.effMap:type_name -> model.GeneralSkill.EffMapEntry
	0,  // 4: model.GeneralConf.quality:type_name -> model.GeneralQuality
	2,  // 5: model.GeneralConf.fSkill:type_name -> model.GeneralSkill
	2,  // 6: model.GeneralConf.skills:type_name -> model.GeneralSkill
	11, // 7: model.GeneralConf.link:type_name -> model.GeneralConf.LinkEntry
	12, // 8: model.GeneralConf.skillMap:type_name -> model.GeneralConf.SkillMapEntry
	13, // 9: model.ResGeneralConf.genConfMap:type_name -> model.ResGeneralConf.GenConfMapEntry
	5,  // 10: model.ResCommonConf.conf:type_name -> model.CommonConf
	7,  // 11: model.ResWhiteListConf.conf:type_name -> model.WhiteListConf
	1,  // 12: model.GeneralSkill.EffMapEntry.value:type_name -> model.GeneralEffect
	2,  // 13: model.GeneralConf.SkillMapEntry.value:type_name -> model.GeneralSkill
	3,  // 14: model.ResGeneralConf.GenConfMapEntry.value:type_name -> model.GeneralConf
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_test_res_proto_init() }
func file_test_res_proto_init() {
	if File_test_res_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_res_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralEffect); i {
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
		file_test_res_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralSkill); i {
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
		file_test_res_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralConf); i {
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
		file_test_res_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGeneralConf); i {
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
		file_test_res_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonConf); i {
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
		file_test_res_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResCommonConf); i {
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
		file_test_res_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhiteListConf); i {
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
		file_test_res_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResWhiteListConf); i {
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
			RawDescriptor: file_test_res_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_res_proto_goTypes,
		DependencyIndexes: file_test_res_proto_depIdxs,
		EnumInfos:         file_test_res_proto_enumTypes,
		MessageInfos:      file_test_res_proto_msgTypes,
	}.Build()
	File_test_res_proto = out.File
	file_test_res_proto_rawDesc = nil
	file_test_res_proto_goTypes = nil
	file_test_res_proto_depIdxs = nil
}
