package engine

import "fmt"

type dimensions struct {
	width  int
	height int
}

func (d *dimensions) midway() (int, int) {
	return d.width / 2, d.height / 2
}

func (d *dimensions) String() string {
	return fmt.Sprintf("[%dx%d]", d.width, d.height)
}
