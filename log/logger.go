package logger

import (
	"fmt"
	"os"
)

/***
	loggers should implimente log
***/

// log type
const (
	LOG_CONN = iota //connection logs
)

type log struct {
	name string
	file *File
}

func Newlog(str string) *log {
	return &log{name: str}
}

func (lg *log) checkOpen() {
	if lg.file == nil || lg.name == "" {
		lg.file, err = os.Open(lg.name)
	}
}

func (lg *log) Close() {
	if lg.file != nil {
		lg.file.Close()
	}
}

func (lg *log) info(format string, a ...interface{}) {
	lg.checkOpen()
	fmt.Fprintf(lg.file, format, a...)
	fmt.Println(format, a...)
}

type Logger struct {
	logs map[string]*log
}

func (lgr *Logger) GetLogger(logname string) *log {
	if _, ok := lgr.logs[logname]; !ok {
		lgr.logs[logname] = Newlog(logname)
	}
	return lgr.logs[logname]
}
