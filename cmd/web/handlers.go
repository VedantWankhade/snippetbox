package main

import (
    "fmt"
    "net/http"
    "strconv"
    "errors"
    // "html/template"
    "github.com/vedantwankhade/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w);
        return
    }
    app.infoLog.Println("GET Request at /")
    snippets, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    for _, snippet := range snippets {
        fmt.Fprintf(w, "%+v\n", snippet)
    }
/*
    files := []string{
        "./ui/html/base.tmpl.html",
        "./ui/html/partials/nav.tmpl.html",
        "./ui/html/pages/home.tmpl.html",
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
    */
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
    if err != nil || id < 0 {
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

    app.infoLog.Println("Valid Query Parameter at /snippet/view")
    fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    title := "0 snail"
    content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := 7

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    app.infoLog.Println("Valid POST Request at /snippet/create")
    http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
