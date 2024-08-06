// This file originates from the SatoshiLabs Trezor `common` repository at:
//   https://github.com/trezor/trezor-common/blob/master/protob/messages-ethereum.proto
// dated 28.05.2019, commit 893fd219d4a01bcffa0cd9cfa631856371ec5aa9.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: messages-ethereum.proto

package trezor

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

// *
// Request: Ask device for public key corresponding to address_n path
// @start
// @next EthereumPublicKey
// @next Failure
type EthereumGetPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN    []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`          // BIP-32 path to derive the key from master node
	ShowDisplay *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"` // optionally show on display before sending the result
}

func (x *EthereumGetPublicKey) Reset() {
	*x = EthereumGetPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumGetPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumGetPublicKey) ProtoMessage() {}

func (x *EthereumGetPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumGetPublicKey.ProtoReflect.Descriptor instead.
func (*EthereumGetPublicKey) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{0}
}

func (x *EthereumGetPublicKey) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *EthereumGetPublicKey) GetShowDisplay() bool {
	if x != nil && x.ShowDisplay != nil {
		return *x.ShowDisplay
	}
	return false
}

// *
// Response: Contains public key derived from device private seed
// @end
type EthereumPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *HDNodeType `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"` // BIP32 public node
	Xpub *string     `protobuf:"bytes,2,opt,name=xpub" json:"xpub,omitempty"` // serialized form of public node
}

func (x *EthereumPublicKey) Reset() {
	*x = EthereumPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumPublicKey) ProtoMessage() {}

func (x *EthereumPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumPublicKey.ProtoReflect.Descriptor instead.
func (*EthereumPublicKey) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{1}
}

func (x *EthereumPublicKey) GetNode() *HDNodeType {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *EthereumPublicKey) GetXpub() string {
	if x != nil && x.Xpub != nil {
		return *x.Xpub
	}
	return ""
}

// *
// Request: Ask device for Gori address corresponding to address_n path
// @start
// @next EthereumAddress
// @next Failure
type EthereumGetAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN    []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`          // BIP-32 path to derive the key from master node
	ShowDisplay *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"` // optionally show on display before sending the result
}

func (x *EthereumGetAddress) Reset() {
	*x = EthereumGetAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumGetAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumGetAddress) ProtoMessage() {}

func (x *EthereumGetAddress) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumGetAddress.ProtoReflect.Descriptor instead.
func (*EthereumGetAddress) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{2}
}

func (x *EthereumGetAddress) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *EthereumGetAddress) GetShowDisplay() bool {
	if x != nil && x.ShowDisplay != nil {
		return *x.ShowDisplay
	}
	return false
}

// *
// Response: Contains an Gori address derived from device private seed
// @end
type EthereumAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressBin []byte  `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"` // Gori address as 20 bytes (legacy firmwares)
	AddressHex *string `protobuf:"bytes,2,opt,name=addressHex" json:"addressHex,omitempty"` // Gori address as hex string (newer firmwares)
}

func (x *EthereumAddress) Reset() {
	*x = EthereumAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumAddress) ProtoMessage() {}

func (x *EthereumAddress) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumAddress.ProtoReflect.Descriptor instead.
func (*EthereumAddress) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{3}
}

func (x *EthereumAddress) GetAddressBin() []byte {
	if x != nil {
		return x.AddressBin
	}
	return nil
}

func (x *EthereumAddress) GetAddressHex() string {
	if x != nil && x.AddressHex != nil {
		return *x.AddressHex
	}
	return ""
}

// *
// Request: Ask device to sign transaction
// All fields are optional from the protocol's point of view. Each field defaults to value `0` if missing.
// Note: the first at most 1024 bytes of data MUST be transmitted as part of this message.
// @start
// @next EthereumTxRequest
// @next Failure
type EthereumSignTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`                          // BIP-32 path to derive the key from master node
	Nonce            []byte   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`                                                 // <=256 bit unsigned big endian
	GasPrice         []byte   `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`                           // <=256 bit unsigned big endian (in wei)
	GasLimit         []byte   `protobuf:"bytes,4,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`                           // <=256 bit unsigned big endian
	ToBin            []byte   `protobuf:"bytes,5,opt,name=toBin" json:"toBin,omitempty"`                                                 // recipient address (20 bytes, legacy firmware)
	ToHex            *string  `protobuf:"bytes,11,opt,name=toHex" json:"toHex,omitempty"`                                                // recipient address (hex string, newer firmware)
	Value            []byte   `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`                                                 // <=256 bit unsigned big endian (in wei)
	DataInitialChunk []byte   `protobuf:"bytes,7,opt,name=data_initial_chunk,json=dataInitialChunk" json:"data_initial_chunk,omitempty"` // The initial data chunk (<= 1024 bytes)
	DataLength       *uint32  `protobuf:"varint,8,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`                    // Length of transaction payload
	ChainId          *uint32  `protobuf:"varint,9,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`                             // Chain Id for EIP 155
	TxType           *uint32  `protobuf:"varint,10,opt,name=tx_type,json=txType" json:"tx_type,omitempty"`                               // (only for Wanchain)
}

func (x *EthereumSignTx) Reset() {
	*x = EthereumSignTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumSignTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumSignTx) ProtoMessage() {}

func (x *EthereumSignTx) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumSignTx.ProtoReflect.Descriptor instead.
func (*EthereumSignTx) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{4}
}

func (x *EthereumSignTx) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *EthereumSignTx) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EthereumSignTx) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *EthereumSignTx) GetGasLimit() []byte {
	if x != nil {
		return x.GasLimit
	}
	return nil
}

func (x *EthereumSignTx) GetToBin() []byte {
	if x != nil {
		return x.ToBin
	}
	return nil
}

func (x *EthereumSignTx) GetToHex() string {
	if x != nil && x.ToHex != nil {
		return *x.ToHex
	}
	return ""
}

func (x *EthereumSignTx) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EthereumSignTx) GetDataInitialChunk() []byte {
	if x != nil {
		return x.DataInitialChunk
	}
	return nil
}

func (x *EthereumSignTx) GetDataLength() uint32 {
	if x != nil && x.DataLength != nil {
		return *x.DataLength
	}
	return 0
}

func (x *EthereumSignTx) GetChainId() uint32 {
	if x != nil && x.ChainId != nil {
		return *x.ChainId
	}
	return 0
}

func (x *EthereumSignTx) GetTxType() uint32 {
	if x != nil && x.TxType != nil {
		return *x.TxType
	}
	return 0
}

// *
// Response: Device asks for more data from transaction payload, or returns the signature.
// If data_length is set, device awaits that many more bytes of payload.
// Otherwise, the signature_* fields contain the computed transaction signature. All three fields will be present.
// @end
// @next EthereumTxAck
type EthereumTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataLength *uint32 `protobuf:"varint,1,opt,name=data_length,json=dataLength" json:"data_length,omitempty"` // Number of bytes being requested (<= 1024)
	SignatureV *uint32 `protobuf:"varint,2,opt,name=signature_v,json=signatureV" json:"signature_v,omitempty"` // Computed signature (recovery parameter, limited to 27 or 28)
	SignatureR []byte  `protobuf:"bytes,3,opt,name=signature_r,json=signatureR" json:"signature_r,omitempty"`  // Computed signature R component (256 bit)
	SignatureS []byte  `protobuf:"bytes,4,opt,name=signature_s,json=signatureS" json:"signature_s,omitempty"`  // Computed signature S component (256 bit)
}

func (x *EthereumTxRequest) Reset() {
	*x = EthereumTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTxRequest) ProtoMessage() {}

func (x *EthereumTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTxRequest.ProtoReflect.Descriptor instead.
func (*EthereumTxRequest) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{5}
}

func (x *EthereumTxRequest) GetDataLength() uint32 {
	if x != nil && x.DataLength != nil {
		return *x.DataLength
	}
	return 0
}

func (x *EthereumTxRequest) GetSignatureV() uint32 {
	if x != nil && x.SignatureV != nil {
		return *x.SignatureV
	}
	return 0
}

func (x *EthereumTxRequest) GetSignatureR() []byte {
	if x != nil {
		return x.SignatureR
	}
	return nil
}

func (x *EthereumTxRequest) GetSignatureS() []byte {
	if x != nil {
		return x.SignatureS
	}
	return nil
}

// *
// Request: Transaction payload data.
// @next EthereumTxRequest
type EthereumTxAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataChunk []byte `protobuf:"bytes,1,opt,name=data_chunk,json=dataChunk" json:"data_chunk,omitempty"` // Bytes from transaction payload (<= 1024 bytes)
}

func (x *EthereumTxAck) Reset() {
	*x = EthereumTxAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTxAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTxAck) ProtoMessage() {}

func (x *EthereumTxAck) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTxAck.ProtoReflect.Descriptor instead.
func (*EthereumTxAck) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{6}
}

func (x *EthereumTxAck) GetDataChunk() []byte {
	if x != nil {
		return x.DataChunk
	}
	return nil
}

// *
// Request: Ask device to sign message
// @start
// @next EthereumMessageSignature
// @next Failure
type EthereumSignMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressN []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"` // BIP-32 path to derive the key from master node
	Message  []byte   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`                    // message to be signed
}

func (x *EthereumSignMessage) Reset() {
	*x = EthereumSignMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumSignMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumSignMessage) ProtoMessage() {}

func (x *EthereumSignMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumSignMessage.ProtoReflect.Descriptor instead.
func (*EthereumSignMessage) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{7}
}

func (x *EthereumSignMessage) GetAddressN() []uint32 {
	if x != nil {
		return x.AddressN
	}
	return nil
}

func (x *EthereumSignMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// *
// Response: Signed message
// @end
type EthereumMessageSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressBin []byte  `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"` // address used to sign the message (20 bytes, legacy firmware)
	Signature  []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`   // signature of the message
	AddressHex *string `protobuf:"bytes,3,opt,name=addressHex" json:"addressHex,omitempty"` // address used to sign the message (hex string, newer firmware)
}

func (x *EthereumMessageSignature) Reset() {
	*x = EthereumMessageSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumMessageSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumMessageSignature) ProtoMessage() {}

func (x *EthereumMessageSignature) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumMessageSignature.ProtoReflect.Descriptor instead.
func (*EthereumMessageSignature) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{8}
}

func (x *EthereumMessageSignature) GetAddressBin() []byte {
	if x != nil {
		return x.AddressBin
	}
	return nil
}

func (x *EthereumMessageSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *EthereumMessageSignature) GetAddressHex() string {
	if x != nil && x.AddressHex != nil {
		return *x.AddressHex
	}
	return ""
}

// *
// Request: Ask device to verify message
// @start
// @next Success
// @next Failure
type EthereumVerifyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressBin []byte  `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"` // address to verify (20 bytes, legacy firmware)
	Signature  []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`   // signature to verify
	Message    []byte  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`       // message to verify
	AddressHex *string `protobuf:"bytes,4,opt,name=addressHex" json:"addressHex,omitempty"` // address to verify (hex string, newer firmware)
}

func (x *EthereumVerifyMessage) Reset() {
	*x = EthereumVerifyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ethereum_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumVerifyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumVerifyMessage) ProtoMessage() {}

func (x *EthereumVerifyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ethereum_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumVerifyMessage.ProtoReflect.Descriptor instead.
func (*EthereumVerifyMessage) Descriptor() ([]byte, []int) {
	return file_messages_ethereum_proto_rawDescGZIP(), []int{9}
}

func (x *EthereumVerifyMessage) GetAddressBin() []byte {
	if x != nil {
		return x.AddressBin
	}
	return nil
}

func (x *EthereumVerifyMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *EthereumVerifyMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *EthereumVerifyMessage) GetAddressHex() string {
	if x != nil && x.AddressHex != nil {
		return *x.AddressHex
	}
	return ""
}

var File_messages_ethereum_proto protoreflect.FileDescriptor

var file_messages_ethereum_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x68, 0x77, 0x2e, 0x74, 0x72,
	0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x1a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a,
	0x14, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x62, 0x0a, 0x11, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x77, 0x2e, 0x74, 0x72,
	0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x44, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x70, 0x75, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x78, 0x70, 0x75, 0x62, 0x22, 0x54, 0x0a, 0x12, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x68, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x22,
	0x51, 0x0a, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48,
	0x65, 0x78, 0x22, 0xc2, 0x02, 0x0a, 0x0e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53,
	0x69, 0x67, 0x6e, 0x54, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x42, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x74, 0x6f, 0x42, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x48, 0x65,
	0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x48, 0x65, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x22, 0x2e, 0x0a, 0x0d, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x54, 0x78, 0x41,
	0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x4c, 0x0a, 0x13, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x53, 0x69, 0x67,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x78, 0x0a, 0x18, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x65, 0x78, 0x42, 0x77, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x74,
	0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x42, 0x15, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x62, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x74, 0x72,
	0x65, 0x7a, 0x6f, 0x72,
}

var (
	file_messages_ethereum_proto_rawDescOnce sync.Once
	file_messages_ethereum_proto_rawDescData = file_messages_ethereum_proto_rawDesc
)

func file_messages_ethereum_proto_rawDescGZIP() []byte {
	file_messages_ethereum_proto_rawDescOnce.Do(func() {
		file_messages_ethereum_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_ethereum_proto_rawDescData)
	})
	return file_messages_ethereum_proto_rawDescData
}

var file_messages_ethereum_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_messages_ethereum_proto_goTypes = []any{
	(*EthereumGetPublicKey)(nil),     // 0: hw.trezor.messages.ethereum.EthereumGetPublicKey
	(*EthereumPublicKey)(nil),        // 1: hw.trezor.messages.ethereum.EthereumPublicKey
	(*EthereumGetAddress)(nil),       // 2: hw.trezor.messages.ethereum.EthereumGetAddress
	(*EthereumAddress)(nil),          // 3: hw.trezor.messages.ethereum.EthereumAddress
	(*EthereumSignTx)(nil),           // 4: hw.trezor.messages.ethereum.EthereumSignTx
	(*EthereumTxRequest)(nil),        // 5: hw.trezor.messages.ethereum.EthereumTxRequest
	(*EthereumTxAck)(nil),            // 6: hw.trezor.messages.ethereum.EthereumTxAck
	(*EthereumSignMessage)(nil),      // 7: hw.trezor.messages.ethereum.EthereumSignMessage
	(*EthereumMessageSignature)(nil), // 8: hw.trezor.messages.ethereum.EthereumMessageSignature
	(*EthereumVerifyMessage)(nil),    // 9: hw.trezor.messages.ethereum.EthereumVerifyMessage
	(*HDNodeType)(nil),               // 10: hw.trezor.messages.common.HDNodeType
}
var file_messages_ethereum_proto_depIdxs = []int32{
	10, // 0: hw.trezor.messages.ethereum.EthereumPublicKey.node:type_name -> hw.trezor.messages.common.HDNodeType
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_messages_ethereum_proto_init() }
func file_messages_ethereum_proto_init() {
	if File_messages_ethereum_proto != nil {
		return
	}
	file_messages_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_ethereum_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumGetPublicKey); i {
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
		file_messages_ethereum_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumPublicKey); i {
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
		file_messages_ethereum_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumGetAddress); i {
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
		file_messages_ethereum_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumAddress); i {
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
		file_messages_ethereum_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumSignTx); i {
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
		file_messages_ethereum_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumTxRequest); i {
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
		file_messages_ethereum_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumTxAck); i {
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
		file_messages_ethereum_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumSignMessage); i {
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
		file_messages_ethereum_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumMessageSignature); i {
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
		file_messages_ethereum_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*EthereumVerifyMessage); i {
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
			RawDescriptor: file_messages_ethereum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_ethereum_proto_goTypes,
		DependencyIndexes: file_messages_ethereum_proto_depIdxs,
		MessageInfos:      file_messages_ethereum_proto_msgTypes,
	}.Build()
	File_messages_ethereum_proto = out.File
	file_messages_ethereum_proto_rawDesc = nil
	file_messages_ethereum_proto_goTypes = nil
	file_messages_ethereum_proto_depIdxs = nil
}
