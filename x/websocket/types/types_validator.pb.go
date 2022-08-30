// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: websocket/types/types_validator.proto

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

type Validator struct {
	// Address is the address who is the original validator to store information of those that execute the oScript
	Address     github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"address,omitempty"`
	VotingPower int64                                         `protobuf:"varint,2,opt,name=votingPower,proto3" json:"votingPower,omitempty"`
	Status      string                                        `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_534f1558886b5fca, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *Validator) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ValResult struct {
	Validator    *Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty" json:"validator"`
	Result       []byte     `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty" json:"result"`
	ResultStatus string     `protobuf:"bytes,3,opt,name=resultStatus,proto3" json:"resultStatus,omitempty" json:"result_status"`
}

func (m *ValResult) Reset()         { *m = ValResult{} }
func (m *ValResult) String() string { return proto.CompactTextString(m) }
func (*ValResult) ProtoMessage()    {}
func (*ValResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_534f1558886b5fca, []int{1}
}
func (m *ValResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValResult.Merge(m, src)
}
func (m *ValResult) XXX_Size() int {
	return m.Size()
}
func (m *ValResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ValResult.DiscardUnknown(m)
}

var xxx_messageInfo_ValResult proto.InternalMessageInfo

func (m *ValResult) GetValidator() *Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *ValResult) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ValResult) GetResultStatus() string {
	if m != nil {
		return m.ResultStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*Validator)(nil), "oraichain.orai.websocket.Validator")
	proto.RegisterType((*ValResult)(nil), "oraichain.orai.websocket.ValResult")
}

func init() {
	proto.RegisterFile("websocket/types/types_validator.proto", fileDescriptor_534f1558886b5fca)
}

var fileDescriptor_534f1558886b5fca = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4f, 0x4d, 0x2a,
	0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x90, 0xf1, 0x65, 0x89,
	0x39, 0x99, 0x29, 0x89, 0x25, 0xf9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x12, 0xf9,
	0x45, 0x89, 0x99, 0xc9, 0x19, 0x89, 0x99, 0x79, 0x7a, 0x20, 0x96, 0x1e, 0x5c, 0x97, 0x94, 0x48,
	0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x91, 0x3e, 0x88, 0x05, 0x51, 0xaf, 0x34, 0x83, 0x91, 0x8b, 0x33,
	0x0c, 0x66, 0x86, 0x90, 0x37, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x8f, 0x93, 0xe1, 0xaf, 0x7b, 0xf2, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25,
	0x1b, 0xe2, 0x0c, 0xbd, 0xb0, 0xc4, 0x1c, 0x47, 0x88, 0xc6, 0x20, 0x98, 0x09, 0x42, 0x0a, 0x5c,
	0xdc, 0x65, 0xf9, 0x25, 0x99, 0x79, 0xe9, 0x01, 0xf9, 0xe5, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0xcc, 0x41, 0xc8, 0x42, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12,
	0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x8c, 0x4a, 0x97,
	0x21, 0x4e, 0x0b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0x0a, 0xe7, 0xe2, 0x84, 0xfb, 0x15, 0xec,
	0x38, 0x6e, 0x23, 0x65, 0x3d, 0x5c, 0x9e, 0xd5, 0x83, 0x7b, 0xc9, 0x49, 0xe4, 0xd3, 0x3d, 0x79,
	0x81, 0xac, 0xe2, 0xfc, 0x3c, 0x2b, 0x25, 0xb8, 0x7e, 0xa5, 0x20, 0x84, 0x59, 0x42, 0x9a, 0x5c,
	0x6c, 0x45, 0x60, 0x2b, 0xc0, 0x2e, 0xe4, 0x71, 0x12, 0xfc, 0x74, 0x4f, 0x9e, 0x17, 0xa2, 0x01,
	0x22, 0xae, 0x14, 0x04, 0x55, 0x20, 0x64, 0xc3, 0xc5, 0x03, 0x61, 0x05, 0x23, 0xb9, 0xda, 0x49,
	0xe2, 0xd3, 0x3d, 0x79, 0x11, 0x64, 0x0d, 0xf1, 0x10, 0x6f, 0x28, 0x05, 0xa1, 0xa8, 0x86, 0xf8,
	0xca, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0x91, 0xc2,
	0x19, 0xee, 0x31, 0x30, 0x4b, 0xbf, 0x42, 0x1f, 0x2d, 0xf6, 0x93, 0xd8, 0xc0, 0xd1, 0x67, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x27, 0xe8, 0x03, 0x95, 0x17, 0x02, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTypesValidator(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VotingPower != 0 {
		i = encodeVarintTypesValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypesValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResultStatus) > 0 {
		i -= len(m.ResultStatus)
		copy(dAtA[i:], m.ResultStatus)
		i = encodeVarintTypesValidator(dAtA, i, uint64(len(m.ResultStatus)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintTypesValidator(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x12
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypesValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypesValidator(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovTypesValidator(uint64(m.VotingPower))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTypesValidator(uint64(l))
	}
	return n
}

func (m *ValResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovTypesValidator(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovTypesValidator(uint64(l))
	}
	l = len(m.ResultStatus)
	if l > 0 {
		n += 1 + l + sovTypesValidator(uint64(l))
	}
	return n
}

func sovTypesValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesValidator(x uint64) (n int) {
	return sovTypesValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
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
				return ErrInvalidLengthTypesValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
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
				return ErrInvalidLengthTypesValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesValidator
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
func (m *ValResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesValidator
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
			return fmt.Errorf("proto: ValResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
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
				return ErrInvalidLengthTypesValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &Validator{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
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
				return ErrInvalidLengthTypesValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesValidator
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
				return ErrInvalidLengthTypesValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesValidator
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
func skipTypesValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesValidator
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
					return 0, ErrIntOverflowTypesValidator
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
					return 0, ErrIntOverflowTypesValidator
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
				return 0, ErrInvalidLengthTypesValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesValidator = fmt.Errorf("proto: unexpected end of group")
)