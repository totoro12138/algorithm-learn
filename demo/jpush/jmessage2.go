package main

import (
	_struct "algorithm-learn/demo/jpush/struct"
	"fmt"
	"log"
	"strconv"
)

func main() {
	message := _struct.New("single", "2389d976-9f7a-4ae5-9ba4-7cafc929e14c", "moyrn",
		"user", "1377b39c-eaea-45e0-b814-3f48560aef31", "初末", false, false)
	result, err := message.SendMessage("text",
		_struct.NotificationMessage{
			Title: "初末",
			Alert: "向你送出" + strconv.Itoa(1) + "个" + "飞机",
		},
		&_struct.MsgBodyText{
			Text: "初末" + "向你送出" + strconv.Itoa(1) + "个" + "飞机",
		})
	if err != nil {
		msg, _ := message.Marshal()
		log.Printf("jmessage error:%s   context:%v\n", err.Error(), msg)
		return
	}
	fmt.Println(result)
}
