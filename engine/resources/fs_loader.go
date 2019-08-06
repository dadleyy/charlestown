package resources

import "io"
import "os"
import "fmt"
import "log"
import "path/filepath"

type fsLoader struct {
	root string
}

func (loader *fsLoader) Load(key string) (io.ReadCloser, error) {
	cwd, e := os.Getwd()

	if e != nil {
		return nil, e
	}

	name := filepath.Join(cwd, "assets", "thing.jpg")
	log.Printf("loading '%s'", name)
	return os.OpenFile(name, os.O_RDONLY, os.ModePerm)
}

// NewFilesystemLoader returns a resource loader that uses the filesystem.
func NewFilesystemLoader(path string) (Loader, error) {
	return &fsLoader{fmt.Sprintf("%s", path)}, nil
}
