package _struct

import (
	"algorithm-learn/demo/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

type Result struct {
	MsgId     int `json:"msg_id"`
	MsgCtime  int `json:"msg_ctime"`
	ErrorCode int `json:"error_code"`
}

type JMessage struct {
	Version    int    `json:"version"`     // 版本号 目前是1 （必填）
	TargetType string `json:"target_type"` // 发送目标类型 single - 个人，group - 群组 chatroom - 聊天室（必填）
	TargetId   string `json:"target_id"`   // 目标id single填username group 填Group id chatroom 填chatroomid（必填）
	TargetName string `json:"target_name"` // 接受者展示名（选填）
	//TargetAppkey   string              `json:"target_appkey"`   //  跨应用目标appkey（选填）
	FromType       string               `json:"from_type"`       //	发送消息者的身份，可为“admin”，“user” （必填）
	FromId         string               `json:"from_id"`         // 发送者的username （必填
	FromName       string               `json:"from_name"`       // 发送者展示名（选填）
	MsgType        string               `json:"msg_type"`        //	发消息类型 text - 文本，image - 图片, custom - 自定义消息（msg_body为json对象即可，服务端不做校验）voice - 语音 （必填）
	NoOffline      bool                 `json:"no_offline"`      // 消息是否离线存储 true或者false，默认为false，表示需要离线存储（选填）
	NoNotification bool                 `json:"no_notification"` // 消息是否在通知栏展示 true或者false，默认为false，表示在通知栏展示（选填）
	Notification   *NotificationMessage `json:"notification"`    //  自定义通知栏展示（选填）
	MsgBody        MsgBody              `json:"msg_body"`        // Json对象的消息体 限制为4096byte
}

func (j *JMessage) Marshal() (string, error) {
	if bytes, err := json.Marshal(j); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

type NotificationMessage struct {
	Title string `json:"title"` // 通知的标题（选填）
	Alert string `json:"alert"` // 通知的内容（选填）
}

const (
	JMessageText  = "text"
	JMessageImage = "image"
	JMessageVoice = "voice"
)

type MsgBody interface {
	Return() (value string, isText bool)
}

type MsgBodyText struct {
	// msg_type为text时，msg_body的格式如下
	Text string `json:"text"` // msg_body -> text    消息内容 （必填）
	//Extras string // msg_body-> extras    选填的json对象 开发者可以自定义extras里面的key value（选填）
}

func (m *MsgBodyText) Return() (value string, isText bool) {
	return m.Text, true
}

type MsgBodyImage struct {
	// msg_type 为 image 时, msg_body 为上传图片返回的json，格式如下
	MediaID    string `json:"media_id"`    // msg_body->media_id    String 文件上传之后服务器端所返回的key，用于之后生成下载的url（必填）
	MediaCrc32 int64  `json:"media_crc32"` // msg_body->media_crc32    long 文件的crc32校验码，用于下载大图的校验 （必填）
	Hash       string `json:"hash"`        // msg_body->hash    String 图片hash值（可选）
	Fsize      int    `json:"fsize"`       // msg_body->fsize    int 文件大小（字节数）（必填）
	Width      int    `json:"width"`       // msg_body->width    int 图片原始宽度（必填）
	Height     int    `json:"height"`      // msg_body->height    int 图片原始高度（必填）
	Format     string `json:"format"`      // msg_body->format    String 图片格式（必填）
}

func (m *MsgBodyImage) Return() (value string, isText bool) {
	return m.MediaID, false
}

type MsgBodyVoice struct {
	// msg_type 为 voice 时, msg_body 为上传语音返回的json，格式如下
	MediaID    string `json:"media_id"`    // msg_body->media_id    String 文件上传之后服务器端所返回的key，用于之后生成下载的url（必填）
	MediaCrc32 int64  `json:"media_crc32"` // msg_body->media_crc32    long 文件的crc32校验码，用于下载大图的校验 （必填）
	Hash       string `json:"hash"`        // msg_body->hash    String 图片hash值（可选）
	Fsize      int    `json:"fsize"`       // msg_body->fsize    int 文件大小（字节数）（必填）
	Duration   int    `json:"duration"`    // msg_body->duration    int 音频时长（必填）
}

func (m *MsgBodyVoice) Return() (value string, isText bool) {
	return m.MediaID, true
}

func New(targetType, targetID, targetName, fromType, fromID, fromName string,
	noOffline, noNotification bool) JMessage {
	return JMessage{
		Version:        1,
		TargetType:     targetType,
		TargetId:       targetID,
		TargetName:     targetName,
		FromType:       fromType,
		FromId:         fromID,
		FromName:       fromName,
		NoOffline:      noOffline,
		NoNotification: noNotification,
		Notification:   &NotificationMessage{},
		MsgBody:        &MsgBodyText{},
	}
}

func (j *JMessage) SendMessage(msgType string, notificationMessage NotificationMessage, body MsgBody) (*Result, error) {
	// https://api.im.jpush.cn/v1/messages
	defer func() {
		j.MsgType = ""
		j.Notification = nil
		j.MsgBody = nil
	}()

	j.MsgType = msgType
	j.Notification = &notificationMessage
	j.MsgBody = body

	s, _ := j.Marshal()
	fmt.Println(s)

	request, err := httplib.Post("https://api.im.jpush.cn/v1/messages").
		Header("Authorization", " Basic "+base64.URLEncoding.EncodeToString([]byte(config.JPushAppKey+":"+config.JPushMasterSecret))).
		Header("Content-Type", "application/json").
		JSONBody(j)
	if err != nil {
		return nil, err
	}
	var result Result
	//if err := request.ToJSON(&result); err != nil {
	str, err := request.String()
	if err != nil {
		return nil, err
	}
	fmt.Println(str)
	return &result, nil
}
