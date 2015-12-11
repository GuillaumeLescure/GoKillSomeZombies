package idl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Button struct {
		Widget
		
		textureHover *sdl.Texture
		textureClick *sdl.Texture
	}
)

// ----------> Package's functions <----------

func NewButton(newParent UiBaser, newPos *Coordinate, newSize *Dimension, newTextureEnable *sdl.Texture, newTextureDisable *sdl.Texture, newTextureHover *sdl.Texture, newTextureClick *sdl.Texture) *Button {
	return &Button{Widget: *NewWidget(newParent, newPos, newSize, newTextureEnable, newTextureDisable), textureHover: newTextureHover, textureClick: newTextureClick}
}

// ----------> Button's methods <----------

func (self *Button) TextureHover() *sdl.Texture {
	return self.textureHover
}

func (self *Button) SetTextureHover(newTextureHover *sdl.Texture) {
	self.textureHover = newTextureHover
}

func (self *Button) TextureClick() *sdl.Texture {
	return self.textureClick
}

func (self *Button) SetTextureClick(newTextureClick *sdl.Texture) {
	self.textureClick = newTextureClick
}
