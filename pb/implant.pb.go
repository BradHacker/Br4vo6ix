// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: implant.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task_TaskType int32

const (
	Task_NOOP   Task_TaskType = 0
	Task_CMD    Task_TaskType = 1
	Task_SCRIPT Task_TaskType = 2
)

// Enum value maps for Task_TaskType.
var (
	Task_TaskType_name = map[int32]string{
		0: "NOOP",
		1: "CMD",
		2: "SCRIPT",
	}
	Task_TaskType_value = map[string]int32{
		"NOOP":   0,
		"CMD":    1,
		"SCRIPT": 2,
	}
)

func (x Task_TaskType) Enum() *Task_TaskType {
	p := new(Task_TaskType)
	*p = x
	return p
}

func (x Task_TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_implant_proto_enumTypes[0].Descriptor()
}

func (Task_TaskType) Type() protoreflect.EnumType {
	return &file_implant_proto_enumTypes[0]
}

func (x Task_TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_TaskType.Descriptor instead.
func (Task_TaskType) EnumDescriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{1, 0}
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Pid       int64                  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	SentAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{0}
}

func (x *Heartbeat) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *Heartbeat) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Heartbeat) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	HeartbeatUuid string        `protobuf:"bytes,2,opt,name=heartbeat_uuid,json=heartbeatUuid,proto3" json:"heartbeat_uuid,omitempty"`
	Type          Task_TaskType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.Task_TaskType" json:"type,omitempty"`
	Payload       string        `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Task) GetHeartbeatUuid() string {
	if x != nil {
		return x.HeartbeatUuid
	}
	return ""
}

func (x *Task) GetType() Task_TaskType {
	if x != nil {
		return x.Type
	}
	return Task_NOOP
}

func (x *Task) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Stdout string `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TaskResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *TaskResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

var File_implant_proto protoreflect.FileDescriptor

var file_implant_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x29, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x02, 0x22, 0x52, 0x0a,
	0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_implant_proto_rawDescOnce sync.Once
	file_implant_proto_rawDescData = file_implant_proto_rawDesc
)

func file_implant_proto_rawDescGZIP() []byte {
	file_implant_proto_rawDescOnce.Do(func() {
		file_implant_proto_rawDescData = protoimpl.X.CompressGZIP(file_implant_proto_rawDescData)
	})
	return file_implant_proto_rawDescData
}

var file_implant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_implant_proto_goTypes = []interface{}{
	(Task_TaskType)(0),            // 0: proto.Task.TaskType
	(*Heartbeat)(nil),             // 1: proto.Heartbeat
	(*Task)(nil),                  // 2: proto.Task
	(*TaskResponse)(nil),          // 3: proto.TaskResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_implant_proto_depIdxs = []int32{
	4, // 0: proto.Heartbeat.sent_at:type_name -> google.protobuf.Timestamp
	0, // 1: proto.Task.type:type_name -> proto.Task.TaskType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_implant_proto_init() }
func file_implant_proto_init() {
	if File_implant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_implant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_implant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_implant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_implant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_implant_proto_goTypes,
		DependencyIndexes: file_implant_proto_depIdxs,
		EnumInfos:         file_implant_proto_enumTypes,
		MessageInfos:      file_implant_proto_msgTypes,
	}.Build()
	File_implant_proto = out.File
	file_implant_proto_rawDesc = nil
	file_implant_proto_goTypes = nil
	file_implant_proto_depIdxs = nil
}
