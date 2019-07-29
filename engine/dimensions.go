package engine

type dimensions struct {
	width  int
	height int
}

func (d *dimensions) midway() (int, int) {
	return d.width / 2, d.height / 2
}
