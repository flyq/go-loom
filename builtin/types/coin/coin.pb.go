// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

package coin

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

type Economy struct {
	TotalSupply          *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Economy) Reset()         { *m = Economy{} }
func (m *Economy) String() string { return proto.CompactTextString(m) }
func (*Economy) ProtoMessage()    {}
func (*Economy) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{0}
}
func (m *Economy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Economy.Unmarshal(m, b)
}
func (m *Economy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Economy.Marshal(b, m, deterministic)
}
func (dst *Economy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Economy.Merge(dst, src)
}
func (m *Economy) XXX_Size() int {
	return xxx_messageInfo_Economy.Size(m)
}
func (m *Economy) XXX_DiscardUnknown() {
	xxx_messageInfo_Economy.DiscardUnknown(m)
}

var xxx_messageInfo_Economy proto.InternalMessageInfo

func (m *Economy) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type Account struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance              *types.BigUInt `protobuf:"bytes,2,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Account) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type Allowance struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender              *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Allowance) Reset()         { *m = Allowance{} }
func (m *Allowance) String() string { return proto.CompactTextString(m) }
func (*Allowance) ProtoMessage()    {}
func (*Allowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{2}
}
func (m *Allowance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Allowance.Unmarshal(m, b)
}
func (m *Allowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Allowance.Marshal(b, m, deterministic)
}
func (dst *Allowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allowance.Merge(dst, src)
}
func (m *Allowance) XXX_Size() int {
	return xxx_messageInfo_Allowance.Size(m)
}
func (m *Allowance) XXX_DiscardUnknown() {
	xxx_messageInfo_Allowance.DiscardUnknown(m)
}

var xxx_messageInfo_Allowance proto.InternalMessageInfo

func (m *Allowance) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Allowance) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *Allowance) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type InitialAccount struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance              uint64         `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitialAccount) Reset()         { *m = InitialAccount{} }
func (m *InitialAccount) String() string { return proto.CompactTextString(m) }
func (*InitialAccount) ProtoMessage()    {}
func (*InitialAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{3}
}
func (m *InitialAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialAccount.Unmarshal(m, b)
}
func (m *InitialAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialAccount.Marshal(b, m, deterministic)
}
func (dst *InitialAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialAccount.Merge(dst, src)
}
func (m *InitialAccount) XXX_Size() int {
	return xxx_messageInfo_InitialAccount.Size(m)
}
func (m *InitialAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialAccount.DiscardUnknown(m)
}

var xxx_messageInfo_InitialAccount proto.InternalMessageInfo

func (m *InitialAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *InitialAccount) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type InitRequest struct {
	Accounts             []*InitialAccount `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{4}
}
func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (dst *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(dst, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetAccounts() []*InitialAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// read only
type TotalSupplyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalSupplyRequest) Reset()         { *m = TotalSupplyRequest{} }
func (m *TotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyRequest) ProtoMessage()    {}
func (*TotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{5}
}
func (m *TotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalSupplyRequest.Unmarshal(m, b)
}
func (m *TotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalSupplyRequest.Marshal(b, m, deterministic)
}
func (dst *TotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyRequest.Merge(dst, src)
}
func (m *TotalSupplyRequest) XXX_Size() int {
	return xxx_messageInfo_TotalSupplyRequest.Size(m)
}
func (m *TotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyRequest proto.InternalMessageInfo

type TotalSupplyResponse struct {
	TotalSupply          *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TotalSupplyResponse) Reset()         { *m = TotalSupplyResponse{} }
func (m *TotalSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyResponse) ProtoMessage()    {}
func (*TotalSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{6}
}
func (m *TotalSupplyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalSupplyResponse.Unmarshal(m, b)
}
func (m *TotalSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalSupplyResponse.Marshal(b, m, deterministic)
}
func (dst *TotalSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyResponse.Merge(dst, src)
}
func (m *TotalSupplyResponse) XXX_Size() int {
	return xxx_messageInfo_TotalSupplyResponse.Size(m)
}
func (m *TotalSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyResponse proto.InternalMessageInfo

func (m *TotalSupplyResponse) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type BalanceOfRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceOfRequest) Reset()         { *m = BalanceOfRequest{} }
func (m *BalanceOfRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceOfRequest) ProtoMessage()    {}
func (*BalanceOfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{7}
}
func (m *BalanceOfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceOfRequest.Unmarshal(m, b)
}
func (m *BalanceOfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceOfRequest.Marshal(b, m, deterministic)
}
func (dst *BalanceOfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceOfRequest.Merge(dst, src)
}
func (m *BalanceOfRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceOfRequest.Size(m)
}
func (m *BalanceOfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceOfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceOfRequest proto.InternalMessageInfo

func (m *BalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type BalanceOfResponse struct {
	Balance              *types.BigUInt `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceOfResponse) Reset()         { *m = BalanceOfResponse{} }
func (m *BalanceOfResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceOfResponse) ProtoMessage()    {}
func (*BalanceOfResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{8}
}
func (m *BalanceOfResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceOfResponse.Unmarshal(m, b)
}
func (m *BalanceOfResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceOfResponse.Marshal(b, m, deterministic)
}
func (dst *BalanceOfResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceOfResponse.Merge(dst, src)
}
func (m *BalanceOfResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceOfResponse.Size(m)
}
func (m *BalanceOfResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceOfResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceOfResponse proto.InternalMessageInfo

func (m *BalanceOfResponse) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AllowanceRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender              *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllowanceRequest) Reset()         { *m = AllowanceRequest{} }
func (m *AllowanceRequest) String() string { return proto.CompactTextString(m) }
func (*AllowanceRequest) ProtoMessage()    {}
func (*AllowanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{9}
}
func (m *AllowanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowanceRequest.Unmarshal(m, b)
}
func (m *AllowanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowanceRequest.Marshal(b, m, deterministic)
}
func (dst *AllowanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceRequest.Merge(dst, src)
}
func (m *AllowanceRequest) XXX_Size() int {
	return xxx_messageInfo_AllowanceRequest.Size(m)
}
func (m *AllowanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceRequest proto.InternalMessageInfo

func (m *AllowanceRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AllowanceRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

type AllowanceResponse struct {
	Amount               *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllowanceResponse) Reset()         { *m = AllowanceResponse{} }
func (m *AllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*AllowanceResponse) ProtoMessage()    {}
func (*AllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{10}
}
func (m *AllowanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowanceResponse.Unmarshal(m, b)
}
func (m *AllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowanceResponse.Marshal(b, m, deterministic)
}
func (dst *AllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceResponse.Merge(dst, src)
}
func (m *AllowanceResponse) XXX_Size() int {
	return xxx_messageInfo_AllowanceResponse.Size(m)
}
func (m *AllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceResponse proto.InternalMessageInfo

func (m *AllowanceResponse) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// volatile
type ApproveRequest struct {
	Spender              *types.Address `protobuf:"bytes,1,opt,name=spender" json:"spender,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApproveRequest) Reset()         { *m = ApproveRequest{} }
func (m *ApproveRequest) String() string { return proto.CompactTextString(m) }
func (*ApproveRequest) ProtoMessage()    {}
func (*ApproveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{11}
}
func (m *ApproveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApproveRequest.Unmarshal(m, b)
}
func (m *ApproveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApproveRequest.Marshal(b, m, deterministic)
}
func (dst *ApproveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApproveRequest.Merge(dst, src)
}
func (m *ApproveRequest) XXX_Size() int {
	return xxx_messageInfo_ApproveRequest.Size(m)
}
func (m *ApproveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApproveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApproveRequest proto.InternalMessageInfo

func (m *ApproveRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *ApproveRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type ApproveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApproveResponse) Reset()         { *m = ApproveResponse{} }
func (m *ApproveResponse) String() string { return proto.CompactTextString(m) }
func (*ApproveResponse) ProtoMessage()    {}
func (*ApproveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{12}
}
func (m *ApproveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApproveResponse.Unmarshal(m, b)
}
func (m *ApproveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApproveResponse.Marshal(b, m, deterministic)
}
func (dst *ApproveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApproveResponse.Merge(dst, src)
}
func (m *ApproveResponse) XXX_Size() int {
	return xxx_messageInfo_ApproveResponse.Size(m)
}
func (m *ApproveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApproveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApproveResponse proto.InternalMessageInfo

type TransferRequest struct {
	To                   *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{13}
}
func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (dst *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(dst, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{14}
}
func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (dst *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(dst, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

type TransferFromRequest struct {
	From                 *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To                   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferFromRequest) Reset()         { *m = TransferFromRequest{} }
func (m *TransferFromRequest) String() string { return proto.CompactTextString(m) }
func (*TransferFromRequest) ProtoMessage()    {}
func (*TransferFromRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{15}
}
func (m *TransferFromRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFromRequest.Unmarshal(m, b)
}
func (m *TransferFromRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFromRequest.Marshal(b, m, deterministic)
}
func (dst *TransferFromRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFromRequest.Merge(dst, src)
}
func (m *TransferFromRequest) XXX_Size() int {
	return xxx_messageInfo_TransferFromRequest.Size(m)
}
func (m *TransferFromRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFromRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFromRequest proto.InternalMessageInfo

func (m *TransferFromRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferFromRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferFromRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferFromResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferFromResponse) Reset()         { *m = TransferFromResponse{} }
func (m *TransferFromResponse) String() string { return proto.CompactTextString(m) }
func (*TransferFromResponse) ProtoMessage()    {}
func (*TransferFromResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_8ca771d27ef71903, []int{16}
}
func (m *TransferFromResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFromResponse.Unmarshal(m, b)
}
func (m *TransferFromResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFromResponse.Marshal(b, m, deterministic)
}
func (dst *TransferFromResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFromResponse.Merge(dst, src)
}
func (m *TransferFromResponse) XXX_Size() int {
	return xxx_messageInfo_TransferFromResponse.Size(m)
}
func (m *TransferFromResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFromResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFromResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Economy)(nil), "Economy")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Allowance)(nil), "Allowance")
	proto.RegisterType((*InitialAccount)(nil), "InitialAccount")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
	proto.RegisterType((*TotalSupplyRequest)(nil), "TotalSupplyRequest")
	proto.RegisterType((*TotalSupplyResponse)(nil), "TotalSupplyResponse")
	proto.RegisterType((*BalanceOfRequest)(nil), "BalanceOfRequest")
	proto.RegisterType((*BalanceOfResponse)(nil), "BalanceOfResponse")
	proto.RegisterType((*AllowanceRequest)(nil), "AllowanceRequest")
	proto.RegisterType((*AllowanceResponse)(nil), "AllowanceResponse")
	proto.RegisterType((*ApproveRequest)(nil), "ApproveRequest")
	proto.RegisterType((*ApproveResponse)(nil), "ApproveResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*TransferFromRequest)(nil), "TransferFromRequest")
	proto.RegisterType((*TransferFromResponse)(nil), "TransferFromResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto", fileDescriptor_coin_8ca771d27ef71903)
}

var fileDescriptor_coin_8ca771d27ef71903 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0x49, 0x5a, 0xbb, 0xf5, 0x46, 0xba, 0xbb, 0x69, 0x91, 0x20, 0x22, 0xcb, 0x3c, 0x15,
	0x8a, 0x89, 0xac, 0xa8, 0x20, 0xbe, 0x64, 0x41, 0xa1, 0x42, 0x11, 0x62, 0xed, 0xab, 0x24, 0xd9,
	0xd9, 0x35, 0x38, 0x99, 0x3b, 0x9d, 0x99, 0xb8, 0xec, 0xb7, 0x97, 0xfc, 0x4f, 0x5a, 0xe2, 0xc6,
	0x97, 0x81, 0xb9, 0x77, 0xee, 0xf9, 0x9d, 0xc9, 0x49, 0x02, 0x9f, 0xb6, 0x89, 0xfe, 0x95, 0x45,
	0x6e, 0x8c, 0xa9, 0xc7, 0x10, 0x53, 0x4e, 0xf5, 0x0e, 0xe5, 0x6f, 0x6f, 0x8b, 0xaf, 0xf3, 0xad,
	0x17, 0x65, 0x09, 0xd3, 0x09, 0xf7, 0xf4, 0x5e, 0x50, 0xe5, 0xc5, 0x98, 0xf0, 0x62, 0x71, 0x85,
	0x44, 0x8d, 0x2f, 0xde, 0x1c, 0x98, 0x2e, 0xa7, 0x8a, 0xb5, 0x9c, 0x20, 0xef, 0x61, 0xf2, 0x39,
	0x46, 0x8e, 0xe9, 0xde, 0xbe, 0x82, 0x67, 0x1a, 0x75, 0xc8, 0x7e, 0xaa, 0x4c, 0x08, 0xb6, 0x77,
	0x8c, 0x85, 0x71, 0x69, 0x2d, 0x4f, 0xdd, 0x55, 0xb2, 0xfd, 0x71, 0xcd, 0x75, 0x60, 0x15, 0xdd,
	0xef, 0x45, 0x93, 0xdc, 0xc0, 0xc4, 0x8f, 0x63, 0xcc, 0xb8, 0xb6, 0x5f, 0xc1, 0x13, 0xdc, 0x71,
	0x2a, 0x9b, 0x01, 0x7f, 0xbd, 0x96, 0x54, 0xa9, 0xa0, 0x2c, 0xdb, 0x04, 0x26, 0x51, 0xc8, 0x42,
	0x1e, 0x53, 0xc7, 0x7c, 0x20, 0x59, 0x37, 0xc8, 0x3d, 0x3c, 0xf5, 0x19, 0xc3, 0x5d, 0xbe, 0x19,
	0x23, 0xa8, 0x04, 0xe5, 0x6b, 0x2a, 0x1b, 0xc1, 0xfa, 0x44, 0xdd, 0xb0, 0x17, 0x70, 0x12, 0xa6,
	0xb9, 0x3d, 0xe7, 0xe8, 0x01, 0xb3, 0xaa, 0x93, 0xaf, 0x70, 0x76, 0xcd, 0x13, 0x9d, 0x84, 0x6c,
	0xec, 0x45, 0x9c, 0xfe, 0x45, 0x8e, 0x5b, 0xfb, 0x1f, 0xc1, 0xca, 0xb5, 0x02, 0x7a, 0x9f, 0x51,
	0xa5, 0xed, 0x2b, 0x38, 0x0d, 0x4b, 0x4d, 0xe5, 0x18, 0x8b, 0xa3, 0x4b, 0x6b, 0x39, 0x75, 0xfb,
	0xac, 0xa0, 0x39, 0x40, 0x2e, 0xc0, 0xbe, 0x6d, 0x1f, 0x6c, 0x25, 0x41, 0x56, 0x70, 0xde, 0xab,
	0x2a, 0x81, 0x5c, 0xd1, 0xff, 0xcb, 0x68, 0x09, 0xb3, 0x55, 0x69, 0xf0, 0xdb, 0xa6, 0xb6, 0x76,
	0xe0, 0x8e, 0xe4, 0x03, 0xcc, 0x3b, 0x33, 0x15, 0xb5, 0x93, 0xa0, 0x31, 0x94, 0xe0, 0x1d, 0xcc,
	0x9a, 0x04, 0x47, 0xc2, 0xc6, 0x04, 0x49, 0xde, 0xc1, 0xbc, 0xa3, 0x5b, 0x19, 0x6a, 0xd3, 0x35,
	0x06, 0xd2, 0xbd, 0x83, 0x33, 0x5f, 0x08, 0x89, 0x7f, 0x1a, 0x33, 0x1d, 0x98, 0x71, 0xf8, 0xad,
	0x31, 0x07, 0x74, 0xe7, 0x30, 0x6d, 0x74, 0x4b, 0x33, 0xe4, 0x06, 0xa6, 0xb7, 0x32, 0xe4, 0x6a,
	0x43, 0x65, 0xcd, 0x72, 0xc0, 0xd4, 0xf8, 0x08, 0x63, 0x6a, 0x1c, 0x41, 0xb0, 0x61, 0xd6, 0xca,
	0x55, 0x08, 0x84, 0xf3, 0xba, 0xf6, 0x45, 0x62, 0x5a, 0x63, 0x5e, 0xc2, 0xf1, 0x46, 0x62, 0xfa,
	0x08, 0x54, 0x54, 0x2b, 0x13, 0xe6, 0x3f, 0x4d, 0x0c, 0x7d, 0x1c, 0xcf, 0xe1, 0xa2, 0x0f, 0x2c,
	0x8d, 0x44, 0x27, 0xc5, 0x5f, 0xe3, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x03, 0xcb,
	0xb2, 0xa7, 0x04, 0x00, 0x00,
}
