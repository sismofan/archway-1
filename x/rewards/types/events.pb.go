// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/rewards/v1beta1/events.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// ContractMetadataSetEvent is emitted when the contract metadata is created or updated.
type ContractMetadataSetEvent struct {
	// contract_address defines the contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// metadata defines the new contract metadata state.
	Metadata ContractMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
}

func (m *ContractMetadataSetEvent) Reset()         { *m = ContractMetadataSetEvent{} }
func (m *ContractMetadataSetEvent) String() string { return proto.CompactTextString(m) }
func (*ContractMetadataSetEvent) ProtoMessage()    {}
func (*ContractMetadataSetEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{0}
}
func (m *ContractMetadataSetEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractMetadataSetEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractMetadataSetEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractMetadataSetEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractMetadataSetEvent.Merge(m, src)
}
func (m *ContractMetadataSetEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractMetadataSetEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractMetadataSetEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractMetadataSetEvent proto.InternalMessageInfo

func (m *ContractMetadataSetEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractMetadataSetEvent) GetMetadata() ContractMetadata {
	if m != nil {
		return m.Metadata
	}
	return ContractMetadata{}
}

// ContractRewardCalculationEvent is emitted when the contract reward is calculated.
type ContractRewardCalculationEvent struct {
	// contract_address defines the contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// gas_consumed defines the total gas consumption by all WASM operations within one transaction.
	GasConsumed uint64 `protobuf:"varint,2,opt,name=gas_consumed,json=gasConsumed,proto3" json:"gas_consumed,omitempty"`
	// inflation_rewards defines the inflation rewards portions of the rewards.
	InflationRewards types.Coin `protobuf:"bytes,3,opt,name=inflation_rewards,json=inflationRewards,proto3" json:"inflation_rewards"`
	// fee_rebate_rewards defines the fee rebate rewards portions of the rewards.
	FeeRebateRewards []types.Coin `protobuf:"bytes,4,rep,name=fee_rebate_rewards,json=feeRebateRewards,proto3" json:"fee_rebate_rewards"`
	// metadata defines the contract metadata (if set).
	Metadata *ContractMetadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *ContractRewardCalculationEvent) Reset()         { *m = ContractRewardCalculationEvent{} }
func (m *ContractRewardCalculationEvent) String() string { return proto.CompactTextString(m) }
func (*ContractRewardCalculationEvent) ProtoMessage()    {}
func (*ContractRewardCalculationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{1}
}
func (m *ContractRewardCalculationEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRewardCalculationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRewardCalculationEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRewardCalculationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRewardCalculationEvent.Merge(m, src)
}
func (m *ContractRewardCalculationEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractRewardCalculationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRewardCalculationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRewardCalculationEvent proto.InternalMessageInfo

func (m *ContractRewardCalculationEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractRewardCalculationEvent) GetGasConsumed() uint64 {
	if m != nil {
		return m.GasConsumed
	}
	return 0
}

func (m *ContractRewardCalculationEvent) GetInflationRewards() types.Coin {
	if m != nil {
		return m.InflationRewards
	}
	return types.Coin{}
}

func (m *ContractRewardCalculationEvent) GetFeeRebateRewards() []types.Coin {
	if m != nil {
		return m.FeeRebateRewards
	}
	return nil
}

func (m *ContractRewardCalculationEvent) GetMetadata() *ContractMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// ContractRewardDistributionEvent is emitted when the contract reward is distributed to the corresponding rewards address.
// This event might not follow the ContractRewardCalculationEvent if the contract has no metadata set or rewards address is empty.
type ContractRewardDistributionEvent struct {
	// contract_address defines the contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// rewards_address defines the rewards address rewards are distributed to.
	RewardAddress string `protobuf:"bytes,2,opt,name=reward_address,json=rewardAddress,proto3" json:"reward_address,omitempty"`
	// rewards defines the total rewards being distributed.
	Rewards []types.Coin `protobuf:"bytes,3,rep,name=rewards,proto3" json:"rewards"`
}

func (m *ContractRewardDistributionEvent) Reset()         { *m = ContractRewardDistributionEvent{} }
func (m *ContractRewardDistributionEvent) String() string { return proto.CompactTextString(m) }
func (*ContractRewardDistributionEvent) ProtoMessage()    {}
func (*ContractRewardDistributionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad2689b4f7dc3cd8, []int{2}
}
func (m *ContractRewardDistributionEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRewardDistributionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRewardDistributionEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRewardDistributionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRewardDistributionEvent.Merge(m, src)
}
func (m *ContractRewardDistributionEvent) XXX_Size() int {
	return m.Size()
}
func (m *ContractRewardDistributionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRewardDistributionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRewardDistributionEvent proto.InternalMessageInfo

func (m *ContractRewardDistributionEvent) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractRewardDistributionEvent) GetRewardAddress() string {
	if m != nil {
		return m.RewardAddress
	}
	return ""
}

func (m *ContractRewardDistributionEvent) GetRewards() []types.Coin {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractMetadataSetEvent)(nil), "archway.rewards.v1beta1.ContractMetadataSetEvent")
	proto.RegisterType((*ContractRewardCalculationEvent)(nil), "archway.rewards.v1beta1.ContractRewardCalculationEvent")
	proto.RegisterType((*ContractRewardDistributionEvent)(nil), "archway.rewards.v1beta1.ContractRewardDistributionEvent")
}

func init() {
	proto.RegisterFile("archway/rewards/v1beta1/events.proto", fileDescriptor_ad2689b4f7dc3cd8)
}

var fileDescriptor_ad2689b4f7dc3cd8 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbd, 0x69, 0xf8, 0xb7, 0xe5, 0x4f, 0xb1, 0x90, 0x30, 0x3d, 0x6c, 0x43, 0x45, 0xa5,
	0xf6, 0xc0, 0x5a, 0x2d, 0x27, 0x8e, 0x34, 0xf4, 0x44, 0x72, 0x31, 0x37, 0x2e, 0xd6, 0x7a, 0x3d,
	0x71, 0x2c, 0xe2, 0xdd, 0x68, 0x77, 0x9d, 0x90, 0xb7, 0x40, 0x3c, 0x08, 0xcf, 0x91, 0x63, 0xc4,
	0x89, 0x13, 0x42, 0xc9, 0x8b, 0xa0, 0xac, 0xd7, 0x26, 0x20, 0x45, 0x4a, 0x7a, 0xb3, 0x67, 0x7f,
	0xf3, 0xcd, 0x37, 0x9f, 0x06, 0xbf, 0x62, 0x8a, 0x0f, 0xa7, 0x6c, 0x16, 0x2a, 0x98, 0x32, 0x95,
	0xea, 0x70, 0x72, 0x99, 0x80, 0x61, 0x97, 0x21, 0x4c, 0x40, 0x18, 0x4d, 0xc7, 0x4a, 0x1a, 0xe9,
	0x3f, 0x77, 0x14, 0x75, 0x14, 0x75, 0xd4, 0xf1, 0xb3, 0x4c, 0x66, 0xd2, 0x32, 0xe1, 0xfa, 0xab,
	0xc2, 0x8f, 0x09, 0x97, 0xba, 0x90, 0x3a, 0x4c, 0x98, 0x86, 0x46, 0x90, 0xcb, 0x5c, 0xb8, 0xf7,
	0xb3, 0x6d, 0x43, 0x6b, 0x79, 0x8b, 0x9d, 0x7e, 0x43, 0x38, 0xe8, 0x4a, 0x61, 0x14, 0xe3, 0xa6,
	0x0f, 0x86, 0xa5, 0xcc, 0xb0, 0x8f, 0x60, 0x6e, 0xd6, 0xce, 0xfc, 0x0b, 0x7c, 0xc4, 0xdd, 0x5b,
	0xcc, 0xd2, 0x54, 0x81, 0xd6, 0x01, 0xea, 0xa0, 0xf3, 0x07, 0xd1, 0x93, 0xba, 0xfe, 0xae, 0x2a,
	0xfb, 0x1f, 0xf0, 0xfd, 0xc2, 0xb5, 0x07, 0xad, 0x0e, 0x3a, 0x3f, 0xbc, 0xba, 0xa0, 0x5b, 0x16,
	0xa2, 0xff, 0xcf, 0xbb, 0x6e, 0xcf, 0x7f, 0x9d, 0x78, 0x51, 0x23, 0x70, 0xfa, 0xa3, 0x85, 0x49,
	0x0d, 0x45, 0xb6, 0xb9, 0xcb, 0x46, 0xbc, 0x1c, 0x31, 0x93, 0x4b, 0xb1, 0xb7, 0xb5, 0x97, 0xf8,
	0x61, 0xc6, 0x74, 0xcc, 0xa5, 0xd0, 0x65, 0x01, 0xa9, 0xb5, 0xd7, 0x8e, 0x0e, 0x33, 0xa6, 0xbb,
	0xae, 0xe4, 0xf7, 0xf0, 0xd3, 0x5c, 0x0c, 0x2a, 0xfd, 0xd8, 0xd9, 0x0d, 0x0e, 0xec, 0x1a, 0x2f,
	0x68, 0x15, 0x34, 0x5d, 0x07, 0xbd, 0xb1, 0x42, 0x2e, 0x9c, 0xed, 0xa3, 0xa6, 0xb3, 0xb2, 0xaa,
	0xfd, 0x3e, 0xf6, 0x07, 0x00, 0xb1, 0x82, 0x84, 0x19, 0x68, 0xe4, 0xda, 0x9d, 0x83, 0x9d, 0xe4,
	0x06, 0x00, 0x91, 0xed, 0xac, 0xe5, 0x6e, 0x36, 0xa2, 0xbd, 0xb3, 0x67, 0xb4, 0x1b, 0xa1, 0x7e,
	0x47, 0xf8, 0xe4, 0xdf, 0x50, 0xdf, 0xe7, 0xda, 0xa8, 0x3c, 0x29, 0x6f, 0x95, 0xea, 0x19, 0x7e,
	0x5c, 0x0d, 0x6f, 0xc0, 0x96, 0x05, 0x1f, 0x55, 0xd5, 0x1a, 0x7b, 0x8b, 0xef, 0xfd, 0xcd, 0x73,
	0xa7, 0x00, 0x6a, 0xfe, 0xba, 0x37, 0x5f, 0x12, 0xb4, 0x58, 0x12, 0xf4, 0x7b, 0x49, 0xd0, 0xd7,
	0x15, 0xf1, 0x16, 0x2b, 0xe2, 0xfd, 0x5c, 0x11, 0xef, 0xd3, 0x55, 0x96, 0x9b, 0x61, 0x99, 0x50,
	0x2e, 0x8b, 0xd0, 0x25, 0xf1, 0x5a, 0x80, 0x99, 0x4a, 0xf5, 0xb9, 0xfe, 0x0f, 0xbf, 0x34, 0x87,
	0x6f, 0x66, 0x63, 0xd0, 0xc9, 0x5d, 0x7b, 0xef, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x80, 0xef, 0x41, 0x8d, 0x03, 0x00, 0x00,
}

func (m *ContractMetadataSetEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractMetadataSetEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractMetadataSetEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractRewardCalculationEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRewardCalculationEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRewardCalculationEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FeeRebateRewards) > 0 {
		for iNdEx := len(m.FeeRebateRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeRebateRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.InflationRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.GasConsumed != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GasConsumed))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractRewardDistributionEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRewardDistributionEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRewardDistributionEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractMetadataSetEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *ContractRewardCalculationEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.GasConsumed != 0 {
		n += 1 + sovEvents(uint64(m.GasConsumed))
	}
	l = m.InflationRewards.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.FeeRebateRewards) > 0 {
		for _, e := range m.FeeRebateRewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *ContractRewardDistributionEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractMetadataSetEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ContractMetadataSetEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractMetadataSetEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *ContractRewardCalculationEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ContractRewardCalculationEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRewardCalculationEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasConsumed", wireType)
			}
			m.GasConsumed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasConsumed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRebateRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeRebateRewards = append(m.FeeRebateRewards, types.Coin{})
			if err := m.FeeRebateRewards[len(m.FeeRebateRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ContractMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *ContractRewardDistributionEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ContractRewardDistributionEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRewardDistributionEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
