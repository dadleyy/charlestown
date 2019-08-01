package constants

import "time"

const (
	// IdleDelay defines how often we want to send idle updates that are ultimately indended to trigger updates.
	IdleDelay = time.Millisecond * 50

	// IncomeDelay is how often an income mutaion gets piped into the update queue.
	IncomeDelay = time.Second
)
