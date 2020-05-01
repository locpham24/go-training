package log

import (
	"github.com/labstack/gommon/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

var myLogger = &MyLogger{}
var Log *logrus.Logger

type MyLogger struct {
	Logger *logrus.Logger
}

func NewLogger() *MyLogger {
	if Log != nil {
		myLogger = &MyLogger{
			Logger: Log,
		}
		return myLogger
	}
	Log = logrus.New()

	writerInfo, err := rotatelogs.New(
		"./log_files/info/"+"%Y%m%d_info.log",
		rotatelogs.WithMaxAge(time.Duration(60)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	if err != nil {
		log.Printf("Failed to create rotatelogs: %s", err)
		return nil
	}

	writerError, err := rotatelogs.New(
		"./log_files/error/"+"%Y%m%d_error.log",
		rotatelogs.WithMaxAge(time.Duration(60)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	if err != nil {
		log.Printf("Failed to create rotatelogs: %s", err)
		return nil
	}

	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writerInfo,
			logrus.WarnLevel:  writerInfo,
			logrus.ErrorLevel: writerError,
		},
		&logrus.TextFormatter{
			TimestampFormat:  time.RFC3339Nano,
			QuoteEmptyFields: true,
		},
	))

	myLogger = &MyLogger{
		Logger: Log,
	}
	return myLogger
}

func Error(i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	myLogger.Logger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Error(i...)
}

func Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	myLogger.Logger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Errorf(format, args...)
}

func Info(i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	myLogger.Logger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Info(i...)
}
