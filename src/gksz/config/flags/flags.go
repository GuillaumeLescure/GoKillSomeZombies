package flags

import (
	"flag"
)

type (
	pkg struct {
		verbose uint
		color   bool
	}
)

const (
	DefaultVerboseLevel uint = 1
	DefaultColor bool = true
)

var (
	this pkg
)

func Parse() {
	flag.UintVar(&this.verbose, "v", DefaultVerboseLevel, "verbose mode (0 to 3)")
	flag.BoolVar(&this.color, "c", DefaultColor, "color mode (true or false)")
	flag.Parse()
}

func Verbose() uint {
	return this.verbose
}

func Color() bool {
	return this.color
}
