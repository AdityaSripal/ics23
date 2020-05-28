package ics23

import (
	"bytes"

	"github.com/pkg/errors"
)

// IavlSpec constrains the format from proofs-iavl (iavl merkle proofs)
var IavlSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		MinPrefixLength: 4,
		MaxPrefixLength: 12,
		ChildSize:       33, // (with length byte)
		Hash:            HashOp_SHA256,
	},
}

// TendermintSpec constrains the format from proofs-tendermint (crypto/merkle SimpleProof)
var TendermintSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		MinPrefixLength: 1,
		MaxPrefixLength: 1,
		ChildSize:       32, // (no length byte)
		Hash:            HashOp_SHA256,
	},
}

// Verify does all checks to ensure this proof is a valid existence proof for this root
// NOTE: In order to prove a value, we must call VerifyItem
func (p *ExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot) error {
	if err := p.CheckAgainstSpec(spec); err != nil {
		return err
	}

	calc, err := p.Calculate()
	if err != nil {
		return errors.Wrap(err, "Error calculating root")
	}
	if !bytes.Equal(root, calc) {
		return errors.Errorf("Calculcated root doesn't match provided root")
	}

	return nil

}

// VerifyItem checks that the proof is a valid proof of a (key, value) for this root
func (p *ExistenceProof) VerifyItem(spec *ProofSpec, root CommitmentRoot, key, value []byte) error {

	if !bytes.Equal(key, p.Key) {
		return errors.Errorf("Provided key doesn't match proof")
	}
	vhash, err := prepareLeafData(p.Leaf.PrehashValue, p.Leaf.Length, value)
	if err != nil {
		return errors.Wrap(err, "prehash value")
	}
	if !bytes.Equal(p.ValueHash, vhash) {
		return errors.Errorf("Provided value doesn't match proof")
	}

	return p.Verify(spec, root)
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() (CommitmentRoot, error) {
	if p.GetLeaf() == nil {
		return nil, errors.New("Existence Proof needs defined LeafOp")
	}

	// leaf step takes the key and value hash as input
	res, err := p.Leaf.Apply(p.Key, p.ValueHash)
	if err != nil {
		return nil, errors.WithMessage(err, "leaf")
	}

	// the rest just take the output of the last step (reducing it)
	for _, step := range p.Path {
		res, err = step.Apply(res)
		if err != nil {
			return nil, errors.WithMessage(err, "inner")
		}
	}
	return res, nil
}

// CheckAgainstSpec will verify the leaf and all path steps are in the format defined in spec
func (p *ExistenceProof) CheckAgainstSpec(spec *ProofSpec) error {
	if p.GetLeaf() == nil {
		return errors.New("Existence Proof needs defined LeafOp")
	}
	err := p.Leaf.CheckAgainstSpec(spec)
	if err != nil {
		return errors.WithMessage(err, "leaf")
	}
	if spec.MinDepth > 0 && len(p.Path) < int(spec.MinDepth) {
		return errors.Errorf("InnerOps depth too short: %d", len(p.Path))
	}
	if spec.MaxDepth > 0 && len(p.Path) > int(spec.MaxDepth) {
		return errors.Errorf("InnerOps depth too long: %d", len(p.Path))
	}

	for _, inner := range p.Path {
		if err := inner.CheckAgainstSpec(spec); err != nil {
			return errors.WithMessage(err, "inner")
		}
	}
	return nil
}

// Verify does all checks to ensure the proof has valid non-existence proofs,
// and they ensure the given key is not in the CommitmentState
func (p *NonExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot, key []byte) error {
	// ensure the existence proofs are valid
	var leftKey, rightKey []byte
	if p.Left != nil {
		if err := p.Left.Verify(spec, root); err != nil {
			return errors.Wrap(err, "left proof")
		}
		leftKey = p.Left.Key
	}
	if p.Right != nil {
		if err := p.Right.Verify(spec, root); err != nil {
			return errors.Wrap(err, "right proof")
		}
		rightKey = p.Right.Key
	}

	// If both proofs are missing, this is not a valid proof
	if leftKey == nil && rightKey == nil {
		return errors.New("both left and right proofs missing")
	}

	// Ensure in valid range
	if rightKey != nil {
		if bytes.Compare(key, rightKey) >= 0 {
			return errors.New("key is not left of right proof")
		}
	}
	if leftKey != nil {
		if bytes.Compare(key, leftKey) <= 0 {
			return errors.New("key is not right of left proof")
		}
	}

	if leftKey == nil {
		if !IsLeftMost(spec.InnerSpec, p.Right.Path) {
			return errors.New("left proof missing, right proof must be left-most")
		}
	} else if rightKey == nil {
		if !IsRightMost(spec.InnerSpec, p.Left.Path) {
			return errors.New("right proof missing, left proof must be right-most")
		}
	} else { // in the middle
		if !IsLeftNeighbor(spec.InnerSpec, p.Left.Path, p.Right.Path) {
			return errors.New("right proof missing, left proof must be right-most")
		}
	}
	return nil
}

// IsLeftMost returns true if this is the left-most path in the tree
func IsLeftMost(spec *InnerSpec, path []*InnerOp) bool {
	minPrefix, maxPrefix, suffix := getPadding(spec, 0)

	// ensure every step has a prefix and suffix defined to be leftmost
	for _, step := range path {
		if !hasPadding(step, minPrefix, maxPrefix, suffix) {
			return false
		}
	}
	return true
}

// IsRightMost returns true if this is the left-most path in the tree
func IsRightMost(spec *InnerSpec, path []*InnerOp) bool {
	last := len(spec.ChildOrder) - 1
	minPrefix, maxPrefix, suffix := getPadding(spec, int32(last))

	// ensure every step has a prefix and suffix defined to be rightmost
	for _, step := range path {
		if !hasPadding(step, minPrefix, maxPrefix, suffix) {
			return false
		}
	}
	return true
}

// IsLeftNeighbor returns true if `right` is the next possible path right of `left`
//
//   Find the common suffix from the Left.Path and Right.Path and remove it. We have LPath and RPath now, which must be neighbors.
//   Validate that LPath[len-1] is the left neighbor of RPath[len-1]
//   For step in LPath[0..len-1], validate step is right-most node
//   For step in RPath[0..len-1], validate step is left-most node
func IsLeftNeighbor(spec *InnerSpec, left []*InnerOp, right []*InnerOp) bool {
	// count common tail (from end, near root)
	left, topleft := left[:len(left)-1], left[len(left)-1]
	right, topright := right[:len(right)-1], right[len(right)-1]
	for bytes.Equal(topleft.Prefix, topright.Prefix) && bytes.Equal(topleft.Suffix, topright.Suffix) {
		left, topleft = left[:len(left)-1], left[len(left)-1]
		right, topright = right[:len(right)-1], right[len(right)-1]
	}

	// now topleft and topright are the first divergent nodes
	// make sure they are left and right of each other
	if !isLeftStep(spec, topleft, topright) {
		return false
	}

	// left and right are remaining children below the split,
	// ensure left child is the rightmost path, and visa versa
	if !IsRightMost(spec, left) {
		return false
	}
	if !IsLeftMost(spec, right) {
		return false
	}
	return true
}

// isLeftStep assumes left and right have common parents
// checks if left is exactly one slot to the left of right
func isLeftStep(spec *InnerSpec, left *InnerOp, right *InnerOp) bool {
	leftidx, err := orderFromPadding(spec, left)
	if err != nil {
		panic(err)
	}
	rightidx, err := orderFromPadding(spec, right)
	if err != nil {
		panic(err)
	}

	// TODO: is it possible there are empty (nil) children???
	return rightidx == leftidx+1
}

func hasPadding(op *InnerOp, minPrefix, maxPrefix, suffix int) bool {
	if len(op.Prefix) < minPrefix {
		return false
	}
	if len(op.Prefix) > maxPrefix {
		return false
	}
	return len(op.Suffix) == suffix
}

// getPadding determines prefix and suffix with the given spec and position in the tree
func getPadding(spec *InnerSpec, branch int32) (minPrefix, maxPrefix, suffix int) {
	idx := getPosition(spec.ChildOrder, branch)

	// count how many children are in the prefix
	prefix := idx * int(spec.ChildSize)
	minPrefix = prefix + int(spec.MinPrefixLength)
	maxPrefix = prefix + int(spec.MaxPrefixLength)

	// count how many children are in the suffix
	suffix = (len(spec.ChildOrder) - 1 - idx) * int(spec.ChildSize)
	return
}

// getPosition checks where the branch is in the order and returns
// the index of this branch
func getPosition(order []int32, branch int32) int {
	if branch < 0 || int(branch) >= len(order) {
		panic(errors.Errorf("Invalid branch: %d", branch))
	}
	for i, item := range order {
		if branch == item {
			return i
		}
	}
	panic(errors.Errorf("Branch %d not found in order %v", branch, order))
}

// This will look at the proof and determine which order it is...
// So we can see if it is branch 0, 1, 2 etc... to determine neighbors
func orderFromPadding(spec *InnerSpec, inner *InnerOp) (int32, error) {
	maxbranch := int32(len(spec.ChildOrder))
	for branch := int32(0); branch < maxbranch; branch++ {
		minp, maxp, suffix := getPadding(spec, branch)
		if hasPadding(inner, minp, maxp, suffix) {
			return branch, nil
		}
	}
	return 0, errors.New("Cannot find any valid spacing for this node")
}
