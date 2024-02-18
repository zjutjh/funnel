package logs

import (
	"fmt"
	"funnel/config/config"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

var logLevels = map[string]logrus.Level{
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func init() {
	Log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	}

	os.Mkdir("log", os.ModePerm)

	logName := "funnel"
	if config.Config.IsSet("log.name") {
		logName = config.Config.GetString("log.name")
	}

	logLevel := "info"
	if config.Config.IsSet("log.level") {
		logLevel = config.Config.GetString("log.level")
	}
	if level, ok := logLevels[logLevel]; ok {
		Log.SetLevel(level)
		if logLevel != "debug" {
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}

	maxAge := uint(14)
	if config.Config.IsSet("log.maxAge") {
		maxAge = config.Config.GetUint("log.maxage")
	}

	rotationTime := uint(1)
	if config.Config.IsSet("log.rotationTime") {
		maxAge = config.Config.GetUint("log.rotationTime")
	}

	hook := newHook("./log/"+logName, maxAge, rotationTime)
	Log.AddHook(hook)
}

func newHook(logName string, maxAge, rotationTime uint) logrus.Hook {
	writer, err := rotatelogs.New(
		logName+"-%Y%m%d.log",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour*24),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour*24),
	)
	if err != nil {
		fmt.Println(err)
	}

	Log.Out = writer

	lfHook := lfshook.NewHook(lfshook.WriterMap{}, &logrus.JSONFormatter{})
	return lfHook
}
