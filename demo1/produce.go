package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	// 创建 一个话题
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig()) // 这个是直接打到  指定的nsqd上了
	if err != nil {
		log.Fatalln(err)
	}

	//
	//err = producer.Publish("Topic1", []byte(fmt.Sprintf("test2")))   // 对话不存在时就创建
	//if err != nil {
	//	log.Fatalln(err)
	//}

	for i := 0; i < 9999; i++ {
		err = producer.Publish("Topic1", []byte(fmt.Sprintf("test%d", i))) // 对话不存在时就创建
		if err != nil {
			log.Fatalln(err)
		}
	}
}
