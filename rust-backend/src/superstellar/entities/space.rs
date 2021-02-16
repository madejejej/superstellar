use crate::superstellar::constants::ASTEROID_COUNT_LIMIT;
use crate::superstellar::entities::spaceship::Spaceship;
use crate::superstellar::Asteroid;
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

        for spaceship in self.spaceships.values_mut() {
            spaceship.update();
        }

        self.spawn_asteroids();

        for asteroid in self.asteroids.values_mut() {
            asteroid.update();
        }
    }

    pub fn add_spaceship(&mut self, id: u32) {
        self.spaceships.insert(id, Spaceship::new(id));
    }

    pub fn remove_spaceship(&mut self, id: &u32) {
        self.spaceships.remove(id);
    }

    pub fn spaceship_thrust(&mut self, id: &u32) {
        self.spaceships
            .get_mut(id)
            .map(|spaceship| spaceship.input_thrust(true));
    }

    pub fn spaceship_turn(&mut self, id: &u32, turn: i32) {
        self.spaceships
            .get_mut(id)
            .map(|spaceship| spaceship.turn(turn));
    }

    pub fn spaceship_no_thrust(&mut self, id: &u32) {
        self.spaceships
            .get_mut(id)
            .map(|spaceship| spaceship.input_thrust(false));
    }

    pub fn to_proto(&self) -> crate::superstellar::Space {
        crate::superstellar::Space {
            physics_frame_id: self.physics_frame_id,
            spaceships: self
                .spaceships
                .values()
                .map(|spaceship| spaceship.to_proto())
                .collect(),
            asteroids: self.asteroids.values().cloned().collect(),
        }
    }

    fn spawn_asteroids(&mut self) {
        if self.asteroids.len() < ASTEROID_COUNT_LIMIT {
            let asteroid = Asteroid::spawn_random();
            self.asteroids.insert(asteroid.id, asteroid);
        }
    }
}
