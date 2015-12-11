package idl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	MainWindower interface {
		UiBaser

		Window() *sdl.Window
		Renderer() *sdl.Renderer
		WindowManager() *WindowManager
		Destroy()
		Update() error
	}

	MainWindow struct {
		UiBase

		win      *sdl.Window
		renderer *sdl.Renderer
		wm       WindowManager
	}
)

//----------> Package's functions <----------

func NewMainWindow(title string, x uint, y uint, w uint, h uint, flags uint32) (*MainWindow, error) {
	newWin, err := sdl.CreateWindow(title, int(x), int(y), int(w), int(h), flags)
	if err != nil {
		return nil, err
	} else {
		newRenderer, err := sdl.CreateRenderer(newWin, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return nil, err
		}
		return &MainWindow{UiBase: *NewUiBase(), win: newWin, renderer: newRenderer}, nil
	}
}

func (self *MainWindow) Window() *sdl.Window {
	return self.win
}

func (self *MainWindow) Renderer() *sdl.Renderer {
	return self.renderer
}

func (self *MainWindow) WindowManager() *WindowManager {
	return &self.wm
}

func (self *MainWindow) Destroy() {
	self.win.Destroy()
	self.renderer.Destroy()
}

func (self *MainWindow) Update() error {
	if err := self.renderer.Clear(); err != nil {
		return err
	}

	size := len(self.children)
	for c := 0; c < size; c++ {
		if err := self.children[c].Draw(); err != nil {
			return err
		}
	}

	self.renderer.Present()
	return nil
}
