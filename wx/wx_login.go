package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
)

func main() {
	type AccessTokenRet struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Openid       string `json:"openid"`
		Scope        string `json:"scope"`

		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	code := "123456"
	var accessToken AccessTokenRet
	// appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	if err := httplib.Get("https://api.weixin.qq.com/sns/oauth2/access_token").
		Param("appid", "").
		Param("secret", "").
		Param("code", code).
		Param("grant_type", "authorization_code").ToJSON(&accessToken); err != nil {
		log.Printf("get wx accesstoken error: %v\n", err)
		return
	}

	fmt.Println(accessToken)
}
