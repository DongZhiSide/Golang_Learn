package logb

import (
	"os"
	"testing"
)

func TestBuildLogger(t *testing.T) {
	lgb := LoggerBuilder{}

	t.Log("0 =================================")
	lgb.Level = StringToLevel("0")
	l := lgb.BuildTextStdout()
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")

	t.Log("1 =================================")
	lgb.Level = StringToLevel("1")
	l = lgb.BuildTextStdout()
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")

	t.Log("0 =================================")
	lgb.Level = StringToLevel("0")
	file, err := OpenLog("/etc/test_logger_output.log")
	lgb.W = file
	l = lgb.BuildTextWriter()
	if err != nil {
		l.Warn("failed to open logger output file", "error", err)
	}
	if file != os.Stdout {
		defer file.Close()
	}
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")
}
