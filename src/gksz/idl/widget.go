package idl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Widgeter interface {
		MainWindow() MainWindower
		Parent() UiBaser
		SetParent(UiBaser)
		Position() *Coordinate
		MoveTo(*Coordinate)
		Move(*Vector)
		TextureEnable() *sdl.Texture
		SetTextureEnable(*sdl.Texture)
		TextureDisable() *sdl.Texture
		SetTextureDisable(*sdl.Texture)
		Focus() bool
		SetFocus(bool)
		Enable() bool
		SetEnable(bool)
		Visible() bool
		SetVisible(bool)
		Contains(Coordinate) bool
		ClickAt(Coordinate)
		Draw() error
	}

	Widget struct {
		UiBase

		mainWindow MainWindower
		parent   UiBaser
		position Coordinate
		size     Dimension
		focus    bool
		enable   bool
		visible  bool
		textureEnable  *sdl.Texture
		textureDisable *sdl.Texture
		currentTexture *sdl.Texture
	}
)

const (
	DefaultFocus        bool = false
	DefaultEnable       bool = true
	DefaultVisible      bool = true
)


// ----------> Package's functions <----------

func NewWidget(newParent UiBaser, newPos *Coordinate, newSize *Dimension, newTextureEnable *sdl.Texture, newTextureDisable *sdl.Texture) *Widget {
	widget := &Widget{UiBase: *NewUiBase(), parent: newParent, position: *newPos, size: *newSize, textureEnable: newTextureEnable, textureDisable: newTextureDisable}
	widget.SetFocus(DefaultFocus)
	widget.SetEnable(DefaultEnable)
	widget.SetVisible(DefaultVisible)
	widget.mainWindow = searchMainWindower(newParent)

	widget.mainWindow.WindowManager().AddWindow(widget)

	return widget
}

func searchMainWindower(base UiBaser) MainWindower {
	for true {
		if win, ok := base.(*MainWindow); ok == true {
			return win
		} else if wid, ok := base.(*Widget); ok == true {
			base = wid.Parent()
		} else {
			panic("base is not a MainWindow nor a Widget")
		}
	}
	panic("MainWindow not found")
}


// ----------> Widgeter's methods <----------

func (self *Widget) MainWindow() MainWindower {
	return self.mainWindow
}

func (self *Widget) Parent() UiBaser {
	return self.parent
}

func (self *Widget) SetParent(newParent UiBaser) {
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

func (self *Widget) TextureEnable() *sdl.Texture {
	return self.textureEnable
}

func (self *Widget) SetTextureEnable(newTextureEnable *sdl.Texture) {
	self.textureEnable = newTextureEnable
}

func (self *Widget) TextureDisable() *sdl.Texture {
	return self.textureDisable
}

func (self *Widget) SetTextureDisable(newTextureDisable *sdl.Texture) {
	self.textureDisable = newTextureDisable
}

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
		self.currentTexture = self.textureEnable
	} else {
		self.currentTexture = self.textureDisable
	}
}

func (self *Widget) Visible() bool {
	return self.visible
}

func (self *Widget) SetVisible(newVisible bool) {
	self.visible = newVisible
	if self.visible == true {
		if self.enable == true {
			self.currentTexture = self.textureEnable
		} else {
			self.currentTexture = self.textureDisable
		}
	} else {
		self.currentTexture = nil
	}
}

func (self *Widget) Contains(point Coordinate) bool {
	return point.X >= self.position.X && point.X <= (self.position.X + self.size.Width) &&
		point.Y >= self.position.Y && point.Y <= (self.position.Y + self.size.Height)
}

func (self *Widget) ClickAt(point Coordinate) {
	
}

func (self *Widget) Draw() error {
	return self.mainWindow.Renderer().Copy(self.currentTexture, nil,
		&sdl.Rect{X: int32(self.position.X), Y: int32(self.position.Y), W: int32(self.position.X + self.size.Width), H: int32(self.position.Y + self.size.Height)})
}
