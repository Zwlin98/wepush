package main

import (
	"log"
	"zwlinc.com/wepush/application"
	"zwlinc.com/wepush/config"
)


func main() {
	err := config.AllConfig.Parse("./config.json")

	if err!=nil{
		log.Fatalln(err)
	}

	agentId:= config.AllConfig.AgentId
	user := config.AllConfig.User

	pusher := application.NewPusher(agentId,user)

	err = pusher.PushMessage("测试消息！\nSender by Pusher")
}
