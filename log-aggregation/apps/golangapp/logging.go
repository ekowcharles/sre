package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogType int

const (
	Debug LogType = iota
	Info
	Warn
	Error
)

func (lt LogType) String() string {
	switch lt {
	case Debug:
		return "[DEBUG] "
	case Info:
		return "[INFO] "
	case Warn:
		return "[WARN] "
	case Error:
		return "[ERROR] "
	}

	return "[****] "
}

var (
	fn   = fmt.Sprintf("golangapp-%d.log", time.Now().UnixNano())
	f, _ = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	lg     = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)
	_lgout = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)
	_lgerr = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)

	logf = func(lt LogType, m string, v ...interface{}) {
		ms := lt.String() + m

		lg.Printf(ms, v...)

		if lt != Error {
			_lgout.Printf(ms, v...)
		} else {
			_lgerr.Printf(ms, v...)
		}
	}

	logln = func(lt LogType, m string) {
		ms := lt.String() + m

		lg.Println(ms)

		if lt != Error {
			_lgout.Println(ms)
		} else {
			_lgerr.Println(ms)
		}
	}
)
