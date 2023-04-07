package logs

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"sync"
	"time"
)

func NewLoggerWithOpt(opt LogOpt) *logrus.Logger {
	ll := logrus.New()
	switch opt.Level {
	case DEBUG:
		ll.SetLevel(logrus.DebugLevel)
	case INFO:
		ll.SetLevel(logrus.InfoLevel)
	case WARN:
		ll.SetLevel(logrus.WarnLevel)
	case ERROR:
		ll.SetLevel(logrus.ErrorLevel)
	}
	switch opt.Format {
	case TXT:
		ll.SetFormatter(&logrus.TextFormatter{})
	case JSON:
		ll.SetFormatter(&logrus.JSONFormatter{})
	default:
		ll.SetFormatter(&logrus.TextFormatter{})
	}
	writer, err := rotatelogs.New(
		// 切割后日志文件名称
		path.Join(opt.LogDir, opt.LogFile+".%Y-%m-%d"),
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		//rotatelogs.WithRotationCount(3),
		//rotatelogs.WithRotationTime(time.Minute), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(writer, ll.Formatter)
	ll.AddHook(lfHook)
	return ll
}

var logrusDef LogrusLogger

type LogrusLogger struct {
	hold *logrus.Logger
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Error(format)
	} else {
		l.hold.Errorf(format, args...)
	}
}
func (l *LogrusLogger) Info(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Info(format)
	} else {
		l.hold.Infof(format, args...)
	}

}
func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Info(format)
	} else {
		l.hold.Infof(format, args...)
	}

}
func (l *LogrusLogger) Warn(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Warn(format)
	} else {
		l.hold.Warnf(format, args...)
	}

}
func (l *LogrusLogger) Debug(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Debug(format)
	} else {
		l.hold.Debugf(format, args...)
	}
}

var once sync.Once

func InitLogrus(opt LogOpt) Logger {
	once.Do(func() {
		l := NewLoggerWithOpt(opt)
		logrusDef = LogrusLogger{hold: l}
	})
	return &logrusDef
}
