use uuid::Uuid;

use futures_util::StreamExt;
use tokio::io::{AsyncRead, AsyncWrite};
use tokio::net::{TcpListener, TcpStream};
use tokio_tungstenite::WebSocketStream;
use tracing::{debug, error};

pub struct Client<S>
where
    S: AsyncRead + AsyncWrite + Unpin,
{
    id: Uuid,
    stream: WebSocketStream<S>,
}

impl<S> Client<S>
where
    S: AsyncRead + AsyncWrite + Unpin,
{
    pub async fn new(stream: S) -> Client<S>
    where
        S: AsyncRead + AsyncWrite + Unpin,
    {
        let stream = tokio_tungstenite::accept_async(stream).await;
        let stream = stream.expect("Failed to handshake with connection");
        let id = Uuid::new_v4();

        Client { stream, id }
    }

    pub async fn run(&mut self) -> () {
        while let Some(message) = self.stream.next().await {
            debug!("Message received from client {}", self.id);
        }
    }
}
