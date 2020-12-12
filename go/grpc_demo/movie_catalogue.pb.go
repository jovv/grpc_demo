// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: movie_catalogue.proto

package grpc_demo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CastMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character string `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *CastMember) Reset() {
	*x = CastMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_catalogue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CastMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CastMember) ProtoMessage() {}

func (x *CastMember) ProtoReflect() protoreflect.Message {
	mi := &file_movie_catalogue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CastMember.ProtoReflect.Descriptor instead.
func (*CastMember) Descriptor() ([]byte, []int) {
	return file_movie_catalogue_proto_rawDescGZIP(), []int{0}
}

func (x *CastMember) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *CastMember) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CastMember) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ProductionYear int32         `protobuf:"varint,3,opt,name=productionYear,proto3" json:"productionYear,omitempty"`
	Genre          string        `protobuf:"bytes,4,opt,name=genre,proto3" json:"genre,omitempty"`
	Duration       int32         `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	CastMembers    []*CastMember `protobuf:"bytes,6,rep,name=castMembers,proto3" json:"castMembers,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_catalogue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_catalogue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_catalogue_proto_rawDescGZIP(), []int{1}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetProductionYear() int32 {
	if x != nil {
		return x.ProductionYear
	}
	return 0
}

func (x *Movie) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Movie) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Movie) GetCastMembers() []*CastMember {
	if x != nil {
		return x.CastMembers
	}
	return nil
}

type MovieCatalogue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *MovieCatalogue) Reset() {
	*x = MovieCatalogue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_catalogue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieCatalogue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieCatalogue) ProtoMessage() {}

func (x *MovieCatalogue) ProtoReflect() protoreflect.Message {
	mi := &file_movie_catalogue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieCatalogue.ProtoReflect.Descriptor instead.
func (*MovieCatalogue) Descriptor() ([]byte, []int) {
	return file_movie_catalogue_proto_rawDescGZIP(), []int{2}
}

func (x *MovieCatalogue) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

var File_movie_catalogue_proto protoreflect.FileDescriptor

var file_movie_catalogue_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x22, 0x64, 0x0a, 0x0a, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x0b, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a,
	0x0e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_catalogue_proto_rawDescOnce sync.Once
	file_movie_catalogue_proto_rawDescData = file_movie_catalogue_proto_rawDesc
)

func file_movie_catalogue_proto_rawDescGZIP() []byte {
	file_movie_catalogue_proto_rawDescOnce.Do(func() {
		file_movie_catalogue_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_catalogue_proto_rawDescData)
	})
	return file_movie_catalogue_proto_rawDescData
}

var file_movie_catalogue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_movie_catalogue_proto_goTypes = []interface{}{
	(*CastMember)(nil),     // 0: grpc_demo.CastMember
	(*Movie)(nil),          // 1: grpc_demo.Movie
	(*MovieCatalogue)(nil), // 2: grpc_demo.MovieCatalogue
}
var file_movie_catalogue_proto_depIdxs = []int32{
	0, // 0: grpc_demo.Movie.castMembers:type_name -> grpc_demo.CastMember
	1, // 1: grpc_demo.MovieCatalogue.movies:type_name -> grpc_demo.Movie
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_movie_catalogue_proto_init() }
func file_movie_catalogue_proto_init() {
	if File_movie_catalogue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_catalogue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CastMember); i {
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
		file_movie_catalogue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_movie_catalogue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieCatalogue); i {
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
			RawDescriptor: file_movie_catalogue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_movie_catalogue_proto_goTypes,
		DependencyIndexes: file_movie_catalogue_proto_depIdxs,
		MessageInfos:      file_movie_catalogue_proto_msgTypes,
	}.Build()
	File_movie_catalogue_proto = out.File
	file_movie_catalogue_proto_rawDesc = nil
	file_movie_catalogue_proto_goTypes = nil
	file_movie_catalogue_proto_depIdxs = nil
}
