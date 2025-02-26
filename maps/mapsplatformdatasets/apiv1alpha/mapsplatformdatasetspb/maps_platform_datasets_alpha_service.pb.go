// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/maps/mapsplatformdatasets/v1alpha/maps_platform_datasets_alpha_service.proto

package mapsplatformdatasetspb

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto protoreflect.FileDescriptor

var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d,
	0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xf0, 0x0b, 0x0a, 0x1b, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x56, 0x31, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x12, 0xcb, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x47, 0xda, 0x41, 0x0e, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x2c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30,
	0x3a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x12, 0xe8, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x54, 0xda, 0x41, 0x13, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x38, 0x3a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x32, 0x2d, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xb2, 0x01, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x3b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x34, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xe5, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x6c, 0x69, 0x73, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xc5, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xb9, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x2a, 0x33, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x57, 0xca, 0x41,
	0x23, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xa2, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x27, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x56, 0x31, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x66, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x70, 0x62, 0x3b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x4d,
	0x44, 0x56, 0x31, 0x41, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61,
	0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca,
	0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x4d, 0x61,
	0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_goTypes = []interface{}{
	(*CreateDatasetRequest)(nil),         // 0: google.maps.mapsplatformdatasets.v1alpha.CreateDatasetRequest
	(*UpdateDatasetMetadataRequest)(nil), // 1: google.maps.mapsplatformdatasets.v1alpha.UpdateDatasetMetadataRequest
	(*GetDatasetRequest)(nil),            // 2: google.maps.mapsplatformdatasets.v1alpha.GetDatasetRequest
	(*ListDatasetVersionsRequest)(nil),   // 3: google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsRequest
	(*ListDatasetsRequest)(nil),          // 4: google.maps.mapsplatformdatasets.v1alpha.ListDatasetsRequest
	(*DeleteDatasetRequest)(nil),         // 5: google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetRequest
	(*DeleteDatasetVersionRequest)(nil),  // 6: google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetVersionRequest
	(*Dataset)(nil),                      // 7: google.maps.mapsplatformdatasets.v1alpha.Dataset
	(*ListDatasetVersionsResponse)(nil),  // 8: google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsResponse
	(*ListDatasetsResponse)(nil),         // 9: google.maps.mapsplatformdatasets.v1alpha.ListDatasetsResponse
	(*emptypb.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_depIdxs = []int32{
	0,  // 0: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.CreateDataset:input_type -> google.maps.mapsplatformdatasets.v1alpha.CreateDatasetRequest
	1,  // 1: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.UpdateDatasetMetadata:input_type -> google.maps.mapsplatformdatasets.v1alpha.UpdateDatasetMetadataRequest
	2,  // 2: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.GetDataset:input_type -> google.maps.mapsplatformdatasets.v1alpha.GetDatasetRequest
	3,  // 3: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.ListDatasetVersions:input_type -> google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsRequest
	4,  // 4: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.ListDatasets:input_type -> google.maps.mapsplatformdatasets.v1alpha.ListDatasetsRequest
	5,  // 5: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.DeleteDataset:input_type -> google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetRequest
	6,  // 6: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.DeleteDatasetVersion:input_type -> google.maps.mapsplatformdatasets.v1alpha.DeleteDatasetVersionRequest
	7,  // 7: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.CreateDataset:output_type -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	7,  // 8: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.UpdateDatasetMetadata:output_type -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	7,  // 9: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.GetDataset:output_type -> google.maps.mapsplatformdatasets.v1alpha.Dataset
	8,  // 10: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.ListDatasetVersions:output_type -> google.maps.mapsplatformdatasets.v1alpha.ListDatasetVersionsResponse
	9,  // 11: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.ListDatasets:output_type -> google.maps.mapsplatformdatasets.v1alpha.ListDatasetsResponse
	10, // 12: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.DeleteDataset:output_type -> google.protobuf.Empty
	10, // 13: google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha.DeleteDatasetVersion:output_type -> google.protobuf.Empty
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_init()
}
func file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_init() {
	if File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto != nil {
		return
	}
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_init()
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_goTypes,
		DependencyIndexes: file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_depIdxs,
	}.Build()
	File_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto = out.File
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_rawDesc = nil
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_goTypes = nil
	file_google_maps_mapsplatformdatasets_v1alpha_maps_platform_datasets_alpha_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MapsPlatformDatasetsV1AlphaClient is the client API for MapsPlatformDatasetsV1Alpha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MapsPlatformDatasetsV1AlphaClient interface {
	// Create a new dataset for the specified project.
	CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Update the metadata for the dataset. To update the data use: UploadDataset.
	UpdateDatasetMetadata(ctx context.Context, in *UpdateDatasetMetadataRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Get the published or latest version of the dataset.
	GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// List all the versions of a dataset.
	ListDatasetVersions(ctx context.Context, in *ListDatasetVersionsRequest, opts ...grpc.CallOption) (*ListDatasetVersionsResponse, error)
	// List all the datasets for the specified project.
	ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error)
	// Delete the specified dataset and optionally all its corresponding
	// versions.
	DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete a specific version of the dataset.
	DeleteDatasetVersion(ctx context.Context, in *DeleteDatasetVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mapsPlatformDatasetsV1AlphaClient struct {
	cc grpc.ClientConnInterface
}

func NewMapsPlatformDatasetsV1AlphaClient(cc grpc.ClientConnInterface) MapsPlatformDatasetsV1AlphaClient {
	return &mapsPlatformDatasetsV1AlphaClient{cc}
}

func (c *mapsPlatformDatasetsV1AlphaClient) CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) UpdateDatasetMetadata(ctx context.Context, in *UpdateDatasetMetadataRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/UpdateDatasetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/GetDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) ListDatasetVersions(ctx context.Context, in *ListDatasetVersionsRequest, opts ...grpc.CallOption) (*ListDatasetVersionsResponse, error) {
	out := new(ListDatasetVersionsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/ListDatasetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error) {
	out := new(ListDatasetsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/ListDatasets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/DeleteDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsPlatformDatasetsV1AlphaClient) DeleteDatasetVersion(ctx context.Context, in *DeleteDatasetVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/DeleteDatasetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapsPlatformDatasetsV1AlphaServer is the server API for MapsPlatformDatasetsV1Alpha service.
type MapsPlatformDatasetsV1AlphaServer interface {
	// Create a new dataset for the specified project.
	CreateDataset(context.Context, *CreateDatasetRequest) (*Dataset, error)
	// Update the metadata for the dataset. To update the data use: UploadDataset.
	UpdateDatasetMetadata(context.Context, *UpdateDatasetMetadataRequest) (*Dataset, error)
	// Get the published or latest version of the dataset.
	GetDataset(context.Context, *GetDatasetRequest) (*Dataset, error)
	// List all the versions of a dataset.
	ListDatasetVersions(context.Context, *ListDatasetVersionsRequest) (*ListDatasetVersionsResponse, error)
	// List all the datasets for the specified project.
	ListDatasets(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error)
	// Delete the specified dataset and optionally all its corresponding
	// versions.
	DeleteDataset(context.Context, *DeleteDatasetRequest) (*emptypb.Empty, error)
	// Delete a specific version of the dataset.
	DeleteDatasetVersion(context.Context, *DeleteDatasetVersionRequest) (*emptypb.Empty, error)
}

// UnimplementedMapsPlatformDatasetsV1AlphaServer can be embedded to have forward compatible implementations.
type UnimplementedMapsPlatformDatasetsV1AlphaServer struct {
}

func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) CreateDataset(context.Context, *CreateDatasetRequest) (*Dataset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataset not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) UpdateDatasetMetadata(context.Context, *UpdateDatasetMetadataRequest) (*Dataset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatasetMetadata not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) GetDataset(context.Context, *GetDatasetRequest) (*Dataset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataset not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) ListDatasetVersions(context.Context, *ListDatasetVersionsRequest) (*ListDatasetVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatasetVersions not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) ListDatasets(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatasets not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) DeleteDataset(context.Context, *DeleteDatasetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataset not implemented")
}
func (*UnimplementedMapsPlatformDatasetsV1AlphaServer) DeleteDatasetVersion(context.Context, *DeleteDatasetVersionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatasetVersion not implemented")
}

func RegisterMapsPlatformDatasetsV1AlphaServer(s *grpc.Server, srv MapsPlatformDatasetsV1AlphaServer) {
	s.RegisterService(&_MapsPlatformDatasetsV1Alpha_serviceDesc, srv)
}

func _MapsPlatformDatasetsV1Alpha_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).CreateDataset(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_UpdateDatasetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatasetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).UpdateDatasetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/UpdateDatasetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).UpdateDatasetMetadata(ctx, req.(*UpdateDatasetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_GetDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).GetDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/GetDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).GetDataset(ctx, req.(*GetDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_ListDatasetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasetVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).ListDatasetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/ListDatasetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).ListDatasetVersions(ctx, req.(*ListDatasetVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_ListDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).ListDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/ListDatasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).ListDatasets(ctx, req.(*ListDatasetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_DeleteDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).DeleteDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/DeleteDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).DeleteDataset(ctx, req.(*DeleteDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MapsPlatformDatasetsV1Alpha_DeleteDatasetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsPlatformDatasetsV1AlphaServer).DeleteDatasetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha/DeleteDatasetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsPlatformDatasetsV1AlphaServer).DeleteDatasetVersion(ctx, req.(*DeleteDatasetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MapsPlatformDatasetsV1Alpha_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.mapsplatformdatasets.v1alpha.MapsPlatformDatasetsV1Alpha",
	HandlerType: (*MapsPlatformDatasetsV1AlphaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDataset",
			Handler:    _MapsPlatformDatasetsV1Alpha_CreateDataset_Handler,
		},
		{
			MethodName: "UpdateDatasetMetadata",
			Handler:    _MapsPlatformDatasetsV1Alpha_UpdateDatasetMetadata_Handler,
		},
		{
			MethodName: "GetDataset",
			Handler:    _MapsPlatformDatasetsV1Alpha_GetDataset_Handler,
		},
		{
			MethodName: "ListDatasetVersions",
			Handler:    _MapsPlatformDatasetsV1Alpha_ListDatasetVersions_Handler,
		},
		{
			MethodName: "ListDatasets",
			Handler:    _MapsPlatformDatasetsV1Alpha_ListDatasets_Handler,
		},
		{
			MethodName: "DeleteDataset",
			Handler:    _MapsPlatformDatasetsV1Alpha_DeleteDataset_Handler,
		},
		{
			MethodName: "DeleteDatasetVersion",
			Handler:    _MapsPlatformDatasetsV1Alpha_DeleteDatasetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/maps/mapsplatformdatasets/v1alpha/maps_platform_datasets_alpha_service.proto",
}
