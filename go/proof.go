package proofs

import (
	"bytes"
	"fmt"
)

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() ([]byte, error) {
	if len(p.Steps) == 0 {
		return nil, fmt.Errorf("Existence Proof needs at least one step")
	}

	leaf, inners := p.Steps[0], p.Steps[1:]
	// leaf step takes the key and value as input
	res, err := leaf.Apply(p.Key, p.Value)
	if err != nil {
		return nil, err
	}

	// the rest just take the output of the last step (reducing it)
	for _, step := range inners {
		res, err = step.Apply(res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (p *ExistenceProof) ValidSpec(spec *ProofSpec) error {
	if len(p.Steps) == 0 {
		return fmt.Errorf("Existence Proof needs at least one step")
	}

	leafOp, inners := p.Steps[0], p.Steps[1:]
	leaf, err := asLeaf(leafOp)
	if err != nil {
		return err
	}
	err = p.checkLeaf(leaf, spec.LeafSpec)
	if err != nil {
		return err
	}
	for _, innerOp := range inners {
		inner, err := asInner(innerOp)
		if err != nil {
			return err
		}
		if err := p.checkInner(inner, spec.LeafSpec.Prefix); err != nil {
			return err
		}
	}
	return nil
}

func (p *ExistenceProof) checkLeaf(leaf *LeafOp, spec *LeafOp) error {
	if leaf.Hash != spec.Hash {
		return fmt.Errorf("Unexpected HashOp: %d", leaf.Hash)
	}
	if leaf.PrehashKey != spec.PrehashKey {
		return fmt.Errorf("Unexpected PrehashKey: %d", leaf.PrehashKey)
	}
	if leaf.PrehashValue != spec.PrehashValue {
		return fmt.Errorf("Unexpected PrehashValue: %d", leaf.PrehashValue)
	}
	if leaf.Length != spec.Length {
		return fmt.Errorf("Unexpected LengthOp: %d", leaf.Length)
	}
	if !bytes.HasPrefix(leaf.Prefix, spec.Prefix) {
		return fmt.Errorf("Leaf Prefix doesn't start with %X", spec.Prefix)
	}
	return nil
}

func (p *ExistenceProof) checkInner(inner *InnerOp, leafPrefix []byte) error {
	if bytes.HasPrefix(inner.Prefix, leafPrefix) {
		return fmt.Errorf("Inner Prefix starts with %X", leafPrefix)
	}
	return nil
}
