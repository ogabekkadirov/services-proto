// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: logistics_staff/service.proto

package logistics_staff

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_logistics_staff_service_proto protoreflect.FileDescriptor

var file_logistics_staff_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x1a, 0x1c, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89,
	0x04, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x61, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x26, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x67, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x70, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x67, 0x61, 0x62, 0x65, 0x6b, 0x6b,
	0x61, 0x64, 0x69, 0x72, 0x6f, 0x76, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_logistics_staff_service_proto_goTypes = []interface{}{
	(*RegisterDriverRequest)(nil),       // 0: logistics_staff.RegisterDriverRequest
	(*LoginDriverRequest)(nil),          // 1: logistics_staff.LoginDriverRequest
	(*ConfirmSmsCodeRequest)(nil),       // 2: logistics_staff.ConfirmSmsCodeRequest
	(*GetDriverProfileRequest)(nil),     // 3: logistics_staff.GetDriverProfileRequest
	(*UpdateDriverProfileRequest)(nil),  // 4: logistics_staff.UpdateDriverProfileRequest
	(*RegisterDriverResponse)(nil),      // 5: logistics_staff.RegisterDriverResponse
	(*LoginDriverResponse)(nil),         // 6: logistics_staff.LoginDriverResponse
	(*ConfirmSmsCodeResponse)(nil),      // 7: logistics_staff.ConfirmSmsCodeResponse
	(*GetDriverProfileResponse)(nil),    // 8: logistics_staff.GetDriverProfileResponse
	(*UpdateDriverProfileResponse)(nil), // 9: logistics_staff.UpdateDriverProfileResponse
}
var file_logistics_staff_service_proto_depIdxs = []int32{
	0, // 0: logistics_staff.StaffService.RegisterDriver:input_type -> logistics_staff.RegisterDriverRequest
	1, // 1: logistics_staff.StaffService.LoginDriver:input_type -> logistics_staff.LoginDriverRequest
	2, // 2: logistics_staff.StaffService.ConfirmSmsCode:input_type -> logistics_staff.ConfirmSmsCodeRequest
	3, // 3: logistics_staff.StaffService.GetDriverProfile:input_type -> logistics_staff.GetDriverProfileRequest
	4, // 4: logistics_staff.StaffService.UpdatpDriverProfile:input_type -> logistics_staff.UpdateDriverProfileRequest
	5, // 5: logistics_staff.StaffService.RegisterDriver:output_type -> logistics_staff.RegisterDriverResponse
	6, // 6: logistics_staff.StaffService.LoginDriver:output_type -> logistics_staff.LoginDriverResponse
	7, // 7: logistics_staff.StaffService.ConfirmSmsCode:output_type -> logistics_staff.ConfirmSmsCodeResponse
	8, // 8: logistics_staff.StaffService.GetDriverProfile:output_type -> logistics_staff.GetDriverProfileResponse
	9, // 9: logistics_staff.StaffService.UpdatpDriverProfile:output_type -> logistics_staff.UpdateDriverProfileResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logistics_staff_service_proto_init() }
func file_logistics_staff_service_proto_init() {
	if File_logistics_staff_service_proto != nil {
		return
	}
	file_logistics_staff_driver_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logistics_staff_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logistics_staff_service_proto_goTypes,
		DependencyIndexes: file_logistics_staff_service_proto_depIdxs,
	}.Build()
	File_logistics_staff_service_proto = out.File
	file_logistics_staff_service_proto_rawDesc = nil
	file_logistics_staff_service_proto_goTypes = nil
	file_logistics_staff_service_proto_depIdxs = nil
}