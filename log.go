package log

import (
	"github.com/sirupsen/logrus"
)

var log = NewLog()

type Log struct {
	*logrus.Entry
	depth int
}

func NewLog() *Log {
	logger := logrus.New()
	// logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:03:04",
		// CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		// 	//处理文件名
		// 	fileName := path.Base(frame.File)
		// 	return frame.Function, fileName
		// },
	})

	return &Log{
		Entry: logrus.NewEntry(logger),
		depth: 0,
	}
}

func (l *Log) Infof(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Infof(format, args...)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	l.Entry.WithField("call", stack(l.depth+1)).Fatalf(format, args...)
}

func Field(key string, val interface{}) *Log {
	return &Log{Entry: log.Logger.WithField(key, val)}
}

func Fields(fields map[string]interface{}) *Log {
	return &Log{Entry: log.Logger.WithFields(fields)}
}

func Infof(layout string, args ...interface{}) {
	log.Logger.Infof(layout, args...)
}

func Fatalf(layout string, args ...interface{}) {
	log.Logger.Fatalf(layout, args...)
}
