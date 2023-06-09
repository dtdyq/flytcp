// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0-devel
// 	protoc        v3.21.12
// source: entity_person.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

import "github.com/golang/protobuf/proto"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Area float32 `protobuf:"fixed32,2,opt,name=area,proto3" json:"area,omitempty"`
	Addr string  `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`

	_idx_Id int

	_idx_Area int

	_idx_Addr int

	_dirty_flag [3]byte

	_prm protoreflect.Message

	_parent        Entity // current struct 's parent
	_idx_in_parent int    // current 's index in parent struct
	_dt            byte   // struct in parent type:1 direct  2  slice  3 map value
}

// entity block start

func (x *House) MarkDirty(idx int, dt byte) {
	if x != nil {
		x._dirty_flag[idx] = dt
		if x._parent != nil && dt > 0 {
			x._parent.MarkDirty(x._idx_in_parent, x._dt)
		}
	}
}

func (x *House) SetParent(entity Entity, idx int, dt byte) {
	if x != nil {
		x._parent = entity
		x._idx_in_parent = idx
		x._dt = dt
	}
}
func (x *House) BuildEntityInfo() {
	if x != nil {
		x._idx_Id = 0
		x._idx_Area = 1
		x._idx_Addr = 2
	}
}
func (x *House) MarkAllDirty(dt byte) {
	if x != nil {
		for i, _ := range x._dirty_flag {
			x._dirty_flag[i] = dt
		}
		if dt == 0 {
		}
	}
}

// entity block start

func (x *House) GetProtoReflect() protoreflect.Message {
	if x._prm == nil {
		x._prm = x.ProtoReflect()
	}
	return x._prm
}
func (x *House) hasField(fn protoreflect.Name) bool {
	return x.GetProtoReflect().Has(x.GetProtoReflect().Descriptor().Fields().ByName(fn))
}
func (x *House) MergeFrom(s *House) {
	if x == nil || s == nil {
		return
	}
	if s.hasField(`Id`) {
		x.SetId(s.GetId())
	}
	if s.hasField(`Area`) {
		x.SetArea(s.GetArea())
	}
	if s.hasField(`Addr`) {
		x.SetAddr(s.GetAddr())
	}
}
func (x *House) IsDirty() bool {
	for _, b := range x._dirty_flag {
		if b > 0 {
			return true
		}
	}
	return false
}
func (x *House) CopyAll(r *House) {
	if x != nil && r != nil {
		bys, _ := proto.Marshal(x)
		proto.Unmarshal(bys, r)
	}
}
func (x *House) CollectDirty(r *House, clear bool) {
	if x != nil && r != nil {
		r.BuildEntityInfo()
		if x._dirty_flag[x._idx_Id] > 0 {
			r.SetId(x.Id)
		}
		if x._dirty_flag[x._idx_Area] > 0 {
			r.SetArea(x.Area)
		}
		if x._dirty_flag[x._idx_Addr] > 0 {
			r.SetAddr(x.Addr)
		}
		if clear {
			x._dirty_flag = [3]byte{}
		}
	}
}

// entity block start

func (x *House) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
func (x *House) SetId(v string) {
	if x != nil {
		x.Id = v
		x.MarkDirty(x._idx_Id, 1)
	}
}
func (x *House) GetArea() float32 {
	if x != nil {
		return x.Area
	}
	return 0
}
func (x *House) SetArea(v float32) {
	if x != nil {
		x.Area = v
		x.MarkDirty(x._idx_Area, 1)
	}
}
func (x *House) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}
func (x *House) SetAddr(v string) {
	if x != nil {
		x.Addr = v
		x.MarkDirty(x._idx_Addr, 1)
	}
}

func (x *House) Reset() {
	*x = House{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House) ProtoMessage() {}

func (x *House) ProtoReflect() protoreflect.Message {
	mi := &file_entity_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House.ProtoReflect.Descriptor instead.
func (*House) Descriptor() ([]byte, []int) {
	return file_entity_person_proto_rawDescGZIP(), []int{0}
}

type Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`

	_idx_Name int

	_idx_Age int

	_dirty_flag [2]byte

	_prm protoreflect.Message

	_parent        Entity // current struct 's parent
	_idx_in_parent int    // current 's index in parent struct
	_dt            byte   // struct in parent type:1 direct  2  slice  3 map value
}

// entity block start

func (x *Child) MarkDirty(idx int, dt byte) {
	if x != nil {
		x._dirty_flag[idx] = dt
		if x._parent != nil && dt > 0 {
			x._parent.MarkDirty(x._idx_in_parent, x._dt)
		}
	}
}

func (x *Child) SetParent(entity Entity, idx int, dt byte) {
	if x != nil {
		x._parent = entity
		x._idx_in_parent = idx
		x._dt = dt
	}
}
func (x *Child) BuildEntityInfo() {
	if x != nil {
		x._idx_Name = 0
		x._idx_Age = 1
	}
}
func (x *Child) MarkAllDirty(dt byte) {
	if x != nil {
		for i, _ := range x._dirty_flag {
			x._dirty_flag[i] = dt
		}
		if dt == 0 {
		}
	}
}

// entity block start

func (x *Child) GetProtoReflect() protoreflect.Message {
	if x._prm == nil {
		x._prm = x.ProtoReflect()
	}
	return x._prm
}
func (x *Child) hasField(fn protoreflect.Name) bool {
	return x.GetProtoReflect().Has(x.GetProtoReflect().Descriptor().Fields().ByName(fn))
}
func (x *Child) MergeFrom(s *Child) {
	if x == nil || s == nil {
		return
	}
	if s.hasField(`Name`) {
		x.SetName(s.GetName())
	}
	if s.hasField(`Age`) {
		x.SetAge(s.GetAge())
	}
}
func (x *Child) IsDirty() bool {
	for _, b := range x._dirty_flag {
		if b > 0 {
			return true
		}
	}
	return false
}
func (x *Child) CopyAll(r *Child) {
	if x != nil && r != nil {
		bys, _ := proto.Marshal(x)
		proto.Unmarshal(bys, r)
	}
}
func (x *Child) CollectDirty(r *Child, clear bool) {
	if x != nil && r != nil {
		r.BuildEntityInfo()
		if x._dirty_flag[x._idx_Name] > 0 {
			r.SetName(x.Name)
		}
		if x._dirty_flag[x._idx_Age] > 0 {
			r.SetAge(x.Age)
		}
		if clear {
			x._dirty_flag = [2]byte{}
		}
	}
}

// entity block start

func (x *Child) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *Child) SetName(v string) {
	if x != nil {
		x.Name = v
		x.MarkDirty(x._idx_Name, 1)
	}
}
func (x *Child) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}
func (x *Child) SetAge(v int32) {
	if x != nil {
		x.Age = v
		x.MarkDirty(x._idx_Age, 1)
	}
}

func (x *Child) Reset() {
	*x = Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Child) ProtoMessage() {}

func (x *Child) ProtoReflect() protoreflect.Message {
	mi := &file_entity_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Child.ProtoReflect.Descriptor instead.
func (*Child) Descriptor() ([]byte, []int) {
	return file_entity_person_proto_rawDescGZIP(), []int{1}
}

type ExtMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo        string `protobuf:"bytes,1,opt,name=Todo,proto3" json:"Todo,omitempty"`
	PreAllocate string `protobuf:"bytes,2,opt,name=PreAllocate,proto3" json:"PreAllocate,omitempty"`

	_idx_Todo int

	_idx_PreAllocate int

	_dirty_flag [2]byte

	_prm protoreflect.Message

	_parent        Entity // current struct 's parent
	_idx_in_parent int    // current 's index in parent struct
	_dt            byte   // struct in parent type:1 direct  2  slice  3 map value
}

// entity block start

func (x *ExtMsg) MarkDirty(idx int, dt byte) {
	if x != nil {
		x._dirty_flag[idx] = dt
		if x._parent != nil && dt > 0 {
			x._parent.MarkDirty(x._idx_in_parent, x._dt)
		}
	}
}

func (x *ExtMsg) SetParent(entity Entity, idx int, dt byte) {
	if x != nil {
		x._parent = entity
		x._idx_in_parent = idx
		x._dt = dt
	}
}
func (x *ExtMsg) BuildEntityInfo() {
	if x != nil {
		x._idx_Todo = 0
		x._idx_PreAllocate = 1
	}
}
func (x *ExtMsg) MarkAllDirty(dt byte) {
	if x != nil {
		for i, _ := range x._dirty_flag {
			x._dirty_flag[i] = dt
		}
		if dt == 0 {
		}
	}
}

// entity block start

func (x *ExtMsg) GetProtoReflect() protoreflect.Message {
	if x._prm == nil {
		x._prm = x.ProtoReflect()
	}
	return x._prm
}
func (x *ExtMsg) hasField(fn protoreflect.Name) bool {
	return x.GetProtoReflect().Has(x.GetProtoReflect().Descriptor().Fields().ByName(fn))
}
func (x *ExtMsg) MergeFrom(s *ExtMsg) {
	if x == nil || s == nil {
		return
	}
	if s.hasField(`Todo`) {
		x.SetTodo(s.GetTodo())
	}
	if s.hasField(`PreAllocate`) {
		x.SetPreAllocate(s.GetPreAllocate())
	}
}
func (x *ExtMsg) IsDirty() bool {
	for _, b := range x._dirty_flag {
		if b > 0 {
			return true
		}
	}
	return false
}
func (x *ExtMsg) CopyAll(r *ExtMsg) {
	if x != nil && r != nil {
		bys, _ := proto.Marshal(x)
		proto.Unmarshal(bys, r)
	}
}
func (x *ExtMsg) CollectDirty(r *ExtMsg, clear bool) {
	if x != nil && r != nil {
		r.BuildEntityInfo()
		if x._dirty_flag[x._idx_Todo] > 0 {
			r.SetTodo(x.Todo)
		}
		if x._dirty_flag[x._idx_PreAllocate] > 0 {
			r.SetPreAllocate(x.PreAllocate)
		}
		if clear {
			x._dirty_flag = [2]byte{}
		}
	}
}

// entity block start

func (x *ExtMsg) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}
func (x *ExtMsg) SetTodo(v string) {
	if x != nil {
		x.Todo = v
		x.MarkDirty(x._idx_Todo, 1)
	}
}
func (x *ExtMsg) GetPreAllocate() string {
	if x != nil {
		return x.PreAllocate
	}
	return ""
}
func (x *ExtMsg) SetPreAllocate(v string) {
	if x != nil {
		x.PreAllocate = v
		x.MarkDirty(x._idx_PreAllocate, 1)
	}
}

func (x *ExtMsg) Reset() {
	*x = ExtMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtMsg) ProtoMessage() {}

func (x *ExtMsg) ProtoReflect() protoreflect.Message {
	mi := &file_entity_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtMsg.ProtoReflect.Descriptor instead.
func (*ExtMsg) Descriptor() ([]byte, []int) {
	return file_entity_person_proto_rawDescGZIP(), []int{2}
}

type PersonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locate string  `protobuf:"bytes,1,opt,name=Locate,proto3" json:"Locate,omitempty"`
	Url    string  `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
	Ext    *ExtMsg `protobuf:"bytes,3,opt,name=Ext,proto3" json:"Ext,omitempty"`

	_idx_Locate int

	_idx_Url int

	_idx_Ext int

	_dirty_flag [3]byte

	_prm protoreflect.Message

	_parent        Entity // current struct 's parent
	_idx_in_parent int    // current 's index in parent struct
	_dt            byte   // struct in parent type:1 direct  2  slice  3 map value
}

// entity block start

func (x *PersonInfo) MarkDirty(idx int, dt byte) {
	if x != nil {
		x._dirty_flag[idx] = dt
		if x._parent != nil && dt > 0 {
			x._parent.MarkDirty(x._idx_in_parent, x._dt)
		}
	}
}

func (x *PersonInfo) SetParent(entity Entity, idx int, dt byte) {
	if x != nil {
		x._parent = entity
		x._idx_in_parent = idx
		x._dt = dt
	}
}
func (x *PersonInfo) BuildEntityInfo() {
	if x != nil {
		x._idx_Locate = 0
		x._idx_Url = 1
		x._idx_Ext = 2
		if x.Ext != nil {
			x.Ext.SetParent(x, x._idx_Ext, 1)
			x.Ext.BuildEntityInfo()
		}
	}
}
func (x *PersonInfo) MarkAllDirty(dt byte) {
	if x != nil {
		if x.Ext != nil {
			x.Ext.MarkAllDirty(dt)
		}
		for i, _ := range x._dirty_flag {
			x._dirty_flag[i] = dt
		}
		if dt == 0 {
		}
	}
}

// entity block start

func (x *PersonInfo) GetProtoReflect() protoreflect.Message {
	if x._prm == nil {
		x._prm = x.ProtoReflect()
	}
	return x._prm
}
func (x *PersonInfo) hasField(fn protoreflect.Name) bool {
	return x.GetProtoReflect().Has(x.GetProtoReflect().Descriptor().Fields().ByName(fn))
}
func (x *PersonInfo) MergeFrom(s *PersonInfo) {
	if x == nil || s == nil {
		return
	}
	if s.hasField(`Locate`) {
		x.SetLocate(s.GetLocate())
	}
	if s.hasField(`Url`) {
		x.SetUrl(s.GetUrl())
	}
	if s.hasField(`Ext`) {
		x.GetExt().MergeFrom(s.GetExt())
	}
}
func (x *PersonInfo) IsDirty() bool {
	for _, b := range x._dirty_flag {
		if b > 0 {
			return true
		}
	}
	return false
}
func (x *PersonInfo) CopyAll(r *PersonInfo) {
	if x != nil && r != nil {
		bys, _ := proto.Marshal(x)
		proto.Unmarshal(bys, r)
	}
}
func (x *PersonInfo) CollectDirty(r *PersonInfo, clear bool) {
	if x != nil && r != nil {
		r.BuildEntityInfo()
		if x._dirty_flag[x._idx_Locate] > 0 {
			r.SetLocate(x.Locate)
		}
		if x._dirty_flag[x._idx_Url] > 0 {
			r.SetUrl(x.Url)
		}
		if x._dirty_flag[x._idx_Ext] > 0 {
			if x.Ext == nil {
				r.Ext = nil
			} else {
				if r.Ext == nil {
					r.Ext = &ExtMsg{}
				}
				x.Ext.CollectDirty(r.Ext, clear)
			}
		}
		if clear {
			x._dirty_flag = [3]byte{}
		}
	}
}

// entity block start

func (x *PersonInfo) GetLocate() string {
	if x != nil {
		return x.Locate
	}
	return ""
}
func (x *PersonInfo) SetLocate(v string) {
	if x != nil {
		x.Locate = v
		x.MarkDirty(x._idx_Locate, 1)
	}
}
func (x *PersonInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}
func (x *PersonInfo) SetUrl(v string) {
	if x != nil {
		x.Url = v
		x.MarkDirty(x._idx_Url, 1)
	}
}
func (x *PersonInfo) GetExt() *ExtMsg {
	if x != nil {
		if x.Ext == nil {
			x.Ext = &ExtMsg{}
		}
		x.Ext.SetParent(x, x._idx_Ext, 1)
		x.Ext.BuildEntityInfo()
		return x.Ext
	}
	return nil
}
func (x *PersonInfo) SetExt(v *ExtMsg) {
	if x != nil {
		x.Ext = v
		x.MarkDirty(x._idx_Ext, 1)
		if x.Ext != nil {
			x.Ext.SetParent(x, x._idx_Ext, 1)
			x.Ext.BuildEntityInfo()
			x.Ext.MarkAllDirty(1)
		}
	}
}

func (x *PersonInfo) Reset() {
	*x = PersonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfo) ProtoMessage() {}

func (x *PersonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_entity_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfo.ProtoReflect.Descriptor instead.
func (*PersonInfo) Descriptor() ([]byte, []int) {
	return file_entity_person_proto_rawDescGZIP(), []int{3}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age           int32             `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Favor         []string          `protobuf:"bytes,3,rep,name=Favor,proto3" json:"Favor,omitempty"` //爱好
	Children      []*Child          `protobuf:"bytes,4,rep,name=Children,proto3" json:"Children,omitempty"`
	Thing         map[string]int32  `protobuf:"bytes,5,rep,name=Thing,proto3" json:"Thing,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`  //物品-数量
	Houses        map[string]*House `protobuf:"bytes,6,rep,name=Houses,proto3" json:"Houses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //住宅
	SelfInfo      *PersonInfo       `protobuf:"bytes,7,opt,name=SelfInfo,proto3" json:"SelfInfo,omitempty"`
	ThingDeleted  []string          `protobuf:"bytes,1000,rep,name=ThingDeleted,proto3" json:"ThingDeleted,omitempty"`
	HousesDeleted []string          `protobuf:"bytes,1001,rep,name=HousesDeleted,proto3" json:"HousesDeleted,omitempty"`

	_idx_Name int

	_idx_Age int

	_idx_Favor int

	_idx_Children int

	_idx_Thing int

	_idx_Houses int

	_idx_SelfInfo int

	_dirty_flag [7]byte

	_prm protoreflect.Message

	_parent        Entity // current struct 's parent
	_idx_in_parent int    // current 's index in parent struct
	_dt            byte   // struct in parent type:1 direct  2  slice  3 map value
}

// entity block start

func (x *Person) MarkDirty(idx int, dt byte) {
	if x != nil {
		x._dirty_flag[idx] = dt
		if x._parent != nil && dt > 0 {
			x._parent.MarkDirty(x._idx_in_parent, x._dt)
		}
	}
}

func (x *Person) SetParent(entity Entity, idx int, dt byte) {
	if x != nil {
		x._parent = entity
		x._idx_in_parent = idx
		x._dt = dt
	}
}
func (x *Person) BuildEntityInfo() {
	if x != nil {
		x._idx_Name = 0
		x._idx_Age = 1
		x._idx_Favor = 2
		x._idx_Children = 3
		for _, v := range x.Children {
			v.SetParent(x, x._idx_Children, 3)
			v.BuildEntityInfo()
		}
		x._idx_Thing = 4
		x._idx_Houses = 5
		for _, v := range x.Houses {
			v.SetParent(x, x._idx_Houses, 3)
			v.BuildEntityInfo()
		}
		x._idx_SelfInfo = 6
		if x.SelfInfo != nil {
			x.SelfInfo.SetParent(x, x._idx_SelfInfo, 1)
			x.SelfInfo.BuildEntityInfo()
		}
	}
}
func (x *Person) MarkAllDirty(dt byte) {
	if x != nil {
		for _, v := range x.Children {
			v.MarkAllDirty(dt)
		}
		for _, v := range x.Houses {
			v.MarkAllDirty(dt)
		}
		if x.SelfInfo != nil {
			x.SelfInfo.MarkAllDirty(dt)
		}
		for i, _ := range x._dirty_flag {
			x._dirty_flag[i] = dt
		}
		if dt == 0 {
			x.ThingDeleted = make([]string, 0)
			x.HousesDeleted = make([]string, 0)
		}
	}
}

// entity block start

func (x *Person) GetProtoReflect() protoreflect.Message {
	if x._prm == nil {
		x._prm = x.ProtoReflect()
	}
	return x._prm
}
func (x *Person) hasField(fn protoreflect.Name) bool {
	return x.GetProtoReflect().Has(x.GetProtoReflect().Descriptor().Fields().ByName(fn))
}
func (x *Person) MergeFrom(s *Person) {
	if x == nil || s == nil {
		return
	}
	if s.hasField(`Name`) {
		x.SetName(s.GetName())
	}
	if s.hasField(`Age`) {
		x.SetAge(s.GetAge())
	}
	if s.hasField(`SelfInfo`) {
		x.GetSelfInfo().MergeFrom(s.GetSelfInfo())
	}
}
func (x *Person) IsDirty() bool {
	for _, b := range x._dirty_flag {
		if b > 0 {
			return true
		}
	}
	return false
}
func (x *Person) CopyAll(r *Person) {
	if x != nil && r != nil {
		bys, _ := proto.Marshal(x)
		proto.Unmarshal(bys, r)
	}
}
func (x *Person) CollectDirty(r *Person, clear bool) {
	if x != nil && r != nil {
		r.BuildEntityInfo()
		if x._dirty_flag[x._idx_Name] > 0 {
			r.SetName(x.Name)
		}
		if x._dirty_flag[x._idx_Age] > 0 {
			r.SetAge(x.Age)
		}
		if x._dirty_flag[x._idx_Favor] > 0 {
			r.Favor = make([]string, 0)
			if x.Favor != nil {
				r.SetFavor(x.Favor)
			}
		}
		if x._dirty_flag[x._idx_Children] > 0 {
			r.Children = make([]*Child, 0)
			if x.Children != nil {
				for _, i := range x.Children {
					if i != nil {
						vi := &Child{}
						i.CopyAll(vi)
						r.AppendChildren(vi)
					}
				}
			}
		}
		if x._dirty_flag[x._idx_Thing] > 0 {
			r.MarkDirty(r._idx_Thing, 1)
			if x.Thing != nil {
				if r.ThingDeleted == nil {
					r.ThingDeleted = make([]string, 0)
				}
				if r.Thing == nil {
					r.Thing = make(map[string]int32, 0)
				}
				r.ThingDeleted = MergeNotExistInto(r.ThingDeleted, x.ThingDeleted)
				for k, v := range x.Thing {
					var _, ok = r.Thing[k]
					if !ok {
						r.ThingDeleted = RemoveValueOf(r.ThingDeleted, k)
					}
					r.PutThingByKey(k, v)
				}
				for _, d := range r.ThingDeleted {
					delete(r.Thing, d)
				}
			} else {
				r.Thing = nil
			}
		}
		if x._dirty_flag[x._idx_Houses] > 0 {
			r.MarkDirty(r._idx_Houses, 1)
			if x.Houses != nil {
				if r.HousesDeleted == nil {
					r.HousesDeleted = make([]string, 0)
				}
				if r.Houses == nil {
					r.Houses = make(map[string]*House, 0)
				}
				r.HousesDeleted = MergeNotExistInto(r.HousesDeleted, x.HousesDeleted)
				for k, v := range x.Houses {
					if v != nil && v.IsDirty() {
						var vr, ok = r.Houses[k]
						if !ok {
							vr = &House{}
							r.HousesDeleted = RemoveValueOf(r.HousesDeleted, k)
						}
						v.CollectDirty(vr, clear)
						r.Houses[k] = vr
					}
				}
				for _, d := range r.HousesDeleted {
					delete(r.Houses, d)
				}
			} else {
				r.Houses = nil
			}
		}
		if x._dirty_flag[x._idx_SelfInfo] > 0 {
			if x.SelfInfo == nil {
				r.SelfInfo = nil
			} else {
				if r.SelfInfo == nil {
					r.SelfInfo = &PersonInfo{}
				}
				x.SelfInfo.CollectDirty(r.SelfInfo, clear)
			}
		}
		if clear {
			x._dirty_flag = [7]byte{}
		}
	}
}

// entity block start

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *Person) SetName(v string) {
	if x != nil {
		x.Name = v
		x.MarkDirty(x._idx_Name, 1)
	}
}
func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}
func (x *Person) SetAge(v int32) {
	if x != nil {
		x.Age = v
		x.MarkDirty(x._idx_Age, 1)
	}
}
func (x *Person) GetFavor() []string {
	if x != nil {
		if x.Favor == nil {
			x.Favor = make([]string, 0)
		}
		return x.Favor
	}
	return nil
}
func (x *Person) SetFavor(v []string) {
	if x != nil {
		x.Favor = v
		x.MarkDirty(x._idx_Favor, 1)
	}
}
func (x *Person) AppendFavor(v string) {
	if x != nil {
		if x.Favor == nil {
			x.Favor = make([]string, 0)
		}
		x.Favor = append(x.Favor, v)
		x.MarkDirty(x._idx_Favor, 1)
	}
}
func (x *Person) GetFavorByIdx(i int) string {
	if x != nil {
		if x.Favor == nil {
			x.Favor = make([]string, 0)
		}
		if len(x.Favor) > i {
			return x.Favor[i]
		}
	}
	return ""
}
func (x *Person) GetChildren() []*Child {
	if x != nil {
		if x.Children == nil {
			x.Children = make([]*Child, 0)
		}
		return x.Children
	}
	return nil
}
func (x *Person) SetChildren(v []*Child) {
	if x != nil {
		x.Children = v
		x.MarkDirty(x._idx_Children, 1)
		if v == nil {
			x.Children = make([]*Child, 0)
		}
		for i, _ := range x.Children {
			if x.Children[i] != nil {
				x.Children[i].SetParent(x, x._idx_Children, 1)
				x.Children[i].BuildEntityInfo()
				x.Children[i].MarkAllDirty(1)
			}
		}
	}
}
func (x *Person) GetChildrenByIdx(i int) *Child {
	if x != nil {
		if x.Children == nil {
			x.Children = make([]*Child, 0)
		}
		if len(x.Children) > i {
			return x.Children[i]
		}
		return &Child{}
	}
	return nil
}
func (x *Person) AppendChildren(v *Child) {
	if x != nil {
		if x.Children == nil {
			x.Children = make([]*Child, 0)
		}
		x.Children = append(x.Children, v)
		x.MarkDirty(x._idx_Children, 1)
		v.SetParent(x, x._idx_Children, 1)
		v.BuildEntityInfo()
		v.MarkAllDirty(1)
	}
}
func (x *Person) GetThing() map[string]int32 {
	if x != nil {
		if x.Thing == nil {
			x.Thing = make(map[string]int32)
		}
		return x.Thing
	}
	return nil
}
func (x *Person) GetThingByKey(k string) (int32, bool) {
	if x != nil {
		if x.Thing == nil {
			x.Thing = make(map[string]int32)
		}
		ret, ok := x.Thing[k]
		return ret, ok
	}
	return 0, false
}
func (x *Person) SetThing(v map[string]int32) {
	if x != nil {
		if x.Thing == nil {
			x.Thing = make(map[string]int32)
		}
		x.MarkDirty(x._idx_Thing, 1)
		for o, _ := range x.Thing {
			x.ThingDeleted = append(x.ThingDeleted, o)
		}
		x.Thing = make(map[string]int32)
		if v != nil {
			for k, _ := range v {
				x.PutThingByKey(k, v[k])
			}
		}
	}
}
func (x *Person) PutThingByKey(k string, v int32) {
	if x != nil {
		x.MarkDirty(x._idx_Thing, 1)
		if x.Thing == nil {
			x.Thing = make(map[string]int32)
		}
		x.Thing[k] = v
		x.ThingDeleted = RemoveValueOf(x.ThingDeleted, k)
	}
}
func (x *Person) DelThingByKey(k string) {
	if x != nil {
		if x.Thing == nil {
			x.Thing = make(map[string]int32)
		}
		_, ok := x.Thing[k]
		if ok {
			x.MarkDirty(x._idx_Thing, 1)
			delete(x.Thing, k)
			x.ThingDeleted = AppendIfNotExist(x.ThingDeleted, k)

		}
	}
}
func (x *Person) GetHouses() map[string]*House {
	if x != nil {
		if x.Houses == nil {
			x.Houses = make(map[string]*House)
		}
		return x.Houses
	}
	return nil
}
func (x *Person) GetHousesByKey(k string) (*House, bool) {
	if x != nil {
		if x.Houses == nil {
			x.Houses = make(map[string]*House)
		}
		ret, ok := x.Houses[k]
		return ret, ok
	}
	return nil, false
}
func (x *Person) SetHouses(v map[string]*House) {
	if x != nil {
		if x.Houses == nil {
			x.Houses = make(map[string]*House)
		}
		x.MarkDirty(x._idx_Houses, 1)
		for o, _ := range x.Houses {
			x.HousesDeleted = append(x.HousesDeleted, o)
		}
		x.Houses = make(map[string]*House)
		if v != nil {
			for k, _ := range v {
				x.PutHousesByKey(k, v[k])
			}
		}
	}
}
func (x *Person) PutHousesByKey(k string, v *House) {
	if x != nil {
		x.MarkDirty(x._idx_Houses, 1)
		if x.Houses == nil {
			x.Houses = make(map[string]*House)
		}
		x.Houses[k] = v
		x.HousesDeleted = RemoveValueOf(x.HousesDeleted, k)
		if x.Houses[k] != nil {
			x.Houses[k].SetParent(x, x._idx_Houses, 1)
			x.Houses[k].BuildEntityInfo()
			x.Houses[k].MarkAllDirty(1)
		}
	}
}
func (x *Person) DelHousesByKey(k string) {
	if x != nil {
		if x.Houses == nil {
			x.Houses = make(map[string]*House)
		}
		_, ok := x.Houses[k]
		if ok {
			x.MarkDirty(x._idx_Houses, 1)
			delete(x.Houses, k)
			x.HousesDeleted = AppendIfNotExist(x.HousesDeleted, k)

		}
	}
}
func (x *Person) GetSelfInfo() *PersonInfo {
	if x != nil {
		if x.SelfInfo == nil {
			x.SelfInfo = &PersonInfo{}
		}
		x.SelfInfo.SetParent(x, x._idx_SelfInfo, 1)
		x.SelfInfo.BuildEntityInfo()
		return x.SelfInfo
	}
	return nil
}
func (x *Person) SetSelfInfo(v *PersonInfo) {
	if x != nil {
		x.SelfInfo = v
		x.MarkDirty(x._idx_SelfInfo, 1)
		if x.SelfInfo != nil {
			x.SelfInfo.SetParent(x, x._idx_SelfInfo, 1)
			x.SelfInfo.BuildEntityInfo()
			x.SelfInfo.MarkAllDirty(1)
		}
	}
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_entity_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_entity_person_proto_rawDescGZIP(), []int{4}
}

var File_entity_person_proto protoreflect.FileDescriptor

var file_entity_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x3f, 0x0a, 0x05,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x2d, 0x0a,
	0x05, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x06,
	0x45, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72,
	0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x22, 0x57, 0x0a, 0x0a,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x03, 0x45, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x4d, 0x73, 0x67,
	0x52, 0x03, 0x45, 0x78, 0x74, 0x22, 0xcf, 0x03, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x08,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x08, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x53, 0x65, 0x6c,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x53, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0c, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0d, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0xe9,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47,
	0x0a, 0x0b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_person_proto_rawDescOnce sync.Once
	file_entity_person_proto_rawDescData = file_entity_person_proto_rawDesc
)

func file_entity_person_proto_rawDescGZIP() []byte {
	file_entity_person_proto_rawDescOnce.Do(func() {
		file_entity_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_person_proto_rawDescData)
	})
	return file_entity_person_proto_rawDescData
}

var file_entity_person_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_entity_person_proto_goTypes = []interface{}{
	(*House)(nil),      // 0: model.House
	(*Child)(nil),      // 1: model.Child
	(*ExtMsg)(nil),     // 2: model.ExtMsg
	(*PersonInfo)(nil), // 3: model.PersonInfo
	(*Person)(nil),     // 4: model.Person
	nil,                // 5: model.Person.ThingEntry
	nil,                // 6: model.Person.HousesEntry
}
var file_entity_person_proto_depIdxs = []int32{
	2, // 0: model.PersonInfo.Ext:type_name -> model.ExtMsg
	1, // 1: model.Person.Children:type_name -> model.Child
	5, // 2: model.Person.Thing:type_name -> model.Person.ThingEntry
	6, // 3: model.Person.Houses:type_name -> model.Person.HousesEntry
	3, // 4: model.Person.SelfInfo:type_name -> model.PersonInfo
	0, // 5: model.Person.HousesEntry.value:type_name -> model.House
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_entity_person_proto_init() }
func file_entity_person_proto_init() {
	if File_entity_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*House); i {
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
		file_entity_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Child); i {
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
		file_entity_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtMsg); i {
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
		file_entity_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfo); i {
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
		file_entity_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_entity_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_person_proto_goTypes,
		DependencyIndexes: file_entity_person_proto_depIdxs,
		MessageInfos:      file_entity_person_proto_msgTypes,
	}.Build()
	File_entity_person_proto = out.File
	file_entity_person_proto_rawDesc = nil
	file_entity_person_proto_goTypes = nil
	file_entity_person_proto_depIdxs = nil
}
