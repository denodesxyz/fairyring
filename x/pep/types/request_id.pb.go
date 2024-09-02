// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/request_id.proto

package types

import (
	fmt "fmt"
	types1 "github.com/Fairblock/fairyring/x/common/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type RequestId struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ReqId   string `protobuf:"bytes,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
}

func (m *RequestId) Reset()         { *m = RequestId{} }
func (m *RequestId) String() string { return proto.CompactTextString(m) }
func (*RequestId) ProtoMessage()    {}
func (*RequestId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e457d2e8ff0411e, []int{0}
}
func (m *RequestId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestId.Merge(m, src)
}
func (m *RequestId) XXX_Size() int {
	return m.Size()
}
func (m *RequestId) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestId.DiscardUnknown(m)
}

var xxx_messageInfo_RequestId proto.InternalMessageInfo

func (m *RequestId) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RequestId) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

type PrivateRequest struct {
	Creator            string                          `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ReqId              string                          `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	Pubkey             string                          `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Amount             types.Coin                      `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
	EncryptedKeyshares map[string]*types1.KeyshareList `protobuf:"bytes,5,rep,name=encrypted_keyshares,json=encryptedKeyshares,proto3" json:"encrypted_keyshares,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *PrivateRequest) Reset()         { *m = PrivateRequest{} }
func (m *PrivateRequest) String() string { return proto.CompactTextString(m) }
func (*PrivateRequest) ProtoMessage()    {}
func (*PrivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e457d2e8ff0411e, []int{1}
}
func (m *PrivateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrivateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateRequest.Merge(m, src)
}
func (m *PrivateRequest) XXX_Size() int {
	return m.Size()
}
func (m *PrivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateRequest proto.InternalMessageInfo

func (m *PrivateRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PrivateRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *PrivateRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *PrivateRequest) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *PrivateRequest) GetEncryptedKeyshares() map[string]*types1.KeyshareList {
	if m != nil {
		return m.EncryptedKeyshares
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestId)(nil), "fairyring.pep.RequestId")
	proto.RegisterType((*PrivateRequest)(nil), "fairyring.pep.PrivateRequest")
	proto.RegisterMapType((map[string]*types1.KeyshareList)(nil), "fairyring.pep.PrivateRequest.EncryptedKeysharesEntry")
}

func init() { proto.RegisterFile("fairyring/pep/request_id.proto", fileDescriptor_3e457d2e8ff0411e) }

var fileDescriptor_3e457d2e8ff0411e = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x6d, 0xda, 0x6d, 0xd0, 0xba, 0x02, 0x21, 0xb3, 0x40, 0xe8, 0xc1, 0x54, 0xcb, 0xa5, 0xe2,
	0x60, 0x6b, 0x0b, 0x08, 0x04, 0xb7, 0x85, 0x45, 0x5a, 0xc1, 0x01, 0xe5, 0xc8, 0x25, 0x72, 0x92,
	0xd9, 0xac, 0xd5, 0x4d, 0xec, 0xda, 0x4e, 0x45, 0xfe, 0x82, 0xcf, 0xda, 0xe3, 0x1e, 0x39, 0x21,
	0x68, 0x7f, 0x04, 0xc5, 0x4e, 0x28, 0x15, 0x82, 0xdb, 0xcc, 0xf8, 0xcd, 0x9b, 0x99, 0xf7, 0x8c,
	0xc8, 0x05, 0x17, 0xba, 0xd1, 0xa2, 0x2a, 0x98, 0x02, 0xc5, 0x34, 0xac, 0x6a, 0x30, 0x36, 0x11,
	0x39, 0x55, 0x5a, 0x5a, 0x89, 0x6f, 0xff, 0x7e, 0xa7, 0x0a, 0xd4, 0x94, 0x64, 0xd2, 0x94, 0xd2,
	0xb0, 0x94, 0x1b, 0x60, 0xeb, 0x93, 0x14, 0x2c, 0x3f, 0x61, 0x99, 0x14, 0x95, 0x87, 0x4f, 0x8f,
	0x0a, 0x59, 0x48, 0x17, 0xb2, 0x36, 0xea, 0xaa, 0x4f, 0x76, 0x43, 0x32, 0x59, 0x96, 0xb2, 0x62,
	0xe6, 0x92, 0x6b, 0xc8, 0x13, 0xdb, 0x28, 0x30, 0x1e, 0x74, 0xfc, 0x06, 0x1d, 0xc6, 0x7e, 0xfa,
	0x79, 0x8e, 0x23, 0x74, 0x2b, 0xd3, 0xc0, 0xad, 0xd4, 0x51, 0x30, 0x0b, 0xe6, 0x87, 0x71, 0x9f,
	0xe2, 0x23, 0x34, 0xd6, 0xb0, 0x3a, 0xcf, 0xa3, 0xa1, 0xab, 0xfb, 0xe4, 0xf8, 0xe7, 0x10, 0xdd,
	0xf9, 0xa4, 0xc5, 0x9a, 0x5b, 0xe8, 0x48, 0xfe, 0x43, 0x71, 0x1f, 0x85, 0x1a, 0x56, 0x89, 0xd8,
	0xe7, 0xc0, 0x0f, 0x50, 0xa8, 0xea, 0x74, 0x09, 0x4d, 0x34, 0x72, 0xe5, 0x2e, 0xc3, 0x2f, 0x51,
	0xc8, 0x4b, 0x59, 0x57, 0x36, 0x3a, 0x98, 0x05, 0xf3, 0xc9, 0xe2, 0x11, 0xf5, 0x22, 0xd0, 0x56,
	0x04, 0xda, 0x89, 0x40, 0xdf, 0x4a, 0x51, 0x9d, 0x1e, 0x5c, 0x7f, 0x7f, 0x3c, 0x88, 0x3b, 0x38,
	0xbe, 0x40, 0xf7, 0xa0, 0xca, 0x74, 0xa3, 0x2c, 0xe4, 0xc9, 0x12, 0x1a, 0x77, 0xb4, 0x89, 0xc6,
	0xb3, 0xd1, 0x7c, 0xb2, 0x78, 0x41, 0xf7, 0x94, 0xa5, 0xfb, 0xdb, 0xd3, 0xb3, 0xbe, 0xf1, 0x43,
	0xdf, 0x77, 0x56, 0x59, 0xdd, 0xc4, 0x18, 0xfe, 0x7a, 0x98, 0x02, 0x7a, 0xf8, 0x0f, 0x38, 0xbe,
	0x8b, 0x46, 0xed, 0x41, 0x5e, 0x80, 0x36, 0xc4, 0xcf, 0xd1, 0x78, 0xcd, 0xaf, 0x6a, 0x70, 0xb7,
	0x4f, 0x16, 0xe4, 0x8f, 0x35, 0xbc, 0x37, 0xb4, 0xa7, 0xf8, 0x28, 0x8c, 0x8d, 0x3d, 0xf8, 0xf5,
	0xf0, 0x55, 0x70, 0xfa, 0xee, 0x7a, 0x43, 0x82, 0x9b, 0x0d, 0x09, 0x7e, 0x6c, 0x48, 0xf0, 0x75,
	0x4b, 0x06, 0x37, 0x5b, 0x32, 0xf8, 0xb6, 0x25, 0x83, 0xcf, 0x4f, 0x0b, 0x61, 0x2f, 0xeb, 0xb4,
	0x25, 0x60, 0xef, 0xb9, 0xd0, 0xe9, 0x95, 0xcc, 0x96, 0x6c, 0x67, 0xfa, 0x17, 0xf7, 0xb7, 0x9c,
	0xd9, 0x69, 0xe8, 0xdc, 0x7e, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x46, 0xe7, 0x5d, 0x79,
	0x02, 0x00, 0x00,
}

func (m *RequestId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrivateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedKeyshares) > 0 {
		for k := range m.EncryptedKeyshares {
			v := m.EncryptedKeyshares[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintRequestId(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintRequestId(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintRequestId(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequestId(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequestId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequestId(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequestId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	return n
}

func (m *PrivateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovRequestId(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovRequestId(uint64(l))
	if len(m.EncryptedKeyshares) > 0 {
		for k, v := range m.EncryptedKeyshares {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovRequestId(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovRequestId(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovRequestId(uint64(mapEntrySize))
		}
	}
	return n
}

func sovRequestId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequestId(x uint64) (n int) {
	return sovRequestId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestId
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
			return fmt.Errorf("proto: RequestId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestId
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
func (m *PrivateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestId
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
			return fmt.Errorf("proto: PrivateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKeyshares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestId
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
				return ErrInvalidLengthRequestId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EncryptedKeyshares == nil {
				m.EncryptedKeyshares = make(map[string]*types1.KeyshareList)
			}
			var mapkey string
			var mapvalue *types1.KeyshareList
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRequestId
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRequestId
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthRequestId
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthRequestId
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRequestId
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthRequestId
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthRequestId
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &types1.KeyshareList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipRequestId(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthRequestId
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.EncryptedKeyshares[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestId
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
func skipRequestId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestId
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
					return 0, ErrIntOverflowRequestId
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
					return 0, ErrIntOverflowRequestId
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
				return 0, ErrInvalidLengthRequestId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequestId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequestId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequestId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequestId = fmt.Errorf("proto: unexpected end of group")
)
