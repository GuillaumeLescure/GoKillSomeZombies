package main

import (
	"github.com/veandco/go-sdl2/sdl"
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

	err := win.Update()
	if err != nil {
		logs.Error(err)
	}

	for game.IsRunning == true {
		event := sdl.WaitEvent()
		logs.Debug("event received: ", logs.VarDetails(event))

		switch event.(type) {
		case *sdl.QuitEvent:
			game.IsRunning = false
		}
	}
}
