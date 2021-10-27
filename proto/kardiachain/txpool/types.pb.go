// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kardiachain/txpool/types.proto

package state

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Txs struct {
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *Txs) Reset()         { *m = Txs{} }
func (m *Txs) String() string { return proto.CompactTextString(m) }
func (*Txs) ProtoMessage()    {}
func (*Txs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb40ebb23df31a9b, []int{0}
}
func (m *Txs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Txs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Txs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Txs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Txs.Merge(m, src)
}
func (m *Txs) XXX_Size() int {
	return m.Size()
}
func (m *Txs) XXX_DiscardUnknown() {
	xxx_messageInfo_Txs.DiscardUnknown(m)
}

var xxx_messageInfo_Txs proto.InternalMessageInfo

func (m *Txs) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type PooledTransactions struct {
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *PooledTransactions) Reset()         { *m = PooledTransactions{} }
func (m *PooledTransactions) String() string { return proto.CompactTextString(m) }
func (*PooledTransactions) ProtoMessage()    {}
func (*PooledTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb40ebb23df31a9b, []int{1}
}
func (m *PooledTransactions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PooledTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PooledTransactions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PooledTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PooledTransactions.Merge(m, src)
}
func (m *PooledTransactions) XXX_Size() int {
	return m.Size()
}
func (m *PooledTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_PooledTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_PooledTransactions proto.InternalMessageInfo

func (m *PooledTransactions) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

type PooledTransactionHashes struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (m *PooledTransactionHashes) Reset()         { *m = PooledTransactionHashes{} }
func (m *PooledTransactionHashes) String() string { return proto.CompactTextString(m) }
func (*PooledTransactionHashes) ProtoMessage()    {}
func (*PooledTransactionHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb40ebb23df31a9b, []int{2}
}
func (m *PooledTransactionHashes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PooledTransactionHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PooledTransactionHashes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PooledTransactionHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PooledTransactionHashes.Merge(m, src)
}
func (m *PooledTransactionHashes) XXX_Size() int {
	return m.Size()
}
func (m *PooledTransactionHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_PooledTransactionHashes.DiscardUnknown(m)
}

var xxx_messageInfo_PooledTransactionHashes proto.InternalMessageInfo

func (m *PooledTransactionHashes) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type RequestPooledTransactions struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (m *RequestPooledTransactions) Reset()         { *m = RequestPooledTransactions{} }
func (m *RequestPooledTransactions) String() string { return proto.CompactTextString(m) }
func (*RequestPooledTransactions) ProtoMessage()    {}
func (*RequestPooledTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb40ebb23df31a9b, []int{3}
}
func (m *RequestPooledTransactions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestPooledTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestPooledTransactions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestPooledTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPooledTransactions.Merge(m, src)
}
func (m *RequestPooledTransactions) XXX_Size() int {
	return m.Size()
}
func (m *RequestPooledTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPooledTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPooledTransactions proto.InternalMessageInfo

func (m *RequestPooledTransactions) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type Message struct {
	// Types that are valid to be assigned to Sum:
	//	*Message_Txs
	//	*Message_PooledTransactionHashes
	//	*Message_PooledTransactions
	//	*Message_RequestPooledTransactions
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb40ebb23df31a9b, []int{4}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Sum interface {
	isMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_Txs struct {
	Txs *Txs `protobuf:"bytes,1,opt,name=txs,proto3,oneof" json:"txs,omitempty"`
}
type Message_PooledTransactionHashes struct {
	PooledTransactionHashes *PooledTransactionHashes `protobuf:"bytes,2,opt,name=pooledTransactionHashes,proto3,oneof" json:"pooledTransactionHashes,omitempty"`
}
type Message_PooledTransactions struct {
	PooledTransactions *PooledTransactions `protobuf:"bytes,3,opt,name=pooledTransactions,proto3,oneof" json:"pooledTransactions,omitempty"`
}
type Message_RequestPooledTransactions struct {
	RequestPooledTransactions *RequestPooledTransactions `protobuf:"bytes,4,opt,name=requestPooledTransactions,proto3,oneof" json:"requestPooledTransactions,omitempty"`
}

func (*Message_Txs) isMessage_Sum()                       {}
func (*Message_PooledTransactionHashes) isMessage_Sum()   {}
func (*Message_PooledTransactions) isMessage_Sum()        {}
func (*Message_RequestPooledTransactions) isMessage_Sum() {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetTxs() *Txs {
	if x, ok := m.GetSum().(*Message_Txs); ok {
		return x.Txs
	}
	return nil
}

func (m *Message) GetPooledTransactionHashes() *PooledTransactionHashes {
	if x, ok := m.GetSum().(*Message_PooledTransactionHashes); ok {
		return x.PooledTransactionHashes
	}
	return nil
}

func (m *Message) GetPooledTransactions() *PooledTransactions {
	if x, ok := m.GetSum().(*Message_PooledTransactions); ok {
		return x.PooledTransactions
	}
	return nil
}

func (m *Message) GetRequestPooledTransactions() *RequestPooledTransactions {
	if x, ok := m.GetSum().(*Message_RequestPooledTransactions); ok {
		return x.RequestPooledTransactions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Txs)(nil),
		(*Message_PooledTransactionHashes)(nil),
		(*Message_PooledTransactions)(nil),
		(*Message_RequestPooledTransactions)(nil),
	}
}

func init() {
	proto.RegisterType((*Txs)(nil), "kardiachain.state.Txs")
	proto.RegisterType((*PooledTransactions)(nil), "kardiachain.state.PooledTransactions")
	proto.RegisterType((*PooledTransactionHashes)(nil), "kardiachain.state.PooledTransactionHashes")
	proto.RegisterType((*RequestPooledTransactions)(nil), "kardiachain.state.RequestPooledTransactions")
	proto.RegisterType((*Message)(nil), "kardiachain.state.Message")
}

func init() { proto.RegisterFile("kardiachain/txpool/types.proto", fileDescriptor_cb40ebb23df31a9b) }

var fileDescriptor_cb40ebb23df31a9b = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x37, 0xcd, 0xf7, 0x55, 0x18, 0x3d, 0xe8, 0x1e, 0xda, 0xf4, 0xb2, 0x48, 0x41, 0x29,
	0x45, 0x13, 0xb4, 0x17, 0xcf, 0x3d, 0xe5, 0x22, 0x48, 0x28, 0x08, 0xde, 0xb6, 0xed, 0xda, 0x04,
	0xd3, 0x6c, 0xcc, 0x6c, 0x20, 0xbe, 0x85, 0x8f, 0xe5, 0xb1, 0x47, 0x8f, 0x92, 0x1c, 0x7d, 0x09,
	0xe9, 0x5a, 0xa1, 0x92, 0x5d, 0xf0, 0x36, 0xc3, 0x7f, 0xe6, 0x37, 0xff, 0x19, 0x06, 0xd8, 0x13,
	0x2f, 0x96, 0x09, 0x5f, 0xc4, 0x3c, 0xc9, 0x02, 0x55, 0xe5, 0x52, 0xa6, 0x81, 0x7a, 0xc9, 0x05,
	0xfa, 0x79, 0x21, 0x95, 0xa4, 0x27, 0x7b, 0xba, 0x8f, 0x8a, 0x2b, 0x31, 0xec, 0x83, 0x3b, 0xab,
	0x90, 0x1e, 0x83, 0xab, 0x2a, 0xf4, 0x9c, 0x53, 0x77, 0x74, 0x14, 0x6d, 0xc3, 0xe1, 0x39, 0xd0,
	0x3b, 0x29, 0x53, 0xb1, 0x9c, 0x15, 0x3c, 0x43, 0xbe, 0x50, 0x89, 0xcc, 0x4c, 0x75, 0x57, 0xd0,
	0x6f, 0xd5, 0x85, 0x1c, 0x63, 0x81, 0xb4, 0x07, 0xdd, 0x58, 0x47, 0xbb, 0xfa, 0x5d, 0x36, 0x9c,
	0xc0, 0x20, 0x12, 0xcf, 0xa5, 0x40, 0x65, 0x98, 0x60, 0x6b, 0xfa, 0xec, 0xc0, 0xc1, 0xad, 0x40,
	0xe4, 0x2b, 0x41, 0xc7, 0x3f, 0x2e, 0x9c, 0xd1, 0xe1, 0x75, 0xcf, 0x6f, 0x6d, 0xe5, 0xcf, 0x2a,
	0x0c, 0x89, 0xf6, 0x47, 0x1f, 0xa1, 0x9f, 0x9b, 0xfd, 0x79, 0x1d, 0xdd, 0x3f, 0x36, 0xf4, 0x5b,
	0x36, 0x0a, 0x49, 0x64, 0x83, 0xd1, 0x7b, 0xa0, 0x2d, 0x09, 0x3d, 0x57, 0x8f, 0x38, 0xfb, 0xcb,
	0x88, 0x2d, 0xdd, 0x80, 0xa0, 0x29, 0x0c, 0x0a, 0xdb, 0xb5, 0xbc, 0x7f, 0x9a, 0x7f, 0x61, 0xe0,
	0x5b, 0x2f, 0x1c, 0x92, 0xc8, 0x0e, 0x9c, 0xfe, 0x07, 0x17, 0xcb, 0xf5, 0x34, 0x7a, 0xab, 0x99,
	0xb3, 0xa9, 0x99, 0xf3, 0x51, 0x33, 0xe7, 0xb5, 0x61, 0x64, 0xd3, 0x30, 0xf2, 0xde, 0x30, 0xf2,
	0x70, 0xb3, 0x4a, 0x54, 0x5c, 0xce, 0xfd, 0x85, 0x5c, 0x07, 0xfb, 0xef, 0xb6, 0x92, 0x97, 0xdf,
	0x69, 0xa0, 0x7f, 0xed, 0x97, 0xa6, 0x1d, 0xcd, 0xbb, 0x5a, 0x98, 0x7c, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0xc6, 0x62, 0xaa, 0xa6, 0x02, 0x00, 0x00,
}

func (m *Txs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Txs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Txs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PooledTransactions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PooledTransactions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PooledTransactions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PooledTransactionHashes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PooledTransactionHashes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PooledTransactionHashes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for iNdEx := len(m.Hashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hashes[iNdEx])
			copy(dAtA[i:], m.Hashes[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Hashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestPooledTransactions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestPooledTransactions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestPooledTransactions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for iNdEx := len(m.Hashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hashes[iNdEx])
			copy(dAtA[i:], m.Hashes[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Hashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_Txs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Txs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Txs != nil {
		{
			size, err := m.Txs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Message_PooledTransactionHashes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PooledTransactionHashes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PooledTransactionHashes != nil {
		{
			size, err := m.PooledTransactionHashes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Message_PooledTransactions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PooledTransactions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PooledTransactions != nil {
		{
			size, err := m.PooledTransactions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Message_RequestPooledTransactions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_RequestPooledTransactions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestPooledTransactions != nil {
		{
			size, err := m.RequestPooledTransactions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Txs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *PooledTransactions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *PooledTransactionHashes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *RequestPooledTransactions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Message_Txs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Txs != nil {
		l = m.Txs.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_PooledTransactionHashes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PooledTransactionHashes != nil {
		l = m.PooledTransactionHashes.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_PooledTransactions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PooledTransactions != nil {
		l = m.PooledTransactions.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_RequestPooledTransactions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestPooledTransactions != nil {
		l = m.RequestPooledTransactions.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Txs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Txs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Txs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, postIndex-iNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PooledTransactions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PooledTransactions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PooledTransactions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, postIndex-iNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PooledTransactionHashes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PooledTransactionHashes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PooledTransactionHashes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestPooledTransactions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestPooledTransactions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestPooledTransactions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Txs{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_Txs{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PooledTransactionHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PooledTransactionHashes{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PooledTransactionHashes{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PooledTransactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PooledTransactions{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PooledTransactions{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPooledTransactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestPooledTransactions{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_RequestPooledTransactions{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
