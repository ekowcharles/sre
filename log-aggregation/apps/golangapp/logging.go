package main

import (
	"bytes"
	"log"
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
	buf    bytes.Buffer
	logger = log.New(&buf, "", log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Lmsgprefix)

	logf = func(lt LogType, m string, v ...interface{}) {
		logger.Printf(lt.String()+m, v...)
	}

	logln = func(lt LogType, m string) {
		logger.Println(lt.String() + m)
	}
)
