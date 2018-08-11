package logger

import (
	"encoding/json"
	"flag"
	"github.com/ameykpatil/go-uploader/utils/helper"
	"log"
	"runtime"
	"strings"
)

func init() {
	if flag.Lookup("test.v") != nil {
		logLine = func(assetID, message string, err error, logLevel string) {
			return
		}
	}
}

//LogObj is struct with all the log info
type LogObj struct {
	Timestamp int64  `json:"timestamp"`
	LogLevel  string `json:"logLevel"`
	Class     string `json:"class"`
	AssetID   string `json:"assetId"`
	Message   string `json:"msg"`
	Error     string `json:"error,omitempty"`
}

//Info just logs with logLevel INFO
func Info(assetID, message string) {
	logLine(assetID, message, nil, "")
}

//Err just logs with logLevel ERROR with error
func Err(assetID, message string, err error) {
	logLine(assetID, message, err, "ERROR")
}

//Warn just logs with logLevel WARNING with error
func Warn(assetID, message string, err error) {
	logLine(assetID, message, err, "WARNING")
}

var logLine = func(assetID, message string, err error, logLevel string) {
	_, className, _, _ := runtime.Caller(2)
	parts := strings.Split(className, "/")
	part := parts[len(parts)-1]
	arr := strings.Split(part, ".")
	class := arr[0]
	logObj := LogObj{
		Timestamp: helper.GetCurrentTime(),
		AssetID:   assetID,
		Message:   message,
		Class:     class,
		LogLevel:  "INFO",
	}
	if err != nil {
		logObj.Error = err.Error()
		logObj.LogLevel = logLevel
	}
	json, _ := json.Marshal(logObj)
	log.Println(string(json))
}
