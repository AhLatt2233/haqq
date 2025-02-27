// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: haqq/coinomics/v1/coinomics.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// coinomics mint distribution's of the minted denom
// https://islamiccoin.net/wp
// - 10%(evergreen_dao) goes to Evergreen DAO.
// - 1%(block_proposer_min) to 5%(block_proposer_max) goes to a block proposer and its delegators.
// - (staking_rewards) The remainder is distributed proportionally to all bonded validators and their delegators.
type MintDistribution struct {
	EvergreenDao     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=evergreen_dao,json=evergreenDao,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"evergreen_dao"`
	BlockProposerMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=block_proposer_min,json=blockProposerMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"block_proposer_min"`
	BlockProposerMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=block_proposer_max,json=blockProposerMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"block_proposer_max"`
}

func (m *MintDistribution) Reset()         { *m = MintDistribution{} }
func (m *MintDistribution) String() string { return proto.CompactTextString(m) }
func (*MintDistribution) ProtoMessage()    {}
func (*MintDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30e636a42ba7e8a, []int{0}
}
func (m *MintDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintDistribution.Merge(m, src)
}
func (m *MintDistribution) XXX_Size() int {
	return m.Size()
}
func (m *MintDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_MintDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_MintDistribution proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MintDistribution)(nil), "haqq.coinomics.v1.MintDistribution")
}

func init() { proto.RegisterFile("haqq/coinomics/v1/coinomics.proto", fileDescriptor_a30e636a42ba7e8a) }

var fileDescriptor_a30e636a42ba7e8a = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x48, 0x2c, 0x2c,
	0xd4, 0x4f, 0xce, 0xcf, 0xcc, 0xcb, 0xcf, 0xcd, 0x4c, 0x2e, 0xd6, 0x2f, 0x33, 0x44, 0x70, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x41, 0x4a, 0xf4, 0x10, 0xa2, 0x65, 0x86, 0x52, 0x22,
	0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0x69, 0x31, 0x13, 0x97, 0x80,
	0x6f, 0x66, 0x5e, 0x89, 0x4b, 0x66, 0x71, 0x49, 0x51, 0x66, 0x52, 0x69, 0x49, 0x66, 0x7e, 0x9e,
	0x50, 0x30, 0x17, 0x6f, 0x6a, 0x59, 0x6a, 0x51, 0x7a, 0x51, 0x6a, 0x6a, 0x5e, 0x7c, 0x4a, 0x62,
	0xbe, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xde, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9,
	0xab, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7,
	0xe6, 0x17, 0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0x97,
	0xd4, 0xe4, 0x20, 0x1e, 0xb8, 0x21, 0x2e, 0x89, 0xf9, 0x42, 0x31, 0x5c, 0x42, 0x49, 0x39, 0xf9,
	0xc9, 0xd9, 0xf1, 0x05, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0xa9, 0x45, 0xf1, 0xb9, 0x99, 0x79, 0x12,
	0x4c, 0x64, 0x99, 0x2c, 0x00, 0x36, 0x29, 0x00, 0x6a, 0x90, 0x6f, 0x66, 0x1e, 0x36, 0xd3, 0x13,
	0x2b, 0x24, 0x98, 0xa9, 0x61, 0x7a, 0x62, 0x85, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0xe9, 0x21, 0x99, 0x09, 0x0a, 0x73, 0xdd, 0xbc, 0xd4, 0x92, 0xf2, 0xfc,
	0xa2, 0x6c, 0x30, 0x47, 0xbf, 0x02, 0x29, 0x96, 0xc0, 0xe6, 0x27, 0xb1, 0x81, 0x83, 0xdd, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x51, 0xc6, 0x68, 0xc4, 0x01, 0x00, 0x00,
}

func (m *MintDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlockProposerMax.Size()
		i -= size
		if _, err := m.BlockProposerMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinomics(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BlockProposerMin.Size()
		i -= size
		if _, err := m.BlockProposerMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinomics(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EvergreenDao.Size()
		i -= size
		if _, err := m.EvergreenDao.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinomics(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCoinomics(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinomics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EvergreenDao.Size()
	n += 1 + l + sovCoinomics(uint64(l))
	l = m.BlockProposerMin.Size()
	n += 1 + l + sovCoinomics(uint64(l))
	l = m.BlockProposerMax.Size()
	n += 1 + l + sovCoinomics(uint64(l))
	return n
}

func sovCoinomics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinomics(x uint64) (n int) {
	return sovCoinomics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinomics
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
			return fmt.Errorf("proto: MintDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvergreenDao", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinomics
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
				return ErrInvalidLengthCoinomics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinomics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EvergreenDao.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockProposerMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinomics
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
				return ErrInvalidLengthCoinomics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinomics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockProposerMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockProposerMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinomics
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
				return ErrInvalidLengthCoinomics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinomics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockProposerMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinomics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinomics
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
func skipCoinomics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinomics
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
					return 0, ErrIntOverflowCoinomics
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
					return 0, ErrIntOverflowCoinomics
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
				return 0, ErrInvalidLengthCoinomics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinomics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinomics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinomics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinomics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinomics = fmt.Errorf("proto: unexpected end of group")
)
