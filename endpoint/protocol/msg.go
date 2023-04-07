package protocol

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

type MsgType int32

const (
	MsgType_CS  MsgType = 0
	MsgType_SS  MsgType = 1
	MsgType_NTF MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "CS",
		1: "SS",
		2: "NTF",
	}
	MsgType_value = map[string]int32{
		"CS":  0,
		"SS":  1,
		"NTF": 2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type MsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgID string            `protobuf:"bytes,1,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	ID    string            `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	SEQ   uint32            `protobuf:"varint,3,opt,name=SEQ,proto3" json:"SEQ,omitempty"` //unique for every msg
	MT    MsgType           `protobuf:"varint,4,opt,name=MT,proto3,enum=model.MsgType" json:"MT,omitempty"`
	Param map[string]string `protobuf:"bytes,5,rep,name=Param,proto3" json:"Param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgHead) Reset() {
	*x = MsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHead) ProtoMessage() {}

func (x *MsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHead.ProtoReflect.Descriptor instead.
func (*MsgHead) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *MsgHead) GetMsgID() string {
	if x != nil {
		return x.MsgID
	}
	return ""
}

func (x *MsgHead) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MsgHead) GetSEQ() uint32 {
	if x != nil {
		return x.SEQ
	}
	return 0
}

func (x *MsgHead) GetMT() MsgType {
	if x != nil {
		return x.MT
	}
	return MsgType_CS
}

func (x *MsgHead) GetParam() map[string]string {
	if x != nil {
		return x.Param
	}
	return nil
}

type TcpEstablish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointType uint32            `protobuf:"varint,1,opt,name=EndpointType,proto3" json:"EndpointType,omitempty"` // 0 cs 1 ss
	ID           string            `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`                      // tcp channel identifier : cs : uid ; ss : svrId
	Ext          map[string]string `protobuf:"bytes,3,rep,name=Ext,proto3" json:"Ext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TcpEstablish) Reset() {
	*x = TcpEstablish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpEstablish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpEstablish) ProtoMessage() {}

func (x *TcpEstablish) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpEstablish.ProtoReflect.Descriptor instead.
func (*TcpEstablish) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *TcpEstablish) GetEndpointType() uint32 {
	if x != nil {
		return x.EndpointType
	}
	return 0
}

func (x *TcpEstablish) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TcpEstablish) GetExt() map[string]string {
	if x != nil {
		return x.Ext
	}
	return nil
}

type ErrMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode    string `protobuf:"bytes,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrInfo    string `protobuf:"bytes,2,opt,name=ErrInfo,proto3" json:"ErrInfo,omitempty"`
	Disconnect bool   `protobuf:"varint,3,opt,name=Disconnect,proto3" json:"Disconnect,omitempty"`
}

func (x *ErrMsg) Reset() {
	*x = ErrMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrMsg) ProtoMessage() {}

func (x *ErrMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrMsg.ProtoReflect.Descriptor instead.
func (*ErrMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *ErrMsg) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

func (x *ErrMsg) GetErrInfo() string {
	if x != nil {
		return x.ErrInfo
	}
	return ""
}

func (x *ErrMsg) GetDisconnect() bool {
	if x != nil {
		return x.Disconnect
	}
	return false
}

type SystemErr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SystemErr) Reset() {
	*x = SystemErr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemErr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemErr) ProtoMessage() {}

func (x *SystemErr) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemErr.ProtoReflect.Descriptor instead.
func (*SystemErr) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *SystemErr) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d,
	0x73, 0x67, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x45, 0x51, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x53, 0x45, 0x51, 0x12, 0x1e, 0x0a, 0x02, 0x4d, 0x54, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x02, 0x4d, 0x54, 0x12, 0x2f, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x73,
	0x67, 0x48, 0x65, 0x61, 0x64, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x54, 0x63, 0x70, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x03, 0x45, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x63, 0x70, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x03, 0x45, 0x78, 0x74, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c,
	0x0a, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0x1d, 0x0a, 0x09,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x2a, 0x22, 0x0a, 0x07, 0x4d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x53, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x54, 0x46, 0x10, 0x02, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_msg_proto_goTypes = []interface{}{
	(MsgType)(0),         // 0: model.MsgType
	(*MsgHead)(nil),      // 1: model.MsgHead
	(*TcpEstablish)(nil), // 2: model.TcpEstablish
	(*ErrMsg)(nil),       // 3: model.ErrMsg
	(*SystemErr)(nil),    // 4: model.SystemErr
	nil,                  // 5: model.MsgHead.ParamEntry
	nil,                  // 6: model.TcpEstablish.ExtEntry
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: model.MsgHead.MT:type_name -> model.MsgType
	5, // 1: model.MsgHead.Param:type_name -> model.MsgHead.ParamEntry
	6, // 2: model.TcpEstablish.Ext:type_name -> model.TcpEstablish.ExtEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHead); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpEstablish); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrMsg); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemErr); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
