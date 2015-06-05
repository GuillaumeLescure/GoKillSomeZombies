package logs

import (
	"fmt"
	"gksz/base/colors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

const (
	logDebug       = "[D]"
	logInformation = "[I]"
	logWarning     = "[W]"
	logError       = "[E]"
)

type self struct {
	color   bool
	verbose uint

	loggerDebug *log.Logger
	loggerInformation *log.Logger
	loggerWarning *log.Logger
	loggerError *log.Logger

	coloredLoggerDebug *log.Logger
	coloredLoggerInformation *log.Logger
	coloredLoggerWarning *log.Logger
	coloredLoggerError *log.Logger
}

var (
	this self
)

func init() {
	this.loggerDebug = log.New(os.Stdout, logDebug + " ", log.Ldate|log.Ltime)
	this.loggerInformation = log.New(os.Stdout, logInformation + " ", log.Ldate|log.Ltime)
	this.loggerWarning = log.New(os.Stdout, logWarning + " ", log.Ldate|log.Ltime)
	this.loggerError = log.New(os.Stderr, logError + " ", log.Ldate|log.Ltime)

	this.coloredLoggerDebug = log.New(os.Stdout, colors.StringColored(colors.Bold, colors.White, colors.Black, logDebug) + " ", log.Ldate|log.Ltime)
	this.coloredLoggerInformation = log.New(os.Stdout, colors.StringColored(colors.Bold, colors.Green, colors.Black, logInformation) + " ", log.Ldate|log.Ltime)
	this.coloredLoggerWarning = log.New(os.Stdout, colors.StringColored(colors.Bold, colors.Orange, colors.Black, logWarning) + " ", log.Ldate|log.Ltime)
	this.coloredLoggerError = log.New(os.Stderr, colors.StringColored(colors.Bold, colors.Red, colors.Black, logError) + " ", log.Ldate|log.Ltime)

}

func Color() bool {
	return this.color
}

func SetColor(color bool) {
	this.color = color
}

func Verbose() uint {
	return this.verbose
}

func SetVerbose(verbose uint) {
	this.verbose = verbose
}

func Debug(msg ...interface{}) {
	if this.verbose > 2 {
		if this.color == true {
			this.coloredLoggerDebug.Println(fileLineAndMessage(msg...))
		} else {
			this.loggerDebug.Println(fileLineAndMessage(msg...))
		}
	}
}

func Information(msg ...interface{}) {
	if this.verbose > 1 {
		if this.color == true {
			this.coloredLoggerInformation.Println(fileLineAndMessage(msg...))
		} else {
			this.loggerInformation.Println(fileLineAndMessage(msg...))
		}
	}
}

func Warning(msg ...interface{}) {
	if this.verbose > 0 {
		if this.color == true {
			this.coloredLoggerWarning.Println(fileLineAndMessage(msg...))
		} else {
			this.loggerWarning.Println(fileLineAndMessage(msg...))
		}
	}
}

func Error(msg ...interface{}) {
	if this.color == true {
		this.coloredLoggerError.Println(fileLineAndMessage(msg...))
	} else {
		this.loggerError.Println(fileLineAndMessage(msg...))
	}
	debug.PrintStack()
	os.Exit(1)
}

func CurrentFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	return filepath.Base(runtime.FuncForPC(pc).Name())
}

func fileLineAndMessage(msg ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fmt.Sprint(msg...))
}
