package mutations

import "time"
import "github.com/dadleyy/charlestown/engine/objects"

// Expire will iterate over objects, clearing them if they are expired.
func Expire() Mutation {
	return expire{}
}

type expire struct {
}

func (mut expire) Apply(game objects.Game) objects.Game {
	next := game
	now := time.Now()
	flashes := make([]objects.Flash, 0, len(next.Flashes))

	for _, f := range next.Flashes {
		if expired := now.After(f.Expires); !expired {
			flashes = append(flashes, f)
		}
	}

	next.Flashes = flashes

	return next
}
