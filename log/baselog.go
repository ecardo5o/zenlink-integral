package log

type BaseLogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Panic(args ...interface{})
	DPanic(args ...interface{})
	Fatal(args ...interface{})

	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	DPanicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)
