package logs

import (
	"github.com/cihub/seelog"
	//"fmt"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	logger,err := seelog.LoggerFromConfigAsFile("../conf/seelog.xml")
	if err != nil {
		panic(err.Error())
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}


