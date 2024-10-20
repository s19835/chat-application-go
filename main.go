package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Create new type for handling templates
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// Listen to the path using net/http pkg
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write out the HTML when a req is made
		w.Write([]byte(
			`<html>
				<head>
					<title> Chat </title>
				</head>
				<body>
					Let's Chat
				</body>
			</html>
		`))
	})

	// Start the web-server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
