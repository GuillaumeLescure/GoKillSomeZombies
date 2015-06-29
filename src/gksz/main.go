package main

import (
	"gksz/base/logs"
	"gksz/config/flags"
	"gksz/config/userfile"
	"gksz/game"
	"gksz/idlcustom"
)

func main() {
	flags.Parse()
	logs.SetColor(flags.Color())
	logs.SetVerbose(flags.Verbose())

	userfile.ReadOrCreateFile(userfile.DefaultPath)

	idlcustom.InitSDL()
	defer sdl.Quit()

	win := idlcustom.CreateMainWindow()
	defer win.Destroy()
	for game.IsRunning == true {
		event := sdl.WaitEvent()
		logs.Debug("event received: ", logs.VarDetails(event))

		switch event.(type) {
			case *sdl.QuitEvent :
				game.IsRunning = false;
		}
	}
}
