package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"scanner/config"
)

type ZapLogger struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.Level(config.LogCfg.LogLevel))
	logger := zap.New(core, zap.AddCaller())

	return &ZapLogger{
		logger:      logger,
		sugarLogger: logger.Sugar(),
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	filename := ""
	ext := config.LogCfg.Extname
	if ext == "" {
		filename = filepath.Join(config.LogCfg.LogDir, fmt.Sprintf("%s-%s.log", config.LogCfg.Project, config.LogCfg.Name))
	} else {
		filename = filepath.Join(config.LogCfg.LogDir, fmt.Sprintf("%s-%s.log.%s", config.LogCfg.Project, config.LogCfg.Name, ext))
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (zl *ZapLogger) Info(args ...interface{}) {
	zl.sugarLogger.Info(args...)
}

func (zl *ZapLogger) Error(args ...interface{}) {
	zl.sugarLogger.Error(args...)
}

func (zl *ZapLogger) Debug(args ...interface{}) {
	zl.sugarLogger.Debug(args...)
}

func (zl *ZapLogger) Warn(args ...interface{}) {
	zl.sugarLogger.Warn(args...)

}

func (zl *ZapLogger) DPanic(args ...interface{}) {
	zl.sugarLogger.DPanic(args...)
}

func (zl *ZapLogger) Panic(args ...interface{}) {
	zl.sugarLogger.Panic(args...)
}

func (zl *ZapLogger) Fatal(args ...interface{}) {
	zl.sugarLogger.Fatal(args...)
}

func (zl *ZapLogger) Infof(format string, args ...interface{}) {
	zl.sugarLogger.Infof(format, args...)
}

func (zl *ZapLogger) Errorf(format string, args ...interface{}) {
	zl.sugarLogger.Errorf(format, args...)

}

func (zl *ZapLogger) Debugf(format string, args ...interface{}) {
	zl.sugarLogger.Debugf(format, args...)
}

func (zl *ZapLogger) Warnf(format string, args ...interface{}) {
	zl.sugarLogger.Warnf(format, args...)

}

func (zl *ZapLogger) DPanicf(format string, args ...interface{}) {
	zl.sugarLogger.DPanicf(format, args...)
}

func (zl *ZapLogger) Panicf(format string, args ...interface{}) {
	zl.sugarLogger.Panicf(format, args...)
}

func (zl *ZapLogger) Fatalf(format string, args ...interface{}) {
	zl.sugarLogger.Fatalf(format, args...)
}
