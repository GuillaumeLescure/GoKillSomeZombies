package idl

import (
	"errors"
	"strconv"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	MainWindow struct {
		*sdl.Window

		children []*WidgetDrawer
	}
)

//----------> Package's functions <----------

func NewMainWindow(title string, x uint, y uint, w uint, h uint, flags uint32) (*MainWindow, error) {
	if CanCreateMainWindow() == true {
		SetCurrentMainWindowNumber(CurrentMainWindowNumber() + 1)
		win, err := sdl.CreateWindow(title, int(x), int(y), int(w), int(h), flags)
		if err != nil {
			return nil, err
		} else {
			return &MainWindow{Window: win}, nil
		}
	}

	return nil, errors.New("limit of MainWindow are reached (max=" + strconv.Itoa(MainWindowLimit()) + ")")
}
