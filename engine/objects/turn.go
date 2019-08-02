package objects

import "time"
import "github.com/dadleyy/charlestown/engine/constants"

// TurnActions store how many actions a user has taken this turn and how many they are allowed to.
type TurnActions struct {
	Used      int
	Available int
}

// Turn holds the information about the current turn.
type Turn struct {
	Actions TurnActions
	Start   time.Time
}

// Inc consumes an available action.
func (t *Turn) Inc() Turn {
	used, avail := t.Actions.Used, t.Actions.Available
	return Turn{TurnActions{used + 1, avail}, t.Start}
}

// Done is true when the player has used all of the actions available to them.
func (t *Turn) Done() bool {
	return t.Actions.Used == t.Actions.Available
}

// Progress returns two 0-1 floating point numbers representing the completeness of the turn.
func (t *Turn) Progress() (float64, float64) {
	actions := float64(t.Actions.Used) / float64(t.Actions.Available)
	timing := t.TimeElapsed().Seconds() / constants.TurnDuration.Seconds()
	return actions, timing
}

// TimeElapsed calculates how long this turn has been going on.
func (t *Turn) TimeElapsed() time.Duration {
	return time.Now().Sub(t.Start)
}
