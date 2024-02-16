package logcustom

import (
	"log"
	"os"

	color "github.com/programacaoThiagoGarcia/logcustom/colors"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
)

type Type int64

const (
	Info    Type = 0
	Warning Type = 1
	Error   Type = 2
	Debug   Type = 3
)

// Log is main method
func (s Type) Log(msg string) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	switch s {

	case Info:
		infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
		infoLogger.Println(color.Green + msg + color.Reset)
	case Warning:
		warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
		warningLogger.Println(color.Yellow + msg + color.Reset)
	case Error:
		errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
		errorLogger.Println(color.Red + msg + color.Reset)
	case Debug:
		debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime)
		debugLogger.Println(color.Cyan + msg + color.Reset)

	}
}
