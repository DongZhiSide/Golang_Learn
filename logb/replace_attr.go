package logb

import "log/slog"

func ReplaceTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}
	if a.Key == slog.LevelKey {
		if lvl, ok := a.Value.Any().(slog.Level); ok {
			a.Value = slog.StringValue(IntLevelToString(lvl))
		}
	}
	return a
}

func IntLevelToString(l slog.Level) string {
	switch l {
	case slog.LevelDebug - 3:
		return "NOTRACE"
	case slog.LevelDebug - 4:
		return "TRACE"
	}
	return l.String()
}
