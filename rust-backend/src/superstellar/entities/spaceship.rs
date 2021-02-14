use crate::superstellar::constants;
use crate::superstellar::constants::{
    FRICTION_COEFFICIENT, SPACESHIP_ANGULAR_FRICTION, SPACESHIP_LINEAR_ANGULAR_ACCELERATION,
    SPACESHIP_MAX_ANGULAR_VELOCITY, SPACESHIP_NON_LINEAR_ANGULAR_ACCELERATION,
};
use crate::superstellar::{Direction, Point, Vector};
use std::f64::consts::PI;

pub struct Spaceship {
    angular_velocity_delta: f32,
    pb_spaceship: crate::superstellar::Spaceship,
}

impl Spaceship {
    pub fn new(id: u32) -> Spaceship {
        Spaceship {
            angular_velocity_delta: 0.0,
            pb_spaceship: crate::superstellar::Spaceship {
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
            },
        }
    }

    pub fn update(&mut self) {
        let thrust = self.pb_spaceship.input_thrust;
        let facing = self.pb_spaceship.facing;

        self.pb_spaceship.velocity.as_mut().map(|velocity| {
            if thrust {
                let delta = Vector::new(facing.cos() as f32, -facing.sin() as f32)
                    * constants::SPACESHIP_ACCELERATION;
                *velocity += &delta
            } else {
                if !velocity.zero() {
                    *velocity *= 1.0 - FRICTION_COEFFICIENT;

                    if velocity.length() < 1.0 {
                        velocity.set_zero();
                    }
                }
            }
        });

        self.pb_spaceship
            .position
            .as_mut()
            .zip(self.pb_spaceship.velocity.as_ref())
            .map(|(position, velocity)| *position += velocity);

        match self.pb_spaceship.input_direction {
            x if x == Direction::DirLeft as i32 => {
                self.angular_velocity_delta = -self.angular_velocity_delta();
                self.limit_angular_velocity_delta();
            }
            x if x == Direction::DirRight as i32 => {
                self.angular_velocity_delta = self.angular_velocity_delta();
                self.limit_angular_velocity_delta();
            }
            _ => {
                self.pb_spaceship.angular_velocity *= 1.0 - SPACESHIP_ANGULAR_FRICTION as f64;
            }
        }

        self.pb_spaceship.angular_velocity += self.angular_velocity_delta as f64;
        self.pb_spaceship.facing -= self.pb_spaceship.angular_velocity;

        if self.pb_spaceship.facing.abs() > PI {
            self.pb_spaceship.facing -= 2.0 * PI * self.pb_spaceship.facing.signum();
        }

        self.angular_velocity_delta = 0.0;
    }

    pub fn turn(&mut self, val: i32) {
        self.pb_spaceship.input_direction = val;
    }

    pub fn input_thrust(&mut self, val: bool) {
        self.pb_spaceship.input_thrust = val;
    }

    pub fn to_proto(&self) -> crate::superstellar::Spaceship {
        self.pb_spaceship.clone()
    }

    fn angular_velocity_delta(&self) -> f32 {
        let nonlinear_part = SPACESHIP_NON_LINEAR_ANGULAR_ACCELERATION
            * self.pb_spaceship.angular_velocity.abs() as f32;
        let linear_part = SPACESHIP_LINEAR_ANGULAR_ACCELERATION;

        nonlinear_part + linear_part
    }

    fn limit_angular_velocity_delta(&mut self) {
        let potential_angular_velocity =
            self.pb_spaceship.angular_velocity as f32 + self.angular_velocity_delta;
        let diff = potential_angular_velocity.abs() - SPACESHIP_MAX_ANGULAR_VELOCITY;

        if diff > 0.0 {
            self.angular_velocity_delta -=
                diff * self.pb_spaceship.angular_velocity.signum() as f32;
        }
    }
}
