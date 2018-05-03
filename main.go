package main

import (
	"fmt"
	"log"

	"github.com/avalchev94/boolean-evaluator/evaluator"
)

/*func main() {
	postID := flag.String("post", "16735011", "ID referencing \"who is hiring\" post.")
	flag.Parse()

	p, err := QueryPost(*postID)
	if err != nil {
		log.Fatalln(err)
	}

	posts := p.SearchKids("alabala")

	for _, post := range posts {
		fmt.Println(post.Text)
		fmt.Println("---------------------------------")
	}
}*/
func main() {
	eval, err := evaluator.New("A&B|C|(A&C|D)")
	if err != nil {
		log.Fatalln(err)
	}

	eval.Parameters["A"] = true

	result, err := eval.Evaluate()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
