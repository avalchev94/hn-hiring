package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	url     = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
	Story   = "story"
	Comment = "comment"
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

type Searcher interface {
	Search(p Post) bool
}

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

func worker(kids <-chan int64, found chan<- *Post, done chan<- bool, s Searcher) {
	for id := range kids {
		if k, err := QueryPost(id); err == nil {
			if s.Search(*k) {
				found <- k
			}
		}
	}
	done <- true
}

func (p *Post) SearchKids(s Searcher, found chan<- *Post, workers int) {
	done := make(chan bool, workers)
	kids := make(chan int64)

	for i := 0; i < workers; i++ {
		go worker(kids, found, done, s)

	}
	for _, k := range p.Kids {
		kids <- k
	}
	close(kids)

	for i := 0; i < workers; i++ {
		<-done
	}
}

func (p *Post) Search(s Searcher, found chan<- *Post, workers int) {
	switch p.Type {
	case Comment:
		if s.Search(*p) {
			found <- p
		}
	case Story:
		p.SearchKids(s, found, workers)
	}
}
