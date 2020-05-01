package logger

import (
	log "github.com/sirupsen/logrus"
)

func Info(msg string, filed map[string]interface{}) {
	log.WithFields(filed).Info(msg)
}

func Warn(msg string, filed map[string]interface{}) {
	log.WithFields(filed).Warn(msg)
}

func Error(msg string, filed map[string]interface{}) {
	log.WithFields(filed).Error(msg)
}

func Fatal(msg string, filed map[string]interface{}) {
	log.WithFields(filed).Fatal(msg)
}

func Panic(msg string, filed map[string]interface{}) {
	log.WithFields(filed).Panic(msg)
}