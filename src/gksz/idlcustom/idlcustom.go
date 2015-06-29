package idlcustom

import (
	"gksz/base/logs"
	"gksz/config/userfile"
	"gksz/idl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowName string  = "GoKillSomeZombies"
)

func InitSDL() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if (err != nil) {
		logs.Error(err)
	}
}

func CreateMainWindow() *idl.MainWindow {
	flags := uint32(sdl.WINDOW_SHOWN | sdl.WINDOW_INPUT_GRABBED)
	if (userfile.FullScreen() == true) {
//		flags |= sdl.WINDOW_FULLSCREEN
	}

	win, err := idl.NewMainWindow(WindowName, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, userfile.ScreenResolutionWidth(), userfile.ScreenResolutionHeight(), flags)
	if err != nil {
		logs.Error(err)
	}

	return win
}
