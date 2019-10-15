package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
)

const (
	PlatformAndroid  = "android"
	PlatformIos      = "ios"
	PlatformWinPhone = "winphone"
)

type Type struct {
	Platform []string `json:"platform"`
	Audience Audience `json:"audience"`
	//Audience     string      `json:"audience"`
	Notification *Notification `json:"notification"`
	//Message      *Message      `json:"message"`
	//SmsMessage   *SmsMessage   `json:"sms_message"`
	Options Options `json:"options"`
}

type Audience struct {
	RegistrationID []string `json:"registration_id"`
}

type Notification struct {
	Android *Android `json:"android"`
	Ios     *Ios     `json:"ios"`
}

type Android struct {
	Alert string `json:"alert"`
	Title string `json:"title"`
}

type Ios struct {
	Alert string `json:"alert"`
}

type Message struct {
	MsgContent  string `json:"msg_content"`  // 必填 	消息内容本身
	Title       string `json:"title"`        // 可选 	消息标题
	ContentType string `json:"content_type"` // 可选		消息内容类型
}

type Options struct {
	ApnsProduction bool `json:"apns_production"`
}

type SmsMessage struct {
	DelayTime    int    `json:"delay_time"`    // 必填	单位为秒，不能超过 24 小时。设置为 0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
	Signid       int    `json:"signid"`        // 选填	签名ID，该字段为空则使用应用默认签名。
	TempId       int    `json:"temp_id"`       // 必填	短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara     string `json:"temp_para"`     // 可选	短信模板中的参数。
	ActiveFilter bool   `json:"active_filter"` // 可选	active_filter 字段用来控制是否对补发短信的用户进行活跃过滤，默认为 true ，做活跃过滤；为 false，则不做活跃过滤；
}

func JPush(p Type) error {
	//appKey := beego.AppConfig.String("jpush_appkey")
	//masterSecret := beego.AppConfig.String("master_secret")
	appKey := ""
	masterSecret := ""
	//if appKey == "" || masterSecret == "" {
	//	panic("jpush config not found")
	//}
	// https://api.jpush.cn/v3/push
	//fmt.Println(base64.RawURLEncoding.EncodeToString([]byte(appKey+":"+masterSecret)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(appKey + ":" + masterSecret)))
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))

	req, err := httplib.Post("https://api.jpush.cn/v3/push").
		Header("Authorization", " Basic "+base64.URLEncoding.EncodeToString([]byte(appKey+":"+masterSecret))).
		Header("Content-Type", "application/json").
		//Header("Authorization", base64.RawURLEncoding.EncodeToString([]byte(appKey+":"+masterSecret))).
		JSONBody(p)
	if err != nil {
		return err
	}

	resp, err := req.DoRequest()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		result := make(map[string]interface{})
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			return err
		}
		return errors.New(string(bytes))
	}
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(all))
	return nil
}

func main() {
	jpush := Type{
		//Platform: "all",
		Platform: []string{
			PlatformAndroid,
			PlatformIos,
		},
		//Audience:"all",
		Audience: Audience{
			RegistrationID: []string{
				"13165ffa4e6c731a894",
			},
		},
		Notification: &Notification{
			Android: &Android{
				Alert: "123",
				Title: "123",
			},
			Ios: &Ios{
				Alert: "123",
			},
		},
		Options: Options{ApnsProduction: false},
		//Message: &Message{
		//	MsgContent:  "123",
		//	Title:       "123",
		//	ContentType: "text",
		//},
	}

	err := JPush(jpush)
	if err != nil {
		panic(err)
	}
}
