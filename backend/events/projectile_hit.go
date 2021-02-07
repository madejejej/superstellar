package events

import "github.com/u2i/superstellar/backend/state"

type ProjectileHit struct {
	Projectile *state.Projectile
}
