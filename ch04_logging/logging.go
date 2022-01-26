package logging

import (
	"log"
	"os"
)

var (
	info  *log.Logger
	debug *log.Logger
)

func init() {
	info = log.New(os.Stderr, "INFO\t", log.LstdFlags)
	debug = log.New(os.Stderr, "DEBUG\t", log.LstdFlags)
}
