package main

import (
	"testing"
)

const subredditsFileRead string = "test-read.json"

func TestLoad(t *testing.T) {
	data, err := SubredditsData{}.Load(subredditsFileRead)

	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}

	if data.Subreddits == nil {
		t.Error("data.subreddits is nil")
	}

	if l := len(data.Subreddits); l != 4 {
		t.Errorf("invalid subreddits size %v", l)
	}
}
