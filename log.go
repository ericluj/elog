package log

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log = NewLog()

type Log struct {
	*logrus.Entry
	depth int
}

func NewLog() *Log {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	})

	return &Log{
		Entry: logrus.NewEntry(logger),
		depth: 0,
	}
}

func (l *Log) infof(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Infof(format, args...)
}

func (l *Log) fatalf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Fatalf(format, args...)
}

func Field(key string, val interface{}) *Log {
	return &Log{Entry: log.Logger.WithField(key, val)}
}

func Fields(fields map[string]interface{}) *Log {
	return &Log{Entry: log.Logger.WithFields(fields)}
}

func Infof(layout string, args ...interface{}) {
	log.infof(layout, args...)
}

func Fatalf(layout string, args ...interface{}) {
	log.fatalf(layout, args...)
}

func stack(depth int) string {
	pc, file, n, ok := runtime.Caller(2 + depth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v:%v", filepath.Base(file), n, runtime.FuncForPC(pc).Name())
}
