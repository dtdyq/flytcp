package logs

// 安装以下依赖库
// go get -u go.uber.org/zap
// go get -u github.com/natefinch/lumberjack
// go get gopkg.in/alecthomas/kingpin.v2
import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
)

// createZapLogWithOpt 初始化 logger
func createZapLogWithOpt(opt LogOpt) *zap.SugaredLogger {
	ws := getLogWriter(opt)    // 日志文件配置 文件位置和切割
	encoder := getEncoder(opt) // 获取日志输出编码
	var ll = zapcore.InfoLevel
	switch opt.Level {
	case ERROR:
		ll = zapcore.ErrorLevel
	case WARN:
		ll = zapcore.WarnLevel
	case DEBUG:
		ll = zapcore.DebugLevel
	default:
		ll = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, ws, ll)
	log := zap.New(core, zap.AddCaller()) //  zap.addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	return log.Sugar()
}

// 编码器(如何写入日志)
func getEncoder(opt LogOpt) zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder   // looger 时间格式 例如: 2021-09-11t20:05:54.852+0800
	ec.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 info debug error
	ec.EncodeCaller = zapcore.FullCallerEncoder
	if opt.Format == TXT {
		return zapcore.NewConsoleEncoder(ec) // 以logfmt格式写入
	} else {
		return zapcore.NewJSONEncoder(ec) // 以json格式写入
	}
}

var LogFileMaxSize = 200

// 获取日志输出方式  日志文件 控制台
func getLogWriter(opt LogOpt) zapcore.WriteSyncer {
	// 日志文件 与 日志切割 配置
	ljl := &lumberjack.Logger{
		Filename:   filepath.Join(opt.LogDir, opt.LogFile), // 日志文件路径
		MaxSize:    LogFileMaxSize,                         // 单个日志文件最大多少 mb
		MaxBackups: 128,                                    // 日志备份数量
		MaxAge:     30,                                     // 日志最长保留时间
		Compress:   false,                                  // 是否压缩日志
	}
	// 日志只输出到文件
	return zapcore.AddSync(ljl)
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

var zapDef ZapLogger

type ZapLogger struct {
	hold *zap.SugaredLogger
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Error(format)
	} else {
		l.hold.Errorf(format, args...)
	}
}
func (l *ZapLogger) Info(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Info(format)
	} else {
		l.hold.Infof(format, args...)
	}

}
func (l *ZapLogger) Printf(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Info(format)
	} else {
		l.hold.Infof(format, args...)
	}

}
func (l *ZapLogger) Warn(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Warn(format)
	} else {
		l.hold.Warnf(format, args...)
	}

}
func (l *ZapLogger) Debug(format string, args ...interface{}) {
	if len(args) == 0 {
		l.hold.Debug(format)
	} else {
		l.hold.Debugf(format, args...)
	}
}

var zo sync.Once

func InitZap(opt LogOpt) Logger {
	zo.Do(func() {
		sug := createZapLogWithOpt(opt)
		zapDef = ZapLogger{hold: sug}
	})
	return &zapDef
}
