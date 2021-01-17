package web

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func (app *application) routes() http.Handler {
	mux := bone.New()

	tusHandler := app.initTus()
	mux.Handle("/files/", tusHandler.UnroutedHandler.Middleware(http.StripPrefix("/files/", tusHandler)))
	mux.GetFunc("/", app.uppyTest)
	return app.recoverPanic(mux)
}
