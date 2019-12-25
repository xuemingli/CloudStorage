// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	ReqSignup
	RespSignup
	ReqSignin
	RespSignin
	ReqUserInfo
	RespUserInfo
	ReqUserFile
	RespUserFile
	ReqUserFileRename
	RespUserFileRename
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ReqSignup struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *ReqSignup) Reset()                    { *m = ReqSignup{} }
func (m *ReqSignup) String() string            { return proto1.CompactTextString(m) }
func (*ReqSignup) ProtoMessage()               {}
func (*ReqSignup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignup struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RespSignup) Reset()                    { *m = RespSignup{} }
func (m *RespSignup) String() string            { return proto1.CompactTextString(m) }
func (*RespSignup) ProtoMessage()               {}
func (*RespSignup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RespSignup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSignin struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *ReqSignin) Reset()                    { *m = ReqSignin{} }
func (m *ReqSignin) String() string            { return proto1.CompactTextString(m) }
func (*ReqSignin) ProtoMessage()               {}
func (*ReqSignin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqSignin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignin struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RespSignin) Reset()                    { *m = RespSignin{} }
func (m *RespSignin) String() string            { return proto1.CompactTextString(m) }
func (*RespSignin) ProtoMessage()               {}
func (*RespSignin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RespSignin) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RespSignin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUserInfo struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *ReqUserInfo) Reset()                    { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string            { return proto1.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()               {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespUserInfo struct {
	Code         int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	SignupAt     string `protobuf:"bytes,6,opt,name=signupAt" json:"signupAt,omitempty"`
	LastActiveAt string `protobuf:"bytes,7,opt,name=lastActiveAt" json:"lastActiveAt,omitempty"`
	Status       int32  `protobuf:"varint,8,opt,name=status" json:"status,omitempty"`
}

func (m *RespUserInfo) Reset()                    { *m = RespUserInfo{} }
func (m *RespUserInfo) String() string            { return proto1.CompactTextString(m) }
func (*RespUserInfo) ProtoMessage()               {}
func (*RespUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RespUserInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RespUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RespUserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RespUserInfo) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *RespUserInfo) GetLastActiveAt() string {
	if m != nil {
		return m.LastActiveAt
	}
	return ""
}

func (m *RespUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqUserFile struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ReqUserFile) Reset()                    { *m = ReqUserFile{} }
func (m *ReqUserFile) String() string            { return proto1.CompactTextString(m) }
func (*ReqUserFile) ProtoMessage()               {}
func (*ReqUserFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RespUserFile struct {
	Code     int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	FileData []byte `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
}

func (m *RespUserFile) Reset()                    { *m = RespUserFile{} }
func (m *RespUserFile) String() string            { return proto1.CompactTextString(m) }
func (*RespUserFile) ProtoMessage()               {}
func (*RespUserFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RespUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFile) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

type ReqUserFileRename struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Filehash    string `protobuf:"bytes,2,opt,name=filehash" json:"filehash,omitempty"`
	NewFileName string `protobuf:"bytes,3,opt,name=newFileName" json:"newFileName,omitempty"`
}

func (m *ReqUserFileRename) Reset()                    { *m = ReqUserFileRename{} }
func (m *ReqUserFileRename) String() string            { return proto1.CompactTextString(m) }
func (*ReqUserFileRename) ProtoMessage()               {}
func (*ReqUserFileRename) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReqUserFileRename) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFileRename) GetFilehash() string {
	if m != nil {
		return m.Filehash
	}
	return ""
}

func (m *ReqUserFileRename) GetNewFileName() string {
	if m != nil {
		return m.NewFileName
	}
	return ""
}

type RespUserFileRename struct {
	Code     int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	FileData []byte `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
}

func (m *RespUserFileRename) Reset()                    { *m = RespUserFileRename{} }
func (m *RespUserFileRename) String() string            { return proto1.CompactTextString(m) }
func (*RespUserFileRename) ProtoMessage()               {}
func (*RespUserFileRename) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RespUserFileRename) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFileRename) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFileRename) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

func init() {
	proto1.RegisterType((*ReqSignup)(nil), "proto.ReqSignup")
	proto1.RegisterType((*RespSignup)(nil), "proto.RespSignup")
	proto1.RegisterType((*ReqSignin)(nil), "proto.ReqSignin")
	proto1.RegisterType((*RespSignin)(nil), "proto.RespSignin")
	proto1.RegisterType((*ReqUserInfo)(nil), "proto.ReqUserInfo")
	proto1.RegisterType((*RespUserInfo)(nil), "proto.RespUserInfo")
	proto1.RegisterType((*ReqUserFile)(nil), "proto.ReqUserFile")
	proto1.RegisterType((*RespUserFile)(nil), "proto.RespUserFile")
	proto1.RegisterType((*ReqUserFileRename)(nil), "proto.ReqUserFileRename")
	proto1.RegisterType((*RespUserFileRename)(nil), "proto.RespUserFileRename")
}

func init() { proto1.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x4d, 0x77, 0x37, 0xd9, 0x76, 0x5a, 0x21, 0xd6, 0x54, 0xc8, 0xe4, 0xb4, 0xf2, 0x09, 0x2e,
	0x8b, 0x04, 0x82, 0x03, 0x17, 0x54, 0x81, 0x40, 0x5c, 0x10, 0xf2, 0x0a, 0x89, 0x13, 0x92, 0xe9,
	0xce, 0xb6, 0x16, 0x89, 0x93, 0xc6, 0x6e, 0xfb, 0x49, 0xfc, 0x14, 0x1f, 0x83, 0x6c, 0xc7, 0x49,
	0x9a, 0x56, 0x95, 0x8a, 0xf6, 0xd4, 0xbe, 0x37, 0x79, 0xcf, 0xcf, 0xe3, 0x19, 0x80, 0xb5, 0xc6,
	0xea, 0xa6, 0xac, 0x0a, 0x53, 0x90, 0xd8, 0xfd, 0xb0, 0x0f, 0x30, 0xe2, 0xb8, 0xba, 0x95, 0x0b,
	0xb5, 0x2e, 0x49, 0x0a, 0x43, 0xfb, 0x85, 0x12, 0x39, 0xd2, 0xc1, 0xf5, 0xe0, 0xf9, 0x88, 0x37,
	0xd8, 0xd6, 0x4a, 0xa1, 0xf5, 0xb6, 0xa8, 0xee, 0xe8, 0x99, 0xaf, 0x05, 0xcc, 0xde, 0x01, 0x70,
	0xd4, 0x65, 0xed, 0x42, 0xe0, 0x62, 0x5e, 0xdc, 0x79, 0x87, 0x98, 0xbb, 0xff, 0x84, 0xc2, 0x65,
	0x8e, 0x5a, 0x8b, 0x05, 0xd6, 0xe2, 0x00, 0x3b, 0x01, 0xa4, 0xfa, 0xef, 0x00, 0xdf, 0xda, 0x00,
	0x52, 0x1d, 0x0c, 0x30, 0x85, 0xd8, 0x14, 0xbf, 0x51, 0xd5, 0x52, 0x0f, 0xba, 0xb1, 0xce, 0x77,
	0x63, 0xbd, 0x80, 0x31, 0xc7, 0xd5, 0x77, 0x8d, 0xd5, 0x17, 0x75, 0x5f, 0x1c, 0x0b, 0xc6, 0xfe,
	0x0e, 0x60, 0x62, 0x4f, 0x6f, 0x3e, 0x3e, 0xa9, 0x01, 0x3b, 0xd6, 0xe7, 0xbd, 0x3b, 0x4f, 0x21,
	0xc6, 0x5c, 0xc8, 0x8c, 0x5e, 0xf8, 0xd4, 0x0e, 0x58, 0xb6, 0x5c, 0x16, 0x0a, 0x69, 0xec, 0x59,
	0x07, 0xac, 0x8f, 0x76, 0x0f, 0x30, 0x33, 0x34, 0xf1, 0x3e, 0x01, 0x13, 0x06, 0x93, 0x4c, 0x68,
	0x33, 0x9b, 0x1b, 0xb9, 0xc1, 0x99, 0xa1, 0x97, 0xae, 0xbe, 0xc3, 0x91, 0xa7, 0x90, 0x68, 0x23,
	0xcc, 0x5a, 0xd3, 0xa1, 0xcb, 0x5d, 0x23, 0xf6, 0xbe, 0xe9, 0xc4, 0x27, 0x99, 0xe1, 0xd1, 0x27,
	0x9a, 0x42, 0x9c, 0xc9, 0x5c, 0x1a, 0x77, 0xc5, 0x98, 0x7b, 0xc0, 0x7e, 0xb4, 0xed, 0x71, 0x0e,
	0x27, 0xb7, 0xe7, 0x5e, 0x66, 0xf8, 0x51, 0x18, 0xe1, 0xda, 0x33, 0xe1, 0x0d, 0x66, 0x39, 0x5c,
	0x75, 0xa2, 0x71, 0x0c, 0x73, 0x72, 0x6c, 0x86, 0xac, 0x78, 0x29, 0xf4, 0x32, 0xcc, 0x50, 0xc0,
	0xe4, 0x1a, 0xc6, 0x0a, 0xb7, 0xd6, 0xe8, 0x6b, 0xfb, 0x14, 0x5d, 0x8a, 0xfd, 0x04, 0xd2, 0xbd,
	0x48, 0x7d, 0xde, 0x83, 0x5d, 0xe7, 0xd5, 0x9f, 0x33, 0x18, 0x5b, 0xf3, 0x5b, 0xac, 0x36, 0x72,
	0x8e, 0xe4, 0x25, 0x24, 0xf5, 0x4a, 0x3d, 0xf6, 0x4b, 0x7b, 0xd3, 0xac, 0x6a, 0x7a, 0xd5, 0x30,
	0x61, 0xef, 0x58, 0x14, 0x04, 0x52, 0xf5, 0x05, 0x52, 0xed, 0x09, 0xa4, 0x62, 0x11, 0x79, 0x03,
	0xc3, 0x76, 0x6a, 0x5b, 0x49, 0xe0, 0xd2, 0x27, 0x1d, 0x51, 0x20, 0x59, 0x44, 0xde, 0xc2, 0x28,
	0x34, 0x41, 0xf7, 0x75, 0x96, 0xdc, 0xd3, 0x59, 0x92, 0x45, 0xe4, 0x33, 0x3c, 0xea, 0x35, 0x8f,
	0xee, 0x8b, 0x7d, 0x25, 0x7d, 0x76, 0xc0, 0xc2, 0x97, 0x58, 0xf4, 0x2b, 0x71, 0xb5, 0xd7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xe8, 0xde, 0x9b, 0xd1, 0x04, 0x00, 0x00,
}