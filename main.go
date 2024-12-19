package main

import (
	"log"

	"github.com/jarivas/redditmongo"
)

func main() {
	data, err := SubredditsData{}.LoadDefault()

	if err != nil {
		log.Fatal(err)
	}

	if len(data.Subreddits) == 0 {
		log.Fatal("problem decoding json file")
	}
	
	e := make(chan error)

	for _, subreddit := range data.Subreddits {
		processSubreddit(subreddit, e)
	}

	for err := range(e) {
		log.Printf("Error: %v", err)
	}
}

func processSubreddit(subreddit string, e chan<- error) {
	rm, err := redditmongo.RedditMongo{}.FromEnv(subreddit)

	if err != nil {
		e <- err
	}

	go rm.Scrape(e)
}