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
    
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errLog,
        Handler: app.routes(),
    }

    app.infoLog.Println("Server started on", srv.Addr)
    err := srv.ListenAndServe()
    app.errLog.Fatal(err)
}
