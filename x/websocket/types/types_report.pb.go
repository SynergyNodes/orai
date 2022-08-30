// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: websocket/types/types_report.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Report stores the result of the data source when validator executes it
type Report struct {
	RequestID         string                                   `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty" json:"request_id"`
	DataSourceResults []*DataSourceResult                      `protobuf:"bytes,2,rep,name=DataSourceResults,proto3" json:"DataSourceResults,omitempty" json:"data_source_result"`
	TestCaseResults   []*TestCaseResult                        `protobuf:"bytes,3,rep,name=TestCaseResults,proto3" json:"TestCaseResults,omitempty" json:"test_case_results"`
	BlockHeight       int64                                    `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty" json:"block_height"`
	Fees              github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees" json:"report_fee"`
	AggregatedResult  []byte                                   `protobuf:"bytes,6,opt,name=aggregatedResult,proto3" json:"aggregatedResult,omitempty" json:"aggregated_result"`
	ResultStatus      string                                   `protobuf:"bytes,7,opt,name=resultStatus,proto3" json:"resultStatus,omitempty" json:"result_status"`
	Reporter          *Reporter                                `protobuf:"bytes,8,opt,name=reporter,proto3" json:"reporter,omitempty" json:"reporter"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_07dc4fde39a045c2, []int{0}
}
func (m *Report) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Report.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return m.Size()
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *Report) GetDataSourceResults() []*DataSourceResult {
	if m != nil {
		return m.DataSourceResults
	}
	return nil
}

func (m *Report) GetTestCaseResults() []*TestCaseResult {
	if m != nil {
		return m.TestCaseResults
	}
	return nil
}

func (m *Report) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Report) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *Report) GetAggregatedResult() []byte {
	if m != nil {
		return m.AggregatedResult
	}
	return nil
}

func (m *Report) GetResultStatus() string {
	if m != nil {
		return m.ResultStatus
	}
	return ""
}

func (m *Report) GetReporter() *Reporter {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func init() {
	proto.RegisterType((*Report)(nil), "oraichain.orai.websocket.Report")
}

func init() {
	proto.RegisterFile("websocket/types/types_report.proto", fileDescriptor_07dc4fde39a045c2)
}

var fileDescriptor_07dc4fde39a045c2 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x91, 0x34, 0xb4, 0x4e, 0xa5, 0x52, 0xb7, 0x08, 0x37, 0x42, 0xbe, 0xe8, 0xc4, 0x60,
	0x15, 0x61, 0xab, 0xe9, 0x44, 0xc5, 0xe4, 0xb6, 0x52, 0x59, 0x1d, 0x26, 0x16, 0xeb, 0xec, 0xbc,
	0x3a, 0x26, 0x6d, 0x2e, 0xdc, 0x5d, 0xf8, 0x31, 0xb3, 0x32, 0x30, 0x32, 0x32, 0xf3, 0x97, 0x74,
	0xec, 0xc8, 0x64, 0x50, 0x22, 0xf1, 0x07, 0xf8, 0x2f, 0x40, 0xbe, 0x73, 0xf3, 0xb3, 0x61, 0x49,
	0x4e, 0xfe, 0xbe, 0xf7, 0xde, 0x77, 0xdf, 0xfb, 0xce, 0x20, 0x1f, 0x21, 0x12, 0x2c, 0xee, 0x83,
	0xf4, 0xe4, 0xe7, 0x21, 0x08, 0xfd, 0x1b, 0x72, 0x18, 0x32, 0x2e, 0xdd, 0x21, 0x67, 0x92, 0x99,
	0x16, 0xe3, 0x34, 0x8d, 0x7b, 0x34, 0x1d, 0xb8, 0xc5, 0xc9, 0x9d, 0x96, 0x34, 0xf7, 0x13, 0x96,
	0x30, 0x45, 0xf2, 0x8a, 0x93, 0xe6, 0x37, 0xed, 0x98, 0x89, 0x6b, 0x26, 0xbc, 0x88, 0x0a, 0xf0,
	0x3e, 0x1c, 0x45, 0x20, 0xe9, 0x91, 0x17, 0xb3, 0x74, 0x50, 0xe2, 0x87, 0xf7, 0xcf, 0xec, 0x0a,
	0x36, 0xe2, 0x31, 0x84, 0x1c, 0xc4, 0xe8, 0xaa, 0x9c, 0xdd, 0x74, 0xee, 0xe7, 0xca, 0x98, 0x8a,
	0x25, 0xe6, 0xb3, 0xff, 0xdd, 0x04, 0xb8, 0x66, 0x91, 0xbf, 0x1b, 0x46, 0x3d, 0x50, 0x9f, 0xcc,
	0x63, 0x63, 0x8b, 0xc3, 0xfb, 0x11, 0x08, 0xf9, 0xfa, 0xcc, 0x42, 0x2d, 0xe4, 0x6c, 0xf9, 0x8f,
	0xf3, 0x0c, 0xef, 0xbe, 0x13, 0x6c, 0x70, 0x42, 0x4a, 0x28, 0x4c, 0xbb, 0x24, 0x98, 0xf1, 0xcc,
	0xaf, 0xc8, 0xd8, 0x3d, 0xa3, 0x92, 0x76, 0x94, 0xd6, 0x40, 0x09, 0x10, 0xd6, 0x83, 0x56, 0xd5,
	0x69, 0xb4, 0x0f, 0xdd, 0x75, 0x46, 0xb9, 0xcb, 0x25, 0x7e, 0x7b, 0x9c, 0xe1, 0xd5, 0x46, 0x79,
	0x86, 0x0f, 0xf4, 0xf8, 0x2e, 0x95, 0x34, 0x5c, 0x30, 0x84, 0x04, 0xab, 0x7c, 0xf3, 0x0b, 0x32,
	0x76, 0xde, 0x80, 0x90, 0xa7, 0x54, 0x4c, 0xc5, 0x54, 0x95, 0x18, 0x67, 0xbd, 0x98, 0xc5, 0x02,
	0xdf, 0x1b, 0x67, 0x78, 0xb9, 0x49, 0x9e, 0x61, 0x4b, 0x0b, 0x91, 0x85, 0x09, 0x73, 0x6e, 0x0b,
	0x12, 0x2c, 0x93, 0xcd, 0x97, 0x46, 0x23, 0xba, 0x62, 0x71, 0xff, 0x02, 0xd2, 0xa4, 0x27, 0xad,
	0x5a, 0x0b, 0x39, 0x55, 0xff, 0x49, 0x9e, 0xe1, 0x3d, 0xdd, 0x43, 0x81, 0x61, 0x4f, 0xa1, 0x24,
	0x98, 0xe7, 0x9a, 0xd2, 0xa8, 0x5d, 0x02, 0x08, 0x6b, 0x43, 0x89, 0x3e, 0x70, 0x75, 0x74, 0xdc,
	0x22, 0x3a, 0x6e, 0x19, 0x1d, 0xf7, 0x94, 0xa5, 0x03, 0xff, 0xfc, 0x26, 0xc3, 0x95, 0xf9, 0xf5,
	0x14, 0x3b, 0x0c, 0x2f, 0x01, 0xc8, 0xcf, 0xdf, 0xd8, 0x49, 0x52, 0xd9, 0x1b, 0x45, 0x6e, 0xcc,
	0xae, 0xbd, 0x32, 0x7c, 0xfa, 0xef, 0x85, 0xe8, 0xf6, 0x75, 0x10, 0x54, 0x17, 0x11, 0xa8, 0x69,
	0xe6, 0x85, 0xf1, 0x88, 0x26, 0x09, 0x87, 0x84, 0x4a, 0xe8, 0xea, 0x5b, 0x58, 0xf5, 0x16, 0x72,
	0xb6, 0xfd, 0xa7, 0xb3, 0x9b, 0xcf, 0x18, 0xd3, 0x0d, 0xac, 0x54, 0x99, 0xaf, 0x8c, 0x6d, 0x0d,
	0x76, 0x24, 0x95, 0x23, 0x61, 0x3d, 0x54, 0x39, 0xb2, 0xf2, 0x0c, 0xef, 0xdf, 0x09, 0x2d, 0xd0,
	0x50, 0x28, 0x98, 0x04, 0x0b, 0x6c, 0xb3, 0x63, 0x6c, 0xde, 0xe5, 0xd3, 0xda, 0x6c, 0x21, 0xa7,
	0xd1, 0x26, 0xeb, 0xd7, 0x16, 0x94, 0x4c, 0x7f, 0x2f, 0xcf, 0xf0, 0xce, 0xbc, 0x0d, 0xc0, 0x49,
	0x30, 0x6d, 0x74, 0x52, 0xfb, 0xfe, 0x03, 0x23, 0xff, 0xfc, 0x66, 0x6c, 0xa3, 0xdb, 0xb1, 0x8d,
	0xfe, 0x8c, 0x6d, 0xf4, 0x6d, 0x62, 0x57, 0x6e, 0x27, 0x76, 0xe5, 0xd7, 0xc4, 0xae, 0xbc, 0x7d,
	0x3e, 0x67, 0xd6, 0x74, 0x98, 0x3a, 0x79, 0x9f, 0xbc, 0xa5, 0x47, 0x14, 0xd5, 0xd5, 0xb3, 0x39,
	0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x59, 0x5b, 0x60, 0x28, 0x04, 0x00, 0x00,
}

func (m *Report) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Report) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Report) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Reporter != nil {
		{
			size, err := m.Reporter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypesReport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.ResultStatus) > 0 {
		i -= len(m.ResultStatus)
		copy(dAtA[i:], m.ResultStatus)
		i = encodeVarintTypesReport(dAtA, i, uint64(len(m.ResultStatus)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.AggregatedResult) > 0 {
		i -= len(m.AggregatedResult)
		copy(dAtA[i:], m.AggregatedResult)
		i = encodeVarintTypesReport(dAtA, i, uint64(len(m.AggregatedResult)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesReport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypesReport(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TestCaseResults) > 0 {
		for iNdEx := len(m.TestCaseResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TestCaseResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesReport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DataSourceResults) > 0 {
		for iNdEx := len(m.DataSourceResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataSourceResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesReport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RequestID) > 0 {
		i -= len(m.RequestID)
		copy(dAtA[i:], m.RequestID)
		i = encodeVarintTypesReport(dAtA, i, uint64(len(m.RequestID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesReport(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesReport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Report) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestID)
	if l > 0 {
		n += 1 + l + sovTypesReport(uint64(l))
	}
	if len(m.DataSourceResults) > 0 {
		for _, e := range m.DataSourceResults {
			l = e.Size()
			n += 1 + l + sovTypesReport(uint64(l))
		}
	}
	if len(m.TestCaseResults) > 0 {
		for _, e := range m.TestCaseResults {
			l = e.Size()
			n += 1 + l + sovTypesReport(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypesReport(uint64(m.BlockHeight))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovTypesReport(uint64(l))
		}
	}
	l = len(m.AggregatedResult)
	if l > 0 {
		n += 1 + l + sovTypesReport(uint64(l))
	}
	l = len(m.ResultStatus)
	if l > 0 {
		n += 1 + l + sovTypesReport(uint64(l))
	}
	if m.Reporter != nil {
		l = m.Reporter.Size()
		n += 1 + l + sovTypesReport(uint64(l))
	}
	return n
}

func sovTypesReport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesReport(x uint64) (n int) {
	return sovTypesReport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Report) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesReport
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
			return fmt.Errorf("proto: Report: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Report: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSourceResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataSourceResults = append(m.DataSourceResults, &DataSourceResult{})
			if err := m.DataSourceResults[len(m.DataSourceResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestCaseResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestCaseResults = append(m.TestCaseResults, &TestCaseResult{})
			if err := m.TestCaseResults[len(m.TestCaseResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatedResult", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregatedResult = append(m.AggregatedResult[:0], dAtA[iNdEx:postIndex]...)
			if m.AggregatedResult == nil {
				m.AggregatedResult = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesReport
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
				return ErrInvalidLengthTypesReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Reporter == nil {
				m.Reporter = &Reporter{}
			}
			if err := m.Reporter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesReport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesReport
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
func skipTypesReport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesReport
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
					return 0, ErrIntOverflowTypesReport
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
					return 0, ErrIntOverflowTypesReport
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
				return 0, ErrInvalidLengthTypesReport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesReport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesReport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesReport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesReport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesReport = fmt.Errorf("proto: unexpected end of group")
)