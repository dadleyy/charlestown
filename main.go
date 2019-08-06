package main

import "os"
import "fmt"
import "flag"
import "github.com/joho/godotenv"
import "github.com/dadleyy/charlestown/engine"
import "github.com/dadleyy/charlestown/engine/constants"

func main() {
	godotenv.Load()
	config := engine.Configuration{}
	options := flag.NewFlagSet("run", flag.ContinueOnError)
	printVersion := false

	options.Usage = func() {}

	options.StringVar(&config.AssetPath, "assets", "assets", "a filepath for where the image assets are.")
	options.StringVar(&config.Logging.Filename, "logfile", "log/charlestown.log", "decide where the logs go.")
	options.BoolVar(&config.Logging.Truncate, "truncate-log", true, "truncate the log file when opening.")
	options.BoolVar(&printVersion, "version", false, "print the version number.")

	if e := options.Parse(os.Args[1:]); e == flag.ErrHelp || printVersion {
		fmt.Printf("charlestown (version %s)\nUsage:\n", constants.AppVersion)
		options.PrintDefaults()
		return
	}

	if e := engine.Run(config); e != nil {
		fmt.Printf("[error] engine error %s\n", e)
	}
}
