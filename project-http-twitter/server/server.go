package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Repository TweetRepository
}

type Response struct {
	ID int `json:"ID"`
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.Repository.Tweets()
	if err != nil {
		log.Printf("Failed to retrieve tweets : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tweetList := tweetsList{
		Tweets: tweets,
	}

	tweetsResponse, err := json.Marshal(tweetList)

	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(tweetsResponse)
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := Tweet{}
	if err := json.Unmarshal(body, &t); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)

	tweetID, err := s.Repository.AddTweet(t)
	if err != nil {
		log.Printf("Failed to add tweet : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tweetPayload, err := json.Marshal(Response{ID: tweetID})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(tweetPayload)
}

func (s Server) Tweets(w http.ResponseWriter, r *http.Request) {
	var duration time.Duration
	start := time.Now()

	defer func() {
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		s.AddTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.ListTweets(w, r)
	}

	duration = time.Since(start)
}
