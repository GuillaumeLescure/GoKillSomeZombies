package idl

type (
	UiBaser interface {
		Children() []Widgeter
		ToMainWindower() MainWindower
		ToWidgeter() Widgeter
	}

	UiBase struct {
		UiBaser

		children []Widgeter
	}
)

const (
	DefaultChildrenSize uint = 3
)

// ----------> Package's functions <----------

func NewUiBase() *UiBase {
	return &UiBase{children: make([]Widgeter, 0, DefaultChildrenSize)}
}

// ----------> UiBaser's methods <----------

func (self *UiBase) Children() []Widgeter {
	return self.children
}

func (self *UiBase) ToMainWindower() MainWindower {
	if res, ok := UiBaser(self).(*MainWindow); ok {
		return res
	}
	
	return nil
}

func (self *UiBase)ToWidgeter() Widgeter {
	if res, ok := UiBaser(self).(*Widget); ok {
		return res
	}
	
	return nil
}