package objects

import "time"

// Flash objects are used to render out banners to the user.
type Flash struct {
	Text    string
	Expires time.Time
}
