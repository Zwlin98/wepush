package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var AllConfig Config

type Config struct {
	CorpId     string `json:"corpId"`
	CorpSecret string `json:"corpSecret"`
	User       string `json:"user"`
	AgentId    int    `json:"agentid"`
}

func (cfg *Config) Parse(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, cfg)
	if err != nil {
		return err
	}

	return nil
}
