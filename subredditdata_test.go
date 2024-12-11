package main

import (
	"os"
	"testing"
)

const subredditsFileRead string = "test-read.json"
const subredditsFileWrite string = "test-write.json"

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

func TestSave(t *testing.T) {
	if _, err := os.Stat(subredditsFileWrite); err == nil {
		if err = os.Remove(subredditsFileWrite); err != nil {
			t.Error(err)
		}
	}

	data, err := SubredditsData{}.Load(subredditsFileRead)

	if err != nil {
		t.Error(err)
	}

	err = data.Save(subredditsFileWrite)

	if err != nil {
		t.Error(err)
	}
}
