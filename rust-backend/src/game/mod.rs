use crate::superstellar;


use std::collections::HashMap;
use tokio::sync::mpsc::UnboundedSender;
use tokio_stream::wrappers::UnboundedReceiverStream;
use tokio_stream::StreamExt;
use tracing::{debug, info};
use uuid::Uuid;

pub type GameInputReceiverStream = UnboundedReceiverStream<GameInputMessage>;
pub type GameOutputSender = UnboundedSender<GameOutputMessage>;
pub type GameOutputReceiverStream = UnboundedReceiverStream<GameOutputMessage>;
pub type GameSender = UnboundedSender<GameInputMessage>;

struct Player {
    pub id: Uuid,
    pub username: Option<String>,
    sender: GameOutputSender,
}

impl Player {
    pub fn new(id: Uuid, sender: GameOutputSender) -> Player {
        Player {
            id,
            username: None,
            sender,
        }
    }
}

pub enum GameOutputMessage {
    Todo,
}

#[derive(Debug)]
pub enum GameInputMessage {
    PlayerConnected {
        id: Uuid,
        sender: UnboundedSender<GameOutputMessage>,
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

        player.username = Some(username);
    }
}
