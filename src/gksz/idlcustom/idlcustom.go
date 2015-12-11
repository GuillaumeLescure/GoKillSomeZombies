package idlcustom

import (
	"github.com/veandco/go-sdl2/sdl"
	"gksz/base/logs"
	"gksz/config/userfile"
	"gksz/idl"
)

const (
	WindowName string = "GoKillSomeZombies"
)

func InitSDL() {
	logs.Debug(logs.CurrentFunctionName() + "()")

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		logs.Error(err)
	}

	logs.Debug(logs.CurrentFunctionName() + " -- end")
}

func CreateMainWindow() *idl.MainWindow {
	logs.Debug(logs.CurrentFunctionName() + "()")

	flags := uint32(sdl.WINDOW_SHOWN/* | sdl.WINDOW_INPUT_GRABBED*/)
	if userfile.FullScreen() == true {
//		flags |= sdl.WINDOW_FULLSCREEN
	}

	win, err := idl.NewMainWindow(WindowName, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, userfile.ScreenResolutionWidth(), userfile.ScreenResolutionHeight(), flags)
	if err != nil {
		logs.Error(err)
	}

	logs.Debug(logs.CurrentFunctionName() + " -- end with " + logs.VarDetails(win))
	return win
}
