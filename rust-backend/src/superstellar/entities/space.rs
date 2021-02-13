use crate::superstellar::{Asteroid, Spaceship};
use std::collections::HashMap;

pub struct Space {
    physics_frame_id: u32,
    spaceships: HashMap<u32, Spaceship>,
    asteroids: HashMap<u32, Asteroid>,
}

impl Space {
    pub fn new() -> Space {
        Space {
            physics_frame_id: 0,
            spaceships: HashMap::new(),
            asteroids: HashMap::new(),
        }
    }

    pub fn update(&mut self) {
        self.physics_frame_id += 1;
    }

    pub fn add_spaceship(&mut self, id: u32) {
        self.spaceships.insert(id, Spaceship::new(id));
    }

    pub fn remove_spaceship(&mut self, id: &u32) {
        self.spaceships.remove(id);
    }

    pub fn to_proto(&self) -> crate::superstellar::Space {
        crate::superstellar::Space {
            physics_frame_id: self.physics_frame_id,
            spaceships: self.spaceships.values().cloned().collect(),
            asteroids: self.asteroids.values().cloned().collect(),
        }
    }
}
