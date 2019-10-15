package mq

import (
	"github.com/nsqio/go-nsq"
	"log"
)

const (
	topic   = ""
	channel = ""
)

var consumer *nsq.Consumer

func init() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("mq message: %v\n", message.Body)

		return nil
	}))
}

func AddHandler(key string, handler func(body []byte)) {

}
