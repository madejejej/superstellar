use uuid::Uuid;

use crate::superstellar;

use crate::game::{GameInputMessage, GameOutputReceiverStream, GameSender};
use tokio_stream::StreamExt;
use tracing::debug;
use warp::ws::WebSocket;

pub struct Client {
    pub id: Uuid,
    pub username: Option<String>,
    sender: GameSender,
    receiver: GameOutputReceiverStream,
    websocket: WebSocket,
}

impl Client {
    pub fn new(
        websocket: WebSocket,
        sender: GameSender,
        receiver: GameOutputReceiverStream,
    ) -> Client {
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

            self.sender.send(GameInputMessage::PlayerCommand {
                id: self.id,
                message: user_message,
            });
        }
    }
}
