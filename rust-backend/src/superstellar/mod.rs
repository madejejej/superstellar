use std::ops::{Add, AddAssign, Mul};
// Include the `items` module, which is generated from items.proto.
include!(concat!(env!("OUT_DIR"), "/superstellar.rs"));
pub mod constants;
pub mod entities;

impl Point {
    pub fn new(x: i32, y: i32) -> Point {
        Point { x, y }
    }
}

impl Add<Point> for Point {
    type Output = Self;

    fn add(self, other: Self) -> Self {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}
impl Add<Vector> for Point {
    type Output = Self;

    fn add(self, other: Vector) -> Self {
        Self {
            x: self.x + other.x as i32,
            y: self.y + other.y as i32,
        }
    }
}

impl AddAssign<Point> for Point {
    fn add_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}

impl AddAssign<&Vector> for Point {
    fn add_assign(&mut self, other: &Vector) {
        *self = Self {
            x: self.x + other.x as i32,
            y: self.y + other.y as i32,
        };
    }
}

impl Vector {
    pub fn new(x: f32, y: f32) -> Vector {
        Vector { x, y }
    }
}

impl Add for Vector {
    type Output = Self;

    fn add(self, other: Self) -> Self {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl AddAssign<&Vector> for Vector {
    fn add_assign(&mut self, other: &Vector) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}

impl Mul<f32> for Vector {
    type Output = Self;

    fn mul(self, rhs: f32) -> Self {
        Self {
            x: self.x * rhs,
            y: self.y * rhs,
        }
    }
}
