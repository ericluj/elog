package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

var log = NewLog()

type Log struct {
	*logrus.Entry
	depth int
}

func NewLog() *Log {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetOutput(os.Stdout)

	return &Log{
		Entry: logrus.NewEntry(logger),
		depth: 0,
	}
}

func Field(key string, val interface{}) *Log {
	return &Log{Entry: log.Logger.WithField(key, val)}
}

func Fields(fields map[string]interface{}) *Log {
	return &Log{Entry: log.Logger.WithFields(fields)}
}

func stack(depth int) string {
	pc, file, n, ok := runtime.Caller(2 + depth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v:%v", filepath.Base(file), n, runtime.FuncForPC(pc).Name())
}

func (l *Log) setLevel(level string) {
	switch level {
	case DebugLevel:
		l.Entry.Logger.Level = logrus.DebugLevel
	case InfoLevel:
		l.Entry.Logger.Level = logrus.InfoLevel
	case WarnLevel:
		l.Entry.Logger.Level = logrus.WarnLevel
	case ErrorLevel:
		l.Entry.Logger.Level = logrus.ErrorLevel
	case FatalLevel:
		l.Entry.Logger.Level = logrus.FatalLevel
	default:
	}
}

func (l *Log) debugf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Debugf(format, args...)
}

func (l *Log) infof(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Infof(format, args...)
}

func (l *Log) warnf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Warnf(format, args...)
}

func (l *Log) errorf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Errorf(format, args...)
}

func (l *Log) fatalf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Fatalf(format, args...)
}

func SetLevel(level string) {
	log.setLevel(level)
}

func Debugf(layout string, args ...interface{}) {
	log.debugf(layout, args...)
}

func Infof(layout string, args ...interface{}) {
	log.infof(layout, args...)
}

func Warnf(layout string, args ...interface{}) {
	log.warnf(layout, args...)
}

func Errorf(layout string, args ...interface{}) {
	log.errorf(layout, args...)
}

func Fatalf(layout string, args ...interface{}) {
	log.fatalf(layout, args...)
}
