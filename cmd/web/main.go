package main

import (
    "log"
    "flag"
    "net/http"
    "os"
)

type application struct {
    errLog *log.Logger
    infoLog *log.Logger
}

func main() {

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) 

    app := &application{
        errLog: errLog,
        infoLog: infoLog,
    }

    addr := flag.String("addr", ":4000", "HTTP port")
    flag.Parse()
        mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/create", app.snippetCreate)
    mux.HandleFunc("/snippet/view", app.snippetView)
    
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errLog,
        Handler: mux,
    }

    app.infoLog.Println("Server started on", srv.Addr)
    err := srv.ListenAndServe()
    app.errLog.Fatal(err)
}
