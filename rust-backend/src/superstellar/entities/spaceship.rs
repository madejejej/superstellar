use crate::superstellar::constants;
use crate::superstellar::{Direction, Point, Vector};

impl crate::superstellar::Spaceship {
    pub fn new(id: u32) -> Self {
        Self {
            id,
            position: Some(Point::new(0, 0)),
            velocity: Some(Vector::new(0.0, 0.0)),
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

    pub fn update(&mut self) {
        if self.input_thrust {
            let delta = Vector::new(self.facing.cos() as f32, -self.facing.sin() as f32)
                * constants::SPACESHIP_ACCELERATION;
            self.velocity.as_mut().map(|velocity| *velocity += &delta);
        }

        self.position
            .as_mut()
            .zip(self.velocity.as_ref())
            .map(|(position, velocity)| *position += velocity);
    }
}
