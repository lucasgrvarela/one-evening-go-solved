package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Server struct {
	TweetsRepository tweetsRepository
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := Tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		log.Println("Failed to unmarshal paylod:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Location == "" || tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Tweet: `%s` from %s\n", tw.Message, tw.Location)

	id, err := s.TweetsRepository.AddTweet(tw)
	if err != nil {
		log.Println("Failed to add user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := response{
		ID: id,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetsRepository.Tweets()
	if err != nil {
		log.Printf("Failed to get tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to get marshal tweets %s:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}
