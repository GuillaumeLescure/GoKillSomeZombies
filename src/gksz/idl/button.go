package idl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Button struct {
		Widget
		
		surfaceHover *sdl.Surface
		surfaceClick *sdl.Surface
	}
)

// ----------> Package's functions <----------

func NewButton(newParent *WidgetDrawer, newPos *Coordinate, newSize *Dimension, newSurfaceEnable *sdl.Surface, newSurfaceDisable *sdl.Surface, newSurfaceHover *sdl.Surface, newSurfaceClick *sdl.Surface) *Button {
	return &Button{Widget: *NewWidget(newParent, newPos, newSize, newSurfaceEnable, newSurfaceDisable), surfaceHover: newSurfaceHover, surfaceClick: newSurfaceClick}
}

// ----------> Button's methods <----------

func (self *Button) SurfaceHover() *sdl.Surface {
	return self.surfaceHover
}

func (self *Button) SetSurfaceHover(newSurfaceHover *sdl.Surface) {
	self.surfaceHover = newSurfaceHover
}

func (self *Button) SurfaceClick() *sdl.Surface {
	return self.surfaceClick
}

func (self *Button) SetSurfaceClick(newSurfaceClick *sdl.Surface) {
	self.surfaceClick = newSurfaceClick
}
