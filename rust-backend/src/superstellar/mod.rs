// Include the `items` module, which is generated from items.proto.
include!(concat!(env!("OUT_DIR"), "/superstellar.rs"));
pub mod constants;
pub mod entities;
