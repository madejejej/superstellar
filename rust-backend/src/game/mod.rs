use std::collections::HashMap;
use std::sync::Arc;

use tokio::sync::mpsc::UnboundedSender;
use tokio_stream::wrappers::UnboundedReceiverStream;
use tokio_stream::StreamExt;
use tracing::{debug, error, info};

use crate::superstellar;

// Sender and Receiver types coming into Game
pub type GameInputReceiverStream = UnboundedReceiverStream<GameInputMessage>;
pub type GameSender = UnboundedSender<GameInputMessage>;

// Sender and Receiver types for Messages coming out of Game.
// We use an Arc, because Game output needs to send data to every client.
// It should more efficient to send a pointer rather than copying the whole structure.
pub type GameOutputMessage = Arc<superstellar::Message>;
pub type GameOutputSender = UnboundedSender<GameOutputMessage>;
pub type GameOutputReceiverStream = UnboundedReceiverStream<GameOutputMessage>;

#[derive(Debug)]
enum PlayerState {
    Joining,
    Playing,
    Lost,
}

#[derive(Debug)]
struct Player {
    pub id: u32,
    pub username: Option<String>,
    state: PlayerState,
    sender: GameOutputSender,
}

impl Player {
    pub fn new(id: u32, sender: GameOutputSender) -> Player {
        Player {
            id,
            state: PlayerState::Joining,
            username: None,
            sender,
        }
    }

    pub fn send_message(&mut self, message: GameOutputMessage) {
        self.sender
            .send(message)
            .unwrap_or_else(|_| error!("Can't send message to player {:?}", self.id));
    }
}

#[derive(Debug)]
pub enum GameInputMessage {
    PlayerConnected {
        id: u32,
        sender: UnboundedSender<GameOutputMessage>,
    },
    PlayerDisconnected {
        id: u32,
    },
    PlayerCommand {
        id: u32,
        message: superstellar::UserMessage,
    },
}

pub struct Game {
    players: HashMap<u32, Player>,
    receiver: GameInputReceiverStream,
    space: superstellar::entities::space::Space,
}

impl Game {
    pub fn new(receiver: GameInputReceiverStream) -> Game {
        let players = HashMap::new();
        let space = superstellar::entities::space::Space::new();

        Game {
            players,
            receiver,
            space,
        }
    }

    pub async fn run(&mut self) {
        info!("Game loop running");

        let mut update_timer = tokio::time::interval(superstellar::constants::UPDATE_SEND_INTERVAL);
        let mut physics_timer =
            tokio::time::interval(superstellar::constants::PHYSICS_FRAME_DURATION);

        loop {
            tokio::select! {
                message = self.receiver.next() => {
                    message.map(|msg| self.handle_message(msg));
                }
                _ = update_timer.tick() => {
                    self.send_updates();
                }
                _ = physics_timer.tick() => {
                    self.space.update();
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
            GameInputMessage::PlayerDisconnected { id } => {
                self.remove_player(id);
            }
            GameInputMessage::PlayerCommand { id, message } => {
                match message.content {
                    Some(superstellar::user_message::Content::JoinGame(join_game)) => {
                        self.join(id, join_game.username);
                    }
                    Some(superstellar::user_message::Content::UserAction(user_action)) => {
                        self.handle_user_action(id, user_action)
                    }
                    _ => {
                        // TODO
                    }
                }
            }
        }
    }

    fn join(&mut self, id: u32, username: String) {
        debug!("Player {} joining with username {}", id, username);

        let mut player = self.players.get_mut(&id).unwrap();
        player.username = Some(username.clone());

        let message = superstellar::Message {
            content: Some(superstellar::message::Content::JoinGameAck(
                superstellar::JoinGameAck {
                    success: true,
                    error: "".to_string(),
                },
            )),
        };

        player.send_message(Arc::new(message));
        self.space.add_spaceship(player.id);
        self.announce_player_joined(id, username);
    }

    fn remove_player(&mut self, id: u32) {
        self.players
            .remove(&id)
            .expect("Cannot remove player from map");

        self.space.remove_spaceship(&id);
        let message = GameOutputMessage::new(superstellar::Message {
            content: Some(superstellar::message::Content::PlayerLeft(
                superstellar::PlayerLeft { id },
            )),
        });

        self.broadcast(message);
    }

    fn announce_player_joined(&mut self, id: u32, username: String) {
        let message = GameOutputMessage::new(superstellar::Message {
            content: Some(superstellar::message::Content::PlayerJoined(
                superstellar::PlayerJoined { id, username },
            )),
        });

        self.broadcast(message);
    }

    fn handle_user_action(&mut self, id: u32, user_action: superstellar::UserAction) {
        match user_action.user_input {
            x if x == superstellar::UserInput::ThrustOn as i32 => {
                self.space.spaceship_thrust(&id);
            }
            x if x == superstellar::UserInput::ThrustOff as i32 => {
                self.space.spaceship_no_thrust(&id);
            }
            _ => {}
        }
    }

    fn send_updates(&mut self) {
        let space = self.space.to_proto();
        let message = GameOutputMessage::new(superstellar::Message {
            content: Some(superstellar::message::Content::Space(space)),
        });

        self.broadcast(message);
    }

    fn broadcast(&mut self, message: GameOutputMessage) {
        for player in self.players.values_mut() {
            player.send_message(message.clone());
        }
    }
}
