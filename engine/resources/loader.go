package resources

import "io"

import (
	// support jpeg image
	_ "image/jpeg"
	// support png image
	_ "image/png"
)

// Loader defines an interface for loading resources from the filesystem or cloud.
type Loader interface {
	Load(string) (io.ReadCloser, error)
}
