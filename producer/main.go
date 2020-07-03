package main

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		panic(err)
	}

	topic := "notification-topic"

	words := []string{"Hello", "World", "My", "Name", "Dicky"}

	for _, word := range words {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)
	}

	producer.Close()
}
