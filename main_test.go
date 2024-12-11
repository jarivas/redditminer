package main

import (
	"testing"
	"time"
)

func TestProcessSubreddit(t *testing.T) {
	data, err := SubredditsData{}.Load(subredditsFileRead)

	if err != nil {
		t.Error(err)
	}

	if len(data.Subreddits) == 0 {
		t.Error("problem decoding json file")
	}

	index := 0
	subreddit := data.Subreddits[index]

	go processSubreddit(data, subreddit)

	wait(t, subreddit)
}

func wait(t *testing.T, subreddit *SubredditInfo) {
	d, err := time.ParseDuration("300ms")
	if err != nil {
		t.Error(err)
	}

	lastId := subreddit.LastId
	max := 10
	i := 0

	for ; i < max; {
		time.Sleep(d)

		if lastId == subreddit.LastId {
			i = i + 1
		} else {
			i = max
		}
	}

	if lastId == subreddit.LastId {
		t.Error("lastid not updated")
	}
}