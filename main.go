package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const (
	templateFolder = "templates"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(templateFolder, "index.html")
	templ := template.Must(template.ParseFiles(path))

	data := map[string]interface{}{
		"Host": r.Host,
	}

	templ.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hire", hireHandler)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js"))))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalln(err)
	}
}
