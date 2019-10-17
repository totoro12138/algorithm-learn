package main

import (
	"algorithm-learn/demo/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"strconv"
	"time"
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
	Sound string `json:"sound"`
}

type Message struct {
	MsgContent  string `json:"msg_content"`  // 必填 	消息内容本身
	Title       string `json:"title"`        // 可选 	消息标题
	ContentType string `json:"content_type"` // 可选		消息内容类型
}

type Options interface {
	Marshal() ([]byte, error)
}

type OptionsSend struct {
	ApnsProduction bool `json:"apns_production"`
}

func (o *OptionsSend) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type OptionsOverwrite struct {
	ApnsProduction bool   `json:"apns_production"`
	OverrideMsgID  int    `json:"override_msg_id"`
	ApnsCollapseID string `json:"apns_collapse_id"`
}

func (o *OptionsOverwrite) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

type SmsMessage struct {
	DelayTime    int    `json:"delay_time"`    // 必填	单位为秒，不能超过 24 小时。设置为 0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
	Signid       int    `json:"signid"`        // 选填	签名ID，该字段为空则使用应用默认签名。
	TempId       int    `json:"temp_id"`       // 必填	短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara     string `json:"temp_para"`     // 可选	短信模板中的参数。
	ActiveFilter bool   `json:"active_filter"` // 可选	active_filter 字段用来控制是否对补发短信的用户进行活跃过滤，默认为 true ，做活跃过滤；为 false，则不做活跃过滤；
}

type PushResult struct {
	Error *Error `json:"error"`

	SendNo string `json:"sendno"`
	MsgID  string `json:"msg_id"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JPush(p Type) (result PushResult, err error) {
	appKey := config.JPushAppKey
	masterSecret := config.JPushMasterSecret
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
		JSONBody(p)
	if err != nil {
		return PushResult{}, err
	}

	resp, err := req.DoRequest()
	if err != nil {
		return PushResult{}, err
	}

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return PushResult{}, err
	}
	defer resp.Body.Close()
	fmt.Println(string(bytes))
	//result := make(map[string]interface{})
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return PushResult{}, err
	}

	return result, nil
}

func main() {
	result, err := JPush(Type{
		//Platform: "all",
		Platform: []string{
			PlatformAndroid,
			PlatformIos,
		},
		//Audience:"all",
		Audience: Audience{
			RegistrationID: []string{
				"190e35f7e02321351bf",
			},
		},
		Notification: &Notification{
			Android: &Android{
				Alert: "moyrn 发起通话",
				Title: "moyrn",
			},
			Ios: &Ios{
				Alert: "moyrn 发起通话",
				Sound: "call",
			},
		},
		Options: &OptionsSend{
			ApnsProduction: false,
		},
	})
	if err != nil {
		panic(err)
	}
	if result.Error != nil {
		panic(result)
	}
	fmt.Println(result)

	time.Sleep(time.Second * 8)
	msgIDInt, err := strconv.Atoi(result.MsgID)
	if err != nil {
		panic(err)
	}

	result, err = JPush(Type{
		Platform: []string{
			PlatformIos,
			PlatformAndroid,
		},
		Audience: Audience{
			RegistrationID: []string{
				"190e35f7e02321351bf",
			},
		},
		Notification: &Notification{
			Android: &Android{
				Alert: "moyrn 已结束拨号",
				Title: "moyrn 已结束拨号",
			},
			Ios: &Ios{
				Alert: "moyrn 已结束拨号",
				Sound: "moyrn 已结束拨号",
			},
		},
		Options: &OptionsOverwrite{
			ApnsProduction: false,
			OverrideMsgID:  msgIDInt,
			ApnsCollapseID: result.MsgID,
		},
	})

	if err != nil {
		panic(err)
	}
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(result)
}
