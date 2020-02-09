package logrus

import (
	"github.com/jonasi/ctxlog"
	"github.com/sirupsen/logrus"
)

// New returns a ctxlog.Logger backed by a logrus.FieldLogger
func New(l logrus.FieldLogger) ctxlog.Logger {
	if l == nil {
		l = logrus.New()
	}

	return logrusAdapter{l}
}

type logrusAdapter struct {
	logrus.FieldLogger
}

func (l logrusAdapter) KV(k string, v interface{}) ctxlog.Logger {
	return logrusAdapter{l.FieldLogger.WithField(k, v)}
}
