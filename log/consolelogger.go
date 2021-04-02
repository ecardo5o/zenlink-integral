package log

import "fmt"

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) Info(args ...interface{}) {
	fmt.Println(args)
}

func (cl *ConsoleLogger) Error(args ...interface{}) {
	fmt.Println(args)
}

func (cl *ConsoleLogger) Debug(args ...interface{}) {
	fmt.Println(args)
}

func (cl *ConsoleLogger) Warn(args ...interface{}) {
	fmt.Println(args)

}

func (cl *ConsoleLogger) DPanic(args ...interface{}) {
	fmt.Println(args)

}

func (cl *ConsoleLogger) Panic(args ...interface{}) {
	fmt.Println(args)
}

func (cl *ConsoleLogger) Fatal(args ...interface{}) {
	fmt.Println(args)
}

func (cl *ConsoleLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (cl *ConsoleLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (cl *ConsoleLogger) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)

}

func (cl *ConsoleLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (cl *ConsoleLogger) DPanicf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (cl *ConsoleLogger) Panicf(format string, args ...interface{}) {
	fmt.Printf(format, args...)

}

func (cl *ConsoleLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
