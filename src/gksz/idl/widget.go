package idl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Focuser interface {
		SetFocus(bool)
		Focus() bool
	}

	Enabler interface {
		SetEnable(bool)
		Enable() bool
	}

	Visibler interface {
		SetVisible(bool)
		Visible() bool
	}

	Drawer interface {
		Draw() error
	}

	WidgetDrawer interface {
		Focuser
		Enabler
		Visibler
		Drawer
	}

	Widget struct {
		WidgetDrawer

		parent   *WidgetDrawer
		children []*Widget

		position Coordinate
		size     Dimension
		focus    bool
		enable   bool
		visible  bool

		surfaceEnable  *sdl.Surface
		surfaceDisable *sdl.Surface
		currentSurface *sdl.Surface
	}
)

const (
	DefaultChildrenSize uint = 3
	DefaultFocus        bool = false
	DefaultEnable       bool = true
	DefaultVisible      bool = true
)

// ----------> Package's functions <----------

func NewWidget(newParent *WidgetDrawer, newPos *Coordinate, newSize *Dimension, newSurfaceEnable *sdl.Surface, newSurfaceDisable *sdl.Surface) *Widget {
	widget := &Widget{parent: newParent, position: *newPos, size: *newSize, surfaceEnable: newSurfaceEnable, surfaceDisable: newSurfaceDisable}
	widget.children = make([]*Widget, 0, DefaultChildrenSize)
	widget.SetFocus(DefaultFocus)
	widget.SetEnable(DefaultEnable)
	widget.SetVisible(DefaultVisible)
	return widget
}

// ----------> Widget's methods <----------

func (self *Widget) Parent() *WidgetDrawer {
	return self.parent
}

func (self *Widget) SetParent(newParent *WidgetDrawer) {
	self.parent = newParent
}

func (self *Widget) Position() *Coordinate {
	return &self.position
}

func (self *Widget) MoveTo(newCoordinates *Coordinate) {
	for _, child := range self.children {
		child.Move(&Vector{X: self.position.X - newCoordinates.X, Y: self.position.Y - newCoordinates.Y})
	}

	self.position.X = newCoordinates.X
	self.position.Y = newCoordinates.Y
}

func (self *Widget) Move(newVector *Vector) {
	for _, child := range self.children {
		child.Move(newVector)
	}

	self.position.X += newVector.X
	self.position.Y += newVector.Y
}

func (self *Widget) SurfaceEnable() *sdl.Surface {
	return self.surfaceEnable
}

func (self *Widget) SetSurfaceEnable(newSurfaceEnable *sdl.Surface) {
	self.surfaceEnable = newSurfaceEnable
}

func (self *Widget) SurfaceDisable() *sdl.Surface {
	return self.surfaceDisable
}

func (self *Widget) SetSurfaceDisable(newSurfaceDisable *sdl.Surface) {
	self.surfaceDisable = newSurfaceDisable
}

// ----------> WidgetDrawer's methods <----------

func (self *Widget) Focus() bool {
	return self.focus
}

func (self *Widget) SetFocus(newFocus bool) {
	self.focus = newFocus
}

func (self *Widget) Enable() bool {
	return self.enable
}

func (self *Widget) SetEnable(newEnable bool) {
	self.enable = newEnable
	if self.enable == true {
		self.currentSurface = self.surfaceEnable
	} else {
		self.currentSurface = self.surfaceDisable
	}
}

func (self *Widget) Visible() bool {
	return self.visible
}

func (self *Widget) SetVisible(newVisible bool) {
	self.visible = newVisible
	if self.visible == true {
		if self.enable == true {
			self.currentSurface = self.surfaceEnable
		} else {
			self.currentSurface = self.surfaceDisable
		}
	} else {
		self.currentSurface = nil
	}
}

func (self *Widget) Draw(surface *sdl.Surface) error {
	return self.currentSurface.Blit(nil, surface,
		&sdl.Rect{X: int32(self.position.X), Y: int32(self.position.Y), W: int32(self.position.X + self.size.Width), H: int32(self.position.Y + self.size.Height)})
}
