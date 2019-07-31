package objects

import "fmt"

// Dimensions wraps a width and height.
type Dimensions struct {
	Width  int
	Height int
}

// Midway returns the x and y values of the width and height divided by two.
func (d *Dimensions) Midway() (int, int) {
	return d.Width / 2, d.Height / 2
}

// String returns a human readble interpretation of the dimensions.
func (d *Dimensions) String() string {
	return fmt.Sprintf("[%dx%d]", d.Width, d.Height)
}

// Dup returns a fresh copy of the dimensions.
func (d *Dimensions) Dup() Dimensions {
	return Dimensions{Width: d.Width, Height: d.Height}
}

// Area returns the product of the width and height.
func (d *Dimensions) Area() int {
	return d.Width * d.Height
}
