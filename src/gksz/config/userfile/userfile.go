package userfile

import (
	"gksz/config"
	"gksz/base/logs"
	"encoding/xml"
	"os"
	"path/filepath"
)

type (
	pkg struct {
		dataFromFile
	}

	dataFromFile struct {
		XMLName    xml.Name  `xml:"GoKillSomeZombies"`
		Type       string    `xml:"type,attr"`
		Version    string    `xml:"version,attr"`
		Name       string    `xml:"user>name"`
		Mail       string    `xml:"user>mail"`
		Width      uint      `xml:"preferences>screenResolution>width"`
		Height     uint      `xml:"preferences>screenResolution>height"`
		FullScreen bool      `xml:"preferences>fullScreen"`
	}
)

const (
	DefaultName string = "???"
	DefaultMail string = "???"
	DefaultWidth uint = 1024
	DefaultHeight uint = 768
	DefaultFullScreen bool = true
)

var (
	this    pkg

	DefaultPath string
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error(err)
    }
	DefaultPath = dir + "/.gksz.conf"
}

func ReadOrCreateFile(path string) {
	logs.Debug(logs.CurrentFunctionName() + "(path=" + path + ")")

	if _, err := os.Stat(path); err == nil {
		LoadFile(path)
	} else {
		WriteDefaultFile(path)
	}

	logs.Debug(logs.CurrentFunctionName() + " -- end")
}

func LoadFile(path string) {
	logs.Debug(logs.CurrentFunctionName() + "(path=" + path + ")")

	file, err := os.Open(path)
	if err != nil {
		logs.Error(err);
	}
	defer file.Close()

	err = xml.NewDecoder(file).Decode(&this.dataFromFile)
	if err != nil {
		logs.Error(err)
	}

	logs.Information("Configuration file is loaded")

	logs.Debug(logs.CurrentFunctionName() + " -- end")
}

func WriteDefaultFile(path string) {
	logs.Debug(logs.CurrentFunctionName() + "(path=" + path + ")")

	file, err := os.Create(path)
	if err != nil {
		logs.Error(err);
	}
	defer file.Close()

	this.dataFromFile = dataFromFile{XMLName:xml.Name{" ", "GoKillSomeZombies"}, Type:"userConfFile", Version:config.Version(), Name:DefaultName, Mail:DefaultMail, Width:DefaultWidth, Height:DefaultHeight, FullScreen:DefaultFullScreen}
	xmlToWrite, err := xml.MarshalIndent(this.dataFromFile, "", "    ")
	if err != nil {
		logs.Error(err)
	}

	_, err = file.Write(xmlToWrite)
	if err != nil {
		logs.Error(err);
	}

	logs.Information("Default configuration file is written at '" + path + "'")

	logs.Debug(logs.CurrentFunctionName() + " -- end")
}

func Name() string {
	return this.dataFromFile.Name
}

func Mail() string {
	return this.dataFromFile.Mail
}

func ScreenResolutionWidth() uint {
	return this.dataFromFile.Width
}

func ScreenResolutionHeight() uint {
	return this.dataFromFile.Height
}

func FullScreen() bool {
	return this.dataFromFile.FullScreen
}
