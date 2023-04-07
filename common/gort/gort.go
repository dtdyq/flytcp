package gort

import (
	logs "flytcp/common/logs"
	"fmt"
	"runtime"
)

func QuietGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				var buf [1024 * 4]byte
				n := runtime.Stack(buf[:], false)
				fmt.Println(string(buf[:n]))
				logs.Error("<panic> msg=%v type=%T traceback:\n %s.", err, err, string(buf[:n])) // nolint:logcheck
				fmt.Println(err)
			}
		}()
		f()
	}()
}
func Error(f func() error) {
	go f()
}
