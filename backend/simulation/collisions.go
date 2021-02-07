package simulation

import "github.com/u2i/superstellar/backend/state"

type Collision interface {
	collide(state.Object, state.Object)
}
