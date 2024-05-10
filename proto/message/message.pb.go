// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/message/message.proto

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

type ElementMessageType int32

const (
	ElementMessageType_UNKNOWN           ElementMessageType = 0
	ElementMessageType_ERROR             ElementMessageType = 1
	ElementMessageType_CHAT_MESSAGE      ElementMessageType = 2
	ElementMessageType_PLAY              ElementMessageType = 3
	ElementMessageType_PAUSE             ElementMessageType = 4
	ElementMessageType_CHECK             ElementMessageType = 5
	ElementMessageType_TOO_FAST          ElementMessageType = 6
	ElementMessageType_TOO_SLOW          ElementMessageType = 7
	ElementMessageType_CHANGE_RATE       ElementMessageType = 8
	ElementMessageType_CHANGE_SEEK       ElementMessageType = 9
	ElementMessageType_CURRENT_CHANGED   ElementMessageType = 10
	ElementMessageType_MOVIES_CHANGED    ElementMessageType = 11
	ElementMessageType_PEOPLE_CHANGED    ElementMessageType = 12
	ElementMessageType_SYNC_MOVIE_STATUS ElementMessageType = 13
	ElementMessageType_CURRENT_EXPIRED   ElementMessageType = 14
)

// Enum value maps for ElementMessageType.
var (
	ElementMessageType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "ERROR",
		2:  "CHAT_MESSAGE",
		3:  "PLAY",
		4:  "PAUSE",
		5:  "CHECK",
		6:  "TOO_FAST",
		7:  "TOO_SLOW",
		8:  "CHANGE_RATE",
		9:  "CHANGE_SEEK",
		10: "CURRENT_CHANGED",
		11: "MOVIES_CHANGED",
		12: "PEOPLE_CHANGED",
		13: "SYNC_MOVIE_STATUS",
		14: "CURRENT_EXPIRED",
	}
	ElementMessageType_value = map[string]int32{
		"UNKNOWN":           0,
		"ERROR":             1,
		"CHAT_MESSAGE":      2,
		"PLAY":              3,
		"PAUSE":             4,
		"CHECK":             5,
		"TOO_FAST":          6,
		"TOO_SLOW":          7,
		"CHANGE_RATE":       8,
		"CHANGE_SEEK":       9,
		"CURRENT_CHANGED":   10,
		"MOVIES_CHANGED":    11,
		"PEOPLE_CHANGED":    12,
		"SYNC_MOVIE_STATUS": 13,
		"CURRENT_EXPIRED":   14,
	}
)

func (x ElementMessageType) Enum() *ElementMessageType {
	p := new(ElementMessageType)
	*p = x
	return p
}

func (x ElementMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ElementMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_message_proto_enumTypes[0].Descriptor()
}

func (ElementMessageType) Type() protoreflect.EnumType {
	return &file_proto_message_message_proto_enumTypes[0]
}

func (x ElementMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ElementMessageType.Descriptor instead.
func (ElementMessageType) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{0}
}

type ChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender  *Sender `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatResp) Reset() {
	*x = ChatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResp) ProtoMessage() {}

func (x *ChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResp.ProtoReflect.Descriptor instead.
func (*ChatResp) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *ChatResp) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *ChatResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Userid   string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *Sender) Reset() {
	*x = Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sender) ProtoMessage() {}

func (x *Sender) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sender.ProtoReflect.Descriptor instead.
func (*Sender) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *Sender) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Sender) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type MovieStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playing bool    `protobuf:"varint,1,opt,name=playing,proto3" json:"playing,omitempty"`
	Seek    float64 `protobuf:"fixed64,2,opt,name=seek,proto3" json:"seek,omitempty"`
	Rate    float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *MovieStatus) Reset() {
	*x = MovieStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieStatus) ProtoMessage() {}

func (x *MovieStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieStatus.ProtoReflect.Descriptor instead.
func (*MovieStatus) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *MovieStatus) GetPlaying() bool {
	if x != nil {
		return x.Playing
	}
	return false
}

func (x *MovieStatus) GetSeek() float64 {
	if x != nil {
		return x.Seek
	}
	return 0
}

func (x *MovieStatus) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type MovieStatusChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender *Sender      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Status *MovieStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MovieStatusChanged) Reset() {
	*x = MovieStatusChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieStatusChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieStatusChanged) ProtoMessage() {}

func (x *MovieStatusChanged) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieStatusChanged.ProtoReflect.Descriptor instead.
func (*MovieStatusChanged) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{3}
}

func (x *MovieStatusChanged) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *MovieStatusChanged) GetStatus() *MovieStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type CheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *MovieStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ExpireId uint64       `protobuf:"varint,2,opt,name=expireId,proto3" json:"expireId,omitempty"`
}

func (x *CheckReq) Reset() {
	*x = CheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReq) ProtoMessage() {}

func (x *CheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReq.ProtoReflect.Descriptor instead.
func (*CheckReq) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{4}
}

func (x *CheckReq) GetStatus() *MovieStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CheckReq) GetExpireId() uint64 {
	if x != nil {
		return x.ExpireId
	}
	return 0
}

type ElementMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                 ElementMessageType  `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ElementMessageType" json:"type,omitempty"`
	Time                 int64               `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ChatReq              string              `protobuf:"bytes,4,opt,name=chatReq,proto3" json:"chatReq,omitempty"`
	ChatResp             *ChatResp           `protobuf:"bytes,5,opt,name=chatResp,proto3" json:"chatResp,omitempty"`
	ChangeMovieStatusReq *MovieStatus        `protobuf:"bytes,6,opt,name=changeMovieStatusReq,proto3" json:"changeMovieStatusReq,omitempty"`
	MovieStatusChanged   *MovieStatusChanged `protobuf:"bytes,7,opt,name=movieStatusChanged,proto3" json:"movieStatusChanged,omitempty"`
	ChangeSeekReq        float64             `protobuf:"fixed64,8,opt,name=changeSeekReq,proto3" json:"changeSeekReq,omitempty"`
	CheckReq             *CheckReq           `protobuf:"bytes,9,opt,name=checkReq,proto3" json:"checkReq,omitempty"`
	PeopleChanged        int64               `protobuf:"varint,11,opt,name=peopleChanged,proto3" json:"peopleChanged,omitempty"`
	MoviesChanged        *Sender             `protobuf:"bytes,12,opt,name=moviesChanged,proto3" json:"moviesChanged,omitempty"`
	CurrentChanged       *Sender             `protobuf:"bytes,13,opt,name=currentChanged,proto3" json:"currentChanged,omitempty"`
}

func (x *ElementMessage) Reset() {
	*x = ElementMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElementMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElementMessage) ProtoMessage() {}

func (x *ElementMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElementMessage.ProtoReflect.Descriptor instead.
func (*ElementMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{5}
}

func (x *ElementMessage) GetType() ElementMessageType {
	if x != nil {
		return x.Type
	}
	return ElementMessageType_UNKNOWN
}

func (x *ElementMessage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ElementMessage) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ElementMessage) GetChatReq() string {
	if x != nil {
		return x.ChatReq
	}
	return ""
}

func (x *ElementMessage) GetChatResp() *ChatResp {
	if x != nil {
		return x.ChatResp
	}
	return nil
}

func (x *ElementMessage) GetChangeMovieStatusReq() *MovieStatus {
	if x != nil {
		return x.ChangeMovieStatusReq
	}
	return nil
}

func (x *ElementMessage) GetMovieStatusChanged() *MovieStatusChanged {
	if x != nil {
		return x.MovieStatusChanged
	}
	return nil
}

func (x *ElementMessage) GetChangeSeekReq() float64 {
	if x != nil {
		return x.ChangeSeekReq
	}
	return 0
}

func (x *ElementMessage) GetCheckReq() *CheckReq {
	if x != nil {
		return x.CheckReq
	}
	return nil
}

func (x *ElementMessage) GetPeopleChanged() int64 {
	if x != nil {
		return x.PeopleChanged
	}
	return 0
}

func (x *ElementMessage) GetMoviesChanged() *Sender {
	if x != nil {
		return x.MoviesChanged
	}
	return nil
}

func (x *ElementMessage) GetCurrentChanged() *Sender {
	if x != nil {
		return x.CurrentChanged
	}
	return nil
}

var File_proto_message_message_proto protoreflect.FileDescriptor

var file_proto_message_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x25, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22,
	0x4f, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x65, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x22, 0x67, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x49, 0x64, 0x22, 0xa8, 0x04,
	0x0a, 0x0e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x46, 0x0a, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x49, 0x0a, 0x12, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52,
	0x12, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x65,
	0x6b, 0x52, 0x65, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x52, 0x08, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x0d,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x0d, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x35, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2a, 0x85, 0x02, 0x0a, 0x12, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x54, 0x5f,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x41,
	0x59, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4f, 0x4f,
	0x5f, 0x46, 0x41, 0x53, 0x54, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4f, 0x4f, 0x5f, 0x53,
	0x4c, 0x4f, 0x57, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x52, 0x41, 0x54, 0x45, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x53, 0x45, 0x45, 0x4b, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e,
	0x4d, 0x4f, 0x56, 0x49, 0x45, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0b,
	0x12, 0x12, 0x0a, 0x0e, 0x50, 0x45, 0x4f, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x44, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4d, 0x4f, 0x56,
	0x49, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0e,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_message_proto_rawDescOnce sync.Once
	file_proto_message_message_proto_rawDescData = file_proto_message_message_proto_rawDesc
)

func file_proto_message_message_proto_rawDescGZIP() []byte {
	file_proto_message_message_proto_rawDescOnce.Do(func() {
		file_proto_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_message_proto_rawDescData)
	})
	return file_proto_message_message_proto_rawDescData
}

var file_proto_message_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_message_message_proto_goTypes = []interface{}{
	(ElementMessageType)(0),    // 0: proto.ElementMessageType
	(*ChatResp)(nil),           // 1: proto.ChatResp
	(*Sender)(nil),             // 2: proto.Sender
	(*MovieStatus)(nil),        // 3: proto.MovieStatus
	(*MovieStatusChanged)(nil), // 4: proto.MovieStatusChanged
	(*CheckReq)(nil),           // 5: proto.CheckReq
	(*ElementMessage)(nil),     // 6: proto.ElementMessage
}
var file_proto_message_message_proto_depIdxs = []int32{
	2,  // 0: proto.ChatResp.sender:type_name -> proto.Sender
	2,  // 1: proto.MovieStatusChanged.sender:type_name -> proto.Sender
	3,  // 2: proto.MovieStatusChanged.status:type_name -> proto.MovieStatus
	3,  // 3: proto.CheckReq.status:type_name -> proto.MovieStatus
	0,  // 4: proto.ElementMessage.type:type_name -> proto.ElementMessageType
	1,  // 5: proto.ElementMessage.chatResp:type_name -> proto.ChatResp
	3,  // 6: proto.ElementMessage.changeMovieStatusReq:type_name -> proto.MovieStatus
	4,  // 7: proto.ElementMessage.movieStatusChanged:type_name -> proto.MovieStatusChanged
	5,  // 8: proto.ElementMessage.checkReq:type_name -> proto.CheckReq
	2,  // 9: proto.ElementMessage.moviesChanged:type_name -> proto.Sender
	2,  // 10: proto.ElementMessage.currentChanged:type_name -> proto.Sender
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_message_message_proto_init() }
func file_proto_message_message_proto_init() {
	if File_proto_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResp); i {
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
		file_proto_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sender); i {
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
		file_proto_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieStatus); i {
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
		file_proto_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieStatusChanged); i {
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
		file_proto_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckReq); i {
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
		file_proto_message_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElementMessage); i {
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
			RawDescriptor: file_proto_message_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_message_proto_goTypes,
		DependencyIndexes: file_proto_message_message_proto_depIdxs,
		EnumInfos:         file_proto_message_message_proto_enumTypes,
		MessageInfos:      file_proto_message_message_proto_msgTypes,
	}.Build()
	File_proto_message_message_proto = out.File
	file_proto_message_message_proto_rawDesc = nil
	file_proto_message_message_proto_goTypes = nil
	file_proto_message_message_proto_depIdxs = nil
}
