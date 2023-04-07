package logs

import (
	"golang.org/x/exp/slog"
	"os"
)

const defLf = "log/gamesvr.log"

var sf *os.File

func InitSlog(lf string) {
	if lf == "" {
		lf = defLf
	}
	f, err := os.OpenFile(lf, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic("open log file error" + err.Error())
	}
	slog.SetDefault(slog.New(slog.HandlerOptions{}.NewJSONHandler(f)))
}

//func Infof(f string, args ...interface{}) {
//	slog.Info(f, args)
//}
//
//func Error(f string, args ...interface{}) {
//	slog.Error(f, nil, args)
//}
