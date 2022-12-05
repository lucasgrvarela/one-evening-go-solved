package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	go spamTweets()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))

	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() error {
	for {
		addTweetPayload := server.Tweet{
			Message:  "abc",
			Location: "def",
		}
		marshaledPayload, err := json.Marshal(addTweetPayload)
		if err != nil {
			return err
		}

		url := "http://localhost:8080/tweets"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Printf("spam sent %s: ", string(body))
	}
}
