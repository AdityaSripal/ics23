extern crate protoc_rust;

use std::env;
use std::vec::Vec;
use protoc_rust::Customize;

fn main() {
	let args: Vec<_> = env::args().collect();
	let mut root = "../..";
	if args.len() > 1 {
		root = &args[1];
	}

	protoc_rust::run(protoc_rust::Args {
	    out_dir: &format!("{}{}", root, "/rust/src"),
	    input: &[&format!("{}{}", root, "/proofs.proto")],
	    includes: &[root],
	    customize: Customize {
	      ..Default::default()
	    },
	}).expect("protoc");
}
