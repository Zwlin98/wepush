package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"zwlinc.com/wepush/core"
)

type TextMessage struct {
	BaseRequest
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}


func SendTextMessage(token core.AccessToken,message *TextMessage) error{
	if token.Expired(){
		return errors.New("the token has expired")
	}

	url := fmt.Sprintf(apiUrlFormat,token.Token)

	body,err := json.Marshal(message)
	if err!=nil{
		return err
	}

	resp,err := http.Post(url,"application/json",bytes.NewReader(body))

	if err!=nil{
		return err
	}

	var respData ResponseData

	b,err := ioutil.ReadAll(resp.Body)

	if err!=nil{
		return err
	}

	err = json.Unmarshal(b,&respData)

	if err !=nil{
		return err
	}

	if respData.ErrCode != 0{
		return errors.New(respData.ErrMsg)
	}
	//部分无权限（不存在）用户失败不认为是错误
	return nil
}
