package web

import (
	"log"
	"net/http"
	"time"

	"github.com/TommasoAmici/mttaudio/pkg/loggers"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Run starts the web application
func Run() {
	// set up loggers
	infoLog := loggers.NewInfoLog()
	errorLog := loggers.NewErrLog()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:         "127.0.0.1:4000",
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	infoLog.Printf("Starting server on %s", "127.0.0.1:4000")

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Printf("Server listening on %s", "127.0.0.1:4000")
}
