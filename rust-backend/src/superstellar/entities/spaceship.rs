use crate::superstellar::constants;
use crate::superstellar::{Direction, Point, Vector};

impl crate::superstellar::Spaceship {
    pub fn new(id: u32) -> Self {
        Self {
            id,
            position: Some(Point { x: 0, y: 0 }),
            velocity: Some(Vector { x: 0.0, y: 0.0 }),
            facing: 0.0,
            angular_velocity: 0.0,
            input_direction: Direction::DirCenter.into(),
            input_thrust: false,
            input_boost: false,
            max_hp: constants::SPACESHIP_INITIAL_HP,
            hp: constants::SPACESHIP_INITIAL_HP,
            max_energy: 100,
            energy: 100,
            auto_repair_delay: 10,
        }
    }
}
