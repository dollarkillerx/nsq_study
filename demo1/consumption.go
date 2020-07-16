package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	consumer, err := nsq.NewConsumer("Topic1", "channel1", nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}
	//err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatalln(err)
	}

	//consumer.
	for {
		select {}
	}
}
