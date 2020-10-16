package logger

import (
	"demo/global"
	"demo/pkg/config"
	"fmt"
	"runtime"
	"time"
)

type Level int8

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
)

//默认的日志级别
var defaultLevel Level = levelDebug

var levelMap = map[string]Level{
	"debug": levelDebug,
	"info":  levelInfo,
	"warn":  levelWarn,
	"error": levelError,
}

func Debug(a ...interface{}) {
	doPrint("debug", a)
}

func Info(a ...interface{}) {
	doPrint("info", a)
}
func Warn(a ...interface{}) {
	doPrint("warn", a)
}
func Error(a ...interface{}) {
	doPrint("error", a)
}

func DebugF(format string, a ...interface{}) {
	doPrintF("debug", format, a)
}

func InfoF(format string, a ...interface{}) {
	doPrintF("info", format, a)
}
func WarnF(format string, a ...interface{}) {
	doPrintF("warn", format, a)
}
func ErrorF(format string, a ...interface{}) {
	doPrintF("error", format, a)
}
func doPrint(level string, a ...interface{}) {
	if levelMap[level] >= levelMap[global.ServerConfig.MustGet("logger").(config.Configure).MustGet("level").(string)] {
		pc, file, line, _ := runtime.Caller(2)
		msg := fmt.Sprintf("[blog-%s] %v |%s %3s %d|%13s\n",
			level,
			time.Now().Format("2006/01/02 - 15:04:05"),
			file,
			runtime.FuncForPC(pc).Name(),
			line,
			fmt.Sprint(a),
		)
		global.Logger.Print(msg)
	}
}
func doPrintF(level string, format string, a ...interface{}) {
	var l Level
	var ok bool
	l, ok = levelMap[global.ServerConfig.MustGet("logger").(config.Configure).MustGet("level").(string)]
	if !ok {
		l = defaultLevel
	}
	if levelMap[level] > l {
		pc, file, line, _ := runtime.Caller(2)
		msg := fmt.Sprintf("[%9s]| %3v |%s %3s %d|%13s\n",
			"blog-"+level,
			time.Now().Format("2006/01/02 - 15:04:05"),
			file,
			runtime.FuncForPC(pc).Name(),
			line,
			fmt.Sprintf(format, a),
		)
		global.Logger.Print(msg)
	}
}
