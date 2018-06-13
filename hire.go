package main

import (
	"github.com/avalchev94/hn-hiring/hackernews"
	"github.com/avalchev94/hn-hiring/searcher/boolean"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type hireRequest struct {
	Items      []int64 `json:"items"`
	Expression string  `json:"expression"`

	Preferences struct {
		Highlight       bool `json:"highlight"`
		CaseInsensitive bool `json:"case_insensitive"`
	} `json:"preferences"`
}

type hireRespond struct {
	Type    int8        `json:"type"`
	Message interface{} `json:"message"`
}

// Respond types
const (
	foundPost int8 = iota
	nothingFound
	searchFinished
	incorrectExpression
)

var upgrader = websocket.Upgrader{}

func hireHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection on hire handler")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading to socket: ", err)
		return
	}
	defer c.Close()

	for {
		var msg hireRequest
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("error on socket reading: ", err)
			return
		}
		log.Println("socket message recieved:", msg)

		searcher, err := boolean.New(msg.Expression)
		if err != nil {
			log.Println("error on eval creation: ", err)

			c.WriteJSON(hireRespond{
				Type:    incorrectExpression,
				Message: *err,
			})
			continue
		}

		found := make(chan *hackernews.Post, 30)
		go foundPostHandler(found, c)

		for _, postID := range msg.Items {
			if p, err := hackernews.QueryPost(postID); err == nil {
				p.Search(searcher, found, len(p.Kids))
			}
		}

		close(found)
	}
}

func foundPostHandler(found <-chan *hackernews.Post, socket *websocket.Conn) {
	foundPosts := 0

	for p := range found {
		socket.WriteJSON(hireRespond{
			Type:    foundPost,
			Message: *p,
		})
		foundPosts++
	}

	if foundPosts > 0 {
		socket.WriteJSON(hireRespond{
			Type:    searchFinished,
			Message: strconv.Itoa(foundPosts) + "job(s) found. Now what?",
		})
	} else {
		socket.WriteJSON(hireRespond{
			Type:    nothingFound,
			Message: "There aren't available jobs for this search. Try another?",
		})
	}
}
