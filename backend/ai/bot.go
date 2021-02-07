package ai

import "github.com/u2i/superstellar/backend/state"

type Bot interface {
	HandleStateUpdate(space *state.Space, spaceship *state.Spaceship)
}
