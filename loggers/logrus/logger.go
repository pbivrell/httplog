// logrus is the package the implements logger interface. This package reports the logs using logrus a popular text based logging package
package logrus

import (
	"github.com/pbivrell/httplog"
	"github.com/sirupsen/logrus"
)

// Logger is the logrus implementation of httplog.Logger
type Logger struct {
	Logger *logrus.Entry
}

type loggerOpt func(l *Logger)

// WithLogrus allows you to replace the default logrus logger with your own
func WithLogrus(logger *logrus.Entry) loggerOpt {
	return func(l *Logger) {
		l.Logger = logger
	}
}

// NewLogger returns an httplog.Logger. Writing logs using the default logrus logger. You may also providing 0 or more optional configuration to alter the defaults
func NewLogger(opts ...loggerOpt) *Logger {
	l := &Logger{
		Logger: logrus.New().WithFields(logrus.Fields{}),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Write implemets the httplog.Logger interface using the logrus logger to write a Info level log with all of the approriate fields.
func (l *Logger) Write(data httplog.Data) error {
	l.Logger.WithFields(logrus.Fields{
		"Method":           data.Method,
		"Endpoint":         data.Endpoint,
		"Code":             data.Code,
		"Referer":          data.Referer,
		"Duration":         data.Duration,
		"IP":               data.IP,
		"UserAgent":        data.UserAgent,
		"UserAgentVersion": data.UserAgentVersion,
		"OS":               data.OS,
		"OSVersion":        data.OSVersion,
		"Device":           data.Device,
		"DeviceType":       data.DeviceType,
		"Time":             data.Time,
	}).Infof("request")
	return nil
}
