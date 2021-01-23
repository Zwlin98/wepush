package application

import (
	"zwlinc.com/wepush/config"
	"zwlinc.com/wepush/core"
	"zwlinc.com/wepush/message"
)

type Pusher struct {
	agentId int
	token   core.AccessToken
	user    string
}

func NewPusher(agentId int, user string) Pusher {
	return Pusher{
		agentId: agentId,
		user:    user,
	}
}

func (app *Pusher) PushMessage(msg string) error {
	if app.token.Expired() {
		err := app.GetToken()
		if err != nil {
			return err
		}
	}

	var body = message.TextMessage{
		BaseRequest: message.BaseRequest{
			ToUser:  app.user,
			ToParty: "",
			ToTag:   "",
			MsgType: "text",
			AgentId: app.agentId,
		},
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msg,
		},
		Safe:                   0,
		EnableIdTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 1800,
	}

	err := message.SendTextMessage(app.token, &body)
	if err != nil {
		return err
	}

	return nil
}

func (app *Pusher) GetToken() error {
	err := app.token.Get(config.AllConfig.CorpId, config.AllConfig.CorpSecret)
	return err
}
