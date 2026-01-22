package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func BuildLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	hd := slog.NewTextHandler(os.Stdout, opts)
	lg := slog.New(hd)
	return lg
}

func FileBuildLogger(file string) (*slog.Logger, func(), error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return nil, nil, err
	}
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	hd := slog.NewTextHandler(f, opts)
	lg := slog.New(hd)
	return lg, func() { f.Close() }, nil
}

func NoTimeBuildLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug - 8,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			if a.Key == slog.LevelKey {
				if lvl, ok := a.Value.Any().(slog.Level); ok && lvl == -8 {
					a.Value = slog.StringValue("TRACE")
				}
			}
			return a
		},
	}
	hd := slog.NewTextHandler(os.Stdout, opts)
	lg := slog.New(hd)
	return lg
}

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

type Account struct {
	Username string
	Password string
}

func (a Account) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("username", a.Username),
		slog.String("password", "You will never know"),
	)
}

func (a Account) String() string {
	return "account string"
}

func main() {
	a := Account{
		Username: "google",
		Password: "shit",
	}
	lg := BuildLogger()
	lg.LogAttrs(context.Background(), slog.LevelDebug, "create account", slog.Any("account", a))

	glg := lg.WithGroup("google-accounts")
	glg.LogAttrs(context.Background(), slog.LevelDebug, "create", slog.Any("account", a))

	flg, close, err := FileBuildLogger("1.log")
	if err != nil {
		glg.LogAttrs(context.Background(), slog.LevelError, "can not open file", slog.Any("error", err))
		os.Exit(1)
	}
	defer close()
	flg.LogAttrs(context.Background(), slog.LevelDebug, "create", slog.Any("account", a))

	ntlg := NoTimeBuildLogger()
	ntlg.LogAttrs(context.Background(), slog.LevelDebug-4, "create account", slog.Any("account", a))
	ntlg.LogAttrs(context.Background(), slog.LevelDebug-3, "create account", slog.Any("account", a))

	cmg := CustomBuildLogger()
	cmg.LogAttrs(context.Background(), slog.LevelDebug-3, "CustomLoggerHandler", slog.Any("account", a))
}
