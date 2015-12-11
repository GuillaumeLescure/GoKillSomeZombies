package config

import (
	"strconv"
)

type (
	pkgV struct {
		version string
	}
)

const (
	MajorVersion uint = 0
	MinorVersion uint = 1
	BuildVersion uint = 0
)

var (
	this pkgV
)

func init() {
	this.version = strconv.FormatInt((int64)(MajorVersion), 10) + "." +
		strconv.FormatInt((int64)(MinorVersion), 10) + "-" +
		strconv.FormatInt((int64)(BuildVersion), 10)
}

func Version() string {
	return this.version
}
