package flag

import (
	"os"
)

var defaultFlags = []Flag{
	// HelpFlag prints usage of application.
	&BoolFlag{
		Name:    "help",
		Usage:   "--help, show help information",
		Default: false,
		Action: func(name string, fs *FlagSet) {
			if fs.Bool(name) {
				fs.PrintDefaults()
				os.Exit(0)
			}

		},
	},
}
