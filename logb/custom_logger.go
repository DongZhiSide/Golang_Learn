package logb

import (
	"context"
	"fmt"
	"log/slog"
)

type CustomLoggerHandler struct{}

func (CustomLoggerHandler) Enabled(context.Context, slog.Level) bool { return true }
func (CustomLoggerHandler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Print(r.Level.String()+" ", r.Message+" ", r.Time.String()+" ")
	r.Attrs(func(a slog.Attr) bool {
		fmt.Print(a.String())
		return true
	})
	return nil
}
func (clh CustomLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return clh }
func (clh CustomLoggerHandler) WithGroup(name string) slog.Handler       { return clh }

func CustomBuildLogger() *slog.Logger {
	hd := CustomLoggerHandler{}
	lg := slog.New(hd)
	return lg
}
