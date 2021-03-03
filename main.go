package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zwlinc.com/wepush/application"
	"zwlinc.com/wepush/config"
)

func main() {
	err := config.AllConfig.Parse("./config.json")

	if err != nil {
		log.Fatalln(err)
	}

	agentId := config.AllConfig.AgentId
	user := config.AllConfig.User

	pusher := application.NewPusher(agentId, user)
	_ = pusher

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/wepush", func(c *gin.Context) {
		base64Text := c.Query("text")
		decodedText, e := base64.URLEncoding.DecodeString(base64Text)
		if e != nil {
			c.String(http.StatusBadRequest, "Error: %s occurred\n", e)
			return
		}
		//fmt.Println(string(decodedText))
		e = pusher.PushMessage(string(decodedText))
		if e == nil {
			c.String(http.StatusOK, "Successfully send message!\n")
			return
		} else {
			c.String(http.StatusBadRequest, "Error: %s occurred\n", e)
			return
		}
	})

	router.Run(":9854")
}
