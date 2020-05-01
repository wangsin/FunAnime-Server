package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sinblog.cn/FunAnime-Server/util/osUtil"
	"time"
)

type Fields log.Fields

func Init() {
	now := time.Now()
	path := "./log/"
	filename := fmt.Sprintf("fa_run.%d%d%d%d.log", now.Year(), now.Month(), now.Day(), now.Hour())

	file, err := osUtil.TouchFile(filename, path)
	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
}

func TestLog() {
	log.WithFields(log.Fields{
		"type": "warn",
	}).Warn("test_log_warn")

	log.WithFields(log.Fields{
		"type": "info",
	}).Info("test_log_info")
}