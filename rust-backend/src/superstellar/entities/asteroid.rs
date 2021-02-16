use crate::superstellar::constants::{ASTEROID_SPAWN_RADIUS, ASTEROID_VELOCITY};
use crate::superstellar::{Point, Vector};
use rand::Rng;
use std::f32::consts::PI;
use std::sync::atomic::{AtomicU32, Ordering};

static NEXT_ASTEROID_ID: AtomicU32 = AtomicU32::new(1);

impl crate::superstellar::Asteroid {
    pub fn spawn_random() -> Self {
        let mut rnd = rand::thread_rng();

        let angle: f32 = rnd.gen::<f32>() * 2.0 * PI;
        let position = Point::from_polar(angle, ASTEROID_SPAWN_RADIUS);

        let direction_range = (rnd.gen::<f32>() - 0.5) * 0.25 * PI;
        let direction_angle = angle - PI + direction_range;

        let velocity = Vector::new(ASTEROID_VELOCITY, 0.0).rotate(direction_angle);
        let id = NEXT_ASTEROID_ID.fetch_add(1, Ordering::Relaxed);

        return Self {
            id,
            position: Some(position),
            velocity: Some(velocity),
            facing: 0.0,
            angular_velocity: 0.0,
        };
    }

    pub fn update(&mut self) {
        self.position
            .as_mut()
            .zip(self.velocity.as_ref())
            .map(|(position, velocity)| *position += velocity);
    }
}
