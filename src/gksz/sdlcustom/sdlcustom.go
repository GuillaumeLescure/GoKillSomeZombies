package sdlcustom

import (
	"gksz/base/logs"
	"gksz/config/userfile"
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

func CreateMainWindow() *sdl.Window {
	flags := sdl.WINDOW_SHOWN | sdl.WINDOW_INPUT_GRABBED
	if (userfile.FullScreen() == true) {
		flags |= sdl.WINDOW_FULLSCREEN
	}

	win, err := sdl.CreateWindow(WindowName, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int(userfile.ScreenResolutionWidth()), int(userfile.ScreenResolutionHeight()), uint32(flags))
	if err != nil {
		logs.Error(err)
	}

	return win
}
