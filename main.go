package main

import (
	"fmt"
	"log"
	"zwlinc.com/wepush/core"
)

func main() {
	var accessToken core.AccessToken

	var cfg core.Config

	err := cfg.Read("./config.json")

	if err!=nil{
		log.Fatalln(err)
	}

	err = accessToken.Get(cfg.CorpId,cfg.CorpSecret)

	if err!=nil{
		log.Fatalln(err)
	}

	fmt.Println(accessToken.Token,accessToken.ExpireTime)
}
