// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: restaurant.proto

package restaurant_restaurant

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Longitude   float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude    float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Address) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Address) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EntityId    string   `protobuf:"bytes,2,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Type        string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt   string   `protobuf:"bytes,7,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt   string   `protobuf:"bytes,8,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *Restaurant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Restaurant) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Restaurant) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Restaurant) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Restaurant) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Restaurant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AddRestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId string `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddRestaurantRequest) Reset() {
	*x = AddRestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRestaurantRequest) ProtoMessage() {}

func (x *AddRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRestaurantRequest.ProtoReflect.Descriptor instead.
func (*AddRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *AddRestaurantRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *AddRestaurantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddRestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant *Restaurant `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
}

func (x *AddRestaurantResponse) Reset() {
	*x = AddRestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRestaurantResponse) ProtoMessage() {}

func (x *AddRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRestaurantResponse.ProtoReflect.Descriptor instead.
func (*AddRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{3}
}

func (x *AddRestaurantResponse) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

type UpdateRestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant *Restaurant `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
}

func (x *UpdateRestaurantRequest) Reset() {
	*x = UpdateRestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRestaurantRequest) ProtoMessage() {}

func (x *UpdateRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRestaurantRequest.ProtoReflect.Descriptor instead.
func (*UpdateRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRestaurantRequest) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

type UpdateRestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant *Restaurant `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
}

func (x *UpdateRestaurantResponse) Reset() {
	*x = UpdateRestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRestaurantResponse) ProtoMessage() {}

func (x *UpdateRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRestaurantResponse.ProtoReflect.Descriptor instead.
func (*UpdateRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRestaurantResponse) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

type ListRestaurantByEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId string `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
}

func (x *ListRestaurantByEntityRequest) Reset() {
	*x = ListRestaurantByEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestaurantByEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestaurantByEntityRequest) ProtoMessage() {}

func (x *ListRestaurantByEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestaurantByEntityRequest.ProtoReflect.Descriptor instead.
func (*ListRestaurantByEntityRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{6}
}

func (x *ListRestaurantByEntityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

type ListRestaurantByEntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*Restaurant `protobuf:"bytes,1,rep,name=restaurants,proto3" json:"restaurants,omitempty"`
}

func (x *ListRestaurantByEntityResponse) Reset() {
	*x = ListRestaurantByEntityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestaurantByEntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestaurantByEntityResponse) ProtoMessage() {}

func (x *ListRestaurantByEntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestaurantByEntityResponse.ProtoReflect.Descriptor instead.
func (*ListRestaurantByEntityResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{7}
}

func (x *ListRestaurantByEntityResponse) GetRestaurants() []*Restaurant {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

type RemoveRestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string `protobuf:"bytes,1,opt,name=restaurant_id,proto3" json:"restaurant_id,omitempty"`
}

func (x *RemoveRestaurantRequest) Reset() {
	*x = RemoveRestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRestaurantRequest) ProtoMessage() {}

func (x *RemoveRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRestaurantRequest.ProtoReflect.Descriptor instead.
func (*RemoveRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRestaurantRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

type RemoveRestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRestaurantResponse) Reset() {
	*x = RemoveRestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRestaurantResponse) ProtoMessage() {}

func (x *RemoveRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRestaurantResponse.ProtoReflect.Descriptor instead.
func (*RemoveRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{9}
}

var File_restaurant_proto protoreflect.FileDescriptor

var file_restaurant_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x48, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x5a, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x18, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x3f,
	0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x67, 0x61, 0x62, 0x65, 0x6b,
	0x6b, 0x61, 0x64, 0x69, 0x72, 0x6f, 0x76, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_proto_rawDescOnce sync.Once
	file_restaurant_proto_rawDescData = file_restaurant_proto_rawDesc
)

func file_restaurant_proto_rawDescGZIP() []byte {
	file_restaurant_proto_rawDescOnce.Do(func() {
		file_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_proto_rawDescData)
	})
	return file_restaurant_proto_rawDescData
}

var file_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_restaurant_proto_goTypes = []interface{}{
	(*Address)(nil),                        // 0: restaurant_restaurant.Address
	(*Restaurant)(nil),                     // 1: restaurant_restaurant.Restaurant
	(*AddRestaurantRequest)(nil),           // 2: restaurant_restaurant.AddRestaurantRequest
	(*AddRestaurantResponse)(nil),          // 3: restaurant_restaurant.AddRestaurantResponse
	(*UpdateRestaurantRequest)(nil),        // 4: restaurant_restaurant.UpdateRestaurantRequest
	(*UpdateRestaurantResponse)(nil),       // 5: restaurant_restaurant.UpdateRestaurantResponse
	(*ListRestaurantByEntityRequest)(nil),  // 6: restaurant_restaurant.ListRestaurantByEntityRequest
	(*ListRestaurantByEntityResponse)(nil), // 7: restaurant_restaurant.ListRestaurantByEntityResponse
	(*RemoveRestaurantRequest)(nil),        // 8: restaurant_restaurant.RemoveRestaurantRequest
	(*RemoveRestaurantResponse)(nil),       // 9: restaurant_restaurant.RemoveRestaurantResponse
}
var file_restaurant_proto_depIdxs = []int32{
	0, // 0: restaurant_restaurant.Restaurant.address:type_name -> restaurant_restaurant.Address
	1, // 1: restaurant_restaurant.AddRestaurantResponse.restaurant:type_name -> restaurant_restaurant.Restaurant
	1, // 2: restaurant_restaurant.UpdateRestaurantRequest.restaurant:type_name -> restaurant_restaurant.Restaurant
	1, // 3: restaurant_restaurant.UpdateRestaurantResponse.restaurant:type_name -> restaurant_restaurant.Restaurant
	1, // 4: restaurant_restaurant.ListRestaurantByEntityResponse.restaurants:type_name -> restaurant_restaurant.Restaurant
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_restaurant_proto_init() }
func file_restaurant_proto_init() {
	if File_restaurant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_restaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_restaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRestaurantRequest); i {
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
		file_restaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRestaurantResponse); i {
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
		file_restaurant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRestaurantRequest); i {
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
		file_restaurant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRestaurantResponse); i {
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
		file_restaurant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestaurantByEntityRequest); i {
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
		file_restaurant_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestaurantByEntityResponse); i {
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
		file_restaurant_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRestaurantRequest); i {
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
		file_restaurant_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRestaurantResponse); i {
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
			RawDescriptor: file_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restaurant_proto_goTypes,
		DependencyIndexes: file_restaurant_proto_depIdxs,
		MessageInfos:      file_restaurant_proto_msgTypes,
	}.Build()
	File_restaurant_proto = out.File
	file_restaurant_proto_rawDesc = nil
	file_restaurant_proto_goTypes = nil
	file_restaurant_proto_depIdxs = nil
}
