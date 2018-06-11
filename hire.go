package main

import (
	"fmt"
	"github.com/avalchev94/hn-hiring/hackernews"
	"github.com/avalchev94/hn-hiring/searcher/boolean"
	"github.com/gorilla/websocket"
	"net/http"
)

type hireOptions struct {
	Highlight bool `json:"highlight"`
}

type hireData struct {
	Items       []int64     `json:"items"`
	Expression  string      `json:"expression"`
	Preferences hireOptions `json:"preferences"`

	Searcher *boolean.Searcher
}

var upgrader = websocket.Upgrader{}

func hireHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error upgrading to socket: ", err)
		return
	}
	defer c.Close()

	for {
		var data hireData
		err := c.ReadJSON(&data)
		if err != nil {
			fmt.Println("error on socket reading: ", err)
			continue
		}
		fmt.Println("socket recieved:", data)

		data.Searcher, err = boolean.New(data.Expression)
		if err != nil {
			fmt.Println("error on eval creation: ", err)
			continue
		}

		found := make(chan *hackernews.Post, 30)
		go func() {
			for p := range found {
				c.WriteJSON(*p)
			}
		}()

		for _, postID := range data.Items {
			if p, err := hackernews.QueryPost(postID); err == nil {
				p.Search(data.Searcher, found, 20)
			}
		}

		close(found)
	}
}
