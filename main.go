package main

import (
    
	"net/http"
	"log"
)

//templ represents a single template
type templateHandler struct {
	one sync.Once
	filename string
	templ *template.Template
}
// ServeHTTP handles the HTTP request.
func (f *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "<html><head><title>Lets Chat!</title></head><body><h1>Hi there, Hello World!</h1></body></html>" ))
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
   
}
