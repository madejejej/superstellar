fn main() {
    prost_build::compile_protos(
        &["../protobuf/superstellar.proto"],
        &["../protobuf"]).unwrap();
}