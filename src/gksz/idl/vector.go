package idl

type (
	Vector struct {
		X uint
		Y uint
	}
)

// ----------> Package's functions <----------

func NewVector(newX uint, newY uint) *Vector {
	return &Vector{X: newX, Y: newY}
}

// ----------> Vector's methods <----------

func (self *Vector) Add(other *Vector) {
	self.X = self.X + other.X
	self.Y = self.Y + other.Y
}

func (self *Vector) Minus(other *Vector) {
	self.X = self.X - other.X
	self.Y = self.Y - other.Y
}
