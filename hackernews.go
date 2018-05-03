package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	url = "https://hacker-news.firebaseio.com/v0/item/%s.json?print=pretty"
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

// QueryPost gets the data for the input id.
func QueryPost(postID string) (*Post, error) {
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

func (p *Post) SearchKids(expression string) []*Post {
	var posts []*Post

	//eval, err := evaluator.New(expression)

	/*	for _, id := range p.Kids {
		if kid, err := QueryPost(strconv.FormatInt(id, 10)); err == nil {
			for _, k := range keywords {
				if strings.Contains(kid.Text, k) {
					posts = append(posts, kid)
					fmt.Println("Found")
					break
				}
			}
		}
	}*/

	return posts
}
