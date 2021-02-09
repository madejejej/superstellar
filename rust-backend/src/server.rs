use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

use tracing::{error, debug};

pub struct Server {
    pub address: &'static str,
}

impl Server {
    pub fn new(address: &'static str) -> Server {
       Server {
           address
       }
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
            let (mut socket, _) = listener.accept().await?;
            debug!("New connection accepted");

            tokio::spawn(async move {
                let mut buf = [0; 1024];

                // In a loop, read data from the socket and write the data back.
                loop {
                    let n = match socket.read(&mut buf).await {
                        // socket closed
                        Ok(n) if n == 0 => return,
                        Ok(n) => n,
                        Err(e) => {
                            error!("failed to read from socket: {:?}", e);
                            return;
                        }
                    };

                    // Write the data back
                    if let Err(e) = socket.write_all(&buf[0..n]).await {
                        error!("failed to write to socket");
                        return;
                    }
                }
            });
        }
    }
}
