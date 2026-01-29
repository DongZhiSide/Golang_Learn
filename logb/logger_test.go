package logb

import (
	"testing"
)

func TestBuildLogger(t *testing.T) {
	l, c := BuildTextFile(StringToLevel("0"), "")
	defer c()
	t.Log("0 =================================")
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")

	t.Log("1 =================================")
	l, c = BuildTextFile(StringToLevel("1"), "")
	defer c()
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")

	t.Log("0 =================================")
	l, c = BuildTextFile(StringToLevel("0"), "./test_logger_output.log")
	defer c()
	l.Debug("debug level log")
	l.Info("info level log")
	l.Warn("warn level log")
	l.Error("error level log")
}
