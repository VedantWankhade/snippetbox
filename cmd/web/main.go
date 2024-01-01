package main

import (
    "log"
    "flag"
    "net/http"
    "os"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "github.com/vedantwankhade/snippetbox/internal/models"
)

type application struct {
    errLog *log.Logger
    infoLog *log.Logger
    snippets *models.SnippetModel
}

func main() {

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)  
    addr := flag.String("addr", ":4000", "HTTP port")
    dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
    flag.Parse()
    
    db, err := openDB(*dsn)
    if err != nil {
        errLog.Fatal(err)
    }

    defer db.Close()

    app := &application{
        errLog: errLog,
        infoLog: infoLog,
        snippets: &models.SnippetModel{DB: db},
    }

    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errLog,
        Handler: app.routes(),
    }

    app.infoLog.Println("Server started on", srv.Addr)
    err = srv.ListenAndServe()
    app.errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}
