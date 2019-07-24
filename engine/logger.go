package engine

import "os"
import "log"
import "path/filepath"

type engineLogger struct {
	*log.Logger
	file *os.File
}

func (logger *engineLogger) Close() error {
	if logger.file != nil {
		return logger.file.Close()
	}
	return nil
}

func initializeLogger(c Configuration) (*engineLogger, error) {
	wd, e := os.Getwd()

	if e != nil {
		return nil, e
	}

	dir := filepath.Dir(c.Logging.Filename)

	if !filepath.IsAbs(dir) {
		dir = filepath.Join(wd, dir)
	}

	if e := os.MkdirAll(dir, os.ModePerm); e != nil {
		return nil, e
	}

	flags := os.O_RDWR | os.O_CREATE

	if c.Logging.Truncate {
		flags = flags | os.O_TRUNC
	} else {
		flags = flags | os.O_APPEND
	}

	file, e := os.OpenFile(filepath.Join(dir, filepath.Base(c.Logging.Filename)), flags, 0644)

	if e != nil {
		return nil, e
	}

	return &engineLogger{
		Logger: log.New(file, "", log.LstdFlags),
		file:   file,
	}, nil
}
