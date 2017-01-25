// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/build/info.proto
// DO NOT EDIT!

/*
	Package build is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/build/info.proto

	It has these top-level messages:
		Info
*/
package build

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Info describes build information for this CockroachDB binary.
type Info struct {
	GoVersion    string `protobuf:"bytes,1,opt,name=go_version,json=goVersion" json:"go_version"`
	Tag          string `protobuf:"bytes,2,opt,name=tag" json:"tag"`
	Time         string `protobuf:"bytes,3,opt,name=time" json:"time"`
	Revision     string `protobuf:"bytes,4,opt,name=revision" json:"revision"`
	CgoCompiler  string `protobuf:"bytes,5,opt,name=cgo_compiler,json=cgoCompiler" json:"cgo_compiler"`
	Platform     string `protobuf:"bytes,6,opt,name=platform" json:"platform"`
	Distribution string `protobuf:"bytes,7,opt,name=distribution" json:"distribution"`
	Type         string `protobuf:"bytes,8,opt,name=type" json:"type"`
	// dependencies exists to allow tests that run against old clusters
	// to unmarshal JSON containing this field. The tag is unimportant,
	// but the field name must remain unchanged.
	//
	// alternatively, we could set jsonpb.Unmarshaler.AllowUnknownFields
	// to true in httputil.doJSONRequest, but that comes at the expense
	// of run-time type checking, which is nice to have.
	Dependencies *string `protobuf:"bytes,10000,opt,name=dependencies" json:"dependencies,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{0} }

func init() {
	proto.RegisterType((*Info)(nil), "cockroach.build.Info")
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.GoVersion)))
	i += copy(dAtA[i:], m.GoVersion)
	dAtA[i] = 0x12
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Tag)))
	i += copy(dAtA[i:], m.Tag)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Time)))
	i += copy(dAtA[i:], m.Time)
	dAtA[i] = 0x22
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Revision)))
	i += copy(dAtA[i:], m.Revision)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.CgoCompiler)))
	i += copy(dAtA[i:], m.CgoCompiler)
	dAtA[i] = 0x32
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Platform)))
	i += copy(dAtA[i:], m.Platform)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Distribution)))
	i += copy(dAtA[i:], m.Distribution)
	dAtA[i] = 0x42
	i++
	i = encodeVarintInfo(dAtA, i, uint64(len(m.Type)))
	i += copy(dAtA[i:], m.Type)
	if m.Dependencies != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0xf1
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(*m.Dependencies)))
		i += copy(dAtA[i:], *m.Dependencies)
	}
	return i, nil
}

func encodeFixed64Info(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Info(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.GoVersion)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Tag)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Time)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Revision)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.CgoCompiler)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Platform)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Distribution)
	n += 1 + l + sovInfo(uint64(l))
	l = len(m.Type)
	n += 1 + l + sovInfo(uint64(l))
	if m.Dependencies != nil {
		l = len(*m.Dependencies)
		n += 3 + l + sovInfo(uint64(l))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Revision = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CgoCompiler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CgoCompiler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distribution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distribution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Dependencies = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/build/info.proto", fileDescriptorInfo) }

var fileDescriptorInfo = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xe3, 0x36, 0x85, 0xd6, 0x44, 0x42, 0xb2, 0x10, 0xb2, 0x18, 0xdc, 0x8a, 0x0e, 0x74,
	0x6a, 0xee, 0x50, 0x26, 0x16, 0x86, 0x0e, 0x0c, 0x2c, 0x55, 0xea, 0x38, 0xc6, 0x6a, 0x92, 0xdf,
	0x72, 0xdc, 0x4a, 0xbd, 0x45, 0xef, 0xc2, 0x25, 0x32, 0x32, 0x32, 0x21, 0x08, 0x17, 0x41, 0x71,
	0x69, 0xe5, 0x6c, 0xd6, 0xf7, 0x3d, 0x3d, 0x3f, 0xfd, 0x98, 0x71, 0xe0, 0x1b, 0x03, 0x09, 0x7f,
	0x8b, 0xf5, 0x46, 0xc6, 0xeb, 0xad, 0xca, 0xd3, 0x58, 0x95, 0x19, 0xcc, 0xb5, 0x01, 0x0b, 0xe4,
	0xfa, 0xec, 0xe7, 0xce, 0xdd, 0xdd, 0x48, 0x90, 0xe0, 0x5c, 0xdc, 0xbe, 0x8e, 0xb1, 0xfb, 0xf7,
	0x1e, 0x0e, 0x9f, 0xca, 0x0c, 0xc8, 0x14, 0x63, 0x09, 0xab, 0x9d, 0x30, 0x95, 0x82, 0x92, 0xa2,
	0x09, 0x9a, 0x8d, 0x16, 0x61, 0xfd, 0x35, 0x0e, 0x96, 0x23, 0x09, 0x2f, 0x47, 0x4c, 0x6e, 0x71,
	0xdf, 0x26, 0x92, 0xf6, 0x3c, 0xdb, 0x02, 0x42, 0x71, 0x68, 0x55, 0x21, 0x68, 0xdf, 0x13, 0x8e,
	0x90, 0x09, 0x1e, 0x1a, 0xb1, 0x53, 0xae, 0x34, 0xf4, 0xec, 0x99, 0x92, 0x07, 0x1c, 0x71, 0x09,
	0x2b, 0x0e, 0x85, 0x56, 0xb9, 0x30, 0x74, 0xe0, 0xa5, 0xae, 0xb8, 0x84, 0xc7, 0x7f, 0xd1, 0x56,
	0xe9, 0x3c, 0xb1, 0x19, 0x98, 0x82, 0x5e, 0xf8, 0x55, 0x27, 0x4a, 0x66, 0x38, 0x4a, 0x55, 0x65,
	0x8d, 0x5a, 0x6f, 0x6d, 0xfb, 0xe1, 0xa5, 0x97, 0xea, 0x18, 0x37, 0x78, 0xaf, 0x05, 0x1d, 0x76,
	0x06, 0xef, 0xb5, 0x20, 0x53, 0x1c, 0xa5, 0x42, 0x8b, 0x32, 0x15, 0x25, 0x57, 0xa2, 0xa2, 0x87,
	0xe7, 0x36, 0xb2, 0xec, 0xc0, 0xc5, 0xb8, 0xfe, 0x61, 0x41, 0xdd, 0x30, 0xf4, 0xd1, 0x30, 0xf4,
	0xd9, 0x30, 0xf4, 0xdd, 0x30, 0x74, 0xf8, 0x65, 0xc1, 0xeb, 0xc0, 0x1d, 0xfb, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x27, 0xf2, 0xe8, 0x0f, 0x9e, 0x01, 0x00, 0x00,
}
