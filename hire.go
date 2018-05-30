package main

import (
	"fmt"
	"github.com/avalchev94/boolean-evaluator/evaluator"
	"github.com/avalchev94/hn-hiring/hackernews"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

type hireOptions struct {
	Highlight bool `json:"highlight"`
}

type hireData struct {
	Items       []int64     `json:"items"`
	Expression  string      `json:"expression"`
	Preferences hireOptions `json:"preferences"`

	Eval *evaluator.Evaluator
}

var upgrader = websocket.Upgrader{}

func (hd hireData) Search(p hackernews.Post) bool {
	for param := range hd.Eval.Parameters {
		hd.Eval.Parameters[param] = strings.Contains(p.Text, param)
	}
	fmt.Println(hd.Eval)
	result, err := hd.Eval.Evaluate()
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

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

		data.Eval, err = evaluator.New(data.Expression)
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
				p.Search(data, found, 20)
			}
		}

		close(found)
	}
}
