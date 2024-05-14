package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *applcation) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// w.Write([]byte("Hello from Snippetbox"))
	log.Println("got request from /")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {

		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *applcation) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	log.Println("got request from /snippet/view")

}

func (app *applcation) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Add("Cache-Control", "public")

		app.clientError(w, http.StatusMethodNotAllowed)
		return

	}
	w.Write([]byte("Create a new snippet..."))
	log.Println("got request from /snippet/create")

}
