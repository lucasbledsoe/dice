// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifier.proto

package dice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CompareOp_CompareType int32

const (
	CompareOp_UNDEFINED CompareOp_CompareType = 0
	CompareOp_EQL       CompareOp_CompareType = 1
	CompareOp_LSS       CompareOp_CompareType = 2
	CompareOp_GTR       CompareOp_CompareType = 3
	CompareOp_LEQ       CompareOp_CompareType = 4
	CompareOp_GEQ       CompareOp_CompareType = 5
)

var CompareOp_CompareType_name = map[int32]string{
	0: "UNDEFINED",
	1: "EQL",
	2: "LSS",
	3: "GTR",
	4: "LEQ",
	5: "GEQ",
}

var CompareOp_CompareType_value = map[string]int32{
	"UNDEFINED": 0,
	"EQL":       1,
	"LSS":       2,
	"GTR":       3,
	"LEQ":       4,
	"GEQ":       5,
}

func (x CompareOp_CompareType) String() string {
	return proto.EnumName(CompareOp_CompareType_name, int32(x))
}

func (CompareOp_CompareType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{0, 0}
}

type Modifier_ModifierType int32

const (
	Modifier_UNSPECIFIED      Modifier_ModifierType = 0
	Modifier_REROLL           Modifier_ModifierType = 1
	Modifier_DROP_KEEP        Modifier_ModifierType = 2
	Modifier_CRITICAL_FAILURE Modifier_ModifierType = 3
	Modifier_CRITICAL_SUCCESS Modifier_ModifierType = 4
)

var Modifier_ModifierType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "REROLL",
	2: "DROP_KEEP",
	3: "CRITICAL_FAILURE",
	4: "CRITICAL_SUCCESS",
}

var Modifier_ModifierType_value = map[string]int32{
	"UNSPECIFIED":      0,
	"REROLL":           1,
	"DROP_KEEP":        2,
	"CRITICAL_FAILURE": 3,
	"CRITICAL_SUCCESS": 4,
}

func (x Modifier_ModifierType) String() string {
	return proto.EnumName(Modifier_ModifierType_name, int32(x))
}

func (Modifier_ModifierType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{1, 0}
}

type DropKeepModifier_DropKeepMethod int32

const (
	DropKeepModifier_NOT_DEFINED  DropKeepModifier_DropKeepMethod = 0
	DropKeepModifier_DROP         DropKeepModifier_DropKeepMethod = 1
	DropKeepModifier_DROP_LOWEST  DropKeepModifier_DropKeepMethod = 2
	DropKeepModifier_DROP_HIGHEST DropKeepModifier_DropKeepMethod = 3
	DropKeepModifier_KEEP         DropKeepModifier_DropKeepMethod = 4
	DropKeepModifier_KEEP_LOWEST  DropKeepModifier_DropKeepMethod = 5
	DropKeepModifier_KEEP_HIGHEST DropKeepModifier_DropKeepMethod = 6
)

var DropKeepModifier_DropKeepMethod_name = map[int32]string{
	0: "NOT_DEFINED",
	1: "DROP",
	2: "DROP_LOWEST",
	3: "DROP_HIGHEST",
	4: "KEEP",
	5: "KEEP_LOWEST",
	6: "KEEP_HIGHEST",
}

var DropKeepModifier_DropKeepMethod_value = map[string]int32{
	"NOT_DEFINED":  0,
	"DROP":         1,
	"DROP_LOWEST":  2,
	"DROP_HIGHEST": 3,
	"KEEP":         4,
	"KEEP_LOWEST":  5,
	"KEEP_HIGHEST": 6,
}

func (x DropKeepModifier_DropKeepMethod) String() string {
	return proto.EnumName(DropKeepModifier_DropKeepMethod_name, int32(x))
}

func (DropKeepModifier_DropKeepMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{2, 0}
}

type CompareOp struct {
	Type                 CompareOp_CompareType `protobuf:"varint,1,opt,name=type,proto3,enum=dice.CompareOp_CompareType" json:"type,omitempty"`
	Target               int64                 `protobuf:"zigzag64,2,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CompareOp) Reset()         { *m = CompareOp{} }
func (m *CompareOp) String() string { return proto.CompactTextString(m) }
func (*CompareOp) ProtoMessage()    {}
func (*CompareOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{0}
}

func (m *CompareOp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompareOp.Unmarshal(m, b)
}
func (m *CompareOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompareOp.Marshal(b, m, deterministic)
}
func (m *CompareOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompareOp.Merge(m, src)
}
func (m *CompareOp) XXX_Size() int {
	return xxx_messageInfo_CompareOp.Size(m)
}
func (m *CompareOp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompareOp.DiscardUnknown(m)
}

var xxx_messageInfo_CompareOp proto.InternalMessageInfo

func (m *CompareOp) GetType() CompareOp_CompareType {
	if m != nil {
		return m.Type
	}
	return CompareOp_UNDEFINED
}

func (m *CompareOp) GetTarget() int64 {
	if m != nil {
		return m.Target
	}
	return 0
}

type Modifier struct {
	// Types that are valid to be assigned to ModifierDifferentiatingFields:
	//	*Modifier_DropKeep
	//	*Modifier_CriticalSuccess
	//	*Modifier_CriticalFailure
	//	*Modifier_Reroll
	ModifierDifferentiatingFields isModifier_ModifierDifferentiatingFields `protobuf_oneof:"modifier_differentiating_fields"`
	XXX_NoUnkeyedLiteral          struct{}                                 `json:"-"`
	XXX_unrecognized              []byte                                   `json:"-"`
	XXX_sizecache                 int32                                    `json:"-"`
}

func (m *Modifier) Reset()         { *m = Modifier{} }
func (m *Modifier) String() string { return proto.CompactTextString(m) }
func (*Modifier) ProtoMessage()    {}
func (*Modifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{1}
}

func (m *Modifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Modifier.Unmarshal(m, b)
}
func (m *Modifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Modifier.Marshal(b, m, deterministic)
}
func (m *Modifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Modifier.Merge(m, src)
}
func (m *Modifier) XXX_Size() int {
	return xxx_messageInfo_Modifier.Size(m)
}
func (m *Modifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Modifier.DiscardUnknown(m)
}

var xxx_messageInfo_Modifier proto.InternalMessageInfo

type isModifier_ModifierDifferentiatingFields interface {
	isModifier_ModifierDifferentiatingFields()
}

type Modifier_DropKeep struct {
	DropKeep *DropKeepModifier `protobuf:"bytes,2,opt,name=drop_keep,json=dropKeep,proto3,oneof"`
}

type Modifier_CriticalSuccess struct {
	CriticalSuccess *CriticalSuccessModifier `protobuf:"bytes,3,opt,name=critical_success,json=criticalSuccess,proto3,oneof"`
}

type Modifier_CriticalFailure struct {
	CriticalFailure *CriticalFailureModifier `protobuf:"bytes,4,opt,name=critical_failure,json=criticalFailure,proto3,oneof"`
}

type Modifier_Reroll struct {
	Reroll *RerollModifier `protobuf:"bytes,5,opt,name=reroll,proto3,oneof"`
}

func (*Modifier_DropKeep) isModifier_ModifierDifferentiatingFields() {}

func (*Modifier_CriticalSuccess) isModifier_ModifierDifferentiatingFields() {}

func (*Modifier_CriticalFailure) isModifier_ModifierDifferentiatingFields() {}

func (*Modifier_Reroll) isModifier_ModifierDifferentiatingFields() {}

func (m *Modifier) GetModifierDifferentiatingFields() isModifier_ModifierDifferentiatingFields {
	if m != nil {
		return m.ModifierDifferentiatingFields
	}
	return nil
}

func (m *Modifier) GetDropKeep() *DropKeepModifier {
	if x, ok := m.GetModifierDifferentiatingFields().(*Modifier_DropKeep); ok {
		return x.DropKeep
	}
	return nil
}

func (m *Modifier) GetCriticalSuccess() *CriticalSuccessModifier {
	if x, ok := m.GetModifierDifferentiatingFields().(*Modifier_CriticalSuccess); ok {
		return x.CriticalSuccess
	}
	return nil
}

func (m *Modifier) GetCriticalFailure() *CriticalFailureModifier {
	if x, ok := m.GetModifierDifferentiatingFields().(*Modifier_CriticalFailure); ok {
		return x.CriticalFailure
	}
	return nil
}

func (m *Modifier) GetReroll() *RerollModifier {
	if x, ok := m.GetModifierDifferentiatingFields().(*Modifier_Reroll); ok {
		return x.Reroll
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Modifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Modifier_DropKeep)(nil),
		(*Modifier_CriticalSuccess)(nil),
		(*Modifier_CriticalFailure)(nil),
		(*Modifier_Reroll)(nil),
	}
}

type DropKeepModifier struct {
	Method               DropKeepModifier_DropKeepMethod `protobuf:"varint,1,opt,name=method,proto3,enum=dice.DropKeepModifier_DropKeepMethod" json:"method,omitempty"`
	Count                int32                           `protobuf:"zigzag32,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DropKeepModifier) Reset()         { *m = DropKeepModifier{} }
func (m *DropKeepModifier) String() string { return proto.CompactTextString(m) }
func (*DropKeepModifier) ProtoMessage()    {}
func (*DropKeepModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{2}
}

func (m *DropKeepModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropKeepModifier.Unmarshal(m, b)
}
func (m *DropKeepModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropKeepModifier.Marshal(b, m, deterministic)
}
func (m *DropKeepModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropKeepModifier.Merge(m, src)
}
func (m *DropKeepModifier) XXX_Size() int {
	return xxx_messageInfo_DropKeepModifier.Size(m)
}
func (m *DropKeepModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_DropKeepModifier.DiscardUnknown(m)
}

var xxx_messageInfo_DropKeepModifier proto.InternalMessageInfo

func (m *DropKeepModifier) GetMethod() DropKeepModifier_DropKeepMethod {
	if m != nil {
		return m.Method
	}
	return DropKeepModifier_NOT_DEFINED
}

func (m *DropKeepModifier) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CriticalSuccessModifier struct {
	Compare              *CompareOp `protobuf:"bytes,1,opt,name=compare,proto3" json:"compare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CriticalSuccessModifier) Reset()         { *m = CriticalSuccessModifier{} }
func (m *CriticalSuccessModifier) String() string { return proto.CompactTextString(m) }
func (*CriticalSuccessModifier) ProtoMessage()    {}
func (*CriticalSuccessModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{3}
}

func (m *CriticalSuccessModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalSuccessModifier.Unmarshal(m, b)
}
func (m *CriticalSuccessModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalSuccessModifier.Marshal(b, m, deterministic)
}
func (m *CriticalSuccessModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalSuccessModifier.Merge(m, src)
}
func (m *CriticalSuccessModifier) XXX_Size() int {
	return xxx_messageInfo_CriticalSuccessModifier.Size(m)
}
func (m *CriticalSuccessModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalSuccessModifier.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalSuccessModifier proto.InternalMessageInfo

func (m *CriticalSuccessModifier) GetCompare() *CompareOp {
	if m != nil {
		return m.Compare
	}
	return nil
}

type CriticalFailureModifier struct {
	Compare              *CompareOp `protobuf:"bytes,1,opt,name=compare,proto3" json:"compare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CriticalFailureModifier) Reset()         { *m = CriticalFailureModifier{} }
func (m *CriticalFailureModifier) String() string { return proto.CompactTextString(m) }
func (*CriticalFailureModifier) ProtoMessage()    {}
func (*CriticalFailureModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{4}
}

func (m *CriticalFailureModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalFailureModifier.Unmarshal(m, b)
}
func (m *CriticalFailureModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalFailureModifier.Marshal(b, m, deterministic)
}
func (m *CriticalFailureModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalFailureModifier.Merge(m, src)
}
func (m *CriticalFailureModifier) XXX_Size() int {
	return xxx_messageInfo_CriticalFailureModifier.Size(m)
}
func (m *CriticalFailureModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalFailureModifier.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalFailureModifier proto.InternalMessageInfo

func (m *CriticalFailureModifier) GetCompare() *CompareOp {
	if m != nil {
		return m.Compare
	}
	return nil
}

type RerollModifier struct {
	Compare              *CompareOp `protobuf:"bytes,1,opt,name=compare,proto3" json:"compare,omitempty"`
	Once                 bool       `protobuf:"varint,2,opt,name=once,proto3" json:"once,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RerollModifier) Reset()         { *m = RerollModifier{} }
func (m *RerollModifier) String() string { return proto.CompactTextString(m) }
func (*RerollModifier) ProtoMessage()    {}
func (*RerollModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_28005a9423eec0f2, []int{5}
}

func (m *RerollModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RerollModifier.Unmarshal(m, b)
}
func (m *RerollModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RerollModifier.Marshal(b, m, deterministic)
}
func (m *RerollModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RerollModifier.Merge(m, src)
}
func (m *RerollModifier) XXX_Size() int {
	return xxx_messageInfo_RerollModifier.Size(m)
}
func (m *RerollModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_RerollModifier.DiscardUnknown(m)
}

var xxx_messageInfo_RerollModifier proto.InternalMessageInfo

func (m *RerollModifier) GetCompare() *CompareOp {
	if m != nil {
		return m.Compare
	}
	return nil
}

func (m *RerollModifier) GetOnce() bool {
	if m != nil {
		return m.Once
	}
	return false
}

func init() {
	proto.RegisterEnum("dice.CompareOp_CompareType", CompareOp_CompareType_name, CompareOp_CompareType_value)
	proto.RegisterEnum("dice.Modifier_ModifierType", Modifier_ModifierType_name, Modifier_ModifierType_value)
	proto.RegisterEnum("dice.DropKeepModifier_DropKeepMethod", DropKeepModifier_DropKeepMethod_name, DropKeepModifier_DropKeepMethod_value)
	proto.RegisterType((*CompareOp)(nil), "dice.CompareOp")
	proto.RegisterType((*Modifier)(nil), "dice.Modifier")
	proto.RegisterType((*DropKeepModifier)(nil), "dice.DropKeepModifier")
	proto.RegisterType((*CriticalSuccessModifier)(nil), "dice.CriticalSuccessModifier")
	proto.RegisterType((*CriticalFailureModifier)(nil), "dice.CriticalFailureModifier")
	proto.RegisterType((*RerollModifier)(nil), "dice.RerollModifier")
}

func init() { proto.RegisterFile("modifier.proto", fileDescriptor_28005a9423eec0f2) }

var fileDescriptor_28005a9423eec0f2 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x63, 0xc7, 0x4d, 0x26, 0x6d, 0xba, 0xac, 0xaa, 0x62, 0x81, 0x10, 0xc5, 0x12, 0x52,
	0x39, 0x90, 0x4a, 0x45, 0x1c, 0x39, 0xb4, 0xce, 0xa6, 0x35, 0x35, 0x71, 0xbb, 0x4e, 0x84, 0xc4,
	0xc5, 0x72, 0xed, 0x75, 0xba, 0x22, 0x89, 0xad, 0x8d, 0x83, 0x54, 0xf1, 0x1b, 0xdc, 0xf9, 0x39,
	0x3e, 0x04, 0xed, 0xda, 0xa6, 0x49, 0x14, 0x0e, 0xbd, 0xcd, 0xbe, 0xf7, 0xe6, 0x79, 0x67, 0xfc,
	0x6c, 0xe8, 0xce, 0xb2, 0x84, 0xa7, 0x9c, 0x89, 0x5e, 0x2e, 0xb2, 0x22, 0xc3, 0x46, 0xc2, 0x63,
	0x66, 0xff, 0xd6, 0xa0, 0xed, 0x64, 0xb3, 0x3c, 0x12, 0xcc, 0xcf, 0xf1, 0x29, 0x18, 0xc5, 0x43,
	0xce, 0x2c, 0xed, 0x58, 0x3b, 0xe9, 0x9e, 0xbd, 0xec, 0x49, 0x49, 0xef, 0x1f, 0x5d, 0x57, 0xa3,
	0x87, 0x9c, 0x51, 0x25, 0xc4, 0x47, 0x60, 0x16, 0x91, 0x98, 0xb0, 0xc2, 0x6a, 0x1c, 0x6b, 0x27,
	0x98, 0x56, 0x27, 0xdb, 0x85, 0xce, 0x8a, 0x18, 0xef, 0x43, 0x7b, 0x3c, 0xec, 0x93, 0x81, 0x3b,
	0x24, 0x7d, 0xb4, 0x83, 0x77, 0x41, 0x27, 0xb7, 0x1e, 0xd2, 0x64, 0xe1, 0x05, 0x01, 0x6a, 0xc8,
	0xe2, 0x72, 0x44, 0x91, 0xae, 0x10, 0x72, 0x8b, 0x0c, 0x85, 0x90, 0x5b, 0xd4, 0xb4, 0x7f, 0xe9,
	0xd0, 0xfa, 0x52, 0x5d, 0x1d, 0x7f, 0x84, 0x76, 0x22, 0xb2, 0x3c, 0xfc, 0xce, 0x58, 0xae, 0x1e,
	0xd9, 0x39, 0x3b, 0x2a, 0x6f, 0xd9, 0x17, 0x59, 0x7e, 0xcd, 0x58, 0x5e, 0x4b, 0xaf, 0x76, 0x68,
	0x2b, 0xa9, 0x30, 0xfc, 0x19, 0x50, 0x2c, 0x78, 0xc1, 0xe3, 0x68, 0x1a, 0x2e, 0x96, 0x71, 0xcc,
	0x16, 0x0b, 0x4b, 0x57, 0xdd, 0xaf, 0xaa, 0x19, 0x2b, 0x36, 0x28, 0xc9, 0x15, 0x93, 0x83, 0x78,
	0x9d, 0x5a, 0xf3, 0x4a, 0x23, 0x3e, 0x5d, 0x0a, 0x66, 0x19, 0xdb, 0xbc, 0x06, 0x25, 0xb9, 0xcd,
	0xab, 0xa2, 0x70, 0x0f, 0x4c, 0xc1, 0x44, 0x36, 0x9d, 0x5a, 0x4d, 0xe5, 0x70, 0x58, 0x3a, 0x50,
	0x85, 0xad, 0x34, 0x56, 0x2a, 0x3b, 0x85, 0xbd, 0x1a, 0x55, 0x7b, 0x3d, 0x80, 0xce, 0x78, 0x18,
	0xdc, 0x10, 0xc7, 0x1d, 0xb8, 0x6a, 0xb3, 0x00, 0x26, 0x25, 0xd4, 0xf7, 0xe4, 0x72, 0xf7, 0xa1,
	0xdd, 0xa7, 0xfe, 0x4d, 0x78, 0x4d, 0xc8, 0x0d, 0x6a, 0xe0, 0x43, 0x40, 0x0e, 0x75, 0x47, 0xae,
	0x73, 0xee, 0x85, 0x83, 0x73, 0xd7, 0x1b, 0x53, 0x82, 0xf4, 0x35, 0x34, 0x18, 0x3b, 0x0e, 0x09,
	0x02, 0x64, 0x5c, 0xbc, 0x81, 0xd7, 0x75, 0x5a, 0xc2, 0x84, 0xa7, 0x29, 0x13, 0x6c, 0x5e, 0xf0,
	0xa8, 0xe0, 0xf3, 0x49, 0x98, 0x72, 0x36, 0x4d, 0x16, 0xf6, 0x1f, 0x0d, 0xd0, 0xe6, 0xce, 0xf1,
	0x27, 0x30, 0x67, 0xac, 0xb8, 0xcf, 0x92, 0x2a, 0x41, 0x6f, 0xb7, 0xbf, 0x9b, 0x47, 0x40, 0x89,
	0x69, 0xd5, 0x84, 0x0f, 0xa1, 0x19, 0x67, 0xcb, 0x79, 0x19, 0xa6, 0x67, 0xb4, 0x3c, 0xd8, 0x3f,
	0xa1, 0xbb, 0xae, 0x97, 0x63, 0x0f, 0xfd, 0x51, 0xf8, 0x18, 0xa8, 0x16, 0x18, 0x72, 0x54, 0xa4,
	0x49, 0x4a, 0x0d, 0xed, 0xf9, 0x5f, 0x49, 0x30, 0x42, 0x0d, 0x8c, 0x60, 0x4f, 0x01, 0x57, 0xee,
	0xe5, 0x95, 0x44, 0x74, 0x29, 0x56, 0x2b, 0x31, 0xa4, 0x58, 0x56, 0xb5, 0xb8, 0x29, 0xc5, 0x0a,
	0xa8, 0xc5, 0xa6, 0xdd, 0x87, 0xe7, 0xff, 0xc9, 0x06, 0x7e, 0x07, 0xbb, 0x71, 0x99, 0x71, 0x35,
	0x6d, 0xe7, 0xec, 0x60, 0xe3, 0x7b, 0xa1, 0x35, 0xbf, 0xea, 0xb2, 0x91, 0x8a, 0xa7, 0xb8, 0xf8,
	0xd0, 0x5d, 0x4f, 0xc6, 0x13, 0x9a, 0x31, 0x06, 0x23, 0x9b, 0xc7, 0x4c, 0xad, 0xb6, 0x45, 0x55,
	0x7d, 0xf1, 0xe2, 0x9b, 0x35, 0xe1, 0xc5, 0xfd, 0xf2, 0xae, 0x17, 0x67, 0xb3, 0xd3, 0x42, 0x44,
	0x3f, 0xf8, 0xe2, 0xfd, 0xe4, 0x54, 0x5a, 0xdc, 0x99, 0xea, 0x2f, 0xf1, 0xe1, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0xd9, 0xd1, 0x24, 0x37, 0x04, 0x00, 0x00,
}
