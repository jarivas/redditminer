package main

import(
    "encoding/json"
    "os"
)

const subredditsFile string = "subreddits.json"


type SubredditsData struct {
	Subreddits []string `json:"subreddits"`
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
