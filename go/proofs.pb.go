package ics23

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HashOp int32

const (
	// NO_HASH is the default if no data passed. Note this is an illegal argument some places.
	HashOp_NO_HASH   HashOp = 0
	HashOp_SHA256    HashOp = 1
	HashOp_SHA512    HashOp = 2
	HashOp_KECCAK    HashOp = 3
	HashOp_RIPEMD160 HashOp = 4
	HashOp_BITCOIN   HashOp = 5
)

var HashOp_name = map[int32]string{
	0: "NO_HASH",
	1: "SHA256",
	2: "SHA512",
	3: "KECCAK",
	4: "RIPEMD160",
	5: "BITCOIN",
}

var HashOp_value = map[string]int32{
	"NO_HASH":   0,
	"SHA256":    1,
	"SHA512":    2,
	"KECCAK":    3,
	"RIPEMD160": 4,
	"BITCOIN":   5,
}

func (x HashOp) String() string {
	return proto.EnumName(HashOp_name, int32(x))
}

func (HashOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{0}
}

//*
//LengthOp defines how to process the key and value of the LeafOp
//to include length information. After encoding the length with the given
//algorithm, the length will be prepended to the key and value bytes.
//(Each one with it's own encoded length)
type LengthOp int32

const (
	// NO_PREFIX don't include any length info
	LengthOp_NO_PREFIX LengthOp = 0
	// VAR_PROTO uses protobuf (and go-amino) varint encoding of the length
	LengthOp_VAR_PROTO LengthOp = 1
	// VAR_RLP uses rlp int encoding of the length
	LengthOp_VAR_RLP LengthOp = 2
	// FIXED32_BIG uses big-endian encoding of the length as a 32 bit integer
	LengthOp_FIXED32_BIG LengthOp = 3
	// FIXED32_LITTLE uses little-endian encoding of the length as a 32 bit integer
	LengthOp_FIXED32_LITTLE LengthOp = 4
	// FIXED64_BIG uses big-endian encoding of the length as a 64 bit integer
	LengthOp_FIXED64_BIG LengthOp = 5
	// FIXED64_LITTLE uses little-endian encoding of the length as a 64 bit integer
	LengthOp_FIXED64_LITTLE LengthOp = 6
	// REQUIRE_32_BYTES is like NONE, but will fail if the input is not exactly 32 bytes (sha256 output)
	LengthOp_REQUIRE_32_BYTES LengthOp = 7
	// REQUIRE_64_BYTES is like NONE, but will fail if the input is not exactly 64 bytes (sha512 output)
	LengthOp_REQUIRE_64_BYTES LengthOp = 8
)

var LengthOp_name = map[int32]string{
	0: "NO_PREFIX",
	1: "VAR_PROTO",
	2: "VAR_RLP",
	3: "FIXED32_BIG",
	4: "FIXED32_LITTLE",
	5: "FIXED64_BIG",
	6: "FIXED64_LITTLE",
	7: "REQUIRE_32_BYTES",
	8: "REQUIRE_64_BYTES",
}

var LengthOp_value = map[string]int32{
	"NO_PREFIX":        0,
	"VAR_PROTO":        1,
	"VAR_RLP":          2,
	"FIXED32_BIG":      3,
	"FIXED32_LITTLE":   4,
	"FIXED64_BIG":      5,
	"FIXED64_LITTLE":   6,
	"REQUIRE_32_BYTES": 7,
	"REQUIRE_64_BYTES": 8,
}

func (x LengthOp) String() string {
	return proto.EnumName(LengthOp_name, int32(x))
}

func (LengthOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{1}
}

//*
//ExistenceProof takes a key and a value and a set of steps to perform on it.
//The result of peforming all these steps will provide a "root hash", which can
//be compared to the value in a header.
//
//Since it is computationally infeasible to produce a hash collission for any of the used
//cryptographic hash functions, if someone can provide a series of operations to transform
//a given key and value into a root hash that matches some trusted root, these key and values
//must be in the referenced merkle tree.
//
//The only possible issue is maliablity in LeafOp, such as providing extra prefix data,
//which should be controlled by a spec. Eg. with lengthOp as NONE,
//prefix = FOO, key = BAR, value = CHOICE
//and
//prefix = F, key = OOBAR, value = CHOICE
//would produce the same value.
//
//With LengthOp this is tricker but not impossible. Which is why the "leafPrefixEqual" field
//in the ProofSpec is valuable to prevent this mutability. And why all trees should
//length-prefix the data before hashing it.
type ExistenceProof struct {
	Key       []byte     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueHash []byte     `protobuf:"bytes,2,opt,name=vhash,proto3" json:"vhash,omitempty"`
	Leaf      *LeafOp    `protobuf:"bytes,3,opt,name=leaf,proto3" json:"leaf,omitempty"`
	Path      []*InnerOp `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"`
}

func (m *ExistenceProof) Reset()         { *m = ExistenceProof{} }
func (m *ExistenceProof) String() string { return proto.CompactTextString(m) }
func (*ExistenceProof) ProtoMessage()    {}
func (*ExistenceProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{0}
}
func (m *ExistenceProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExistenceProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExistenceProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExistenceProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistenceProof.Merge(m, src)
}
func (m *ExistenceProof) XXX_Size() int {
	return m.Size()
}
func (m *ExistenceProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistenceProof.DiscardUnknown(m)
}

var xxx_messageInfo_ExistenceProof proto.InternalMessageInfo

func (m *ExistenceProof) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ExistenceProof) GetValueHash() []byte {
	if m != nil {
		return m.ValueHash
	}
	return nil
}

func (m *ExistenceProof) GetLeaf() *LeafOp {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *ExistenceProof) GetPath() []*InnerOp {
	if m != nil {
		return m.Path
	}
	return nil
}

//
//NonExistenceProof takes a proof of two neighbors, one left of the desired key,
//one right of the desired key. If both proofs are valid AND they are neighbors,
//then there is no valid proof for the given key.
type NonExistenceProof struct {
	Key   []byte          `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Left  *ExistenceProof `protobuf:"bytes,2,opt,name=left,proto3" json:"left,omitempty"`
	Right *ExistenceProof `protobuf:"bytes,3,opt,name=right,proto3" json:"right,omitempty"`
}

func (m *NonExistenceProof) Reset()         { *m = NonExistenceProof{} }
func (m *NonExistenceProof) String() string { return proto.CompactTextString(m) }
func (*NonExistenceProof) ProtoMessage()    {}
func (*NonExistenceProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{1}
}
func (m *NonExistenceProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonExistenceProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonExistenceProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonExistenceProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonExistenceProof.Merge(m, src)
}
func (m *NonExistenceProof) XXX_Size() int {
	return m.Size()
}
func (m *NonExistenceProof) XXX_DiscardUnknown() {
	xxx_messageInfo_NonExistenceProof.DiscardUnknown(m)
}

var xxx_messageInfo_NonExistenceProof proto.InternalMessageInfo

func (m *NonExistenceProof) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *NonExistenceProof) GetLeft() *ExistenceProof {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *NonExistenceProof) GetRight() *ExistenceProof {
	if m != nil {
		return m.Right
	}
	return nil
}

//
//CommitmentProof is either an ExistenceProof or a NonExistenceProof, or a Batch of such messages
type CommitmentProof struct {
	// Types that are valid to be assigned to Proof:
	//	*CommitmentProof_Exist
	//	*CommitmentProof_Nonexist
	//	*CommitmentProof_Batch
	//	*CommitmentProof_Compressed
	Proof isCommitmentProof_Proof `protobuf_oneof:"proof"`
}

func (m *CommitmentProof) Reset()         { *m = CommitmentProof{} }
func (m *CommitmentProof) String() string { return proto.CompactTextString(m) }
func (*CommitmentProof) ProtoMessage()    {}
func (*CommitmentProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{2}
}
func (m *CommitmentProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitmentProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitmentProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitmentProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitmentProof.Merge(m, src)
}
func (m *CommitmentProof) XXX_Size() int {
	return m.Size()
}
func (m *CommitmentProof) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitmentProof.DiscardUnknown(m)
}

var xxx_messageInfo_CommitmentProof proto.InternalMessageInfo

type isCommitmentProof_Proof interface {
	isCommitmentProof_Proof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CommitmentProof_Exist struct {
	Exist *ExistenceProof `protobuf:"bytes,1,opt,name=exist,proto3,oneof"`
}
type CommitmentProof_Nonexist struct {
	Nonexist *NonExistenceProof `protobuf:"bytes,2,opt,name=nonexist,proto3,oneof"`
}
type CommitmentProof_Batch struct {
	Batch *BatchProof `protobuf:"bytes,3,opt,name=batch,proto3,oneof"`
}
type CommitmentProof_Compressed struct {
	Compressed *CompressedBatchProof `protobuf:"bytes,4,opt,name=compressed,proto3,oneof"`
}

func (*CommitmentProof_Exist) isCommitmentProof_Proof()      {}
func (*CommitmentProof_Nonexist) isCommitmentProof_Proof()   {}
func (*CommitmentProof_Batch) isCommitmentProof_Proof()      {}
func (*CommitmentProof_Compressed) isCommitmentProof_Proof() {}

func (m *CommitmentProof) GetProof() isCommitmentProof_Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *CommitmentProof) GetExist() *ExistenceProof {
	if x, ok := m.GetProof().(*CommitmentProof_Exist); ok {
		return x.Exist
	}
	return nil
}

func (m *CommitmentProof) GetNonexist() *NonExistenceProof {
	if x, ok := m.GetProof().(*CommitmentProof_Nonexist); ok {
		return x.Nonexist
	}
	return nil
}

func (m *CommitmentProof) GetBatch() *BatchProof {
	if x, ok := m.GetProof().(*CommitmentProof_Batch); ok {
		return x.Batch
	}
	return nil
}

func (m *CommitmentProof) GetCompressed() *CompressedBatchProof {
	if x, ok := m.GetProof().(*CommitmentProof_Compressed); ok {
		return x.Compressed
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CommitmentProof) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CommitmentProof_OneofMarshaler, _CommitmentProof_OneofUnmarshaler, _CommitmentProof_OneofSizer, []interface{}{
		(*CommitmentProof_Exist)(nil),
		(*CommitmentProof_Nonexist)(nil),
		(*CommitmentProof_Batch)(nil),
		(*CommitmentProof_Compressed)(nil),
	}
}

func _CommitmentProof_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CommitmentProof)
	// proof
	switch x := m.Proof.(type) {
	case *CommitmentProof_Exist:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Exist); err != nil {
			return err
		}
	case *CommitmentProof_Nonexist:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nonexist); err != nil {
			return err
		}
	case *CommitmentProof_Batch:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Batch); err != nil {
			return err
		}
	case *CommitmentProof_Compressed:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Compressed); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CommitmentProof.Proof has unexpected type %T", x)
	}
	return nil
}

func _CommitmentProof_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CommitmentProof)
	switch tag {
	case 1: // proof.exist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CommitmentProof_Exist{msg}
		return true, err
	case 2: // proof.nonexist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CommitmentProof_Nonexist{msg}
		return true, err
	case 3: // proof.batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BatchProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CommitmentProof_Batch{msg}
		return true, err
	case 4: // proof.compressed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompressedBatchProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CommitmentProof_Compressed{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CommitmentProof_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CommitmentProof)
	// proof
	switch x := m.Proof.(type) {
	case *CommitmentProof_Exist:
		s := proto.Size(x.Exist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CommitmentProof_Nonexist:
		s := proto.Size(x.Nonexist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CommitmentProof_Batch:
		s := proto.Size(x.Batch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CommitmentProof_Compressed:
		s := proto.Size(x.Compressed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//*
//LeafOp represents the raw key-value data we wish to prove, and
//must be flexible to represent the internal transformation from
//the original key-value pairs into the basis hash, for many existing
//merkle trees.
//
//key and valuehash are passed in. So that the signature of this operation is:
//leafOp(key, valuehash) -> output
//
//To process this, first prehash the keys if needed (ANY means no hash in this case):
//hkey = prehashKey(key)
//
//Then combine the bytes, and hash it
//output = hash(prefix || length(hkey) || hkey || valuehash)
type LeafOp struct {
	Hash       HashOp   `protobuf:"varint,1,opt,name=hash,proto3,enum=ics23.HashOp" json:"hash,omitempty"`
	PrehashKey HashOp   `protobuf:"varint,2,opt,name=prehash_key,json=prehashKey,proto3,enum=ics23.HashOp" json:"prehash_key,omitempty"`
	Length     LengthOp `protobuf:"varint,3,opt,name=length,proto3,enum=ics23.LengthOp" json:"length,omitempty"`
	// prefix is a fixed bytes that may optionally be included at the beginning to differentiate
	// a leaf node from an inner node.
	Prefix []byte `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *LeafOp) Reset()         { *m = LeafOp{} }
func (m *LeafOp) String() string { return proto.CompactTextString(m) }
func (*LeafOp) ProtoMessage()    {}
func (*LeafOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{3}
}
func (m *LeafOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeafOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeafOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeafOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafOp.Merge(m, src)
}
func (m *LeafOp) XXX_Size() int {
	return m.Size()
}
func (m *LeafOp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafOp.DiscardUnknown(m)
}

var xxx_messageInfo_LeafOp proto.InternalMessageInfo

func (m *LeafOp) GetHash() HashOp {
	if m != nil {
		return m.Hash
	}
	return HashOp_NO_HASH
}

func (m *LeafOp) GetPrehashKey() HashOp {
	if m != nil {
		return m.PrehashKey
	}
	return HashOp_NO_HASH
}

func (m *LeafOp) GetLength() LengthOp {
	if m != nil {
		return m.Length
	}
	return LengthOp_NO_PREFIX
}

func (m *LeafOp) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

//*
//InnerOp represents a merkle-proof step that is not a leaf.
//It represents concatenating two children and hashing them to provide the next result.
//
//The result of the previous step is passed in, so the signature of this op is:
//innerOp(child) -> output
//
//The result of applying InnerOp should be:
//output = op.hash(op.prefix || child || op.suffix)
//
//where the || operator is concatenation of binary data,
//and child is the result of hashing all the tree below this step.
//
//Any special data, like prepending child with the length, or prepending the entire operation with
//some value to differentiate from leaf nodes, should be included in prefix and suffix.
//If either of prefix or suffix is empty, we just treat it as an empty string
type InnerOp struct {
	Hash   HashOp `protobuf:"varint,1,opt,name=hash,proto3,enum=ics23.HashOp" json:"hash,omitempty"`
	Prefix []byte `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix []byte `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (m *InnerOp) Reset()         { *m = InnerOp{} }
func (m *InnerOp) String() string { return proto.CompactTextString(m) }
func (*InnerOp) ProtoMessage()    {}
func (*InnerOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{4}
}
func (m *InnerOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InnerOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InnerOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InnerOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerOp.Merge(m, src)
}
func (m *InnerOp) XXX_Size() int {
	return m.Size()
}
func (m *InnerOp) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerOp.DiscardUnknown(m)
}

var xxx_messageInfo_InnerOp proto.InternalMessageInfo

func (m *InnerOp) GetHash() HashOp {
	if m != nil {
		return m.Hash
	}
	return HashOp_NO_HASH
}

func (m *InnerOp) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *InnerOp) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

//*
//ProofSpec defines what the expected parameters are for a given proof type.
//This can be stored in the client and used to validate any incoming proofs.
//
//verify(ProofSpec, Proof) -> Proof | Error
//
//As demonstrated in tests, if we don't fix the algorithm used to calculate the
//LeafHash for a given tree, there are many possible key-value pairs that can
//generate a given hash (by interpretting the preimage differently).
//We need this for proper security, requires client knows a priori what
//tree format server uses. But not in code, rather a configuration object.
type ProofSpec struct {
	// any field in the ExistenceProof must be the same as in this spec.
	// except Prefix, which is just the first bytes of prefix (spec can be longer)
	LeafSpec  *LeafOp    `protobuf:"bytes,1,opt,name=leaf_spec,json=leafSpec,proto3" json:"leaf_spec,omitempty"`
	InnerSpec *InnerSpec `protobuf:"bytes,2,opt,name=inner_spec,json=innerSpec,proto3" json:"inner_spec,omitempty"`
	// max_depth (if > 0) is the maximum number of InnerOps allowed (mainly for fixed-depth tries)
	MaxDepth int32 `protobuf:"varint,3,opt,name=max_depth,json=maxDepth,proto3" json:"max_depth,omitempty"`
	// min_depth (if > 0) is the minimum number of InnerOps allowed (mainly for fixed-depth tries)
	MinDepth int32 `protobuf:"varint,4,opt,name=min_depth,json=minDepth,proto3" json:"min_depth,omitempty"`
}

func (m *ProofSpec) Reset()         { *m = ProofSpec{} }
func (m *ProofSpec) String() string { return proto.CompactTextString(m) }
func (*ProofSpec) ProtoMessage()    {}
func (*ProofSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{5}
}
func (m *ProofSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofSpec.Merge(m, src)
}
func (m *ProofSpec) XXX_Size() int {
	return m.Size()
}
func (m *ProofSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProofSpec proto.InternalMessageInfo

func (m *ProofSpec) GetLeafSpec() *LeafOp {
	if m != nil {
		return m.LeafSpec
	}
	return nil
}

func (m *ProofSpec) GetInnerSpec() *InnerSpec {
	if m != nil {
		return m.InnerSpec
	}
	return nil
}

func (m *ProofSpec) GetMaxDepth() int32 {
	if m != nil {
		return m.MaxDepth
	}
	return 0
}

func (m *ProofSpec) GetMinDepth() int32 {
	if m != nil {
		return m.MinDepth
	}
	return 0
}

//
//InnerSpec contains all store-specific structure info to determine if two proofs from a
//given store are neighbors.
//
//This enables:
//
//isLeftMost(spec: InnerSpec, op: InnerOp)
//isRightMost(spec: InnerSpec, op: InnerOp)
//isLeftNeighbor(spec: InnerSpec, left: InnerOp, right: InnerOp)
type InnerSpec struct {
	// Child order is the ordering of the children node, must count from 0
	// iavl tree is [0, 1] (left then right)
	// merk is [0, 2, 1] (left, right, here)
	ChildOrder      []int32 `protobuf:"varint,1,rep,packed,name=child_order,json=childOrder,proto3" json:"child_order,omitempty"`
	ChildSize       int32   `protobuf:"varint,2,opt,name=child_size,json=childSize,proto3" json:"child_size,omitempty"`
	MinPrefixLength int32   `protobuf:"varint,3,opt,name=min_prefix_length,json=minPrefixLength,proto3" json:"min_prefix_length,omitempty"`
	MaxPrefixLength int32   `protobuf:"varint,4,opt,name=max_prefix_length,json=maxPrefixLength,proto3" json:"max_prefix_length,omitempty"`
	// empty child is the prehash image that is used when one child is nil (eg. 20 bytes of 0)
	EmptyChild []byte `protobuf:"bytes,5,opt,name=empty_child,json=emptyChild,proto3" json:"empty_child,omitempty"`
	// hash is the algorithm that must be used for each InnerOp
	Hash HashOp `protobuf:"varint,6,opt,name=hash,proto3,enum=ics23.HashOp" json:"hash,omitempty"`
}

func (m *InnerSpec) Reset()         { *m = InnerSpec{} }
func (m *InnerSpec) String() string { return proto.CompactTextString(m) }
func (*InnerSpec) ProtoMessage()    {}
func (*InnerSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{6}
}
func (m *InnerSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InnerSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InnerSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InnerSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerSpec.Merge(m, src)
}
func (m *InnerSpec) XXX_Size() int {
	return m.Size()
}
func (m *InnerSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerSpec.DiscardUnknown(m)
}

var xxx_messageInfo_InnerSpec proto.InternalMessageInfo

func (m *InnerSpec) GetChildOrder() []int32 {
	if m != nil {
		return m.ChildOrder
	}
	return nil
}

func (m *InnerSpec) GetChildSize() int32 {
	if m != nil {
		return m.ChildSize
	}
	return 0
}

func (m *InnerSpec) GetMinPrefixLength() int32 {
	if m != nil {
		return m.MinPrefixLength
	}
	return 0
}

func (m *InnerSpec) GetMaxPrefixLength() int32 {
	if m != nil {
		return m.MaxPrefixLength
	}
	return 0
}

func (m *InnerSpec) GetEmptyChild() []byte {
	if m != nil {
		return m.EmptyChild
	}
	return nil
}

func (m *InnerSpec) GetHash() HashOp {
	if m != nil {
		return m.Hash
	}
	return HashOp_NO_HASH
}

//
//BatchProof is a group of multiple proof types than can be compressed
type BatchProof struct {
	Entries []*BatchEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (m *BatchProof) Reset()         { *m = BatchProof{} }
func (m *BatchProof) String() string { return proto.CompactTextString(m) }
func (*BatchProof) ProtoMessage()    {}
func (*BatchProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{7}
}
func (m *BatchProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchProof.Merge(m, src)
}
func (m *BatchProof) XXX_Size() int {
	return m.Size()
}
func (m *BatchProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchProof.DiscardUnknown(m)
}

var xxx_messageInfo_BatchProof proto.InternalMessageInfo

func (m *BatchProof) GetEntries() []*BatchEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Use BatchEntry not CommitmentProof, to avoid recursion
type BatchEntry struct {
	// Types that are valid to be assigned to Proof:
	//	*BatchEntry_Exist
	//	*BatchEntry_Nonexist
	Proof isBatchEntry_Proof `protobuf_oneof:"proof"`
}

func (m *BatchEntry) Reset()         { *m = BatchEntry{} }
func (m *BatchEntry) String() string { return proto.CompactTextString(m) }
func (*BatchEntry) ProtoMessage()    {}
func (*BatchEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{8}
}
func (m *BatchEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchEntry.Merge(m, src)
}
func (m *BatchEntry) XXX_Size() int {
	return m.Size()
}
func (m *BatchEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BatchEntry proto.InternalMessageInfo

type isBatchEntry_Proof interface {
	isBatchEntry_Proof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BatchEntry_Exist struct {
	Exist *ExistenceProof `protobuf:"bytes,1,opt,name=exist,proto3,oneof"`
}
type BatchEntry_Nonexist struct {
	Nonexist *NonExistenceProof `protobuf:"bytes,2,opt,name=nonexist,proto3,oneof"`
}

func (*BatchEntry_Exist) isBatchEntry_Proof()    {}
func (*BatchEntry_Nonexist) isBatchEntry_Proof() {}

func (m *BatchEntry) GetProof() isBatchEntry_Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *BatchEntry) GetExist() *ExistenceProof {
	if x, ok := m.GetProof().(*BatchEntry_Exist); ok {
		return x.Exist
	}
	return nil
}

func (m *BatchEntry) GetNonexist() *NonExistenceProof {
	if x, ok := m.GetProof().(*BatchEntry_Nonexist); ok {
		return x.Nonexist
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BatchEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BatchEntry_OneofMarshaler, _BatchEntry_OneofUnmarshaler, _BatchEntry_OneofSizer, []interface{}{
		(*BatchEntry_Exist)(nil),
		(*BatchEntry_Nonexist)(nil),
	}
}

func _BatchEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BatchEntry)
	// proof
	switch x := m.Proof.(type) {
	case *BatchEntry_Exist:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Exist); err != nil {
			return err
		}
	case *BatchEntry_Nonexist:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nonexist); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BatchEntry.Proof has unexpected type %T", x)
	}
	return nil
}

func _BatchEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BatchEntry)
	switch tag {
	case 1: // proof.exist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &BatchEntry_Exist{msg}
		return true, err
	case 2: // proof.nonexist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &BatchEntry_Nonexist{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BatchEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BatchEntry)
	// proof
	switch x := m.Proof.(type) {
	case *BatchEntry_Exist:
		s := proto.Size(x.Exist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchEntry_Nonexist:
		s := proto.Size(x.Nonexist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CompressedBatchProof struct {
	Entries      []*CompressedBatchEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	LookupInners []*InnerOp              `protobuf:"bytes,2,rep,name=lookup_inners,json=lookupInners,proto3" json:"lookup_inners,omitempty"`
}

func (m *CompressedBatchProof) Reset()         { *m = CompressedBatchProof{} }
func (m *CompressedBatchProof) String() string { return proto.CompactTextString(m) }
func (*CompressedBatchProof) ProtoMessage()    {}
func (*CompressedBatchProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{9}
}
func (m *CompressedBatchProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompressedBatchProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompressedBatchProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompressedBatchProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedBatchProof.Merge(m, src)
}
func (m *CompressedBatchProof) XXX_Size() int {
	return m.Size()
}
func (m *CompressedBatchProof) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedBatchProof.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedBatchProof proto.InternalMessageInfo

func (m *CompressedBatchProof) GetEntries() []*CompressedBatchEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *CompressedBatchProof) GetLookupInners() []*InnerOp {
	if m != nil {
		return m.LookupInners
	}
	return nil
}

// Use BatchEntry not CommitmentProof, to avoid recursion
type CompressedBatchEntry struct {
	// Types that are valid to be assigned to Proof:
	//	*CompressedBatchEntry_Exist
	//	*CompressedBatchEntry_Nonexist
	Proof isCompressedBatchEntry_Proof `protobuf_oneof:"proof"`
}

func (m *CompressedBatchEntry) Reset()         { *m = CompressedBatchEntry{} }
func (m *CompressedBatchEntry) String() string { return proto.CompactTextString(m) }
func (*CompressedBatchEntry) ProtoMessage()    {}
func (*CompressedBatchEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{10}
}
func (m *CompressedBatchEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompressedBatchEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompressedBatchEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompressedBatchEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedBatchEntry.Merge(m, src)
}
func (m *CompressedBatchEntry) XXX_Size() int {
	return m.Size()
}
func (m *CompressedBatchEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedBatchEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedBatchEntry proto.InternalMessageInfo

type isCompressedBatchEntry_Proof interface {
	isCompressedBatchEntry_Proof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CompressedBatchEntry_Exist struct {
	Exist *CompressedExistenceProof `protobuf:"bytes,1,opt,name=exist,proto3,oneof"`
}
type CompressedBatchEntry_Nonexist struct {
	Nonexist *CompressedNonExistenceProof `protobuf:"bytes,2,opt,name=nonexist,proto3,oneof"`
}

func (*CompressedBatchEntry_Exist) isCompressedBatchEntry_Proof()    {}
func (*CompressedBatchEntry_Nonexist) isCompressedBatchEntry_Proof() {}

func (m *CompressedBatchEntry) GetProof() isCompressedBatchEntry_Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *CompressedBatchEntry) GetExist() *CompressedExistenceProof {
	if x, ok := m.GetProof().(*CompressedBatchEntry_Exist); ok {
		return x.Exist
	}
	return nil
}

func (m *CompressedBatchEntry) GetNonexist() *CompressedNonExistenceProof {
	if x, ok := m.GetProof().(*CompressedBatchEntry_Nonexist); ok {
		return x.Nonexist
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CompressedBatchEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CompressedBatchEntry_OneofMarshaler, _CompressedBatchEntry_OneofUnmarshaler, _CompressedBatchEntry_OneofSizer, []interface{}{
		(*CompressedBatchEntry_Exist)(nil),
		(*CompressedBatchEntry_Nonexist)(nil),
	}
}

func _CompressedBatchEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CompressedBatchEntry)
	// proof
	switch x := m.Proof.(type) {
	case *CompressedBatchEntry_Exist:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Exist); err != nil {
			return err
		}
	case *CompressedBatchEntry_Nonexist:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nonexist); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CompressedBatchEntry.Proof has unexpected type %T", x)
	}
	return nil
}

func _CompressedBatchEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CompressedBatchEntry)
	switch tag {
	case 1: // proof.exist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompressedExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CompressedBatchEntry_Exist{msg}
		return true, err
	case 2: // proof.nonexist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompressedNonExistenceProof)
		err := b.DecodeMessage(msg)
		m.Proof = &CompressedBatchEntry_Nonexist{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CompressedBatchEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CompressedBatchEntry)
	// proof
	switch x := m.Proof.(type) {
	case *CompressedBatchEntry_Exist:
		s := proto.Size(x.Exist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CompressedBatchEntry_Nonexist:
		s := proto.Size(x.Nonexist)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CompressedExistenceProof struct {
	Key       []byte  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueHash []byte  `protobuf:"bytes,2,opt,name=vhash,proto3" json:"value,omitempty"`
	Leaf      *LeafOp `protobuf:"bytes,3,opt,name=leaf,proto3" json:"leaf,omitempty"`
	// these are indexes into the lookup_inners table in CompressedBatchProof
	Path []int32 `protobuf:"varint,4,rep,packed,name=path,proto3" json:"path,omitempty"`
}

func (m *CompressedExistenceProof) Reset()         { *m = CompressedExistenceProof{} }
func (m *CompressedExistenceProof) String() string { return proto.CompactTextString(m) }
func (*CompressedExistenceProof) ProtoMessage()    {}
func (*CompressedExistenceProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{11}
}
func (m *CompressedExistenceProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompressedExistenceProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompressedExistenceProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompressedExistenceProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedExistenceProof.Merge(m, src)
}
func (m *CompressedExistenceProof) XXX_Size() int {
	return m.Size()
}
func (m *CompressedExistenceProof) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedExistenceProof.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedExistenceProof proto.InternalMessageInfo

func (m *CompressedExistenceProof) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CompressedExistenceProof) GetValueHash() []byte {
	if m != nil {
		return m.ValueHash
	}
	return nil
}

func (m *CompressedExistenceProof) GetLeaf() *LeafOp {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *CompressedExistenceProof) GetPath() []int32 {
	if m != nil {
		return m.Path
	}
	return nil
}

type CompressedNonExistenceProof struct {
	Key   []byte                    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Left  *CompressedExistenceProof `protobuf:"bytes,2,opt,name=left,proto3" json:"left,omitempty"`
	Right *CompressedExistenceProof `protobuf:"bytes,3,opt,name=right,proto3" json:"right,omitempty"`
}

func (m *CompressedNonExistenceProof) Reset()         { *m = CompressedNonExistenceProof{} }
func (m *CompressedNonExistenceProof) String() string { return proto.CompactTextString(m) }
func (*CompressedNonExistenceProof) ProtoMessage()    {}
func (*CompressedNonExistenceProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_855156e15e7b8e99, []int{12}
}
func (m *CompressedNonExistenceProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompressedNonExistenceProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompressedNonExistenceProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompressedNonExistenceProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedNonExistenceProof.Merge(m, src)
}
func (m *CompressedNonExistenceProof) XXX_Size() int {
	return m.Size()
}
func (m *CompressedNonExistenceProof) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedNonExistenceProof.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedNonExistenceProof proto.InternalMessageInfo

func (m *CompressedNonExistenceProof) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CompressedNonExistenceProof) GetLeft() *CompressedExistenceProof {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *CompressedNonExistenceProof) GetRight() *CompressedExistenceProof {
	if m != nil {
		return m.Right
	}
	return nil
}

func init() {
	proto.RegisterEnum("ics23.HashOp", HashOp_name, HashOp_value)
	proto.RegisterEnum("ics23.LengthOp", LengthOp_name, LengthOp_value)
	proto.RegisterType((*ExistenceProof)(nil), "ics23.ExistenceProof")
	proto.RegisterType((*NonExistenceProof)(nil), "ics23.NonExistenceProof")
	proto.RegisterType((*CommitmentProof)(nil), "ics23.CommitmentProof")
	proto.RegisterType((*LeafOp)(nil), "ics23.LeafOp")
	proto.RegisterType((*InnerOp)(nil), "ics23.InnerOp")
	proto.RegisterType((*ProofSpec)(nil), "ics23.ProofSpec")
	proto.RegisterType((*InnerSpec)(nil), "ics23.InnerSpec")
	proto.RegisterType((*BatchProof)(nil), "ics23.BatchProof")
	proto.RegisterType((*BatchEntry)(nil), "ics23.BatchEntry")
	proto.RegisterType((*CompressedBatchProof)(nil), "ics23.CompressedBatchProof")
	proto.RegisterType((*CompressedBatchEntry)(nil), "ics23.CompressedBatchEntry")
	proto.RegisterType((*CompressedExistenceProof)(nil), "ics23.CompressedExistenceProof")
	proto.RegisterType((*CompressedNonExistenceProof)(nil), "ics23.CompressedNonExistenceProof")
}

func init() { proto.RegisterFile("proofs.proto", fileDescriptor_855156e15e7b8e99) }

var fileDescriptor_855156e15e7b8e99 = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0xe3, 0x24, 0xce, 0x8f, 0x97, 0xfe, 0x70, 0x47, 0x05, 0x59, 0xaa, 0x48, 0x8b, 0x2f,
	0x74, 0xbb, 0xa2, 0xb0, 0xc9, 0x36, 0x88, 0x03, 0x12, 0x4d, 0xea, 0x25, 0x56, 0x4b, 0x13, 0x26,
	0xd9, 0xd5, 0x22, 0x21, 0x59, 0xde, 0x74, 0xb2, 0xb1, 0x36, 0xb1, 0x2d, 0xdb, 0x45, 0xc9, 0x0a,
	0x84, 0xc4, 0x5f, 0xc0, 0x19, 0x6e, 0x5c, 0xf9, 0x47, 0x38, 0xf6, 0xc8, 0x11, 0xb5, 0x07, 0x0e,
	0xfc, 0x13, 0xe8, 0xcd, 0x4c, 0xdc, 0xa4, 0x49, 0x68, 0x0f, 0x88, 0x9b, 0xe7, 0xfb, 0x3e, 0x6f,
	0xe6, 0xcd, 0xfb, 0x91, 0x09, 0xac, 0x05, 0xa1, 0xef, 0xf7, 0xa3, 0xc3, 0x20, 0xf4, 0x63, 0x9f,
	0xa8, 0x6e, 0x2f, 0xaa, 0x54, 0x8d, 0x1f, 0x60, 0xc3, 0x1c, 0xbb, 0x51, 0xcc, 0xbc, 0x1e, 0x6b,
	0xa3, 0x9d, 0x68, 0x90, 0x79, 0xc3, 0x26, 0xba, 0xb2, 0xa7, 0xec, 0xaf, 0x51, 0xfc, 0x24, 0xdb,
	0xa0, 0x7e, 0xeb, 0x0c, 0x2f, 0x99, 0x9e, 0xe6, 0x9a, 0x58, 0x90, 0xf7, 0x21, 0x3b, 0x64, 0x4e,
	0x5f, 0xcf, 0xec, 0x29, 0xfb, 0xa5, 0xca, 0xfa, 0x21, 0xdf, 0xef, 0xf0, 0x8c, 0x39, 0xfd, 0x56,
	0x40, 0xb9, 0x89, 0x18, 0x90, 0x0d, 0x9c, 0x78, 0xa0, 0x67, 0xf7, 0x32, 0xfb, 0xa5, 0xca, 0x86,
	0x44, 0x2c, 0xcf, 0x63, 0x21, 0x32, 0x68, 0x33, 0xbe, 0x87, 0xad, 0x73, 0xdf, 0xbb, 0x37, 0x86,
	0x47, 0x78, 0x5a, 0x3f, 0xe6, 0x21, 0x94, 0x2a, 0xef, 0xc8, 0xad, 0xe6, 0xdd, 0x28, 0x47, 0xc8,
	0x63, 0x50, 0x43, 0xf7, 0xf5, 0x20, 0x96, 0x91, 0xad, 0x60, 0x05, 0x63, 0xfc, 0xad, 0xc0, 0x66,
	0xc3, 0x1f, 0x8d, 0xdc, 0x78, 0xc4, 0xbc, 0x58, 0x9c, 0xfe, 0x21, 0xa8, 0x0c, 0x61, 0x7e, 0xfe,
	0xaa, 0x0d, 0x9a, 0x29, 0x2a, 0x28, 0x52, 0x83, 0x82, 0xe7, 0x7b, 0xc2, 0x43, 0x84, 0xa7, 0x4b,
	0x8f, 0x85, 0x8b, 0x35, 0x53, 0x34, 0x61, 0xc9, 0x23, 0x50, 0x5f, 0x39, 0x71, 0x6f, 0x20, 0xe3,
	0xdc, 0x92, 0x4e, 0x75, 0xd4, 0x92, 0x23, 0x38, 0x41, 0x3e, 0x03, 0xe8, 0xf9, 0xa3, 0x20, 0x64,
	0x51, 0xc4, 0x2e, 0xf4, 0x2c, 0xe7, 0x77, 0x24, 0xdf, 0x48, 0x0c, 0x73, 0x9e, 0x33, 0x0e, 0xf5,
	0x3c, 0xa8, 0xbc, 0xf6, 0xc6, 0x95, 0x02, 0x39, 0x51, 0x21, 0x2c, 0xdf, 0xc0, 0x89, 0x06, 0xfc,
	0x8e, 0x1b, 0x49, 0xf9, 0x9a, 0x4e, 0x34, 0xc0, 0xd2, 0xa0, 0x89, 0x1c, 0x42, 0x29, 0x08, 0x19,
	0x7e, 0xda, 0x58, 0x8d, 0xf4, 0x32, 0x12, 0x24, 0x71, 0xca, 0x26, 0xa4, 0x02, 0xeb, 0x53, 0x5e,
	0xf4, 0x4b, 0x66, 0x99, 0xc7, 0x9a, 0x64, 0x5e, 0xf0, 0x2e, 0xfa, 0x00, 0x72, 0x43, 0xe6, 0xbd,
	0xe6, 0x4d, 0x82, 0xf0, 0x66, 0xd2, 0x47, 0x28, 0xb6, 0x02, 0x2a, 0xcd, 0xe4, 0x5d, 0xc8, 0x05,
	0x21, 0xeb, 0xbb, 0x63, 0x5d, 0xe5, 0x5d, 0x21, 0x57, 0xc6, 0x37, 0x90, 0x97, 0x0d, 0xf5, 0x90,
	0x2b, 0xdd, 0xee, 0x92, 0x9e, 0xdd, 0x05, 0xf5, 0xe8, 0xb2, 0x8f, 0x7a, 0x46, 0xe8, 0x62, 0x65,
	0xfc, 0xaa, 0x40, 0x91, 0x67, 0xb4, 0x13, 0xb0, 0x1e, 0x39, 0x80, 0x22, 0xf6, 0xb5, 0x1d, 0x05,
	0xac, 0x27, 0x9b, 0xe3, 0x4e, 0xdf, 0x17, 0xd0, 0xce, 0xd9, 0x8f, 0x00, 0x5c, 0x8c, 0x4b, 0xc0,
	0xa2, 0x2f, 0xb4, 0xd9, 0x09, 0x40, 0x8a, 0x16, 0xdd, 0xe9, 0x27, 0xd9, 0x81, 0xe2, 0xc8, 0x19,
	0xdb, 0x17, 0x2c, 0x88, 0x45, 0x4b, 0xa8, 0xb4, 0x30, 0x72, 0xc6, 0x27, 0xb8, 0xe6, 0x46, 0xd7,
	0x93, 0xc6, 0xac, 0x34, 0xba, 0x1e, 0x37, 0x1a, 0x7f, 0x29, 0x50, 0x4c, 0xb6, 0x24, 0xbb, 0x50,
	0xea, 0x0d, 0xdc, 0xe1, 0x85, 0xed, 0x87, 0x17, 0x2c, 0xd4, 0x95, 0xbd, 0xcc, 0xbe, 0x4a, 0x81,
	0x4b, 0x2d, 0x54, 0xc8, 0x7b, 0x20, 0x56, 0x76, 0xe4, 0xbe, 0x15, 0x33, 0xad, 0xd2, 0x22, 0x57,
	0x3a, 0xee, 0x5b, 0x46, 0x0e, 0x60, 0x0b, 0x8f, 0x12, 0x89, 0xb1, 0x65, 0x71, 0x44, 0x3c, 0x9b,
	0x23, 0xd7, 0x6b, 0x73, 0x5d, 0x94, 0x87, 0xb3, 0xce, 0xf8, 0x0e, 0x9b, 0x95, 0xac, 0x33, 0x9e,
	0x63, 0x77, 0xa1, 0xc4, 0x46, 0x41, 0x3c, 0xb1, 0xf9, 0x51, 0xb2, 0x8a, 0xc0, 0xa5, 0x06, 0x2a,
	0x49, 0xf9, 0x72, 0x2b, 0xcb, 0x67, 0x7c, 0x0a, 0x70, 0xdb, 0xe4, 0xe4, 0x31, 0xe4, 0x99, 0x17,
	0x87, 0x2e, 0x8b, 0xf8, 0x2d, 0xef, 0x8c, 0x90, 0xe9, 0xc5, 0xe1, 0x84, 0x4e, 0x09, 0xe3, 0x3b,
	0xe9, 0xca, 0xe5, 0xff, 0x69, 0xc4, 0x6f, 0x07, 0xef, 0x47, 0x05, 0xb6, 0x97, 0x0d, 0x2a, 0x39,
	0xba, 0x7b, 0x87, 0x15, 0x63, 0x3d, 0x7f, 0x1b, 0x52, 0x85, 0xf5, 0xa1, 0xef, 0xbf, 0xb9, 0x0c,
	0x6c, 0xde, 0x40, 0x91, 0x9e, 0x5e, 0xfa, 0x13, 0xbb, 0x26, 0x20, 0xbe, 0x8c, 0x8c, 0x9f, 0x17,
	0x83, 0x10, 0xd9, 0xf8, 0x64, 0x3e, 0x1b, 0xbb, 0x0b, 0x21, 0xac, 0xca, 0xcb, 0xe7, 0x0b, 0x79,
	0x31, 0x16, 0x7c, 0x1f, 0x98, 0xa1, 0x09, 0xe8, 0xab, 0xce, 0xfb, 0x2f, 0x9f, 0x24, 0x32, 0xf3,
	0x24, 0xa9, 0xf2, 0x09, 0xfa, 0x45, 0x81, 0x9d, 0x7f, 0x89, 0x77, 0xc9, 0xf1, 0xd5, 0xb9, 0xd7,
	0xe8, 0xbe, 0x7c, 0xc9, 0x77, 0xe9, 0x68, 0xfe, 0x5d, 0xba, 0xd7, 0x4b, 0xd0, 0x07, 0xcf, 0x21,
	0x27, 0x66, 0x80, 0x94, 0x20, 0x7f, 0xde, 0xb2, 0x9b, 0xc7, 0x9d, 0xa6, 0x96, 0x22, 0x00, 0xb9,
	0x4e, 0xf3, 0xb8, 0x72, 0x54, 0xd3, 0x14, 0xf9, 0x7d, 0xf4, 0xa4, 0xa2, 0xa5, 0xf1, 0xfb, 0xd4,
	0x6c, 0x34, 0x8e, 0x4f, 0xb5, 0x0c, 0x59, 0x87, 0x22, 0xb5, 0xda, 0xe6, 0x97, 0x27, 0x4f, 0x6a,
	0x1f, 0x6b, 0x59, 0xf4, 0xaf, 0x5b, 0xdd, 0x46, 0xcb, 0x3a, 0xd7, 0xd4, 0x83, 0xdf, 0x14, 0x28,
	0x4c, 0x7f, 0x64, 0x11, 0x3c, 0x6f, 0xd9, 0x6d, 0x6a, 0x3e, 0xb3, 0x5e, 0x6a, 0x29, 0x5c, 0xbe,
	0x38, 0xa6, 0x76, 0x9b, 0xb6, 0xba, 0x2d, 0x4d, 0x41, 0x3f, 0x5c, 0xd2, 0xb3, 0xb6, 0x96, 0x26,
	0x9b, 0x50, 0x7a, 0x66, 0xbd, 0x34, 0x4f, 0xaa, 0x15, 0xbb, 0x6e, 0x7d, 0xa1, 0x65, 0x08, 0x81,
	0x8d, 0xa9, 0x70, 0x66, 0x75, 0xbb, 0x67, 0xa6, 0x96, 0x4d, 0xa0, 0xda, 0x53, 0x0e, 0xa9, 0x09,
	0x54, 0x7b, 0x3a, 0x85, 0x72, 0x64, 0x1b, 0x34, 0x6a, 0x7e, 0xf5, 0xdc, 0xa2, 0xa6, 0x8d, 0x9b,
	0x7d, 0xdd, 0x35, 0x3b, 0x5a, 0x7e, 0x56, 0x45, 0x6f, 0xae, 0x16, 0xea, 0xfa, 0xef, 0xd7, 0x65,
	0xe5, 0xea, 0xba, 0xac, 0xfc, 0x79, 0x5d, 0x56, 0x7e, 0xba, 0x29, 0xa7, 0xae, 0x6e, 0xca, 0xa9,
	0x3f, 0x6e, 0xca, 0xa9, 0x57, 0x39, 0xfe, 0x77, 0xa6, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x20, 0xff, 0xe9, 0x19, 0xde, 0x08, 0x00, 0x00,
}

func (m *ExistenceProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExistenceProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.ValueHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.ValueHash)))
		i += copy(dAtA[i:], m.ValueHash)
	}
	if m.Leaf != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Leaf.Size()))
		n1, err := m.Leaf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Path) > 0 {
		for _, msg := range m.Path {
			dAtA[i] = 0x22
			i++
			i = encodeVarintProofs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NonExistenceProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonExistenceProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Left != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Left.Size()))
		n2, err := m.Left.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Right != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Right.Size()))
		n3, err := m.Right.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *CommitmentProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitmentProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		nn4, err := m.Proof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *CommitmentProof_Exist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Exist != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Exist.Size()))
		n5, err := m.Exist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *CommitmentProof_Nonexist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Nonexist != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Nonexist.Size()))
		n6, err := m.Nonexist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *CommitmentProof_Batch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Batch != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Batch.Size()))
		n7, err := m.Batch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *CommitmentProof_Compressed) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Compressed != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Compressed.Size()))
		n8, err := m.Compressed.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func (m *LeafOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeafOp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hash != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Hash))
	}
	if m.PrehashKey != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.PrehashKey))
	}
	if m.PrehashValue != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.PrehashValue))
	}
	if m.Length != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Length))
	}
	if len(m.Prefix) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	return i, nil
}

func (m *InnerOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InnerOp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hash != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Hash))
	}
	if len(m.Prefix) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Suffix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Suffix)))
		i += copy(dAtA[i:], m.Suffix)
	}
	return i, nil
}

func (m *ProofSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LeafSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.LeafSpec.Size()))
		n9, err := m.LeafSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.InnerSpec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.InnerSpec.Size()))
		n10, err := m.InnerSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.MaxDepth != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.MaxDepth))
	}
	if m.MinDepth != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.MinDepth))
	}
	return i, nil
}

func (m *InnerSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InnerSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChildOrder) > 0 {
		dAtA12 := make([]byte, len(m.ChildOrder)*10)
		var j11 int
		for _, num1 := range m.ChildOrder {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA12[j11] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j11++
			}
			dAtA12[j11] = uint8(num)
			j11++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(j11))
		i += copy(dAtA[i:], dAtA12[:j11])
	}
	if m.ChildSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.ChildSize))
	}
	if m.MinPrefixLength != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.MinPrefixLength))
	}
	if m.MaxPrefixLength != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.MaxPrefixLength))
	}
	if len(m.EmptyChild) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.EmptyChild)))
		i += copy(dAtA[i:], m.EmptyChild)
	}
	if m.Hash != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Hash))
	}
	return i, nil
}

func (m *BatchProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProofs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *BatchEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		nn13, err := m.Proof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn13
	}
	return i, nil
}

func (m *BatchEntry_Exist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Exist != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Exist.Size()))
		n14, err := m.Exist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	return i, nil
}
func (m *BatchEntry_Nonexist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Nonexist != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Nonexist.Size()))
		n15, err := m.Nonexist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	return i, nil
}
func (m *CompressedBatchProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompressedBatchProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProofs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.LookupInners) > 0 {
		for _, msg := range m.LookupInners {
			dAtA[i] = 0x12
			i++
			i = encodeVarintProofs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CompressedBatchEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompressedBatchEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		nn16, err := m.Proof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn16
	}
	return i, nil
}

func (m *CompressedBatchEntry_Exist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Exist != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Exist.Size()))
		n17, err := m.Exist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	return i, nil
}
func (m *CompressedBatchEntry_Nonexist) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Nonexist != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Nonexist.Size()))
		n18, err := m.Nonexist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	return i, nil
}
func (m *CompressedExistenceProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompressedExistenceProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.ValueHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.ValueHash)))
		i += copy(dAtA[i:], m.ValueHash)
	}
	if m.Leaf != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Leaf.Size()))
		n19, err := m.Leaf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if len(m.Path) > 0 {
		dAtA21 := make([]byte, len(m.Path)*10)
		var j20 int
		for _, num1 := range m.Path {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA21[j20] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j20++
			}
			dAtA21[j20] = uint8(num)
			j20++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintProofs(dAtA, i, uint64(j20))
		i += copy(dAtA[i:], dAtA21[:j20])
	}
	return i, nil
}

func (m *CompressedNonExistenceProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompressedNonExistenceProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProofs(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Left != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Left.Size()))
		n22, err := m.Left.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	if m.Right != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProofs(dAtA, i, uint64(m.Right.Size()))
		n23, err := m.Right.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	return i, nil
}

func encodeVarintProofs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExistenceProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	l = len(m.ValueHash)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Leaf != nil {
		l = m.Leaf.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if len(m.Path) > 0 {
		for _, e := range m.Path {
			l = e.Size()
			n += 1 + l + sovProofs(uint64(l))
		}
	}
	return n
}

func (m *NonExistenceProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Left != nil {
		l = m.Left.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Right != nil {
		l = m.Right.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}

func (m *CommitmentProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		n += m.Proof.Size()
	}
	return n
}

func (m *CommitmentProof_Exist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exist != nil {
		l = m.Exist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CommitmentProof_Nonexist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonexist != nil {
		l = m.Nonexist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CommitmentProof_Batch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Batch != nil {
		l = m.Batch.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CommitmentProof_Compressed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Compressed != nil {
		l = m.Compressed.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *LeafOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hash != 0 {
		n += 1 + sovProofs(uint64(m.Hash))
	}
	if m.PrehashKey != 0 {
		n += 1 + sovProofs(uint64(m.PrehashKey))
	}
	if m.PrehashValue != 0 {
		n += 1 + sovProofs(uint64(m.PrehashValue))
	}
	if m.Length != 0 {
		n += 1 + sovProofs(uint64(m.Length))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}

func (m *InnerOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hash != 0 {
		n += 1 + sovProofs(uint64(m.Hash))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	l = len(m.Suffix)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}

func (m *ProofSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LeafSpec != nil {
		l = m.LeafSpec.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.InnerSpec != nil {
		l = m.InnerSpec.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.MaxDepth != 0 {
		n += 1 + sovProofs(uint64(m.MaxDepth))
	}
	if m.MinDepth != 0 {
		n += 1 + sovProofs(uint64(m.MinDepth))
	}
	return n
}

func (m *InnerSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChildOrder) > 0 {
		l = 0
		for _, e := range m.ChildOrder {
			l += sovProofs(uint64(e))
		}
		n += 1 + sovProofs(uint64(l)) + l
	}
	if m.ChildSize != 0 {
		n += 1 + sovProofs(uint64(m.ChildSize))
	}
	if m.MinPrefixLength != 0 {
		n += 1 + sovProofs(uint64(m.MinPrefixLength))
	}
	if m.MaxPrefixLength != 0 {
		n += 1 + sovProofs(uint64(m.MaxPrefixLength))
	}
	l = len(m.EmptyChild)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovProofs(uint64(m.Hash))
	}
	return n
}

func (m *BatchProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovProofs(uint64(l))
		}
	}
	return n
}

func (m *BatchEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		n += m.Proof.Size()
	}
	return n
}

func (m *BatchEntry_Exist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exist != nil {
		l = m.Exist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *BatchEntry_Nonexist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonexist != nil {
		l = m.Nonexist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CompressedBatchProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovProofs(uint64(l))
		}
	}
	if len(m.LookupInners) > 0 {
		for _, e := range m.LookupInners {
			l = e.Size()
			n += 1 + l + sovProofs(uint64(l))
		}
	}
	return n
}

func (m *CompressedBatchEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		n += m.Proof.Size()
	}
	return n
}

func (m *CompressedBatchEntry_Exist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exist != nil {
		l = m.Exist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CompressedBatchEntry_Nonexist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonexist != nil {
		l = m.Nonexist.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}
func (m *CompressedExistenceProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	l = len(m.ValueHash)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Leaf != nil {
		l = m.Leaf.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if len(m.Path) > 0 {
		l = 0
		for _, e := range m.Path {
			l += sovProofs(uint64(e))
		}
		n += 1 + sovProofs(uint64(l)) + l
	}
	return n
}

func (m *CompressedNonExistenceProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Left != nil {
		l = m.Left.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	if m.Right != nil {
		l = m.Right.Size()
		n += 1 + l + sovProofs(uint64(l))
	}
	return n
}

func sovProofs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProofs(x uint64) (n int) {
	return sovProofs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExistenceProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: ExistenceProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExistenceProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueHash = append(m.ValueHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValueHash == nil {
				m.ValueHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leaf == nil {
				m.Leaf = &LeafOp{}
			}
			if err := m.Leaf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, &InnerOp{})
			if err := m.Path[len(m.Path)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *NonExistenceProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: NonExistenceProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonExistenceProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Left == nil {
				m.Left = &ExistenceProof{}
			}
			if err := m.Left.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Right == nil {
				m.Right = &ExistenceProof{}
			}
			if err := m.Right.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *CommitmentProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: CommitmentProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitmentProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CommitmentProof_Exist{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonexist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NonExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CommitmentProof_Nonexist{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BatchProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CommitmentProof_Batch{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compressed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CompressedBatchProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CommitmentProof_Compressed{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *LeafOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: LeafOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeafOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= HashOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrehashKey", wireType)
			}
			m.PrehashKey = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrehashKey |= HashOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrehashValue", wireType)
			}
			m.PrehashValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrehashValue |= HashOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= LengthOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *InnerOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: InnerOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= HashOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suffix = append(m.Suffix[:0], dAtA[iNdEx:postIndex]...)
			if m.Suffix == nil {
				m.Suffix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *ProofSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: ProofSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeafSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeafSpec == nil {
				m.LeafSpec = &LeafOp{}
			}
			if err := m.LeafSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InnerSpec == nil {
				m.InnerSpec = &InnerSpec{}
			}
			if err := m.InnerSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDepth", wireType)
			}
			m.MaxDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDepth |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDepth", wireType)
			}
			m.MinDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDepth |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *InnerSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: InnerSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProofs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ChildOrder = append(m.ChildOrder, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProofs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProofs
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProofs
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ChildOrder) == 0 {
					m.ChildOrder = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProofs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ChildOrder = append(m.ChildOrder, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildOrder", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildSize", wireType)
			}
			m.ChildSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChildSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrefixLength", wireType)
			}
			m.MinPrefixLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPrefixLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrefixLength", wireType)
			}
			m.MaxPrefixLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPrefixLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyChild", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmptyChild = append(m.EmptyChild[:0], dAtA[iNdEx:postIndex]...)
			if m.EmptyChild == nil {
				m.EmptyChild = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= HashOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *BatchProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: BatchProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &BatchEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *BatchEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: BatchEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &BatchEntry_Exist{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonexist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NonExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &BatchEntry_Nonexist{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *CompressedBatchProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: CompressedBatchProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressedBatchProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &CompressedBatchEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LookupInners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LookupInners = append(m.LookupInners, &InnerOp{})
			if err := m.LookupInners[len(m.LookupInners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *CompressedBatchEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: CompressedBatchEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressedBatchEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CompressedExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CompressedBatchEntry_Exist{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonexist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CompressedNonExistenceProof{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Proof = &CompressedBatchEntry_Nonexist{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *CompressedExistenceProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: CompressedExistenceProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressedExistenceProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueHash = append(m.ValueHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValueHash == nil {
				m.ValueHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leaf == nil {
				m.Leaf = &LeafOp{}
			}
			if err := m.Leaf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProofs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Path = append(m.Path, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProofs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProofs
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProofs
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Path) == 0 {
					m.Path = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProofs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Path = append(m.Path, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func (m *CompressedNonExistenceProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProofs
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
			return fmt.Errorf("proto: CompressedNonExistenceProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressedNonExistenceProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Left == nil {
				m.Left = &CompressedExistenceProof{}
			}
			if err := m.Left.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProofs
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
				return ErrInvalidLengthProofs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProofs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Right == nil {
				m.Right = &CompressedExistenceProof{}
			}
			if err := m.Right.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProofs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProofs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProofs
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
func skipProofs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProofs
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
					return 0, ErrIntOverflowProofs
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
					return 0, ErrIntOverflowProofs
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
				return 0, ErrInvalidLengthProofs
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProofs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProofs
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
				next, err := skipProofs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProofs
				}
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
	ErrInvalidLengthProofs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProofs   = fmt.Errorf("proto: integer overflow")
)
