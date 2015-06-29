package idl

type (
	pkgWM struct {
		mainWindowLimit         int
		currentMainWindowNumber int
	}
)

const (
	DefaultMainWindowLimit int = 1
	UnlimitedMainWindow int = -1
)

var (
	thisWM pkgWM
)

//----------> Package's functions <----------

func init() {
	thisWM.mainWindowLimit = DefaultMainWindowLimit
}

func MainWindowLimit() int {
	return thisWM.mainWindowLimit
}

func SetMainWindowLimit(newMainWindowLimit int) {
	thisWM.mainWindowLimit = newMainWindowLimit
}

func CurrentMainWindowNumber() int {
	return thisWM.currentMainWindowNumber
}

func SetCurrentMainWindowNumber(newCurrentMainWindowNumber int) {
	thisWM.currentMainWindowNumber = newCurrentMainWindowNumber
}

func CanCreateMainWindow() bool {
	return thisWM.currentMainWindowNumber < thisWM.mainWindowLimit
}
