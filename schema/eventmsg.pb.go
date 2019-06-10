// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventmsg.protobuf

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protobuf package it is being compiled against.
// A compilation error at this line likely means your copy of the
// protobuf package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the protobuf package

type ReqQ_MsgComType int32

const (
	ReqQ_QUE ReqQ_MsgComType = 0
	ReqQ_SUB ReqQ_MsgComType = 1
)

var ReqQ_MsgComType_name = map[int32]string{
	0: "QUE",
	1: "SUB",
}
var ReqQ_MsgComType_value = map[string]int32{
	"QUE": 0,
	"SUB": 1,
}

func (x ReqQ_MsgComType) String() string {
	return proto.EnumName(ReqQ_MsgComType_name, int32(x))
}
func (ReqQ_MsgComType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{0, 0}
}

type ResEvent_ResultType int32

const (
	ResEvent_SUCCESS ResEvent_ResultType = 0
	ResEvent_FAILURE ResEvent_ResultType = 1
)

var ResEvent_ResultType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
}
var ResEvent_ResultType_value = map[string]int32{
	"SUCCESS": 0,
	"FAILURE": 1,
}

func (x ResEvent_ResultType) String() string {
	return proto.EnumName(ResEvent_ResultType_name, int32(x))
}
func (ResEvent_ResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{2, 0}
}

// ReqQ message in queue
type ReqQ struct {
	ReqId                int64           `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ComType              ReqQ_MsgComType `protobuf:"varint,3,opt,name=comType,proto3,enum=schema.ReqQ_MsgComType" json:"comType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReqQ) Reset()         { *m = ReqQ{} }
func (m *ReqQ) String() string { return proto.CompactTextString(m) }
func (*ReqQ) ProtoMessage()    {}
func (*ReqQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{0}
}
func (m *ReqQ) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReqQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReqQ.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReqQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQ.Merge(dst, src)
}
func (m *ReqQ) XXX_Size() int {
	return m.Size()
}
func (m *ReqQ) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQ.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQ proto.InternalMessageInfo

func (m *ReqQ) GetReqId() int64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *ReqQ) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqQ) GetComType() ReqQ_MsgComType {
	if m != nil {
		return m.ComType
	}
	return ReqQ_QUE
}

type ReqEvent struct {
	ReqId                int64    `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MsgBody              []byte   `protobuf:"bytes,3,opt,name=msgBody,proto3" json:"msgBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEvent) Reset()         { *m = ReqEvent{} }
func (m *ReqEvent) String() string { return proto.CompactTextString(m) }
func (*ReqEvent) ProtoMessage()    {}
func (*ReqEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{1}
}
func (m *ReqEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReqEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReqEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReqEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEvent.Merge(dst, src)
}
func (m *ReqEvent) XXX_Size() int {
	return m.Size()
}
func (m *ReqEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEvent proto.InternalMessageInfo

func (m *ReqEvent) GetReqId() int64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *ReqEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqEvent) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

type ResEvent struct {
	ReqId                int64               `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResultType           ResEvent_ResultType `protobuf:"varint,3,opt,name=resultType,proto3,enum=schema.ResEvent_ResultType" json:"resultType,omitempty"`
	ResultBody           []byte              `protobuf:"bytes,5,opt,name=resultBody,proto3" json:"resultBody,omitempty"`
	Error                *CodeError          `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ResEvent) Reset()         { *m = ResEvent{} }
func (m *ResEvent) String() string { return proto.CompactTextString(m) }
func (*ResEvent) ProtoMessage()    {}
func (*ResEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{2}
}
func (m *ResEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResEvent.Merge(dst, src)
}
func (m *ResEvent) XXX_Size() int {
	return m.Size()
}
func (m *ResEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResEvent proto.InternalMessageInfo

func (m *ResEvent) GetReqId() int64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *ResEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResEvent) GetResultType() ResEvent_ResultType {
	if m != nil {
		return m.ResultType
	}
	return ResEvent_SUCCESS
}

func (m *ResEvent) GetResultBody() []byte {
	if m != nil {
		return m.ResultBody
	}
	return nil
}

func (m *ResEvent) GetError() *CodeError {
	if m != nil {
		return m.Error
	}
	return nil
}

type CodeError struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeError) Reset()         { *m = CodeError{} }
func (m *CodeError) String() string { return proto.CompactTextString(m) }
func (*CodeError) ProtoMessage()    {}
func (*CodeError) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventmsg_ba49838ee186b968, []int{3}
}
func (m *CodeError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CodeError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeError.Merge(dst, src)
}
func (m *CodeError) XXX_Size() int {
	return m.Size()
}
func (m *CodeError) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeError.DiscardUnknown(m)
}

var xxx_messageInfo_CodeError proto.InternalMessageInfo

func (m *CodeError) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CodeError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CodeError) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqQ)(nil), "schema.ReqQ")
	proto.RegisterType((*ReqEvent)(nil), "schema.ReqEvent")
	proto.RegisterType((*ResEvent)(nil), "schema.ResEvent")
	proto.RegisterType((*CodeError)(nil), "schema.CodeError")
	proto.RegisterEnum("schema.ReqQ_MsgComType", ReqQ_MsgComType_name, ReqQ_MsgComType_value)
	proto.RegisterEnum("schema.ResEvent_ResultType", ResEvent_ResultType_name, ResEvent_ResultType_value)
}
func (m *ReqQ) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqQ) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.ReqId))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.ComType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.ComType))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ReqEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.ReqId))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.MsgBody) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.MsgBody)))
		i += copy(dAtA[i:], m.MsgBody)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.ReqId))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.ResultType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.ResultType))
	}
	if m.Error != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ResultBody) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.ResultBody)))
		i += copy(dAtA[i:], m.ResultBody)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CodeError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeError) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Prefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEventmsg(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEventmsg(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReqQ) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqId != 0 {
		n += 1 + sovEventmsg(uint64(m.ReqId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	if m.ComType != 0 {
		n += 1 + sovEventmsg(uint64(m.ComType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReqEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqId != 0 {
		n += 1 + sovEventmsg(uint64(m.ReqId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	l = len(m.MsgBody)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqId != 0 {
		n += 1 + sovEventmsg(uint64(m.ReqId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	if m.ResultType != 0 {
		n += 1 + sovEventmsg(uint64(m.ResultType))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEventmsg(uint64(l))
	}
	l = len(m.ResultBody)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CodeError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovEventmsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEventmsg(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEventmsg(x uint64) (n int) {
	return sovEventmsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReqQ) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventmsg
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
			return fmt.Errorf("protobuf: ReqQ: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("protobuf: ReqQ: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ReqId", wireType)
			}
			m.ReqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ComType", wireType)
			}
			m.ComType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComType |= (ReqQ_MsgComType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEventmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventmsg
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
func (m *ReqEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventmsg
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
			return fmt.Errorf("protobuf: ReqEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("protobuf: ReqEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ReqId", wireType)
			}
			m.ReqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field MsgBody", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgBody = append(m.MsgBody[:0], dAtA[iNdEx:postIndex]...)
			if m.MsgBody == nil {
				m.MsgBody = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventmsg
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
func (m *ResEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventmsg
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
			return fmt.Errorf("protobuf: ResEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("protobuf: ResEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ReqId", wireType)
			}
			m.ReqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ResultType", wireType)
			}
			m.ResultType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultType |= (ResEvent_ResultType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &CodeError{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field ResultBody", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultBody = append(m.ResultBody[:0], dAtA[iNdEx:postIndex]...)
			if m.ResultBody == nil {
				m.ResultBody = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventmsg
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
func (m *CodeError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventmsg
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
			return fmt.Errorf("protobuf: CodeError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("protobuf: CodeError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("protobuf: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventmsg
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
				return ErrInvalidLengthEventmsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventmsg
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
func skipEventmsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventmsg
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
					return 0, ErrIntOverflowEventmsg
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
					return 0, ErrIntOverflowEventmsg
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
				return 0, ErrInvalidLengthEventmsg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEventmsg
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
				next, err := skipEventmsg(dAtA[start:])
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
			return 0, fmt.Errorf("protobuf: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEventmsg = fmt.Errorf("protobuf: negative length found during unmarshaling")
	ErrIntOverflowEventmsg   = fmt.Errorf("protobuf: integer overflow")
)

func init() { proto.RegisterFile("eventmsg.protobuf", fileDescriptor_eventmsg_ba49838ee186b968) }

var fileDescriptor_eventmsg_ba49838ee186b968 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0xec, 0x9a, 0xb6, 0x31, 0xaf, 0x52, 0xe2, 0x22, 0x1a, 0x10, 0x42, 0xc8, 0x41, 0x73, 0x0a,
	0x58, 0x8f, 0x9e, 0x6c, 0x88, 0x50, 0x50, 0xa1, 0x1b, 0xf3, 0x01, 0xb5, 0x79, 0x46, 0xc1, 0x74,
	0xdb, 0xdd, 0x2a, 0xf6, 0xec, 0x4f, 0xf8, 0x49, 0x1e, 0xfd, 0x04, 0x89, 0x3f, 0x22, 0xbb, 0x69,
	0xda, 0x5c, 0x7b, 0x9b, 0x79, 0x6f, 0x76, 0xde, 0x0c, 0x2c, 0xf4, 0xf1, 0x1d, 0x67, 0xcb, 0x42,
	0xe6, 0xe1, 0x5c, 0xf0, 0x25, 0xa7, 0x5d, 0x39, 0x7d, 0xc6, 0x62, 0xe2, 0x7f, 0x12, 0x68, 0x33,
	0x5c, 0x8c, 0xe9, 0x11, 0x74, 0x04, 0x2e, 0x46, 0x99, 0x43, 0x3c, 0x12, 0x18, 0xac, 0x22, 0x94,
	0x42, 0x7b, 0x36, 0x29, 0xd0, 0xd9, 0xf3, 0x48, 0x60, 0x31, 0x8d, 0xe9, 0x05, 0x98, 0x53, 0x5e,
	0x3c, 0xac, 0xe6, 0xe8, 0x18, 0x1e, 0x09, 0xfa, 0x83, 0x93, 0xb0, 0x32, 0x0b, 0x95, 0x51, 0x78,
	0x27, 0xf3, 0xa8, 0x5a, 0xb3, 0x5a, 0xe7, 0xbb, 0x00, 0xdb, 0x31, 0x35, 0xc1, 0x18, 0xa7, 0xb1,
	0xdd, 0x52, 0x20, 0x49, 0x87, 0x36, 0xf1, 0xef, 0x61, 0x9f, 0xe1, 0x22, 0x56, 0x11, 0x77, 0x08,
	0xe2, 0x80, 0x59, 0xc8, 0x7c, 0xc8, 0xb3, 0x95, 0x0e, 0x72, 0xc0, 0x6a, 0xea, 0x97, 0x44, 0x19,
	0xca, 0x5d, 0x0d, 0xaf, 0x00, 0x04, 0xca, 0xb7, 0xd7, 0x65, 0xa3, 0xdc, 0xe9, 0xb6, 0x5c, 0xe5,
	0xa7, 0xc0, 0x5a, 0xc2, 0x1a, 0x72, 0x7a, 0x0e, 0x1d, 0x14, 0x82, 0x0b, 0xa7, 0xed, 0x91, 0xa0,
	0x37, 0x38, 0xac, 0xdf, 0x45, 0x3c, 0xc3, 0x58, 0x2d, 0x58, 0xb5, 0xa7, 0x6e, 0x7d, 0x45, 0x27,
	0xef, 0xe8, 0xe4, 0x8d, 0x89, 0x7f, 0x06, 0xb0, 0x3d, 0x41, 0x7b, 0x60, 0x26, 0x69, 0x14, 0xc5,
	0x49, 0x62, 0xb7, 0x14, 0xb9, 0xb9, 0x1e, 0xdd, 0xa6, 0x2c, 0xb6, 0x89, 0x3f, 0x02, 0x6b, 0xe3,
	0x4d, 0x6d, 0x30, 0x0a, 0x99, 0xeb, 0x8a, 0x16, 0x53, 0x50, 0x15, 0x9c, 0xf2, 0x6c, 0x53, 0x50,
	0x61, 0x7a, 0x0c, 0xdd, 0xb9, 0xc0, 0xa7, 0x97, 0x0f, 0x5d, 0xce, 0x62, 0x6b, 0x36, 0xb4, 0xbf,
	0x4b, 0x97, 0xfc, 0x94, 0x2e, 0xf9, 0x2d, 0x5d, 0xf2, 0xf5, 0xe7, 0xb6, 0x1e, 0xbb, 0xfa, 0x9b,
	0x5c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x9b, 0xe7, 0x58, 0x38, 0x02, 0x00, 0x00,
}
