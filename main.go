package main

import (
	"errors"
	"log"

	"github.com/jarivas/redditmongo"
	"github.com/jarivas/redditscraper"
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
		go processSubreddit(subreddit, e)
	}

	for err := range(e) {
		log.Printf("Error: %v", err)
	}
}

func getMongoReddit(subreddit string) (*redditmongo.RedditMongo, error) {
	ms, err := redditmongo.MongoStorage{}.FromEnv()

	if err != nil {
		log.Fatal(err)
	}

	rs, err := redditscraper.RedditScraper{}.FromEnv(subreddit)

	if err != nil {
		log.Fatal(err)
	}

	return redditmongo.RedditMongo{}.New(ms, rs)
}

func processSubreddit(subreddit string, e chan<- error) {
	rm, err := getMongoReddit(subreddit)

	if err != nil {
		e <- err
	}

	s := make(chan string)

	go rm.Scrape(s, e)

	for lastId := range(s) {
		if (lastId != "") {
			log.Printf("%v - %v\n", subreddit, lastId)
		} else {
			e <- errors.New("empty last id on " + subreddit)
		}
	}
}