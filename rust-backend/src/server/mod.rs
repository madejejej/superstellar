mod client;

use crate::server::client::Client;
use tokio::net::{TcpListener, TcpStream};
use tracing::{debug, error};

pub struct Server {
    pub address: &'static str,
}

impl Server {
    pub fn new(address: &'static str) -> Server {
        Server { address }
    }

    pub async fn run(&self) -> Result<(), Box<dyn std::error::Error>> {
        debug!("Starting server");

        let listener = match TcpListener::bind(&self.address).await {
            Ok(n) => n,
            Err(e) => {
                error!("Failed to start the server: {:?}", e);
                return Err(e.into());
            }
        };

        debug!("Started listening on socket {}", self.address);

        loop {
            let (socket, _) = listener.accept().await.expect("No connections to accept");
            debug!("New connection accepted");

            tokio::spawn(async move {
                let mut client = Client::new(socket).await;
                client.run();
            });
        }
    }
}
