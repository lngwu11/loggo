package loggo

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var logger *Logger

func Custom(server, uid string) {
	logger = New()

	infoWriter := getWriter(fmt.Sprintf("logs/%s/%s_info.log", server, uid))
	writers := []io.Writer{infoWriter, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	logger.SetFormatter(&TextFormatter{
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}

func ShowCaller(isShow bool) {
	if logger != nil {
		logger.SetReportCaller(isShow)
	}
}

func getWriter(filename string) io.Writer {
	// 保存30天内的日志，按天分割日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"_%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func Tracef(format string, args ...interface{}) {
	logger.Logf(TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Logf(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf(InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	entry := logger.newEntry()
	entry.Printf(format, args...)
	logger.releaseEntry(entry)
}

func Warnf(format string, args ...interface{}) {
	logger.Logf(WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Logf(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Logf(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Logf(FatalLevel, format, args...)
	logger.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	logger.Logf(PanicLevel, format, args...)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	logger.Log(TraceLevel, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logger.Log(DebugLevel, args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	entry := logger.newEntry()
	entry.Print(args...)
	logger.releaseEntry(entry)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logger.Log(InfoLevel, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	logger.Log(ErrorLevel, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logger.Log(PanicLevel, args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	logger.Log(FatalLevel, args...)
	logger.Exit(1)
}
