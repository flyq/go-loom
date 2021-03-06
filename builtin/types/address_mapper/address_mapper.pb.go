// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/address_mapper/address_mapper.proto

/*
Package address_mapper is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/address_mapper/address_mapper.proto

It has these top-level messages:
	AddressMapperMapping
	AddressMapperInitRequest
	AddressMapperAddMappingRequest
	AddressMapperRemoveMappingRequest
	AddressMapperGetMappingRequest
	AddressMapperGetMappingResponse
*/
package address_mapper

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Bidirectional mapping between addresses (contracts or accounts) on two different blockchains.
type AddressMapperMapping struct {
	// Address on a blockchain
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Corresponding address on another blockchain
	To *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
}

func (m *AddressMapperMapping) Reset()         { *m = AddressMapperMapping{} }
func (m *AddressMapperMapping) String() string { return proto.CompactTextString(m) }
func (*AddressMapperMapping) ProtoMessage()    {}
func (*AddressMapperMapping) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{0}
}

func (m *AddressMapperMapping) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperMapping) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

type AddressMapperInitRequest struct {
}

func (m *AddressMapperInitRequest) Reset()         { *m = AddressMapperInitRequest{} }
func (m *AddressMapperInitRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperInitRequest) ProtoMessage()    {}
func (*AddressMapperInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{1}
}

type AddressMapperAddMappingRequest struct {
	// Address on a blockchain
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Corresponding address on another blockchain
	To *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
}

func (m *AddressMapperAddMappingRequest) Reset()         { *m = AddressMapperAddMappingRequest{} }
func (m *AddressMapperAddMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperAddMappingRequest) ProtoMessage()    {}
func (*AddressMapperAddMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{2}
}

func (m *AddressMapperAddMappingRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperAddMappingRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

type AddressMapperRemoveMappingRequest struct {
}

func (m *AddressMapperRemoveMappingRequest) Reset()         { *m = AddressMapperRemoveMappingRequest{} }
func (m *AddressMapperRemoveMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperRemoveMappingRequest) ProtoMessage()    {}
func (*AddressMapperRemoveMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{3}
}

type AddressMapperGetMappingRequest struct {
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
}

func (m *AddressMapperGetMappingRequest) Reset()         { *m = AddressMapperGetMappingRequest{} }
func (m *AddressMapperGetMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperGetMappingRequest) ProtoMessage()    {}
func (*AddressMapperGetMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{4}
}

func (m *AddressMapperGetMappingRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

type AddressMapperGetMappingResponse struct {
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
}

func (m *AddressMapperGetMappingResponse) Reset()         { *m = AddressMapperGetMappingResponse{} }
func (m *AddressMapperGetMappingResponse) String() string { return proto.CompactTextString(m) }
func (*AddressMapperGetMappingResponse) ProtoMessage()    {}
func (*AddressMapperGetMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorAddressMapper, []int{5}
}

func (m *AddressMapperGetMappingResponse) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperGetMappingResponse) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressMapperMapping)(nil), "AddressMapperMapping")
	proto.RegisterType((*AddressMapperInitRequest)(nil), "AddressMapperInitRequest")
	proto.RegisterType((*AddressMapperAddMappingRequest)(nil), "AddressMapperAddMappingRequest")
	proto.RegisterType((*AddressMapperRemoveMappingRequest)(nil), "AddressMapperRemoveMappingRequest")
	proto.RegisterType((*AddressMapperGetMappingRequest)(nil), "AddressMapperGetMappingRequest")
	proto.RegisterType((*AddressMapperGetMappingResponse)(nil), "AddressMapperGetMappingResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/address_mapper/address_mapper.proto", fileDescriptorAddressMapper)
}

var fileDescriptorAddressMapper = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x93, 0x4a, 0x33, 0x73, 0x4a, 0x32, 0xf3,
	0xf4, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0xe3, 0x73,
	0x13, 0x0b, 0x0a, 0x52, 0x8b, 0xd0, 0xb8, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x52, 0x06, 0x04,
	0x4c, 0x84, 0x98, 0x04, 0x26, 0x21, 0x3a, 0x94, 0xfc, 0xb8, 0x44, 0x1c, 0x21, 0x26, 0xf9, 0x82,
	0x0d, 0x02, 0x91, 0x99, 0x79, 0xe9, 0x42, 0x32, 0x5c, 0x2c, 0x69, 0x45, 0xf9, 0xb9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x1c, 0x7a, 0x50, 0x45, 0x41, 0x60, 0x51, 0x21, 0x09, 0x2e, 0xa6,
	0x92, 0x7c, 0x09, 0x26, 0x34, 0x39, 0xa6, 0x92, 0x7c, 0x25, 0x29, 0x2e, 0x09, 0x14, 0xf3, 0x3c,
	0xf3, 0x32, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x22, 0xb8, 0xe4, 0x50, 0xe4,
	0x1c, 0x53, 0x52, 0xa0, 0xd6, 0x41, 0x55, 0x90, 0x6d, 0xab, 0x32, 0x97, 0x22, 0x8a, 0xc9, 0x41,
	0xa9, 0xb9, 0xf9, 0x65, 0xa9, 0xa8, 0x86, 0x2b, 0xd9, 0xa1, 0x59, 0xef, 0x9e, 0x5a, 0x42, 0x8a,
	0xf5, 0x4a, 0x91, 0x5c, 0xf2, 0x38, 0xf5, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x92, 0xeb, 0xfe,
	0x24, 0x36, 0x70, 0x64, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x43, 0x88, 0xb9, 0x12,
	0x02, 0x00, 0x00,
}
