package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	url = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
)

// Post is containing the basic data got from hackernews. The data could be
// accessed directly, however there are methods that can help processing it.
type Post struct {
	ID    int64   `json:"id"`
	Title string  `json:"title"`
	Text  string  `json:"text"`
	Kids  []int64 `json:"kids"`
	Type  string  `json:"type"`
}

// SearchFunc is a function type that is used by SearchKids method.
// Input: the title and text of the post
// Output: return if the post meets the requirements
type SearchFunc func(title, text string) bool

// QueryPost gets the data for the input id.
func QueryPost(postID int64) (*Post, error) {
	var post Post

	r, err := http.Get(fmt.Sprintf(url, postID))
	if err != nil {
		return &post, err
	}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return &post, err
	}

	return &post, nil
}

func searchWorker(kids <-chan int64, found chan<- *Post, f SearchFunc) {
	for id := range kids {
		if p, err := QueryPost(id); err == nil {
			if f(p.Title, p.Text) {
				found <- p
			}
		}
	}
}

func (p *Post) SearchKids(f SearchFunc, found chan<- *Post, threads int) {
	kids := make(chan int64)

	for i := 0; i < threads; i++ {
		go searchWorker(kids, found, f)
	}

	for _, k := range p.Kids {
		kids <- k
	}
	close(kids)
}
