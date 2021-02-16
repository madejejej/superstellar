use std::time::Duration;

// How many Asteroids to spawn at most
pub const ASTEROID_COUNT_LIMIT: usize = 20;

// The radius in which the Asteroids are spawned
pub const ASTEROID_SPAWN_RADIUS: f32 = WORLD_RADIUS as f32 * 2.0;

// Initial velocity of an Asteroid
pub const ASTEROID_VELOCITY: f32 = 100.0;

// the amount of energy a spaceship will receive after AUTO_REPAIR_DELAY
pub const AUTO_ENERGY_RECHARGE_AMOUNT: u32 = 3;

// the amount of HP received after AUTO_REPAIR_DELAY
pub const AUTO_REPAIR_AMOUNT: u32 = 2;

// the time after which the spaceship will repair itself
pub const AUTO_REPAIR_DELAY: u32 = 1000;

// the time between repairs when a Spaceship wasn't hit
pub const AUTO_REPAIR_INTERVAL: u32 = 1;

// the width of boundary region (in .01 units), i.e. from WorldRadius till when no more movement is possible
pub const BOUNDARY_ANNULUS_WIDTH: f32 = 40000.0;

// energy cost of frame of boosting
pub const BOOST_PER_FRAME_ENERGY_COST: u32 = 20;

// delta used for checking f32 equality
pub const DELTA: f32 = 0.0000001;

// a timeout measured in frames after which the ship is marked dirty
pub const DIRTY_FRAMES_TIMEOUT: u32 = 50;

// coefficient saying how fast a spaceship will slow down when not using acceleration
pub const FRICTION_COEFFICIENT: f32 = 0.005;

// minimum time between firing
pub const MIN_FIRE_INTERVAL: Duration = Duration::from_millis(250);

// Time between physics frames to be recomputed. We aim for 50 fps.
pub const PHYSICS_FRAME_DURATION: Duration = Duration::from_millis(20);

// the minimum radius around randomized initial position that needs to be free of any objects
pub const RANDOM_POSITION_EMPTY_RADIUS: f32 = 5000.0;

// Acceleration is spaceship's linear acceleration on thruster.
pub const SPACESHIP_ACCELERATION: f32 = 50.0;

// the multiplier for maximum speed when the boost is active
pub const SPACESHIP_BOOST_FACTOR: f32 = 2.5;

pub const SPACESHIP_ANGULAR_FRICTION: f32 = 0.2;

// Initial hit points value for each Spaceship
pub const SPACESHIP_INITIAL_HP: u32 = 500;

// linear angular acceleration of the spaceship
pub const SPACESHIP_LINEAR_ANGULAR_ACCELERATION: f32 = 0.0001;

// Maximum angular velocity added on user input
pub const SPACESHIP_MAX_ANGULAR_VELOCITY: f32 = 0.12;

// maximum speed of each Spaceship
pub const SPACESHIP_MAX_SPEED: u32 = 600;

// non linear spaceship angular acceleration
pub const SPACESHIP_NON_LINEAR_ANGULAR_ACCELERATION: f32 = 2.0;

// How often we want to send updates to clients (if there are any updates)
pub const UPDATE_SEND_INTERVAL: Duration = Duration::from_millis(10);

// The radius of the playable world (in 0.01 units)
pub const WORLD_RADIUS: f32 = 100000.0;
