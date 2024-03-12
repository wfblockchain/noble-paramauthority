// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos_proto/params/v1beta1/params.proto

package proposal

import (
	fmt "fmt"
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

// ParameterChangeProposal defines a proposal to change one or more parameters.
type ParameterChangeProposal struct {
	Title       string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Changes     []ParamChange `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes"`
}

func (m *ParameterChangeProposal) Reset()         { *m = ParameterChangeProposal{} }
func (m *ParameterChangeProposal) String() string { return proto.CompactTextString(m) }
func (*ParameterChangeProposal) ProtoMessage()    {}
func (*ParameterChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d8e783c2a0505, []int{0}
}
func (m *ParameterChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParameterChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParameterChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParameterChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterChangeProposal.Merge(m, src)
}
func (m *ParameterChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *ParameterChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterChangeProposal proto.InternalMessageInfo

func (m *ParameterChangeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ParameterChangeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ParameterChangeProposal) GetChanges() []ParamChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

// ParamChange defines an individual parameter change, for use in
// ParameterChangeProposal.
type ParamChange struct {
	Subspace string `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ParamChange) Reset()         { *m = ParamChange{} }
func (m *ParamChange) String() string { return proto.CompactTextString(m) }
func (*ParamChange) ProtoMessage()    {}
func (*ParamChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d8e783c2a0505, []int{1}
}
func (m *ParamChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamChange.Merge(m, src)
}
func (m *ParamChange) XXX_Size() int {
	return m.Size()
}
func (m *ParamChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamChange.DiscardUnknown(m)
}

var xxx_messageInfo_ParamChange proto.InternalMessageInfo

func (m *ParamChange) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *ParamChange) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ParamChange) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ParameterChangeProposal)(nil), "cosmos_proto.params.v1beta1.ParameterChangeProposal")
	proto.RegisterType((*ParamChange)(nil), "cosmos_proto.params.v1beta1.ParamChange")
}

func init() {
	proto.RegisterFile("cosmos_proto/params/v1beta1/params.proto", fileDescriptor_198d8e783c2a0505)
}

var fileDescriptor_198d8e783c2a0505 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x93, 0x2f, 0x1f, 0xff, 0xdc, 0x05, 0x59, 0x95, 0x88, 0x8a, 0x64, 0xaa, 0x4e, 0x59,
	0x88, 0x55, 0x78, 0x83, 0xc2, 0xc0, 0x58, 0x3a, 0x32, 0x80, 0x6c, 0x63, 0x12, 0xab, 0x69, 0xae,
	0x65, 0x3b, 0x85, 0xbe, 0x05, 0x3b, 0x2f, 0xd4, 0xb1, 0x23, 0x13, 0x42, 0xed, 0x8b, 0xa0, 0x38,
	0x29, 0xca, 0xc4, 0x76, 0x8f, 0xcf, 0xb9, 0xf7, 0xfe, 0x6c, 0xa3, 0x44, 0x80, 0x5d, 0x80, 0x7d,
	0xd2, 0x06, 0x1c, 0x50, 0xcd, 0x0c, 0x5b, 0x58, 0xba, 0x1c, 0x73, 0xe9, 0xd8, 0xb8, 0x95, 0xa9,
	0xf7, 0xf0, 0x79, 0x37, 0x99, 0xb6, 0x56, 0x9b, 0x1c, 0xf4, 0x33, 0xc8, 0xa0, 0x99, 0x51, 0x57,
	0x4d, 0xcb, 0xe8, 0x23, 0x44, 0x67, 0xd3, 0x3a, 0x28, 0x9d, 0x34, 0x37, 0x39, 0x2b, 0x33, 0x39,
	0x35, 0xa0, 0xc1, 0xb2, 0x02, 0xf7, 0xd1, 0x81, 0x53, 0xae, 0x90, 0x71, 0x38, 0x0c, 0x93, 0x93,
	0x59, 0x23, 0xf0, 0x10, 0xf5, 0x9e, 0xa5, 0x15, 0x46, 0x69, 0xa7, 0xa0, 0x8c, 0xff, 0x79, 0xaf,
	0x7b, 0x84, 0xef, 0xd0, 0x91, 0xf0, 0x93, 0x6c, 0x1c, 0x0d, 0xa3, 0xa4, 0x77, 0x95, 0xa4, 0x7f,
	0x80, 0xa5, 0x7e, 0x7d, 0xb3, 0x7a, 0xf2, 0x7f, 0xfd, 0x75, 0x11, 0xcc, 0xf6, 0xed, 0xa3, 0x7b,
	0xd4, 0xeb, 0xb8, 0x78, 0x80, 0x8e, 0x6d, 0xc5, 0xad, 0x66, 0x62, 0xcf, 0xf4, 0xab, 0xf1, 0x29,
	0x8a, 0xe6, 0x72, 0xd5, 0xe2, 0xd4, 0x65, 0x8d, 0xbf, 0x64, 0x45, 0x25, 0xe3, 0xa8, 0xc1, 0xf7,
	0x62, 0xf2, 0xb8, 0xde, 0x92, 0x70, 0xb3, 0x25, 0xe1, 0xf7, 0x96, 0x84, 0xef, 0x3b, 0x12, 0x6c,
	0x76, 0x24, 0xf8, 0xdc, 0x91, 0xe0, 0xe1, 0x36, 0x53, 0x2e, 0xaf, 0x78, 0x2a, 0x60, 0x41, 0x5f,
	0x5f, 0x78, 0x01, 0x62, 0x2e, 0x72, 0xa6, 0x4a, 0x5a, 0x02, 0x2f, 0xe4, 0xa5, 0xa7, 0x66, 0x95,
	0xcb, 0xc1, 0x28, 0xb7, 0xa2, 0x6f, 0xfb, 0x9f, 0x70, 0x2b, 0x2d, 0x2d, 0xd5, 0xed, 0xa3, 0xf1,
	0x43, 0x7f, 0xc7, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x0d, 0x9b, 0xd4, 0xb6, 0x01,
	0x00, 0x00,
}

func (m *ParameterChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParameterChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParameterChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ParamChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParameterChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *ParamChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParameterChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ParameterChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParameterChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, ParamChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *ParamChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ParamChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
