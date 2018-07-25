// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/signal.proto

package sdk // import "github.com/argoproj/argo-events/sdk"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
import empty "github.com/golang/protobuf/ptypes/empty"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignalContext struct {
	Signal               *v1alpha1.Signal `protobuf:"bytes,1,opt,name=signal" json:"signal,omitempty"`
	Done                 bool             `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SignalContext) Reset()         { *m = SignalContext{} }
func (m *SignalContext) String() string { return proto.CompactTextString(m) }
func (*SignalContext) ProtoMessage()    {}
func (*SignalContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_signal_420d13aed200bc99, []int{0}
}
func (m *SignalContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignalContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignalContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SignalContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalContext.Merge(dst, src)
}
func (m *SignalContext) XXX_Size() int {
	return m.Size()
}
func (m *SignalContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalContext.DiscardUnknown(m)
}

var xxx_messageInfo_SignalContext proto.InternalMessageInfo

func (m *SignalContext) GetSignal() *v1alpha1.Signal {
	if m != nil {
		return m.Signal
	}
	return nil
}

func (m *SignalContext) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type EventContext struct {
	Event                *v1alpha1.Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Done                 bool            `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EventContext) Reset()         { *m = EventContext{} }
func (m *EventContext) String() string { return proto.CompactTextString(m) }
func (*EventContext) ProtoMessage()    {}
func (*EventContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_signal_420d13aed200bc99, []int{1}
}
func (m *EventContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventContext.Merge(dst, src)
}
func (m *EventContext) XXX_Size() int {
	return m.Size()
}
func (m *EventContext) XXX_DiscardUnknown() {
	xxx_messageInfo_EventContext.DiscardUnknown(m)
}

var xxx_messageInfo_EventContext proto.InternalMessageInfo

func (m *EventContext) GetEvent() *v1alpha1.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventContext) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterType((*SignalContext)(nil), "sdk.SignalContext")
	proto.RegisterType((*EventContext)(nil), "sdk.EventContext")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SignalService service

type SignalServiceClient interface {
	// Listen to the signal. Events are streamed back to the client.
	//
	// The following describes the mechanics for the stream.
	// The first send MUST contain a non-nil signal.
	// This first send MUST be followed by a first receive which contains
	// a confirmation in the form of done=true to signify that the Signal is set up
	// correctly.
	//
	// Later sends should only signify if the signal is done. Later sends do not
	// update the signal process attached to this channel.
	Listen(ctx context.Context, opts ...grpc.CallOption) (SignalService_ListenClient, error)
	// Ping the signal service.
	// This is used on the client-side to monitor the presence of signal services.
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type signalServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignalServiceClient(cc *grpc.ClientConn) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) Listen(ctx context.Context, opts ...grpc.CallOption) (SignalService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SignalService_serviceDesc.Streams[0], "/sdk.SignalService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &signalServiceListenClient{stream}
	return x, nil
}

type SignalService_ListenClient interface {
	Send(*SignalContext) error
	Recv() (*EventContext, error)
	grpc.ClientStream
}

type signalServiceListenClient struct {
	grpc.ClientStream
}

func (x *signalServiceListenClient) Send(m *SignalContext) error {
	return x.ClientStream.SendMsg(m)
}

func (x *signalServiceListenClient) Recv() (*EventContext, error) {
	m := new(EventContext)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *signalServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/sdk.SignalService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SignalService service

type SignalServiceServer interface {
	// Listen to the signal. Events are streamed back to the client.
	//
	// The following describes the mechanics for the stream.
	// The first send MUST contain a non-nil signal.
	// This first send MUST be followed by a first receive which contains
	// a confirmation in the form of done=true to signify that the Signal is set up
	// correctly.
	//
	// Later sends should only signify if the signal is done. Later sends do not
	// update the signal process attached to this channel.
	Listen(SignalService_ListenServer) error
	// Ping the signal service.
	// This is used on the client-side to monitor the presence of signal services.
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterSignalServiceServer(s *grpc.Server, srv SignalServiceServer) {
	s.RegisterService(&_SignalService_serviceDesc, srv)
}

func _SignalService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SignalServiceServer).Listen(&signalServiceListenServer{stream})
}

type SignalService_ListenServer interface {
	Send(*EventContext) error
	Recv() (*SignalContext, error)
	grpc.ServerStream
}

type signalServiceListenServer struct {
	grpc.ServerStream
}

func (x *signalServiceListenServer) Send(m *EventContext) error {
	return x.ServerStream.SendMsg(m)
}

func (x *signalServiceListenServer) Recv() (*SignalContext, error) {
	m := new(SignalContext)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SignalService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.SignalService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SignalService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _SignalService_Listen_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sdk/signal.proto",
}

func (m *SignalContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignalContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Signal != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSignal(dAtA, i, uint64(m.Signal.Size()))
		n1, err := m.Signal.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Done {
		dAtA[i] = 0x10
		i++
		if m.Done {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EventContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSignal(dAtA, i, uint64(m.Event.Size()))
		n2, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Done {
		dAtA[i] = 0x10
		i++
		if m.Done {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSignal(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SignalContext) Size() (n int) {
	var l int
	_ = l
	if m.Signal != nil {
		l = m.Signal.Size()
		n += 1 + l + sovSignal(uint64(l))
	}
	if m.Done {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EventContext) Size() (n int) {
	var l int
	_ = l
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovSignal(uint64(l))
	}
	if m.Done {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSignal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSignal(x uint64) (n int) {
	return sovSignal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignalContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignal
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
			return fmt.Errorf("proto: SignalContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignalContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignal
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
				return ErrInvalidLengthSignal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signal == nil {
				m.Signal = &v1alpha1.Signal{}
			}
			if err := m.Signal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Done", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignal
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
			m.Done = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSignal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSignal
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
func (m *EventContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignal
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
			return fmt.Errorf("proto: EventContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignal
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
				return ErrInvalidLengthSignal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &v1alpha1.Event{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Done", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignal
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
			m.Done = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSignal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSignal
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
func skipSignal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignal
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
					return 0, ErrIntOverflowSignal
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
					return 0, ErrIntOverflowSignal
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
				return 0, ErrInvalidLengthSignal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSignal
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
				next, err := skipSignal(dAtA[start:])
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
	ErrInvalidLengthSignal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sdk/signal.proto", fileDescriptor_signal_420d13aed200bc99) }

var fileDescriptor_signal_420d13aed200bc99 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0x8d, 0xd6, 0x22, 0x51, 0x41, 0xb3, 0x90, 0x32, 0xc2, 0x50, 0xea, 0xa6, 0x1b, 0x5f,
	0x6c, 0x8b, 0x6e, 0x15, 0xa5, 0xe0, 0xc2, 0x85, 0xb4, 0x08, 0xe2, 0x46, 0xa6, 0x9d, 0x67, 0x3a,
	0x4e, 0x9b, 0x84, 0x24, 0x2d, 0x16, 0xf4, 0x3f, 0xfc, 0x24, 0x97, 0x7e, 0x82, 0xd4, 0x1f, 0x91,
	0x26, 0x1d, 0x54, 0x50, 0x04, 0x77, 0x97, 0xcb, 0x70, 0xce, 0x9b, 0x5c, 0xba, 0x65, 0xd3, 0x9c,
	0xdb, 0x4c, 0xc8, 0x64, 0x08, 0xda, 0x28, 0xa7, 0xd8, 0x8a, 0x4d, 0xf3, 0x68, 0x57, 0x28, 0x25,
	0x86, 0xc8, 0x7d, 0xd5, 0x1b, 0xdf, 0x71, 0x1c, 0x69, 0x37, 0x0d, 0x5f, 0x44, 0xe7, 0x22, 0x73,
	0x83, 0x71, 0x0f, 0xfa, 0x6a, 0xc4, 0x13, 0x23, 0x94, 0x36, 0xea, 0xde, 0x87, 0x7d, 0x9c, 0xa0,
	0x74, 0x96, 0xeb, 0x5c, 0xf0, 0x44, 0x67, 0x96, 0x5b, 0x94, 0x56, 0x19, 0x3e, 0x69, 0x24, 0x43,
	0x3d, 0x48, 0x1a, 0x5c, 0xa0, 0x44, 0x93, 0x38, 0x4c, 0x03, 0xa9, 0xf6, 0x44, 0x37, 0xbb, 0xde,
	0x7d, 0xa6, 0xa4, 0xc3, 0x07, 0xc7, 0xae, 0x69, 0x39, 0x1c, 0x53, 0x21, 0x55, 0x52, 0x5f, 0x6f,
	0x9e, 0xc0, 0xa7, 0x0b, 0x0a, 0x97, 0x0f, 0xb7, 0xc1, 0x05, 0x3a, 0x17, 0x30, 0x77, 0x41, 0x70,
	0x41, 0xe1, 0x82, 0x00, 0xee, 0x2c, 0x78, 0x8c, 0xd1, 0x52, 0xaa, 0x24, 0x56, 0x96, 0xab, 0xa4,
	0xbe, 0xd6, 0xf1, 0xb9, 0x36, 0xa5, 0x1b, 0xed, 0x39, 0xa5, 0xb0, 0x5f, 0xd1, 0x55, 0x4f, 0x5d,
	0xc8, 0x8f, 0xff, 0x2f, 0xf7, 0xd8, 0x4e, 0xa0, 0xfd, 0xa4, 0x6e, 0x3e, 0x16, 0x7f, 0xde, 0x45,
	0x33, 0xc9, 0xfa, 0xc8, 0x5a, 0xb4, 0x7c, 0x91, 0x59, 0x87, 0x92, 0x31, 0xb0, 0x69, 0x0e, 0xdf,
	0xde, 0x25, 0xda, 0xf6, 0xdd, 0xd7, 0x63, 0xeb, 0xe4, 0x80, 0xb0, 0x23, 0x5a, 0xba, 0xcc, 0xa4,
	0x60, 0x3b, 0x10, 0xf6, 0x82, 0x62, 0x2f, 0x68, 0xcf, 0xf7, 0x8a, 0x7e, 0xe9, 0x4f, 0x0f, 0x5f,
	0x66, 0x31, 0x79, 0x9d, 0xc5, 0xe4, 0x6d, 0x16, 0x93, 0xe7, 0xf7, 0x78, 0xe9, 0x66, 0xef, 0xaf,
	0x4d, 0x6d, 0x9a, 0xf7, 0xca, 0x1e, 0xd3, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x12, 0xe7,
	0x6e, 0x35, 0x02, 0x00, 0x00,
}
