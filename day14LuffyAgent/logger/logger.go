package logger

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	logger  *logrus.Logger
	initLog sync.Once
)

func Init() error {
	// 设置日志格式为json格式

	err := errors.New("已经被初始化")
	initLog.Do(func() {
		err = nil
		logger = logrus.New()
		logger.Formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		}
		//filename := settings.Config().Logfile
		//if filename != ""{
		//	filename = "logfile.log"
		//}
		var filename string = "logfile.log"
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		logger.Out = f
		logger.Level = logrus.DebugLevel
	})
	return err
}

func SetLog(l *logrus.Logger) {
	logger = l
}

// WithField 使用全局log返回logrus.Entry指针
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// Debug 使用全局log记录信息
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info 使用全局log记录信息
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Fatal 使用全局log记录信息
func Fatal(args ...interface{}) {
	logger.Fatalln(args...)
}

// 正常启动日志
func StartupInfo(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "startup").Info(msg...)
		return err
	}
	WithField("key", "startup").Info(msg...)
	return nil
}

// 正常启动Debug日志
func StartupDebug(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "startup").Debug(msg...)
		return err
	}
	WithField("key", "startup").Debug(msg...)
	return nil
}

// 启动失败日志
func StartupFatal(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "startup").Fatalln(msg...)
		return err
	}
	WithField("key", "startup").Fatalln(msg...)
	return nil
}

// 监控采集相关DEBUG日志
func ToMOCDebug(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "moc").Debug(msg...)
		return err
	}
	WithField("key", "moc").Debug(msg...)
	return nil
}

// 心跳相关DEBUG日志
func HeartBeatsDebug(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "heartbeat").Debug(msg...)
		return err
	}
	WithField("key", "heartbeat").Debug(msg...)
	return nil
}

// 作业平台相关DEBUG日志
func JobDebug(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "job").Debug(msg...)
		return err
	}
	WithField("key", "job").Debug(msg...)
	return nil
}

func JobFatal(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "job").Fatalln(msg...)
		return err
	}
	WithField("key", "job").Fatalln(msg...)
	return nil
}

func JobInfo(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "job").Info(msg...)
		return err
	}
	WithField("key", "job").Info(msg...)
	return nil
}
