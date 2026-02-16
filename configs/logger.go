package configs

import (
	"io"
	"log/slog"
	"time"
)

func NewLogger(out io.Writer, minLevel slog.Level) *slog.Logger {
	log := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
		AddSource:   false,
		Level:       minLevel,
		ReplaceAttr: replaceTimeFormat,
	}))
	return log
}

func replaceTimeFormat(group []string, a slog.Attr) slog.Attr {
	if a.Key == "time" {
		value := time.Now().Format("2006-01-03T15:04:05")
		return slog.Attr{Key: a.Key, Value: slog.StringValue(value)}
	}
	return slog.Attr{Key: a.Key, Value: a.Value}
}
