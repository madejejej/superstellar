use uuid::Uuid;

use futures_util::StreamExt;
use tokio::io::{AsyncRead, AsyncWrite};
use tokio::net::{TcpListener, TcpStream};
use tokio_tungstenite::WebSocketStream;
use tracing::{debug, error};
use tokio_tungstenite::tungstenite::Message;
use crate::superstellar;

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

        debug!("Websocket handshake done for client {}", id);

        Client { stream, id }
    }

    pub async fn run(&mut self) -> () {
        while let Some(message) = self.stream.next().await {
            let message = message.expect("Cannot read message from client");
            debug!("Message received from client {}", self.id);

            match message {
                Message::Binary(vec) => {
                    let pb_message: superstellar::Message = prost::Message::decode(&*vec).expect("Cannot decode protobuf message");
                    debug!("Received {:?}", pb_message);
                }
                message => {
                    error!("Unexpected message {:?}", message);
                }
            }
        }
    }
}
