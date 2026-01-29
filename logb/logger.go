package logb

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type LoggerBuilder struct {
	slog.HandlerOptions
	W io.Writer
}

func (lb LoggerBuilder) BuildTextStdout() *slog.Logger {
	opts := lb.HandlerOptions
	return slog.New(slog.NewTextHandler(os.Stdout, &opts))
}
func (lb LoggerBuilder) BuildJsonStdout() *slog.Logger {
	opts := lb.HandlerOptions
	return slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}

func (lb LoggerBuilder) BuildTextWriter() *slog.Logger {
	opts := lb.HandlerOptions
	return slog.New(slog.NewTextHandler(lb.W, &opts))
}

func (lb LoggerBuilder) BuildJsonWriter() *slog.Logger {
	opts := lb.HandlerOptions
	return slog.New(slog.NewJSONHandler(lb.W, &opts))
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

// if path is empty, the *os.File will be os.Stdout,
// and error will be nil.
func OpenLog(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	}
	w, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		w = os.Stdout
		err = fmt.Errorf("failed to open logger output file, defaulting to stdout: error: %w", err)
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
