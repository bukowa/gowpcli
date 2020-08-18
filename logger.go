package gowpcli

import (
	"fmt"
	"log"
)


var LogLevel = LevelInfo
var Logger = Logging{}

const (
	LevelDebug = iota
	LevelInfo = iota
)

type Logging struct {}

func (l *Logging) Debug(msg string)  {
	if LogLevel == LevelDebug {
		log.Print(msg)
	}
}

func (l *Logging) Debugf(format string, a ...interface{}){
	l.Debug(fmt.Sprintf(format, a...))
}
