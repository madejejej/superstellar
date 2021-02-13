impl crate::superstellar::Space {
    pub fn new() -> Self {
        Self {
            physics_frame_id: 0,
            spaceships: vec![],
            asteroids: vec![],
        }
    }

    pub fn update(&mut self) {
        self.physics_frame_id += 1;
    }
}
