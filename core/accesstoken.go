package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AccessToken struct {
	Token string

	ExpireTime time.Time

}

const apiUrlFormat = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"

func (at *AccessToken) Get(cropId, cropSecret string) error {
	if !at.Expired(){
		return nil
	}

	url := fmt.Sprintf(apiUrlFormat, cropId, cropSecret)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(resp.Body)

	var responseData struct {
		ErrorCode int    `json:"errorcode"`
		ErrMsg    string `json:"errmsg"`
		Token     string `json:"access_token"`
		Expire    int    `json:"expires_in"`
	}
	err = json.Unmarshal(result, &responseData)
	if err != nil {
		return err
	}

	if responseData.ErrorCode != 0 {
		return errors.New(responseData.ErrMsg)
	}


	at.Token = responseData.Token
	at.ExpireTime = time.Now().Add(time.Duration(responseData.Expire) * time.Second)

	return nil
}


func (at *AccessToken) Expired() bool {
	if time.Now().After(at.ExpireTime){
		return true
	}
	return false
}
