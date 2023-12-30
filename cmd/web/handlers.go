package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r);
        log.Println("Invalid GET Request at /")
        return
    }
    log.Println("GET Request at /")

    files := []string{
        "./ui/html/base.tmpl.html",
        "./ui/html/partials/nav.tmpl.html",
        "./ui/html/pages/home.tmpl.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
    if err != nil || id < 0 {
        http.NotFound(w, r)
        log.Println("Invalid Query Parameter at /snippet/view")
        return
    }
    log.Println("Valid Query Parameter at /snippet/view")
    fmt.Fprintf(w, "Displaying snippet ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        log.Println("Invalid GET Request at /snippet/create")
        return
    }
    log.Println("Valid POST Request at /snippet/create")
    w.Write([]byte("Creating new snippet"))
}
