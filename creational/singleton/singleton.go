package singleton

import (
	"fmt"
	"sync"
)

type Logger struct{}

var loggerInstance *Logger
var loggerOnce sync.Once

func GetLoggerInstance() *Logger {
	loggerOnce.Do(func() {
		fmt.Println("Initializing Logger...")
		loggerInstance = &Logger{}
	})
	return loggerInstance
}

func (l *Logger) Log(message string) {
	fmt.Println("[SINGLETON LOGGER] ", message)
}

func Run() {
	logger1 := GetLoggerInstance()
	logger1.Log("First log message")

	logger2 := GetLoggerInstance()
	logger2.Log("Second log message")

	fmt.Printf("Are loggers the same? %v\n", logger1 == logger2)
}
