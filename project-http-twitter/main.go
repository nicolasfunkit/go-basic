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
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := server.Server{
		Repository: &server.TweetMemoryRepository{},
	}

	r.Get("/tweets", s.ListTweets)

	rateLimit := httprate.LimitByIP(10, 1*time.Minute)
	r.With(rateLimit).Post("/tweets", s.AddTweet)

	go spamTweets()

	http.ListenAndServe(":8080", r)
}

func spamTweets() {
	for {
		time.Sleep(time.Millisecond * 100)
		// Prepare payload
		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}
		// Marshal payload
		marshaledPayload, err := json.Marshal(addTweetPayload)
		if err != nil {
			panic(err)
		}

		// Send HTTP POST request
		url := "http://localhost:8080/tweets"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			panic(err)
		}
		// Close response body
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// (Optionally read and print the response)
		respBody := server.Response{}
		if err := json.Unmarshal(body, &respBody); err != nil {
			log.Printf("body error : %v", string(body))
			log.Println(err)
		} else {
			fmt.Sprintln("Add Tweet %d", respBody.ID)
		}

	}
}
