package main

import (
	"context"
	"golanglearn/logb"
	"log/slog"
)

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
	lgb := logb.LoggerBuilder{}
	lgb.Options.Level = slog.LevelDebug
	lgb.Options.AddSource = true

	a := Account{
		Username: "google",
		Password: "shit",
	}
	lg := lgb.BuildTextStdout()
	lg.LogAttrs(context.Background(), slog.LevelDebug, "create account", slog.Any("account", a))

	glg := lg.WithGroup("google-accounts")
	glg.LogAttrs(context.Background(), slog.LevelDebug, "create", slog.Any("account", a))

	flg, close := lgb.BuildTextFile("1.log")
	defer close()
	flg.LogAttrs(context.Background(), slog.LevelDebug, "create", slog.Any("account", a))

	lgb.Options.Level = slog.LevelDebug - 8
	lgb.Options.ReplaceAttr = logb.ReplaceTime
	lgb.Options.AddSource = false
	ntlg := lgb.BuildTextStdout()
	ntlg.LogAttrs(context.Background(), slog.LevelDebug-4, "create account", slog.Any("account", a))
	ntlg.LogAttrs(context.Background(), slog.LevelDebug-3, "create account", slog.Any("account", a))

	cmg := logb.CustomBuildLogger()
	cmg.LogAttrs(context.Background(), slog.LevelDebug-3, "CustomLoggerHandler", slog.Any("account", a))
}
