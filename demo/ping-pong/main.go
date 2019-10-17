package main

import (
	"algorithm-learn/demo/ping-pong/ws"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gogf/gf/container/gmap"
	"gopkg.in/olahol/melody.v1"
	"time"
)

var Connect *gmap.HashMap

func main() {
	Connect = gmap.New()

	ws.AddHandler("heart_beat", func(session *melody.Session, body []byte) {
		type Type struct {
			Key    string `json:"key"`
			UserID uint   `json:"user_id"`
		}
		var param Type
		if err := json.Unmarshal(body, &param); err != nil {
			fmt.Println(err)
		}
		notRun := Connect.Get(param.UserID)
		if notRun == nil {
			notRun := make(chan struct{})
			go deadline(param.UserID, notRun)
			Connect.Set(param.UserID, notRun)
			fmt.Println("create connection")
		}
		notRunChan, ok := notRun.(chan struct{})
		if ok {
			notRunChan <- struct{}{}
			fmt.Println("notRunChan <- struct{}{}")
		}
	})

	beego.Run()
}

func deadline(userID uint, notRun chan struct{}) {
	for {
		select {
		case <-notRun:
		case <-time.After(time.Second * 6):
			Connect.Remove(userID)
			close(notRun)
			fmt.Println("offline")

			return
		}
	}
}
