package main

import (
	"gksz/base/logs"
	"gksz/config/flags"
	"gksz/config/userfile"
)

func main() {
	flags.Parse()
	logs.SetColor(flags.Color())
	logs.SetVerbose(flags.Verbose())

	userfile.ReadOrCreateFile(userfile.DefaultPath)
}
