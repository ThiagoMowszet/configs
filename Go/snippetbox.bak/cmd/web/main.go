package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


type applicattion struct {
	errorLog *log.Logger
	infoLog *log.Logger
}


func main() {
	addr := flag.String("addr", ":4000", "HTTP network adress")

	flag.Parse()

	// NOTE: if you want to include the full file path in your log output, instead of just the filename, you can use the log.Llongfile flag instead of log.Lshortfile when creating your custom logger.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	app := &applicattion{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))


	// mux.HandleFunc("/static", http.StripPrefix("/static/", fileServer))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// log.Print("Starting server on :4000")
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
