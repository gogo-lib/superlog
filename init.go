package superlog

import (
	"log"

	"github.com/natefinch/lumberjack"
)

// LogType log variation
type LogType int

// LogType value
const (
	Console LogType = iota
	DSPLog
)

// mapping LogType -> LogType instance
var (
	mapStructureLog   = make(map[LogType]structureLogger)
	mapUnStructureLog = make(map[LogType]unStructureLog)
)

func init() {
	{
		writer := &lumberjack.Logger{
			Filename: "",
			MaxSize:  1024, // 1GB
			MaxAge:   1,    // 1 day
			Compress: false,
		}

		_cronjob.AddFunc("*/10 * * * *", func() {
			err := writer.Rotate()
			if err != nil {
				log.Println("rotate DSP Log error: ", err)
			}
		})

		mapUnStructureLog[DSPLog] = unStructureLog{
			writer: writer,
		}
	}
}

// Info log with info level
func Info(logType LogType, msg any) {
	mapStructureLog[logType].info(msg)
}

// Infow log with info level, with key-value pair as option
func Infow(logType LogType, msg string, kv ...interface{}) {
	mapStructureLog[logType].infow(msg, kv...)
}

// Warn log with warn level
func Warn(logType LogType, msg any) {
	mapStructureLog[logType].warn(msg)
}

// Warnw log with warn level, with key-value pair as option
func Warnw(logType LogType, msg string, kv ...interface{}) {
	mapStructureLog[logType].warnw(msg, kv...)
}

// Error log with error level
func Error(logType LogType, err any) {
	mapStructureLog[logType].error(err)
}

// Errorw log with error level, with key-value pair as option
func Errorw(logType LogType, msg string, kv ...interface{}) {
	mapStructureLog[logType].errorw(msg, kv...)
}

// Raw log without level
func Raw(logType LogType, body []byte) {
	mapUnStructureLog[logType].raw(body)
}
