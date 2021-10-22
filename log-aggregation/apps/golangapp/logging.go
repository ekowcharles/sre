package main

import (
	"bytes"
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
	buf bytes.Buffer

	fn   = fmt.Sprintf("golangapp-%d.log", time.Now().UnixNano())
	f, _ = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	logger   = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)
	_blogger = log.New(&buf, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)

	logf = func(lt LogType, m string, v ...interface{}) {
		ms := lt.String() + m

		logger.Printf(ms, v...)

		_blogger.Printf(ms, v...)
		fmt.Print(&buf)
	}

	logln = func(lt LogType, m string) {
		ms := lt.String() + m

		logger.Println(ms)

		_blogger.Println(ms)
		fmt.Print(&buf)
	}
)
