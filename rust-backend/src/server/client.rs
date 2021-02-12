use uuid::Uuid;

use crate::superstellar;

use std::error::Error;
use tokio::sync::mpsc::UnboundedSender;
use tokio_stream::wrappers::UnboundedReceiverStream;
use tokio_stream::StreamExt;
use tracing::debug;
use warp::ws::{Message, WebSocket};

type ReceiverStream = UnboundedReceiverStream<Result<Message, warp::Error>>;
type Sender = UnboundedSender<Result<Message, warp::Error>>;

#[derive(Debug)]
pub struct Client {
    pub id: Uuid,
    pub username: Option<String>,
    sender: Sender,
    receiver: ReceiverStream,
    websocket: WebSocket,
}

impl Client {
    pub fn new(websocket: WebSocket, sender: Sender, receiver: ReceiverStream) -> Client {
        let id = Uuid::new_v4();

        debug!("Initializing client {}", id);

        Client {
            id,
            username: None,
            sender,
            receiver,
            websocket,
        }
    }

    pub async fn run(&mut self) {
        debug!("Client {} running", self.id);
        while let Some(message) = self.websocket.next().await {
            let message = message.expect("Cant read message");

            let user_message: superstellar::UserMessage =
                prost::Message::decode(message.as_bytes()).expect("Cannot decode protobuf message");
            debug!("Received {:?}", user_message);

            match user_message.content {
                Some(superstellar::user_message::Content::JoinGame(join_game)) => {
                    debug!("Player asks to join {:?}", join_game.username);
                    self.username = Some(join_game.username);
                }
                _ => {
                    // TODO
                }
            }
        }
    }
}
