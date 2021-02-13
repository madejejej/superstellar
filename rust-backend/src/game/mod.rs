use crate::superstellar;

use std::collections::HashMap;
use tokio::sync::mpsc::UnboundedSender;
use tokio_stream::wrappers::UnboundedReceiverStream;
use tokio_stream::StreamExt;
use tracing::{debug, info};
use uuid::Uuid;

pub type GameInputReceiverStream = UnboundedReceiverStream<GameInputMessage>;
pub type GameOutputSender = UnboundedSender<superstellar::Message>;
pub type GameOutputReceiverStream = UnboundedReceiverStream<superstellar::Message>;
pub type GameSender = UnboundedSender<GameInputMessage>;

#[derive(Debug)]
enum PlayerState {
    Joining,
    Playing,
    Lost,
}

#[derive(Debug)]
struct Player {
    pub id: Uuid,
    pub username: Option<String>,
    state: PlayerState,
    sender: GameOutputSender,
}

impl Player {
    pub fn new(id: Uuid, sender: GameOutputSender) -> Player {
        Player {
            id,
            state: PlayerState::Joining,
            username: None,
            sender,
        }
    }

    pub fn send_message(&mut self, message: superstellar::Message) {
        self.sender.send(message);
    }
}

#[derive(Debug)]
pub enum GameInputMessage {
    PlayerConnected {
        id: Uuid,
        sender: UnboundedSender<superstellar::Message>,
    },
    PlayerCommand {
        id: Uuid,
        message: superstellar::UserMessage,
    },
}

pub struct Game {
    players: HashMap<Uuid, Player>,
    receiver: GameInputReceiverStream,
}

impl Game {
    pub fn new(receiver: GameInputReceiverStream) -> Game {
        let players = HashMap::new();

        Game { players, receiver }
    }

    pub async fn run(&mut self) {
        info!("Game loop running");

        loop {
            tokio::select! {
                message = self.receiver.next() => {
                    let message = message.expect("cant read message");
                    debug!("Received message: {:?}", message);
                    self.handle_message(message);

                }
            };
        }
    }

    fn handle_message(&mut self, message: GameInputMessage) {
        match message {
            GameInputMessage::PlayerConnected { id, sender } => {
                let player = Player::new(id, sender);
                self.players.insert(id, player);
            }
            GameInputMessage::PlayerCommand { id, message } => {
                match message.content {
                    Some(superstellar::user_message::Content::JoinGame(join_game)) => {
                        self.join(id, join_game.username);
                    }
                    _ => {
                        // TODO
                    }
                }
            }
        }
    }

    pub fn join(&mut self, id: Uuid, username: String) {
        debug!("Player {} joining with username {}", id, username);

        let mut player = self.players.get_mut(&id).unwrap();

        player.username = Some(username.clone());

        let message = superstellar::Message {
            content: Some(superstellar::message::Content::PlayerJoined(
                superstellar::PlayerJoined { id: 1, username },
            )),
        };

        for (id, player) in &mut self.players {
            player.send_message(message.clone());
        }
    }
}
