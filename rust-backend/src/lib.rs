pub mod types;
pub mod server;

// Include the `items` module, which is generated from items.proto.
pub mod superstellar {
    include!(concat!(env!("OUT_DIR"), "/superstellar.rs"));
}