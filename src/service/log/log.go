package log

import (
	"context"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func newLoggerEntry(local bool) *logrus.Entry {
	logger := logrus.StandardLogger()
	if local {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})

		entry := logrus.NewEntry(logger)

		return entry
	}

	logger.SetFormatter(stackdriver.NewFormatter())

	return logrus.NewEntry(logger)
}

var (
	L = newLoggerEntry(true)
)

type logKey struct{}

func InitLog(local bool) {
	L = newLoggerEntry(local)
}

func FromContext(ctx context.Context) *logrus.Entry {
	entry := ctx.Value(logKey{})
	if entry == nil {
		return L
	}

	return entry.(*logrus.Entry)
}

func WithLogger(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, logKey{}, entry)
}
