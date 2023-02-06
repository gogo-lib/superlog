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
