package clog

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// cLog 单例
var cLog *logrus.Logger

// InitLog bootstrap 初始化
func InitLog() {
	cLog = newLog("log/")
}

// Log 日志
func Log() *logrus.Logger {
	return cLog
}

// newLog 新建日志对象
func newLog(logpath string) *logrus.Logger {
	logger := logrus.New()
	level := logrus.InfoLevel
	logger.SetLevel(level)
	logger.SetReportCaller(true)
	fileHook, err := localRunHook(logpath)
	if err != nil {
		logger.Error("config local web file system for Logger error: %v", err)
	}
	logger.AddHook(fileHook)
	return logger
}

// myFormatter 格式
type myFormatter struct{}

// Format 格式
func (m *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// localRunHook 日志切割
func localRunHook(path string) (*lfshook.LfsHook, error) {
	writer, err := rotatelogs.New(
		path+"%Y%m%d"+".log",
		rotatelogs.WithRotationCount(7),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationSize(int64(5*1024*1024)),
	)
	if err != nil {
		return nil, err
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &myFormatter{})
	return lfsHook, err
}
