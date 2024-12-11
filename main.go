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

	for _, subreddit := range data.Subreddits {
		go processSubreddit(data, subreddit)
	}
}

func processSubreddit(data *SubredditsData, subreddit *SubredditInfo) {
	rp, err := redditmongo.RedditParams{}.Default(subreddit.Subreddit)

	if err != nil {
		log.Fatal(err)
	}

	rm, err := redditmongo.RedditMongo{}.FromEnv(rp)

	if err != nil {
		log.Fatal(err)
	}

	s := make(chan string)

	go func() {
		err = rm.Scrape(s)

		if err != nil {
			log.Fatal(err)
		}
	}()

	for lastId := range(s) {
		subreddit.LastId = lastId

		data.SaveDefault()
	}
}