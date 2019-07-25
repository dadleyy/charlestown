package main

import "log"
import "flag"
import "github.com/joho/godotenv"
import "github.com/dadleyy/charlestown/engine"

type cliOptions struct {
}

func main() {
	godotenv.Load()
	config := engine.Configuration{}

	flag.StringVar(&config.Logging.Filename, "logfile", "log/charlestown.log", "decide where the logs go")
	flag.BoolVar(&config.Logging.Truncate, "truncate-log", true, "truncate the log file when opening")
	flag.Parse()

	if e := engine.Run(config); e != nil {
		log.Printf("[error] engine error %s", e)
	}
}
