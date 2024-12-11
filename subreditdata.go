package main

import(
    "encoding/json"
    "os"
)

const subredditsFile string = "subreddits.json"

type SubredditInfo struct {
	Subreddit string `json:"subreddit"`
	LastId    string `json:"lastId"`
}

type SubredditsData struct {
	Subreddits []*SubredditInfo `json:"subreddits"`
}

func (s SubredditsData) Load(path string) (*SubredditsData, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
    }

	defer file.Close()

	data := &SubredditsData{}

	json.NewDecoder(file).Decode(data)

	return data, nil
}

func (s SubredditsData) LoadDefault() (*SubredditsData, error) {
	return s.Load(subredditsFile)
}

func (s SubredditsData) Save(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)

    if err != nil {
        return err
    }

	defer file.Close()

	return json.NewEncoder(file).Encode(s)
}

func (s SubredditsData) SaveDefault() error {
	return s.Save(subredditsFile)
}
