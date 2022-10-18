package main

import (
	"fmt"
	kafka2 "github.com/AdilsonBND/tracking/simulator/infra/kafka"
	"github.com/joho/godotenv"
	"log"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish ("ola", "readtest", producer)
	for {
		_ = 1
	}
	// msgChan := make(chan *ckafka.Message)
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()
	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// 	go kafka2.Produce(msg)
	// }
}