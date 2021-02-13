use std::time::Duration;

// Time between physics frames to be recomputed. We aim for 50 fps.
pub const PHYSICS_FRAME_DURATION: Duration = Duration::from_millis(20);

// How often we want to send updates to clients (if there are any updates)
pub const UPDATE_SEND_INTERVAL: Duration = Duration::from_millis(10);

// Acceleration is spaceship's linear acceleration on thruster.
pub const SPACESHIP_ACCELERATION: f32 = 50.0;

// maximum speed of each Spaceship
pub const SPACESHIP_MAX_SPEED: f64 = 600.0;

// Initial hit points value for each Spaceship
pub const SPACESHIP_INITIAL_HP: u32 = 500;
