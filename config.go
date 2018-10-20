package main

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type botConfig struct {
	Token    string          `json:"token"`
	Services []serviceConfig `json:"services"`
}

type serviceConfig struct {
	Name    string `json:"name"`
	Process string `json:"process"`
}

var sampleBotConfig = botConfig{
	Token: "114514",
	Services: []serviceConfig{
		{
			Name:    "Super evil daemon",
			Process: "evild",
		},
		{
			Name:    "Very innocent agent",
			Process: "nice-agent",
		},
	},
}

func generateSampleBotConfig(filename string) error {
	bs, err := json.Marshal(sampleBotConfig)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bs, 0644)
}

func loadBotConfig(filename string) (*botConfig, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg botConfig
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, err
	}
	//Check required fields
	if !(len(cfg.Token) > 0 && len(cfg.Services) > 0) {
		return nil, errors.New("incomplete configuration file")
	}
	return &cfg, nil
}
