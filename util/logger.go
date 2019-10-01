package util

import (
	"FunAnime-Server/model"
	"fmt"
	"time"
)

// TODO:写入日志文件

const (
	LogRequest  = 1
	LogResponse = 2
	LogDatabase = 3
	LogClassic  = 4

	LogInfo    = 10
	LogWarning = 20
	LogError   = 30
	LogPanic   = 40
)

var levelMap = map[int]string{
	LogInfo:    "INFO",
	LogWarning: "WARNING",
	LogError:   "ERROR",
	LogPanic:   "PANIC",
}

var typeMap = map[int]string{
	LogRequest:  "Request",
	LogResponse: "Response",
	LogDatabase: "Database",
	LogClassic:  "Classic",
}

type logger struct {
	level     int // 错误等级
	forType   int // 日志位置
	logTime   time.Time // 时间
	errno     int64  // 错误码
	message   string // 日志信息
	extraInfo string // 额外信息
}

func (lg *logger) Println() {
	fmt.Printf("[%s]|time=%v|errno=%d|forType=%d|msg=%s|extraInfo=%s", levelMap[lg.level], lg.logTime, lg.errno, typeMap[lg.forType], lg.message, lg.extraInfo)
}

func (lg *logger) SaveToDatabase() {
	traceId, _ := model.SaveLogInfo(&model.Log{
		Level:     lg.level,
		ForType:   lg.forType,
		LogTime:   lg.logTime,
		Errno:     lg.errno,
		Message:   lg.message,
		ExtraInfo: lg.extraInfo,
	})
	if lg.level > LogInfo {
		fmt.Printf("|trace_id=%d\n", traceId)
	} else {
		fmt.Printf("|\n")
	}
}

func Infof(forType int, errno int64, message string, args ...interface{}) {
	logInfo := logger{
		level:   LogInfo,
		logTime: time.Now(),
		forType: forType,
		errno:   errno,
		message: fmt.Sprintf(message, args),
	}

	logInfo.Println()
	logInfo.SaveToDatabase()
}

func Warnf(forType int, errno int64, message string, args ...interface{}) {
	logInfo := logger{
		level:   LogWarning,
		logTime: time.Now(),
		forType: forType,
		errno:   errno,
		message: fmt.Sprintf(message, args),
	}

	logInfo.Println()
	logInfo.SaveToDatabase()
}

func Errnof(forType int, errno int64, message string, args ...interface{}) {
	logInfo := logger{
		level:   LogError,
		logTime: time.Now(),
		forType: forType,
		errno:   errno,
		message: fmt.Sprintf(message, args),
	}

	logInfo.Println()
	logInfo.SaveToDatabase()
}

func Panicf(forType int, errno int64, message string, args ...interface{}) {
	logInfo := logger{
		level:   LogPanic,
		logTime: time.Now(),
		forType: forType,
		errno:   errno,
		message: fmt.Sprintf(message, args),
	}

	logInfo.Println()
	logInfo.SaveToDatabase()
}
