package engine

type Configuration struct {
	Logging struct {
		Filename string
		Truncate bool
	}
}
