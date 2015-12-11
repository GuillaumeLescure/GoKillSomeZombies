package idl

type (
	WindowManager struct {
		windowsStack []*Widget
	}
)

// ----------> Package's functions <----------

func NewWindowManager() *WindowManager {
	return &WindowManager{}
}

// ----------> WindowManager's functions <----------

func (self *WindowManager) Stack() []*Widget {
	return self.windowsStack
}

func (self *WindowManager) AddWindow(win *Widget) {
	self.windowsStack = append([]*Widget{win}, self.windowsStack...)

	self.windowsStack[1].SetFocus(false)
	self.windowsStack[0].SetFocus(true)
}

func (self *WindowManager) ClickAt(point Coordinate) {
	for key, widget := range self.windowsStack {
		if widget.Visible() == true && widget.Contains(point) == true {
			self.windowsStack = append(self.windowsStack[:key], self.windowsStack[key+1:]...)
			self.AddWindow(widget)

			widget.SetFocus(true)
			widget.ClickAt(point)
			break
		}
	}
}
