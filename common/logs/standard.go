package logs

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

const (
	logUse = "logrus"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type LogFormat int

const (
	TXT LogFormat = iota
	JSON
)

type LogOpt struct {
	LogDir  string
	LogFile string
	Level   LogLevel
	Format  LogFormat
}

var logPathVar string
var levelVar string
var formatVar string

func init() {
	flag.StringVar(&logPathVar, "logpath", "log/gamesvr.log", "log path")
	flag.StringVar(&levelVar, "loglevel", "INFO", "log level")
	flag.StringVar(&formatVar, "logformat", "INFO", "log format(TXT/JSON)")
}

type Logger interface {
	Errorf(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

var defLogger Logger

func LogInit() {
	dir, name := filepath.Split(logPathVar)
	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(dir); !exist {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}
	var lv = INFO
	switch strings.ToLower(levelVar) {
	case "info":
		lv = DEBUG
	case "warn":
		lv = WARN
	case "error":
		lv = ERROR
	default:
		lv = INFO
	}
	var format = TXT
	switch strings.ToLower(formatVar) {
	case "txt":
		format = TXT
	case "json":
		format = JSON
	default:
		format = TXT
	}
	opt := LogOpt{
		LogDir:  dir,
		LogFile: name,
		Level:   lv,
		Format:  format,
	}
	switch logUse {
	case "logrus":
		defLogger = InitLogrus(opt)
	case "zap":
		defLogger = InitZap(opt)
	}
}
func GetLogger() Logger {
	return defLogger
}

func Error(template string, args ...interface{}) {
	defLogger.Errorf(template, args...)
}
func Warn(format string, args ...interface{}) {
	defLogger.Warn(format, args...)
}
func Info(format string, args ...interface{}) {
	defLogger.Info(format, args...)
}
func Debug(format string, args ...interface{}) {
	defLogger.Debug(format, args...)
}
