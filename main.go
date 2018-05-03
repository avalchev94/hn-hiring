package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	postID := flag.String("post", "16735011", "ID referencing \"who is hiring\" post.")
	flag.Parse()

	p, err := QueryPost(*postID)
	if err != nil {
		log.Fatalln(err)
	}

	posts := p.SearchKids([]string{"Remote"})

	for _, post := range posts {
		fmt.Println(post.Text)
		fmt.Println("---------------------------------")
	}
}
