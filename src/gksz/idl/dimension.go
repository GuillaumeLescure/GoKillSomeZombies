package idl

type (
	Dimension struct {
		Width  uint
		Height uint
	}
)

//----------> Package's functions <----------

func NewDimension(newWidth uint, newHeight uint) *Dimension {
	return &Dimension{Width: newWidth, Height: newHeight}
}
