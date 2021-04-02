package log

import (
	"errors"
	"fmt"
	"zenlink-integral/config"
)

var logger BaseLogger
var logLevel int

func init() {
	if config.LogCfg.Logger == "zap" {
		logger = NewZapLogger()
	} else {
		logger = &ConsoleLogger{}
	}

	if logger == nil {
		panic(errors.New("Initlogger error"))
	}
}

func Info(args ...interface{}) {
	if logLevel <= InfoLevel {
		logger.Info(args)
	}
}

func Error(args ...interface{}) {
	if logLevel <= ErrorLevel {
		fmt.Println(args)
	}
}

func Debug(args ...interface{}) {
	if logLevel <= DebugLevel {
		fmt.Println(args)
	}
}

func  Warn(args ...interface{}) {
	if logLevel <= WarnLevel {
		fmt.Println(args)
	}
}

func DPanic(args ...interface{}) {
	if logLevel <= DPanicLevel {
		fmt.Println(args)
	}
}

func Panic(args ...interface{}) {
	if logLevel <= PanicLevel {
		fmt.Println(args)
	}
}

func  Fatal(args ...interface{}) {
	if logLevel <= FatalLevel {
		fmt.Println(args)
	}
}