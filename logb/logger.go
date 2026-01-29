package logb

import (
	"io"
	"log/slog"
	"os"
)

type LoggerBuilder struct {
	Options slog.HandlerOptions
}

func (lb *LoggerBuilder) BuildTextStdout() *slog.Logger {
	return lb.BuildText(os.Stdout)
}

func (lb *LoggerBuilder) BuildText(w io.Writer) *slog.Logger {
	opts := lb.Options
	return slog.New(slog.NewTextHandler(w, &opts))
}

func (lb *LoggerBuilder) BuildJson(w io.Writer) *slog.Logger {
	opts := lb.Options
	return slog.New(slog.NewJSONHandler(w, &opts))
}

type Close = func() error

func getCloser(w *os.File) Close {
	return func() error {
		err := error(nil)
		if w != nil && w != os.Stdout {
			err = w.Close()
		}
		return err
	}
}

func (lb *LoggerBuilder) BuildTextFile(path string) (*slog.Logger, Close) {
	w, err := OpenLog(path)
	lg := lb.BuildText(w)
	if err != nil {
		lg.Warn("failed to open logger output file, defaulting to stdout", "error", err)
	}
	return lg, getCloser(w)
}

func (lb *LoggerBuilder) BuildJsonFile(path string) (*slog.Logger, Close) {
	w, err := OpenLog(path)
	lg := lb.BuildJson(w)
	if err != nil {
		lg.Warn("failed to open logger output file, defaulting to stdout", "error", err)
	}
	return lg, getCloser(w)
}

// if path is empty or can not open path,
// the ouput will be os.Stdout.
func BuildTextFile(level slog.Level, path string) (*slog.Logger, Close) {
	lb := LoggerBuilder{}
	lb.Options.Level = level
	return lb.BuildTextFile(path)
}

// if path is empty, the *os.File will be os.Stdout,
// and error will be nil.
func OpenLog(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	}
	w, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		w = os.Stdout
	}
	return w, err
}

func StringToLevel(level string) slog.Level {
	switch level {
	case "0", "debug":
		return slog.LevelDebug
	case "1", "info":
		return slog.LevelInfo
	case "2", "warn":
		return slog.LevelWarn
	case "3", "error":
		return slog.LevelError
	default:
		return slog.LevelWarn
	}
}
