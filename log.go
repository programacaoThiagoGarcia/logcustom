package logcustom

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
)

type Type int64

const (
	Info    Type = 0
	Warning Type = 1
	Error   Type = 2
	Debug   Type = 3
)

func (s Type) Log(msg string) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	switch s {
	case Info:
		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		InfoLogger.Println(msg)
	case Warning:
		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger.Println(msg)
	case Error:
		ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger.Println(msg)

	case Debug:
		DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		DebugLogger.Printf("Debug message: %s  Some debug Info", msg)

	}
}
