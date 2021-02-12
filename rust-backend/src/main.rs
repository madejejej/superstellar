use simplelog::*;

mod server;
mod superstellar;

#[tokio::main]
async fn main() {
    CombinedLogger::init(vec![TermLogger::new(
        LevelFilter::Debug,
        Config::default(),
        TerminalMode::Mixed,
    )])
    .unwrap();

    let server = crate::server::Server::new("127.0.0.1:8081");

    server.run().await
}
