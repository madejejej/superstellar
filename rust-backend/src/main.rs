use simplelog::*;

mod server;
mod superstellar;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    CombinedLogger::init(vec![TermLogger::new(
        LevelFilter::Debug,
        Config::default(),
        TerminalMode::Mixed,
    )])
    .unwrap();

    // TODO: use warp create to handle websockets and http
    let server = crate::server::Server::new("127.0.0.1:8081");

    server.run().await
}
