package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w);
        return
    }
    app.infoLog.Println("GET Request at /")

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
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
    if err != nil || id < 0 {
        app.notFound(w)
        return
    }
    app.infoLog.Println("Valid Query Parameter at /snippet/view")
    fmt.Fprintf(w, "Displaying snippet ID %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    app.infoLog.Println("Valid POST Request at /snippet/create")
    w.Write([]byte("Creating new snippet"))
}
