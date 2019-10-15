package main

import (
	_struct "algorithm-learn/demo/jpush/struct"
	"fmt"
)

func main() {
	var (
		targetType = "single"
		targetID   = "9ba8808f-8779-442a-896d-8d4050c446ef"
		targetName = "lzk"
		fromType   = "user"
		fromID     = "bb7e9e9a-c8e3-4a56-915d-a1438da4060c"
		fromName   = "hw"
		//noOffline      = true
		//noNotification = true
	)
	client := _struct.New(targetType, targetID, targetName, fromType, fromID, fromName, true, true)
	result, err := client.SendMessage(
		_struct.JMessageText,
		_struct.NotificationMessage{
			Title: "hhhh",
			Alert: "哈哈哈哈",
		},
		&_struct.MsgBodyText{Text: "哈哈哈哈"},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
