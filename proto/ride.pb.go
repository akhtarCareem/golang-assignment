// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.2
// source: ride.proto

package proto

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

type Ride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Distance    int32  `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	Cost        int32  `protobuf:"varint,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *Ride) Reset() {
	*x = Ride{}
	mi := &file_ride_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ride) ProtoMessage() {}

func (x *Ride) ProtoReflect() protoreflect.Message {
	mi := &file_ride_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ride.ProtoReflect.Descriptor instead.
func (*Ride) Descriptor() ([]byte, []int) {
	return file_ride_proto_rawDescGZIP(), []int{0}
}

func (x *Ride) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Ride) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Ride) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Ride) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UpdateRideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RideId string `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	Ride   *Ride  `protobuf:"bytes,2,opt,name=ride,proto3" json:"ride,omitempty"`
}

func (x *UpdateRideRequest) Reset() {
	*x = UpdateRideRequest{}
	mi := &file_ride_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRideRequest) ProtoMessage() {}

func (x *UpdateRideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ride_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRideRequest.ProtoReflect.Descriptor instead.
func (*UpdateRideRequest) Descriptor() ([]byte, []int) {
	return file_ride_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRideRequest) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

func (x *UpdateRideRequest) GetRide() *Ride {
	if x != nil {
		return x.Ride
	}
	return nil
}

type UpdateRideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateRideResponse) Reset() {
	*x = UpdateRideResponse{}
	mi := &file_ride_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRideResponse) ProtoMessage() {}

func (x *UpdateRideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRideResponse.ProtoReflect.Descriptor instead.
func (*UpdateRideResponse) Descriptor() ([]byte, []int) {
	return file_ride_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRideResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ride_proto protoreflect.FileDescriptor

var file_ride_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x04, 0x52, 0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x04,
	0x72, 0x69, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x53, 0x0a, 0x0c, 0x52, 0x69, 0x64, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x64, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ride_proto_rawDescOnce sync.Once
	file_ride_proto_rawDescData = file_ride_proto_rawDesc
)

func file_ride_proto_rawDescGZIP() []byte {
	file_ride_proto_rawDescOnce.Do(func() {
		file_ride_proto_rawDescData = protoimpl.X.CompressGZIP(file_ride_proto_rawDescData)
	})
	return file_ride_proto_rawDescData
}

var file_ride_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ride_proto_goTypes = []any{
	(*Ride)(nil),               // 0: proto.Ride
	(*UpdateRideRequest)(nil),  // 1: proto.UpdateRideRequest
	(*UpdateRideResponse)(nil), // 2: proto.UpdateRideResponse
}
var file_ride_proto_depIdxs = []int32{
	0, // 0: proto.UpdateRideRequest.ride:type_name -> proto.Ride
	1, // 1: proto.RidesService.UpdateRide:input_type -> proto.UpdateRideRequest
	2, // 2: proto.RidesService.UpdateRide:output_type -> proto.UpdateRideResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ride_proto_init() }
func file_ride_proto_init() {
	if File_ride_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ride_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ride_proto_goTypes,
		DependencyIndexes: file_ride_proto_depIdxs,
		MessageInfos:      file_ride_proto_msgTypes,
	}.Build()
	File_ride_proto = out.File
	file_ride_proto_rawDesc = nil
	file_ride_proto_goTypes = nil
	file_ride_proto_depIdxs = nil
}
