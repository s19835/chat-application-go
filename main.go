package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// Create new type for handling templates
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

// Listen to the path using net/http pkg
func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// Start the web-server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
