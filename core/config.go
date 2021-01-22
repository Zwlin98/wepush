package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	CorpId string `json:"corpId"`
	CorpSecret string `json:"corpSecret"`
}

func (cfg *Config) Read(path string) error {
	file,err := os.Open(path)
	if err!=nil{
		return err
	}

	b,err := ioutil.ReadAll(file)
	if err!=nil{
		return err
	}

	err =json.Unmarshal(b,cfg)
	if err!=nil{
		return err
	}

	return nil
}