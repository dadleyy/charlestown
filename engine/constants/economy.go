package constants

import "time"

const (
	// EconomyMultiplier offsets building costs and income to make more sense.
	EconomyMultiplier = 1000

	// BaseTurnCardinality the amount of turns offered regardless of pop.
	BaseTurnCardinality = 4

	// TurnDuration is how many seconds turns are.
	TurnDuration = time.Second * 10
)
