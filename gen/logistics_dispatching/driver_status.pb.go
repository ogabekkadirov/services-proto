// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: driver_status.proto

package logistics_dispatching

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

type DriverLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DriverId  string  `protobuf:"bytes,2,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	CreatedAt string  `protobuf:"bytes,5,opt,name=created_at,proto3" json:"created_at,omitempty"`
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocation.ProtoReflect.Descriptor instead.
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{0}
}

func (x *DriverLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverLocation) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DriverLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DriverLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DriverLocation) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type DriverStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId  string `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *DriverStatus) Reset() {
	*x = DriverStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverStatus) ProtoMessage() {}

func (x *DriverStatus) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverStatus.ProtoReflect.Descriptor instead.
func (*DriverStatus) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{1}
}

func (x *DriverStatus) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DriverStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DriverStatus) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DriverStatus) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type TrackDriverLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId  string  `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *TrackDriverLocationRequest) Reset() {
	*x = TrackDriverLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackDriverLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackDriverLocationRequest) ProtoMessage() {}

func (x *TrackDriverLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackDriverLocationRequest.ProtoReflect.Descriptor instead.
func (*TrackDriverLocationRequest) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{2}
}

func (x *TrackDriverLocationRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *TrackDriverLocationRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *TrackDriverLocationRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type TrackDriverLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *DriverLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *TrackDriverLocationResponse) Reset() {
	*x = TrackDriverLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackDriverLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackDriverLocationResponse) ProtoMessage() {}

func (x *TrackDriverLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackDriverLocationResponse.ProtoReflect.Descriptor instead.
func (*TrackDriverLocationResponse) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{3}
}

func (x *TrackDriverLocationResponse) GetLocation() *DriverLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type ListDriverLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId  string `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	StartTime string `protobuf:"bytes,2,opt,name=start_time,proto3" json:"start_time,omitempty"`
	EndTime   string `protobuf:"bytes,3,opt,name=end_time,proto3" json:"end_time,omitempty"`
}

func (x *ListDriverLocationRequest) Reset() {
	*x = ListDriverLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDriverLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriverLocationRequest) ProtoMessage() {}

func (x *ListDriverLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriverLocationRequest.ProtoReflect.Descriptor instead.
func (*ListDriverLocationRequest) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{4}
}

func (x *ListDriverLocationRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *ListDriverLocationRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ListDriverLocationRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ListDriverLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*DriverLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *ListDriverLocationResponse) Reset() {
	*x = ListDriverLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDriverLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriverLocationResponse) ProtoMessage() {}

func (x *ListDriverLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriverLocationResponse.ProtoReflect.Descriptor instead.
func (*ListDriverLocationResponse) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{5}
}

func (x *ListDriverLocationResponse) GetLocations() []*DriverLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type UpdateDriverStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateDriverStatusRequest) Reset() {
	*x = UpdateDriverStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverStatusRequest) ProtoMessage() {}

func (x *UpdateDriverStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateDriverStatusRequest) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDriverStatusRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *UpdateDriverStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateDriverStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDriverStatusResponse) Reset() {
	*x = UpdateDriverStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverStatusResponse) ProtoMessage() {}

func (x *UpdateDriverStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateDriverStatusResponse) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{7}
}

type GetDriverStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
}

func (x *GetDriverStatusRequest) Reset() {
	*x = GetDriverStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverStatusRequest) ProtoMessage() {}

func (x *GetDriverStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverStatusRequest.ProtoReflect.Descriptor instead.
func (*GetDriverStatusRequest) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{8}
}

func (x *GetDriverStatusRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

type GetDriverStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *DriverStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetDriverStatusResponse) Reset() {
	*x = GetDriverStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_status_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverStatusResponse) ProtoMessage() {}

func (x *GetDriverStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_status_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverStatusResponse.ProtoReflect.Descriptor instead.
func (*GetDriverStatusResponse) Descriptor() ([]byte, []int) {
	return file_driver_status_proto_rawDescGZIP(), []int{9}
}

func (x *GetDriverStatusResponse) GetStatus() *DriverStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_driver_status_proto protoreflect.FileDescriptor

var file_driver_status_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x98, 0x01, 0x0a,
	0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x74,
	0x0a, 0x1a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x61, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x51, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x36, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x67, 0x61, 0x62, 0x65, 0x6b, 0x6b, 0x61, 0x64, 0x69, 0x72, 0x6f, 0x76, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_status_proto_rawDescOnce sync.Once
	file_driver_status_proto_rawDescData = file_driver_status_proto_rawDesc
)

func file_driver_status_proto_rawDescGZIP() []byte {
	file_driver_status_proto_rawDescOnce.Do(func() {
		file_driver_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_status_proto_rawDescData)
	})
	return file_driver_status_proto_rawDescData
}

var file_driver_status_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_driver_status_proto_goTypes = []interface{}{
	(*DriverLocation)(nil),              // 0: logistics_dispatching.DriverLocation
	(*DriverStatus)(nil),                // 1: logistics_dispatching.DriverStatus
	(*TrackDriverLocationRequest)(nil),  // 2: logistics_dispatching.TrackDriverLocationRequest
	(*TrackDriverLocationResponse)(nil), // 3: logistics_dispatching.TrackDriverLocationResponse
	(*ListDriverLocationRequest)(nil),   // 4: logistics_dispatching.ListDriverLocationRequest
	(*ListDriverLocationResponse)(nil),  // 5: logistics_dispatching.ListDriverLocationResponse
	(*UpdateDriverStatusRequest)(nil),   // 6: logistics_dispatching.UpdateDriverStatusRequest
	(*UpdateDriverStatusResponse)(nil),  // 7: logistics_dispatching.UpdateDriverStatusResponse
	(*GetDriverStatusRequest)(nil),      // 8: logistics_dispatching.GetDriverStatusRequest
	(*GetDriverStatusResponse)(nil),     // 9: logistics_dispatching.GetDriverStatusResponse
}
var file_driver_status_proto_depIdxs = []int32{
	0, // 0: logistics_dispatching.TrackDriverLocationResponse.location:type_name -> logistics_dispatching.DriverLocation
	0, // 1: logistics_dispatching.ListDriverLocationResponse.locations:type_name -> logistics_dispatching.DriverLocation
	1, // 2: logistics_dispatching.GetDriverStatusResponse.status:type_name -> logistics_dispatching.DriverStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_driver_status_proto_init() }
func file_driver_status_proto_init() {
	if File_driver_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverLocation); i {
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
		file_driver_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverStatus); i {
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
		file_driver_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackDriverLocationRequest); i {
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
		file_driver_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackDriverLocationResponse); i {
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
		file_driver_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDriverLocationRequest); i {
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
		file_driver_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDriverLocationResponse); i {
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
		file_driver_status_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverStatusRequest); i {
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
		file_driver_status_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverStatusResponse); i {
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
		file_driver_status_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverStatusRequest); i {
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
		file_driver_status_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverStatusResponse); i {
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
			RawDescriptor: file_driver_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_driver_status_proto_goTypes,
		DependencyIndexes: file_driver_status_proto_depIdxs,
		MessageInfos:      file_driver_status_proto_msgTypes,
	}.Build()
	File_driver_status_proto = out.File
	file_driver_status_proto_rawDesc = nil
	file_driver_status_proto_goTypes = nil
	file_driver_status_proto_depIdxs = nil
}
