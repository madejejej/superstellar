use std::collections::HashMap;
use uuid::Uuid;

struct Player {
    id: Uuid,
    username: String,
}

impl Player {
    pub fn new(id: Uuid, username: String) -> Player {
        Player { id, username }
    }
}

pub struct Game {
    players: HashMap<Uuid, Player>,
}

impl Game {
    pub fn new() -> Game {
        let players = HashMap::new();

        Game { players }
    }

    pub fn join(&mut self, id: Uuid, username: String) {
        let player = Player::new(id, username);

        self.players.insert(id, player);
    }
}
