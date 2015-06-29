package idl

type (
	Coordinate struct {
		X uint
		Y uint
	}
)

// ----------> Package's functions <----------

func NewCoordinate(newX uint, newY uint) *Coordinate {
	return &Coordinate{X: newX, Y: newY}
}

// ----------> Coordinate's methods <----------

func (self *Coordinate) Add(other *Coordinate) {
	self.X = self.X + other.X
	self.Y = self.Y + other.Y
}

func (self *Coordinate) Minus(other *Coordinate) {
	self.X = self.X - other.X
	self.Y = self.Y - other.Y
}
