package web

import (
	"fmt"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func (app *application) initTus() *tusd.Handler {
	// Create a new FileStore instance which is responsible for
	// storing the uploaded file on disk in the specified directory.
	// This path _must_ exist before tusd will store uploads in it.
	// If you want to save them on a different medium, for example
	// a remote FTP server, you can implement your own storage backend
	// by implementing the tusd.DataStore interface.
	store := filestore.FileStore{
		Path: "./uploads",
	}

	// A storage backend for tusd may consist of multiple different parts which
	// handle upload creation, locking, termination and so on. The composer is a
	// place where all those separated pieces are joined together. In this example
	// we only use the file store but you may plug in multiple.
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	// Create a new HTTP handler for the tusd server by providing a configuration.
	// The StoreComposer property must be set to allow the handler to function.
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/files/",
		StoreComposer:           composer,
		MaxSize:                 8 * 1_000_000 * 1_000,
		RespectForwardedHeaders: true,
		NotifyCompleteUploads:   true,
		Logger:                  app.infoLog,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	// Start another goroutine for receiving events from the handler whenever
	// an upload is completed. The event will contains details about the upload
	// itself and the relevant HTTP request.
	go func() {
		for {
			event := <-handler.CompleteUploads
			md := event.Upload.MetaData
			app.infoLog.Println(md)
		}
	}()
	return handler
}
