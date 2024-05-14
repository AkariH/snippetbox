package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *applcation) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprint("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *applcation) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}

func (app *applcation) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
