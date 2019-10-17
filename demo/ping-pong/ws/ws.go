package ws

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gopkg.in/olahol/melody.v1"
	"sync"
)

var server *melody.Melody
var handlers map[string]func(*melody.Session, []byte)
var handlerMutex sync.Mutex

type Message struct {
	mutex  sync.Mutex             `json:"-"`
	Key    string                 `json:"key"`
	Params map[string]interface{} `json:"params"`
}

func (m *Message) SetParams(key string, value interface{}) {
	if m.Params == nil {
		m.Params = make(map[string]interface{})
	}
	m.mutex.Lock()
	m.Params[key] = value
	m.mutex.Unlock()
}

func (m *Message) GetParams(key string) interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if value, ok := m.Params[key]; ok {
		return value
	}

	return ""
}

func (m *Message) Bytes() []byte {
	m.mutex.Lock()
	body, err := json.Marshal(m)
	m.mutex.Unlock()
	if err != nil {
		return nil
	}
	return body
}

func (m *Message) Unmarshal(body []byte) error {
	if err := json.Unmarshal(body, m); err != nil {
		return err
	}
	return nil
}

func init() {
	server = melody.New()
	beego.Get("/ws", func(c *context.Context) {
		_ = server.HandleRequest(c.ResponseWriter, c.Request)
	})
	server.HandleMessage(func(session *melody.Session, body []byte) {

		beego.Debug("ws message received:", string(body))

		var m Message
		if err := m.Unmarshal(body); err != nil {
			beego.Warning("ws message unmarshal error:", err)
		}
		//find_message_key
		if handler, ok := handlers[m.Key]; ok {
			handler(session, body)
		}

	})
}

func AddHandler(key string, handler func(session *melody.Session, body []byte)) {

	if handlers == nil {
		handlers = make(map[string]func(*melody.Session, []byte))
	}

	handlerMutex.Lock()
	handlers[key] = handler
	handlerMutex.Unlock()
}
