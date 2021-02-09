use simplelog::*;
use crate::server::Server;

mod server;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    CombinedLogger::init(
        vec![
            TermLogger::new(LevelFilter::Debug, Config::default(), TerminalMode::Mixed),
        ]
    ).unwrap();

    let server = Server::new("127.0.0.1:8081");

    server.run().await
}
