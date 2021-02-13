use std::time::Duration;

// Time between physics frames to be recomputed. We aim for 50 fps.
pub const PHYSICS_FRAME_DURATION: Duration = Duration::from_millis(20);

// How often we want to send updates to clients (if there are any updates)
pub const UPDATE_SEND_INTERVAL: Duration = Duration::from_millis(10);
