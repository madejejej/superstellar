package events

import "github.com/u2i/superstellar/backend/state"

type ProjectileFired struct {
	Projectile *state.Projectile
}
