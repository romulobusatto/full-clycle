package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	producer := NewKafkaProducer()
	publish("Transferência", "teste", producer, []byte("transferencia"), deliveryChan)
	go DeliveryReport(deliveryChan) //async
	fmt.Println("Romulo")
	producer.Flush(9000)

	//e := <-deliveryChan
	//msg := e.(*kafka.Message)
	//
	//if msg.TopicPartition.Error != nil{
	//	fmt.Println("Erro ao enviar mensagem: %v", msg.TopicPartition.Error)
	//}else{
	//	fmt.Println("Mensagem enviada", msg.TopicPartition)
	//}
	//producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
		//"delivery.timeout.ms": "0",
		//"acks": "all",
		//"enable.idempotence": "true",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar mensagem")
			} else {
				fmt.Println("Mensagem enviada", ev.TopicPartition)
			}
		}
	}
}
