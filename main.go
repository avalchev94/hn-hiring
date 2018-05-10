package main

import (
	"flag"
	"fmt"
	"github.com/avalchev94/boolean-evaluator/evaluator"
	"github.com/avalchev94/hn-hiring/hackernews"
	"log"
	"strings"
)

func main() {
	postID := flag.Int64("post", 16735011, "ID referencing \"who is hiring\" post.")
	flag.Parse()

	p, err := hackernews.QueryPost(*postID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.Kids)
	found := make(chan *hackernews.Post)
	go func() {
		for p := range found {
			fmt.Printf("Post found with ID %d and Title %s.\n", p.ID, p.Title)
		}
	}()

	eval, err := evaluator.New("Remote & (C|Go|Golang)")
	if err != nil {
		log.Fatalln(err)
	}

	searchFunc := func(title, text string) bool {
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

	fmt.Printf("Asenaki")
}
