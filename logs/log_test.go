package logs

import (
	"testing"
)

func TestLog(t *testing.T){
	Logger.Trace("trace ...")
	Logger.Debug("debug ...")
	Logger.Info("info ...")
	Logger.Warn("warn ...")
	Logger.Error("error ...")
	Logger.Critical("critical ...")
}