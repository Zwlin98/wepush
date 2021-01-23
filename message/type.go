package message

const apiUrlFormat = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"

type BaseRequest struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int    `json:"agentid"`
}

type ResponseData struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InValidUser  string `json:"invaliduser"`
	InValidParty string `json:"invalidparty"`
	InValidTag   string `json:"invalidtag"`
}
