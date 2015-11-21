package main

import (
	"text/template"
	"net/http"
	"log"
	"path/filepath"
	"sync"
)

//templ represents a single template
type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "<html><head><title>Lets Chat!</title></head><body><h1>Hi there, Hello World!</h1></body></html>" ))
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
   
}
