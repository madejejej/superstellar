use crate::superstellar;

use crate::game::{GameInputMessage, GameOutputMessage, GameOutputReceiverStream, GameSender};
use anyhow::Result;
use futures_util::SinkExt;
use prost::Message;
use tokio_stream::StreamExt;
use tracing::{debug, error};
use warp::ws::WebSocket;

pub struct Client {
    pub id: u32,
    pub username: Option<String>,
    sender: GameSender,
    receiver: GameOutputReceiverStream,
    websocket: WebSocket,
}

impl Client {
    pub fn new(
        id: u32,
        websocket: WebSocket,
        sender: GameSender,
        receiver: GameOutputReceiverStream,
    ) -> Client {
        debug!("Initializing client {}", id);

        Client {
            id,
            username: None,
            sender,
            receiver,
            websocket,
        }
    }

    pub async fn run(&mut self) -> Result<()> {
        debug!("Client {} running", self.id);

        loop {
            tokio::select! {
                maybe_message = self.websocket.next() => {
                    match maybe_message {
                        Some(message) => {
                            match message {
                                Ok(message) => self.handle_websocket_message(message).await?,
                                Err(e) => {
                                    error!("Client {} error when reading message: {}", self.id, e);
                                    return self.disconnect_client().await;
                                }
                            }
                        }
                        None => {
                            return self.disconnect_client().await;
                        }
                    }
                }
                message = self.receiver.next() => {
                    let message = message.expect("Cant read message");
                    self.handle_game_message(message).await?;
                }

            }
        }
    }

    async fn handle_websocket_message(&mut self, message: warp::ws::Message) -> Result<()> {
        let user_message: superstellar::UserMessage = prost::Message::decode(message.as_bytes())?;
        debug!("Received {:?}", user_message);

        self.sender.send(GameInputMessage::PlayerCommand {
            id: self.id,
            message: user_message,
        })?;

        Ok(())
    }

    async fn handle_game_message(&mut self, message: GameOutputMessage) -> Result<()> {
        let mut buf = vec![];
        message.encode(&mut buf)?;
        // TODO: we might batch some messages and periodically flush
        self.websocket.send(warp::ws::Message::binary(buf)).await?;

        Ok(())
    }

    async fn disconnect_client(&mut self) -> Result<()> {
        self.sender
            .send(GameInputMessage::PlayerDisconnected { id: self.id })?;

        Ok(())
    }
}
