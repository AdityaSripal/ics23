// we want to name functions verify_* to match ics23
#![allow(clippy::module_name_repetitions)]

use crate::helpers::Result;
use crate::ops::{apply_inner, apply_leaf};
use crate::proofs;
use failure::{bail, ensure};

pub type CommitmentRoot = ::std::vec::Vec<u8>;

pub fn verify_existence(
    proof: &proofs::ExistenceProof,
    spec: &proofs::ProofSpec,
    root: &[u8],
    key: &[u8],
    value: &[u8],
) -> Result<()> {
    check_existence_spec(proof, spec)?;
    ensure!(proof.key == key, "Provided key doesn't match proof");
    ensure!(proof.value == value, "Provided value doesn't match proof");

    let calc = calculate_existence_root(&proof)?;
    ensure!(calc == root, "Root hash doesn't match");
    Ok(())
}

pub fn verify_non_existence(
    proof: &proofs::NonExistenceProof,
    spec: &proofs::ProofSpec,
    root: &[u8],
    key: &[u8],
) -> Result<()> {
    let mut left_key: Option<&[u8]> = None;
    let mut right_key: Option<&[u8]> = None;

//    if let Some(left) = &proof.left {
//        verify_existence(&left, spec, root, &left.key, &left.value)?;
//        ensure!(&left.key < key, "left key isn't before key");
//        left_key = Some(&left.key);
//    }
//    if let Some(right) = &proof.right {
//        verify_existence(&right, spec, root, &right.key, &right.value)?;
//        ensure!(key < &right.key, "right key isn't after key");
//        right_key = Some(&right.key);
//    }

    ensure!(left_key != None || right_key != None, "neither left nor right proof defined");

//    check_existence_spec(proof, spec)?;
//    ensure!(proof.key == key, "Provided key doesn't match proof");
//    ensure!(proof.value == value, "Provided value doesn't match proof");
//
//    let calc = calculate_existence_root(&proof)?;
//    ensure!(calc == root, "Root hash doesn't match");
    Ok(())
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
pub fn calculate_existence_root(proof: &proofs::ExistenceProof) -> Result<CommitmentRoot> {
    ensure!(!proof.key.is_empty(), "Existence proof must have key set");
    ensure!(
        !proof.value.is_empty(),
        "Existence proof must have value set"
    );

    if let Some(leaf_node) = &proof.leaf {
        let mut hash = apply_leaf(leaf_node, &proof.key, &proof.value)?;
        for step in proof.path.iter() {
            hash = apply_inner(step, &hash)?;
        }
        Ok(hash)
    } else {
        bail!("No leaf operation set")
    }
}

fn check_existence_spec(proof: &proofs::ExistenceProof, spec: &proofs::ProofSpec) -> Result<()> {
    if let (Some(leaf), Some(leaf_spec)) = (&proof.leaf, &spec.leaf_spec) {
        ensure_leaf(leaf, leaf_spec)?;
        for step in proof.path.iter() {
            ensure_inner(step, spec)?;
        }
        Ok(())
    } else {
        bail!("Leaf and Leaf Spec must be set")
    }
}

fn ensure_leaf(leaf: &proofs::LeafOp, leaf_spec: &proofs::LeafOp) -> Result<()> {
    ensure!(
        leaf_spec.hash == leaf.hash,
        "Unexpected hashOp: {:?}",
        leaf.hash
    );
    ensure!(
        leaf_spec.prehash_key == leaf.prehash_key,
        "Unexpected prehash_key: {:?}",
        leaf.prehash_key
    );
    ensure!(
        leaf_spec.prehash_value == leaf.prehash_value,
        "Unexpected prehash_value: {:?}",
        leaf.prehash_value
    );
    ensure!(
        leaf_spec.length == leaf.length,
        "Unexpected lengthOp: {:?}",
        leaf.length
    );
    ensure!(
        has_prefix(&leaf_spec.prefix, &leaf.prefix),
        "Incorrect prefix on leaf"
    );
    Ok(())
}

fn has_prefix(prefix: &[u8], data: &[u8]) -> bool {
    if prefix.len() > data.len() {
        return false;
    }
    prefix == &data[..prefix.len()]
}

fn ensure_inner(inner: &proofs::InnerOp, spec: &proofs::ProofSpec) -> Result<()> {
    if let Some(leaf_spec) = &spec.leaf_spec {
        ensure!(
            !has_prefix(&leaf_spec.prefix, &inner.prefix),
            "Inner node with leaf prefix"
        );
        Ok(())
    } else {
        bail!("Spec missing leaf_spec")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use crate::proofs::{HashOp, LengthOp};

    #[test]
    fn calculate_root_from_leaf() -> Result<()> {
        let leaf = proofs::LeafOp{
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![]
        };

        let proof = proofs::ExistenceProof{
            key: b"food".to_vec(),
            value: b"some longer text".to_vec(),
            leaf: Some(leaf),
            path: vec![]
        };

        let expected =
            hex::decode("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265")?;
        ensure!(
            expected == calculate_existence_root(&proof)?,
            "invalid root hash"
        );
        Ok(())
    }

    #[test]
    fn calculate_root_from_leaf_and_inner() -> Result<()> {
        let leaf = proofs::LeafOp{
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![]
        };

        let inner = proofs::InnerOp{
            hash: HashOp::Sha256.into(),
            prefix: hex::decode("deadbeef00cafe00")?,
            suffix: vec![]
        };

        let proof = proofs::ExistenceProof{
            key: b"food".to_vec(),
            value: b"some longer text".to_vec(),
            leaf: Some(leaf),
            path: vec![inner],
        };

        let expected =
            hex::decode("836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9")?;
        ensure!(
            expected == calculate_existence_root(&proof)?,
            "invalid root hash"
        );
        Ok(())
    }
}
