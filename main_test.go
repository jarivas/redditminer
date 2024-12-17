package main

import (
	"testing"
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
	
	e := make(chan error)

	go processSubreddit(subreddit, e)

	err = <- e

	t.Error(err)
}