package adaptor

import "fmt"

// Old Logger (Incompatible interface)
type OldLogger struct {}

func (o *OldLogger) LogMessage(msg string) {
  fmt.Println("[OLD LOG] " + msg)
}


// New LoggerInterface
type Logger interface {
  Log(msg string)
}

// adaptor
type LoggerAdapter struct {
  oldLogger *OldLogger
}

func (l *LoggerAdapter) Log(msg string) {
  l.oldLogger.LogMessage(msg) // adapting old method to new interface
}

func Run() {
  // using old logger directly 
  oldLogger := &OldLogger{}
  oldLogger.LogMessage("This is an old message.")

  // Using the adapter to work with the new interface
  adapter := &LoggerAdapter{oldLogger}
  adapter.Log("This is a new log message using the adapter.")
}
