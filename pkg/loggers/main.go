package loggers

import (
	"log"
	"os"
)

// NewErrLog creates an error log to Stderr
func NewErrLog() (errorLog *log.Logger) {
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return
}

// NewInfoLog creates an info log to Stdout
func NewInfoLog() (infoLog *log.Logger) {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	return
}
