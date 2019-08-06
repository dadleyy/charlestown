package engine

// Configuration represents all of the available options provided before booting the application.
type Configuration struct {
	Logging struct {
		Filename string
		Truncate bool
	}
	AssetPath string
}
