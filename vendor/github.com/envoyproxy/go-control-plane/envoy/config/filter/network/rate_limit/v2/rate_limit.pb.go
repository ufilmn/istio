// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
import v2 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *time.Duration `protobuf:"bytes,4,opt,name=timeout,proto3,stdduration" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	// [#comment:TODO(ramaraochavali): Make this required as part of cleanup of deprecated ratelimit
	// service config in bootstrap.]
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_rate_limit_75a078a1b4a188a4, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(dst, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}
func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if len(m.Domain) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRateLimit(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)))
		n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeDeny {
		dAtA[i] = 0x28
		i++
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RateLimitService != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(m.RateLimitService.Size()))
		n2, err := m.RateLimitService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRateLimit(uint64(l))
		}
	}
	if m.Timeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitService != nil {
		l = m.RateLimitService.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
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
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, &ratelimit.RateLimitDescriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeDeny", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailureModeDeny = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RateLimitService == nil {
				m.RateLimitService = &v2.RateLimitServiceConfig{}
			}
			if err := m.RateLimitService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
				return 0, ErrInvalidLengthRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRateLimit
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
				next, err := skipRateLimit(dAtA[start:])
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
	ErrInvalidLengthRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_rate_limit_75a078a1b4a188a4)
}

var fileDescriptor_rate_limit_75a078a1b4a188a4 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x8f, 0xd3, 0x30,
	0x14, 0x97, 0xd3, 0x5e, 0xb9, 0xba, 0x03, 0x60, 0x21, 0x11, 0x6e, 0x28, 0x01, 0x24, 0x54, 0x0e,
	0xc9, 0x16, 0x61, 0x40, 0x30, 0x96, 0x6e, 0x80, 0x74, 0x0a, 0x13, 0x2c, 0x91, 0xaf, 0x79, 0x89,
	0x2c, 0xd2, 0xbc, 0xc8, 0x71, 0xcd, 0xe5, 0x6b, 0x30, 0xf1, 0x59, 0x98, 0x18, 0x19, 0xf9, 0x06,
	0xa0, 0x6e, 0x7c, 0x08, 0x24, 0x94, 0x38, 0x69, 0x7a, 0x9d, 0xd8, 0x9e, 0xfd, 0xfb, 0xf3, 0xfc,
	0x7e, 0x7e, 0xf4, 0x15, 0x14, 0x16, 0x6b, 0xb1, 0xc6, 0x22, 0x55, 0x99, 0x48, 0x55, 0x6e, 0x40,
	0x8b, 0x02, 0xcc, 0x67, 0xd4, 0x9f, 0x84, 0x96, 0x06, 0xe2, 0x5c, 0x6d, 0x94, 0x11, 0x36, 0x3c,
	0x38, 0xf1, 0x52, 0xa3, 0x41, 0xf6, 0xa4, 0xd5, 0x72, 0xa7, 0xe5, 0x4e, 0xcb, 0x3b, 0x2d, 0x3f,
	0x60, 0xdb, 0xf0, 0xec, 0xb1, 0x6b, 0x23, 0x4b, 0xd5, 0x3b, 0x39, 0xdb, 0x7d, 0xe5, 0x2c, 0xcf,
	0x1e, 0x5d, 0x7b, 0xce, 0xc0, 0x6b, 0x44, 0x79, 0xd5, 0x91, 0xe6, 0x19, 0x62, 0x96, 0x83, 0x68,
	0x4f, 0x97, 0xdb, 0x54, 0x24, 0x5b, 0x2d, 0x8d, 0xc2, 0xa2, 0xc3, 0xef, 0x5a, 0x99, 0xab, 0x44,
	0x1a, 0x10, 0x7d, 0xd1, 0x01, 0x77, 0x32, 0xcc, 0xb0, 0x2d, 0x45, 0x53, 0xb9, 0xdb, 0x87, 0x7f,
	0x3d, 0x3a, 0x8d, 0xa4, 0x81, 0xb7, 0x4d, 0x27, 0x76, 0x4e, 0x67, 0x95, 0x91, 0x26, 0x2e, 0x35,
	0xa4, 0xea, 0xca, 0x27, 0x01, 0x59, 0x4c, 0x97, 0xd3, 0x6f, 0x7f, 0xbe, 0x8f, 0xc6, 0xda, 0x0b,
	0x48, 0x44, 0x1b, 0xf4, 0xa2, 0x05, 0xd9, 0x03, 0x3a, 0x49, 0x70, 0x23, 0x55, 0xe1, 0x7b, 0xc7,
	0xb4, 0x0e, 0x60, 0x1f, 0xe8, 0x2c, 0x81, 0x6a, 0xad, 0x55, 0x69, 0x50, 0x57, 0xfe, 0x28, 0x18,
	0x2d, 0x66, 0xe1, 0x53, 0xee, 0x92, 0x93, 0xa5, 0xe2, 0x36, 0xe4, 0x43, 0x08, 0xfb, 0x67, 0xac,
	0xf6, 0x9a, 0x25, 0x6d, 0x4c, 0x4f, 0xbe, 0x10, 0xef, 0x94, 0x44, 0x87, 0x5e, 0xec, 0x25, 0xbd,
	0x61, 0xd4, 0x06, 0x70, 0x6b, 0xfc, 0x71, 0x40, 0x16, 0xb3, 0xf0, 0x1e, 0x77, 0xc1, 0xf0, 0x3e,
	0x18, 0xbe, 0xea, 0x82, 0x59, 0x8e, 0xbf, 0xfe, 0xba, 0x4f, 0xa2, 0x9e, 0xcf, 0xce, 0xe9, 0xed,
	0x54, 0xaa, 0x7c, 0xab, 0x21, 0xde, 0x60, 0x02, 0x71, 0x02, 0x45, 0xed, 0x9f, 0x04, 0x64, 0x71,
	0x1a, 0xdd, 0xec, 0x80, 0x77, 0x98, 0xc0, 0x0a, 0x8a, 0x9a, 0xc5, 0x94, 0x0d, 0x7f, 0x19, 0x57,
	0xa0, 0xad, 0x5a, 0x83, 0x3f, 0x69, 0x3b, 0x3e, 0xe3, 0xd7, 0x56, 0x60, 0x18, 0xc4, 0x86, 0xc3,
	0x2c, 0xef, 0x9d, 0xe4, 0x75, 0xcb, 0x89, 0x6e, 0xe9, 0xa3, 0xfb, 0xe5, 0x9b, 0x1f, 0xbb, 0x39,
	0xf9, 0xb9, 0x9b, 0x93, 0xdf, 0xbb, 0x39, 0xa1, 0x2f, 0x14, 0x3a, 0xd3, 0x52, 0xe3, 0x55, 0xcd,
	0xff, 0x7b, 0xc5, 0x2e, 0xc8, 0x47, 0xcf, 0x86, 0x97, 0x93, 0x76, 0xf6, 0xe7, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0x0e, 0x23, 0x69, 0xd8, 0x02, 0x00, 0x00,
}
