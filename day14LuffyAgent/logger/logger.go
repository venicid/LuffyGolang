package logger

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	logger *logrus.Logger
	initLog sync.Once
)

func Init() error {
	// 设置日志格式为json
	err := errors.New("初始化")
	initLog.Do(func() {
		err = nil
		logger = logrus.New()
		logger.Formatter = &logrus.TextFormatter{
			FullTimestamp: true,
			TimestampFormat: "2006-01-02 15:04:05",
		}
		var filename string = "logfile.log"
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		logger.Out = f
		logger.Level = logrus.DebugLevel
	})
	return err
}

// WithField使用全局log返回logrus.Entry指针
func WithField(key string, value interface{})  * logrus.Entry{
	return logrus.WithField(key, value)
}


// Info 使用全局log记录信息
func Info(args ...interface{})  {
	logger.Info(args...)
}

func Fatal(args...interface{})  {
	logger.Fatalln(args...)
}

// 正常启动日志
func StartupInfo(msg ...interface{}) error {
	if err := Init();err !=nil{
		WithField("key", "startup").Info(msg...)
		return err
	}
	WithField("key", "startup").Info(msg...)
	return nil

}

// 启动失败日志
func FatalInfo(msg ...interface{}) error {
	if err := Init();err !=nil{
		WithField("key", "startup").Info(msg...)
		return err
	}
	WithField("key", "startup").Info(msg...)
	return nil

}