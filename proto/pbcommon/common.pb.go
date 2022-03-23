// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pbcommon/common.proto

package pbcommon

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// RaftIndex is used to track the index used while creating
// or modifying a given struct type.
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.RaftIndex
// output=common.gen.go
// name=Structs
// ignore-fields=state,sizeCache,unknownFields
type RaftIndex struct {
	// @gotags: bexpr:"-"
	CreateIndex uint64 `protobuf:"varint,1,opt,name=CreateIndex,proto3" json:"CreateIndex,omitempty" bexpr:"-"`
	// @gotags: bexpr:"-"
	ModifyIndex          uint64   `protobuf:"varint,2,opt,name=ModifyIndex,proto3" json:"ModifyIndex,omitempty" bexpr:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaftIndex) Reset()         { *m = RaftIndex{} }
func (m *RaftIndex) String() string { return proto.CompactTextString(m) }
func (*RaftIndex) ProtoMessage()    {}
func (*RaftIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{0}
}

func (m *RaftIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftIndex.Unmarshal(m, b)
}
func (m *RaftIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftIndex.Marshal(b, m, deterministic)
}
func (m *RaftIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftIndex.Merge(m, src)
}
func (m *RaftIndex) XXX_Size() int {
	return xxx_messageInfo_RaftIndex.Size(m)
}
func (m *RaftIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftIndex.DiscardUnknown(m)
}

var xxx_messageInfo_RaftIndex proto.InternalMessageInfo

func (m *RaftIndex) GetCreateIndex() uint64 {
	if m != nil {
		return m.CreateIndex
	}
	return 0
}

func (m *RaftIndex) GetModifyIndex() uint64 {
	if m != nil {
		return m.ModifyIndex
	}
	return 0
}

// TargetDatacenter is intended to be used within other messages used for RPC routing
// amongst the various Consul datacenters
type TargetDatacenter struct {
	Datacenter           string   `protobuf:"bytes,1,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetDatacenter) Reset()         { *m = TargetDatacenter{} }
func (m *TargetDatacenter) String() string { return proto.CompactTextString(m) }
func (*TargetDatacenter) ProtoMessage()    {}
func (*TargetDatacenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{1}
}

func (m *TargetDatacenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetDatacenter.Unmarshal(m, b)
}
func (m *TargetDatacenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetDatacenter.Marshal(b, m, deterministic)
}
func (m *TargetDatacenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetDatacenter.Merge(m, src)
}
func (m *TargetDatacenter) XXX_Size() int {
	return xxx_messageInfo_TargetDatacenter.Size(m)
}
func (m *TargetDatacenter) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetDatacenter.DiscardUnknown(m)
}

var xxx_messageInfo_TargetDatacenter proto.InternalMessageInfo

func (m *TargetDatacenter) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

type WriteRequest struct {
	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{2}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// ReadRequest is a type that may be embedded into any requests for read
// operations.
// It is a replacement for QueryOptions now that we no longer need any of those
// fields because we are moving away from using blocking queries.
// It is also similar to WriteRequest. It is a separate type so that in the
// future we can introduce fields that may only be relevant for reads.
type ReadRequest struct {
	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	// RequireConsistent indicates that the request must be sent to the leader.
	RequireConsistent    bool     `protobuf:"varint,2,opt,name=RequireConsistent,proto3" json:"RequireConsistent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ReadRequest) GetRequireConsistent() bool {
	if m != nil {
		return m.RequireConsistent
	}
	return false
}

// QueryOptions is used to specify various flags for read queries
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.QueryOptions
// output=common.gen.go
// name=Structs
// ignore-fields=StaleIfError,AllowNotModifiedResponse,state,sizeCache,unknownFields
type QueryOptions struct {
	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	// If set, wait until query exceeds given index. Must be provided
	// with MaxQueryTime.
	MinQueryIndex uint64 `protobuf:"varint,2,opt,name=MinQueryIndex,proto3" json:"MinQueryIndex,omitempty"`
	// Provided with MinQueryIndex to wait for change.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxQueryTime *duration.Duration `protobuf:"bytes,3,opt,name=MaxQueryTime,proto3" json:"MaxQueryTime,omitempty"`
	// If set, any follower can service the request. Results
	// may be arbitrarily stale.
	AllowStale bool `protobuf:"varint,4,opt,name=AllowStale,proto3" json:"AllowStale,omitempty"`
	// If set, the leader must verify leadership prior to
	// servicing the request. Prevents a stale read.
	RequireConsistent bool `protobuf:"varint,5,opt,name=RequireConsistent,proto3" json:"RequireConsistent,omitempty"`
	// If set, the local agent may respond with an arbitrarily stale locally
	// cached response. The semantics differ from AllowStale since the agent may
	// be entirely partitioned from the servers and still considered "healthy" by
	// operators. Stale responses from Servers are also arbitrarily stale, but can
	// provide additional bounds on the last contact time from the leader. It's
	// expected that servers that are partitioned are noticed and replaced in a
	// timely way by operators while the same may not be true for client agents.
	UseCache bool `protobuf:"varint,6,opt,name=UseCache,proto3" json:"UseCache,omitempty"`
	// If set and AllowStale is true, will try first a stale
	// read, and then will perform a consistent read if stale
	// read is older than value.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxStaleDuration *duration.Duration `protobuf:"bytes,7,opt,name=MaxStaleDuration,proto3" json:"MaxStaleDuration,omitempty"`
	// MaxAge limits how old a cached value will be returned if UseCache is true.
	// If there is a cached response that is older than the MaxAge, it is treated
	// as a cache miss and a new fetch invoked. If the fetch fails, the error is
	// returned. Clients that wish to allow for stale results on error can set
	// StaleIfError to a longer duration to change this behavior. It is ignored
	// if the endpoint supports background refresh caching. See
	// https://www.consul.io/api/index.html#agent-caching for more details.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxAge *duration.Duration `protobuf:"bytes,8,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
	// MustRevalidate forces the agent to fetch a fresh version of a cached
	// resource or at least validate that the cached version is still fresh. It is
	// implied by either max-age=0 or must-revalidate Cache-Control headers. It
	// only makes sense when UseCache is true. We store it since MaxAge = 0 is the
	// default unset value.
	MustRevalidate bool `protobuf:"varint,9,opt,name=MustRevalidate,proto3" json:"MustRevalidate,omitempty"`
	// StaleIfError specifies how stale the client will accept a cached response
	// if the servers are unavailable to fetch a fresh one. Only makes sense when
	// UseCache is true and MaxAge is set to a lower, non-zero value. It is
	// ignored if the endpoint supports background refresh caching. See
	// https://www.consul.io/api/index.html#agent-caching for more details.
	StaleIfError *duration.Duration `protobuf:"bytes,10,opt,name=StaleIfError,proto3" json:"StaleIfError,omitempty"`
	// Filter specifies the go-bexpr filter expression to be used for
	// filtering the data prior to returning a response
	Filter               string   `protobuf:"bytes,11,opt,name=Filter,proto3" json:"Filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryOptions) Reset()         { *m = QueryOptions{} }
func (m *QueryOptions) String() string { return proto.CompactTextString(m) }
func (*QueryOptions) ProtoMessage()    {}
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{4}
}

func (m *QueryOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOptions.Unmarshal(m, b)
}
func (m *QueryOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOptions.Marshal(b, m, deterministic)
}
func (m *QueryOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOptions.Merge(m, src)
}
func (m *QueryOptions) XXX_Size() int {
	return xxx_messageInfo_QueryOptions.Size(m)
}
func (m *QueryOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOptions proto.InternalMessageInfo

func (m *QueryOptions) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *QueryOptions) GetMinQueryIndex() uint64 {
	if m != nil {
		return m.MinQueryIndex
	}
	return 0
}

func (m *QueryOptions) GetMaxQueryTime() *duration.Duration {
	if m != nil {
		return m.MaxQueryTime
	}
	return nil
}

func (m *QueryOptions) GetAllowStale() bool {
	if m != nil {
		return m.AllowStale
	}
	return false
}

func (m *QueryOptions) GetRequireConsistent() bool {
	if m != nil {
		return m.RequireConsistent
	}
	return false
}

func (m *QueryOptions) GetUseCache() bool {
	if m != nil {
		return m.UseCache
	}
	return false
}

func (m *QueryOptions) GetMaxStaleDuration() *duration.Duration {
	if m != nil {
		return m.MaxStaleDuration
	}
	return nil
}

func (m *QueryOptions) GetMaxAge() *duration.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *QueryOptions) GetMustRevalidate() bool {
	if m != nil {
		return m.MustRevalidate
	}
	return false
}

func (m *QueryOptions) GetStaleIfError() *duration.Duration {
	if m != nil {
		return m.StaleIfError
	}
	return nil
}

func (m *QueryOptions) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// QueryMeta allows a query response to include potentially
// useful metadata about a query
//
// mog annotation:
//
// target=github.com/hashicorp/consul/agent/structs.QueryMeta
// output=common.gen.go
// name=Structs
// ignore-fields=NotModified,Backend,state,sizeCache,unknownFields
type QueryMeta struct {
	// This is the index associated with the read
	Index uint64 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	// If AllowStale is used, this is time elapsed since
	// last contact between the follower and leader. This
	// can be used to gauge staleness.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	LastContact *duration.Duration `protobuf:"bytes,2,opt,name=LastContact,proto3" json:"LastContact,omitempty"`
	// Used to indicate if there is a known leader node
	KnownLeader bool `protobuf:"varint,3,opt,name=KnownLeader,proto3" json:"KnownLeader,omitempty"`
	// Consistencylevel returns the consistency used to serve the query
	// Having `discovery_max_stale` on the agent can affect whether
	// the request was served by a leader.
	ConsistencyLevel string `protobuf:"bytes,4,opt,name=ConsistencyLevel,proto3" json:"ConsistencyLevel,omitempty"`
	// ResultsFilteredByACLs is true when some of the query's results were
	// filtered out by enforcing ACLs. It may be false because nothing was
	// removed, or because the endpoint does not yet support this flag.
	ResultsFilteredByACLs bool     `protobuf:"varint,7,opt,name=ResultsFilteredByACLs,proto3" json:"ResultsFilteredByACLs,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *QueryMeta) Reset()         { *m = QueryMeta{} }
func (m *QueryMeta) String() string { return proto.CompactTextString(m) }
func (*QueryMeta) ProtoMessage()    {}
func (*QueryMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{5}
}

func (m *QueryMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMeta.Unmarshal(m, b)
}
func (m *QueryMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMeta.Marshal(b, m, deterministic)
}
func (m *QueryMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMeta.Merge(m, src)
}
func (m *QueryMeta) XXX_Size() int {
	return xxx_messageInfo_QueryMeta.Size(m)
}
func (m *QueryMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMeta.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMeta proto.InternalMessageInfo

func (m *QueryMeta) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *QueryMeta) GetLastContact() *duration.Duration {
	if m != nil {
		return m.LastContact
	}
	return nil
}

func (m *QueryMeta) GetKnownLeader() bool {
	if m != nil {
		return m.KnownLeader
	}
	return false
}

func (m *QueryMeta) GetConsistencyLevel() string {
	if m != nil {
		return m.ConsistencyLevel
	}
	return ""
}

func (m *QueryMeta) GetResultsFilteredByACLs() bool {
	if m != nil {
		return m.ResultsFilteredByACLs
	}
	return false
}

// EnterpriseMeta contains metadata that is only used by the Enterprise version
// of Consul.
type EnterpriseMeta struct {
	// Namespace in which the entity exists.
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// Partition in which the entity exists.
	Partition            string   `protobuf:"bytes,2,opt,name=Partition,proto3" json:"Partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterpriseMeta) Reset()         { *m = EnterpriseMeta{} }
func (m *EnterpriseMeta) String() string { return proto.CompactTextString(m) }
func (*EnterpriseMeta) ProtoMessage()    {}
func (*EnterpriseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f5ac44994d718c, []int{6}
}

func (m *EnterpriseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseMeta.Unmarshal(m, b)
}
func (m *EnterpriseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseMeta.Marshal(b, m, deterministic)
}
func (m *EnterpriseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseMeta.Merge(m, src)
}
func (m *EnterpriseMeta) XXX_Size() int {
	return xxx_messageInfo_EnterpriseMeta.Size(m)
}
func (m *EnterpriseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseMeta proto.InternalMessageInfo

func (m *EnterpriseMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *EnterpriseMeta) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func init() {
	proto.RegisterType((*RaftIndex)(nil), "common.RaftIndex")
	proto.RegisterType((*TargetDatacenter)(nil), "common.TargetDatacenter")
	proto.RegisterType((*WriteRequest)(nil), "common.WriteRequest")
	proto.RegisterType((*ReadRequest)(nil), "common.ReadRequest")
	proto.RegisterType((*QueryOptions)(nil), "common.QueryOptions")
	proto.RegisterType((*QueryMeta)(nil), "common.QueryMeta")
	proto.RegisterType((*EnterpriseMeta)(nil), "common.EnterpriseMeta")
}

func init() {
	proto.RegisterFile("proto/pbcommon/common.proto", fileDescriptor_a6f5ac44994d718c)
}

var fileDescriptor_a6f5ac44994d718c = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xb7, 0x2e, 0x4b, 0xae, 0x65, 0x2a, 0x16, 0xa0, 0x30, 0xd0, 0x54, 0x45, 0x13, 0x9a,
	0xa6, 0xa9, 0x11, 0x83, 0x37, 0xc4, 0x43, 0xd7, 0x15, 0x69, 0xa3, 0x61, 0xcc, 0x14, 0x21, 0xf1,
	0xe6, 0x26, 0xd7, 0xd6, 0x22, 0x8d, 0x83, 0xed, 0x6c, 0xed, 0x7f, 0x46, 0xfc, 0x06, 0x14, 0xa7,
	0xed, 0x52, 0xba, 0xad, 0x4f, 0xd1, 0xf7, 0xdd, 0xe7, 0xf3, 0xdd, 0x7d, 0xe7, 0xc0, 0xab, 0x54,
	0x0a, 0x2d, 0xfc, 0x74, 0x10, 0x8a, 0xc9, 0x44, 0x24, 0x7e, 0xf1, 0x69, 0x19, 0x96, 0x58, 0x05,
	0xda, 0x3f, 0x18, 0x09, 0x31, 0x8a, 0xd1, 0x37, 0xec, 0x20, 0x1b, 0xfa, 0x51, 0x26, 0x99, 0xe6,
	0x0b, 0x9d, 0x77, 0x05, 0x0e, 0x65, 0x43, 0x7d, 0x91, 0x44, 0x38, 0x25, 0x4d, 0xa8, 0x75, 0x24,
	0x32, 0x8d, 0x06, 0xba, 0x95, 0x66, 0xe5, 0xa8, 0x4a, 0xcb, 0x54, 0xae, 0x08, 0x44, 0xc4, 0x87,
	0xb3, 0x42, 0xb1, 0x55, 0x28, 0x4a, 0x94, 0x77, 0x0a, 0x8d, 0x3e, 0x93, 0x23, 0xd4, 0xe7, 0x4c,
	0xb3, 0x10, 0x13, 0x8d, 0x92, 0x1c, 0x00, 0xdc, 0x21, 0x93, 0xd6, 0xa1, 0x25, 0xc6, 0x3b, 0x84,
	0xfa, 0x0f, 0xc9, 0x35, 0x52, 0xfc, 0x9d, 0xa1, 0xd2, 0xe4, 0x19, 0xec, 0xf4, 0xc5, 0x2f, 0x4c,
	0xe6, 0xd2, 0x02, 0x78, 0xd7, 0x50, 0xa3, 0xc8, 0xa2, 0x47, 0x45, 0xe4, 0x04, 0x9e, 0xe6, 0x02,
	0x2e, 0xb1, 0x23, 0x12, 0xc5, 0x95, 0xc6, 0x44, 0x9b, 0x32, 0x6d, 0xba, 0x1e, 0xf0, 0xfe, 0x6c,
	0x43, 0xfd, 0x3a, 0x43, 0x39, 0xbb, 0x4a, 0xf3, 0x99, 0xa8, 0x07, 0x92, 0x1e, 0xc2, 0x93, 0x80,
	0x27, 0x46, 0x58, 0xee, 0x7b, 0x95, 0x24, 0x1f, 0xa1, 0x1e, 0xb0, 0xa9, 0x21, 0xfa, 0x7c, 0x82,
	0xee, 0x76, 0xb3, 0x72, 0x54, 0x3b, 0x7d, 0xd9, 0x2a, 0x1c, 0x68, 0x2d, 0x1c, 0x68, 0x9d, 0xcf,
	0x1d, 0xa0, 0x2b, 0xf2, 0x7c, 0x48, 0xed, 0x38, 0x16, 0xb7, 0xdf, 0x34, 0x8b, 0xd1, 0xad, 0x9a,
	0x92, 0x4b, 0xcc, 0xfd, 0x9d, 0xed, 0x3c, 0xd0, 0x19, 0xd9, 0x07, 0xfb, 0xbb, 0xc2, 0x0e, 0x0b,
	0xc7, 0xe8, 0x5a, 0x46, 0xb4, 0xc4, 0xa4, 0x0b, 0x8d, 0x80, 0x4d, 0x4d, 0xd6, 0x45, 0x2d, 0xee,
	0xee, 0xa6, 0x62, 0xd7, 0x8e, 0x90, 0xb7, 0x60, 0x05, 0x6c, 0xda, 0x1e, 0xa1, 0x6b, 0x6f, 0x3a,
	0x3c, 0x17, 0x92, 0x37, 0xb0, 0x17, 0x64, 0x4a, 0x53, 0xbc, 0x61, 0x31, 0x8f, 0x98, 0x46, 0xd7,
	0x31, 0xb5, 0xfd, 0xc7, 0xe6, 0xa3, 0x34, 0x77, 0x5d, 0x0c, 0xbb, 0x52, 0x0a, 0xe9, 0xc2, 0xc6,
	0x51, 0x96, 0xe5, 0xe4, 0x05, 0x58, 0x9f, 0x78, 0x9c, 0xef, 0x5a, 0xcd, 0xd8, 0x38, 0x47, 0xde,
	0xdf, 0x0a, 0x38, 0x66, 0xe0, 0x01, 0x6a, 0x96, 0x7b, 0x5d, 0xde, 0xf3, 0x02, 0x90, 0x0f, 0x50,
	0xeb, 0x31, 0xa5, 0x3b, 0x22, 0xd1, 0x2c, 0x2c, 0x56, 0xe7, 0xd1, 0x9b, 0xcb, 0xea, 0xfc, 0x79,
	0x7c, 0x4e, 0xc4, 0x6d, 0xd2, 0x43, 0x16, 0xa1, 0x34, 0x1b, 0x60, 0xd3, 0x32, 0x45, 0x8e, 0xa1,
	0xb1, 0x74, 0x29, 0x9c, 0xf5, 0xf0, 0x06, 0x63, 0xe3, 0xb5, 0x43, 0xd7, 0x78, 0xf2, 0x1e, 0x9e,
	0x53, 0x54, 0x59, 0xac, 0x55, 0x51, 0x3f, 0x46, 0x67, 0xb3, 0x76, 0xa7, 0xa7, 0x8c, 0x59, 0x36,
	0xbd, 0x3f, 0x78, 0x59, 0xb5, 0x77, 0x1a, 0xd6, 0x65, 0xd5, 0xb6, 0x1a, 0xbb, 0x5e, 0x0f, 0xf6,
	0xba, 0xf9, 0x0b, 0x4b, 0x25, 0x57, 0x68, 0x9a, 0x7e, 0x0d, 0xce, 0x17, 0x36, 0x41, 0x95, 0xb2,
	0x10, 0xe7, 0x4b, 0x7e, 0x47, 0xe4, 0xd1, 0xaf, 0x4c, 0x6a, 0x6e, 0x56, 0x62, 0xab, 0x88, 0x2e,
	0x89, 0xb3, 0x93, 0x9f, 0xc7, 0x23, 0xae, 0xc7, 0xd9, 0xa0, 0x15, 0x8a, 0x89, 0x3f, 0x66, 0x6a,
	0xcc, 0x43, 0x21, 0x53, 0x3f, 0x14, 0x89, 0xca, 0x62, 0x7f, 0xf5, 0x77, 0x34, 0xb0, 0x0c, 0x7e,
	0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x35, 0xe5, 0x62, 0x05, 0xa7, 0x04, 0x00, 0x00,
}
