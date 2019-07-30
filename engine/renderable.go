package engine

import "fmt"

type renderable struct {
	location point
	symbol   rune
}

func (r *renderable) String() string {
	return fmt.Sprintf("<'%s'@%c>", r.location, r.symbol)
}
