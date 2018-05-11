package main

import (
	"flag"
	"github.com/avalchev94/boolean-evaluator/evaluator"
	"github.com/avalchev94/hn-hiring/hackernews"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	templateFolder = "templates"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(templateFolder, "hire.html")
	templ := template.Must(template.ParseFiles(path))

	data := map[string]interface{}{
		"Host": r.Host,
	}

	templ.Execute(w, data)
}

var postID = flag.Int64("post", 16735011, "ID referencing \"who is hiring\" post.")
var upgrader = websocket.Upgrader{}

func hireHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	p, err := hackernews.QueryPost(*postID)
	if err != nil {
		log.Fatalln(err)
	}

	found := make(chan *hackernews.Post, 10)
	go func() {
		for p := range found {
			c.WriteJSON(*p)
		}
	}()

	searchFunc := func(title, text string) bool {
		eval, err := evaluator.New("Remote & (Go|Golang)")
		if err != nil {
			log.Fatalln(err)
		}

		for param, _ := range eval.Parameters {
			eval.Parameters[param] = strings.Contains(text, param)
		}
		result, err := eval.Evaluate()
		if err != nil {
			log.Fatalln(err)
		}

		return result
	}

	p.SearchKids(searchFunc, found, 4)
	close(found)
}

func main() {

	flag.Parse()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hire", hireHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
