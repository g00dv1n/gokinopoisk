package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	MongoURL string `json:"mongoUrl"`
	StartID  uint64 `json:"startId,string"`
	EndID    uint64 `json:"endId"`
}

func (c *config) GetConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)

	return err
}
