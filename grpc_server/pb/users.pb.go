// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: users.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
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
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *UsersRequest) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    uint64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	User    *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *UserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UsersStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRep []*UserResponse `protobuf:"bytes,1,rep,name=userRep,proto3" json:"userRep,omitempty"`
}

func (x *UsersStreamResponse) Reset() {
	*x = UsersStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersStreamResponse) ProtoMessage() {}

func (x *UsersStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersStreamResponse.ProtoReflect.Descriptor instead.
func (*UsersStreamResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *UsersStreamResponse) GetUserRep() []*UserResponse {
	if x != nil {
		return x.UserRep
	}
	return nil
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x13, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0d, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x36, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x71,
	0x6c, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x71, 0x6c, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_users_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: pbfiles.User
	(*UsersRequest)(nil),        // 1: pbfiles.UsersRequest
	(*UserResponse)(nil),        // 2: pbfiles.UserResponse
	(*UsersStreamResponse)(nil), // 3: pbfiles.UsersStreamResponse
}
var file_users_proto_depIdxs = []int32{
	0, // 0: pbfiles.UsersRequest.users:type_name -> pbfiles.User
	0, // 1: pbfiles.UserResponse.user:type_name -> pbfiles.User
	2, // 2: pbfiles.UsersStreamResponse.userRep:type_name -> pbfiles.UserResponse
	0, // 3: pbfiles.UserService.LoginUser:input_type -> pbfiles.User
	1, // 4: pbfiles.UserService.GetUserStream:input_type -> pbfiles.UsersRequest
	0, // 5: pbfiles.UserService.LoginUserSql:input_type -> pbfiles.User
	1, // 6: pbfiles.UserService.GetUserStreamSql:input_type -> pbfiles.UsersRequest
	2, // 7: pbfiles.UserService.LoginUser:output_type -> pbfiles.UserResponse
	3, // 8: pbfiles.UserService.GetUserStream:output_type -> pbfiles.UsersStreamResponse
	2, // 9: pbfiles.UserService.LoginUserSql:output_type -> pbfiles.UserResponse
	3, // 10: pbfiles.UserService.GetUserStreamSql:output_type -> pbfiles.UsersStreamResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRequest); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersStreamResponse); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserStream(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUserStreamClient, error)
	LoginUserSql(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserStreamSql(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUserStreamSqlClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pbfiles.UserService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStream(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUserStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/pbfiles.UserService/GetUserStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserStreamClient interface {
	Recv() (*UsersStreamResponse, error)
	grpc.ClientStream
}

type userServiceGetUserStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserStreamClient) Recv() (*UsersStreamResponse, error) {
	m := new(UsersStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) LoginUserSql(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pbfiles.UserService/LoginUserSql", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStreamSql(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUserStreamSqlClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[1], "/pbfiles.UserService/GetUserStreamSql", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserStreamSqlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserStreamSqlClient interface {
	Recv() (*UsersStreamResponse, error)
	grpc.ClientStream
}

type userServiceGetUserStreamSqlClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserStreamSqlClient) Recv() (*UsersStreamResponse, error) {
	m := new(UsersStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	LoginUser(context.Context, *User) (*UserResponse, error)
	GetUserStream(*UsersRequest, UserService_GetUserStreamServer) error
	LoginUserSql(context.Context, *User) (*UserResponse, error)
	GetUserStreamSql(*UsersRequest, UserService_GetUserStreamSqlServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) LoginUser(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUserStream(*UsersRequest, UserService_GetUserStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserStream not implemented")
}
func (*UnimplementedUserServiceServer) LoginUserSql(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUserSql not implemented")
}
func (*UnimplementedUserServiceServer) GetUserStreamSql(*UsersRequest, UserService_GetUserStreamSqlServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserStreamSql not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.UserService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserStream(m, &userServiceGetUserStreamServer{stream})
}

type UserService_GetUserStreamServer interface {
	Send(*UsersStreamResponse) error
	grpc.ServerStream
}

type userServiceGetUserStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserStreamServer) Send(m *UsersStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_LoginUserSql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUserSql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.UserService/LoginUserSql",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUserSql(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStreamSql_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserStreamSql(m, &userServiceGetUserStreamSqlServer{stream})
}

type UserService_GetUserStreamSqlServer interface {
	Send(*UsersStreamResponse) error
	grpc.ServerStream
}

type userServiceGetUserStreamSqlServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserStreamSqlServer) Send(m *UsersStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfiles.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _UserService_LoginUser_Handler,
		},
		{
			MethodName: "LoginUserSql",
			Handler:    _UserService_LoginUserSql_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserStream",
			Handler:       _UserService_GetUserStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetUserStreamSql",
			Handler:       _UserService_GetUserStreamSql_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users.proto",
}