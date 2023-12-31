// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: crew.proto

package planetexpress

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

type Crew_Person_Department int32

const (
	Crew_Person_CAPTAIN Crew_Person_Department = 0
	Crew_Person_DECK    Crew_Person_Department = 1
	Crew_Person_ENGINE  Crew_Person_Department = 2
)

// Enum value maps for Crew_Person_Department.
var (
	Crew_Person_Department_name = map[int32]string{
		0: "CAPTAIN",
		1: "DECK",
		2: "ENGINE",
	}
	Crew_Person_Department_value = map[string]int32{
		"CAPTAIN": 0,
		"DECK":    1,
		"ENGINE":  2,
	}
)

func (x Crew_Person_Department) Enum() *Crew_Person_Department {
	p := new(Crew_Person_Department)
	*p = x
	return p
}

func (x Crew_Person_Department) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Crew_Person_Department) Descriptor() protoreflect.EnumDescriptor {
	return file_crew_proto_enumTypes[0].Descriptor()
}

func (Crew_Person_Department) Type() protoreflect.EnumType {
	return &file_crew_proto_enumTypes[0]
}

func (x Crew_Person_Department) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Crew_Person_Department.Descriptor instead.
func (Crew_Person_Department) EnumDescriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Crew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members map[string]*Crew_Person `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Crew) Reset() {
	*x = Crew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crew) ProtoMessage() {}

func (x *Crew) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crew.ProtoReflect.Descriptor instead.
func (*Crew) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{0}
}

func (x *Crew) GetMembers() map[string]*Crew_Person {
	if x != nil {
		return x.Members
	}
	return nil
}

type GetCrewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crew *Crew `protobuf:"bytes,1,opt,name=crew,proto3" json:"crew,omitempty"`
}

func (x *GetCrewResponse) Reset() {
	*x = GetCrewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrewResponse) ProtoMessage() {}

func (x *GetCrewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrewResponse.ProtoReflect.Descriptor instead.
func (*GetCrewResponse) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{1}
}

func (x *GetCrewResponse) GetCrew() *Crew {
	if x != nil {
		return x.Crew
	}
	return nil
}

type Crew_Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age        int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Department Crew_Person_Department `protobuf:"varint,3,opt,name=department,proto3,enum=planetexpress.Crew_Person_Department" json:"department,omitempty"`
}

func (x *Crew_Person) Reset() {
	*x = Crew_Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crew_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crew_Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crew_Person) ProtoMessage() {}

func (x *Crew_Person) ProtoReflect() protoreflect.Message {
	mi := &file_crew_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crew_Person.ProtoReflect.Descriptor instead.
func (*Crew_Person) Descriptor() ([]byte, []int) {
	return file_crew_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Crew_Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Crew_Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Crew_Person) GetDepartment() Crew_Person_Department {
	if x != nil {
		return x.Department
	}
	return Crew_Person_CAPTAIN
}

var File_crew_proto protoreflect.FileDescriptor

var file_crew_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc3, 0x02, 0x0a, 0x04,
	0x43, 0x72, 0x65, 0x77, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x1a, 0xa6, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x50, 0x54, 0x41, 0x49,
	0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x1a, 0x56, 0x0a, 0x0c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x77, 0x52, 0x04, 0x63, 0x72, 0x65, 0x77, 0x42, 0x13, 0x5a,
	0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crew_proto_rawDescOnce sync.Once
	file_crew_proto_rawDescData = file_crew_proto_rawDesc
)

func file_crew_proto_rawDescGZIP() []byte {
	file_crew_proto_rawDescOnce.Do(func() {
		file_crew_proto_rawDescData = protoimpl.X.CompressGZIP(file_crew_proto_rawDescData)
	})
	return file_crew_proto_rawDescData
}

var file_crew_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crew_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crew_proto_goTypes = []interface{}{
	(Crew_Person_Department)(0), // 0: planetexpress.Crew.Person.Department
	(*Crew)(nil),                // 1: planetexpress.Crew
	(*GetCrewResponse)(nil),     // 2: planetexpress.GetCrewResponse
	(*Crew_Person)(nil),         // 3: planetexpress.Crew.Person
	nil,                         // 4: planetexpress.Crew.MembersEntry
}
var file_crew_proto_depIdxs = []int32{
	4, // 0: planetexpress.Crew.members:type_name -> planetexpress.Crew.MembersEntry
	1, // 1: planetexpress.GetCrewResponse.crew:type_name -> planetexpress.Crew
	0, // 2: planetexpress.Crew.Person.department:type_name -> planetexpress.Crew.Person.Department
	3, // 3: planetexpress.Crew.MembersEntry.value:type_name -> planetexpress.Crew.Person
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_crew_proto_init() }
func file_crew_proto_init() {
	if File_crew_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crew_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crew); i {
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
		file_crew_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrewResponse); i {
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
		file_crew_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crew_Person); i {
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
			RawDescriptor: file_crew_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crew_proto_goTypes,
		DependencyIndexes: file_crew_proto_depIdxs,
		EnumInfos:         file_crew_proto_enumTypes,
		MessageInfos:      file_crew_proto_msgTypes,
	}.Build()
	File_crew_proto = out.File
	file_crew_proto_rawDesc = nil
	file_crew_proto_goTypes = nil
	file_crew_proto_depIdxs = nil
}
