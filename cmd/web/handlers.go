package main

import (
	"errors"
	"fmt"

	// "html/template"
	"log"
	"net/http"
	"strconv"

	"moe.akari.best/internal/models"
)

func (app *applcation) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// w.Write([]byte("Hello from Snippetbox"))
	log.Println("got request from /")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/pages/home.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {

	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

}

func (app *applcation) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	log.Println("got request from /snippet/view")

}

func (app *applcation) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return

	}

	// dummy varible
	title := "0 snail"
	content := "0 nail\nClimb Mount Fuji,\nbut slowly,slowly!\n\n- Kobayashi Issa"
	expires := 7

	// pass the data into snippetmodel.Insert() method

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Create a new snippet..."))
	log.Println("got request from /snippet/create")

}
