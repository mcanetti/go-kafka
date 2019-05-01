package main

import (
	"context"
	"fmt"
	"time"
	"github.com/segmentio/kafka-go"
)


func main() {
	fmt.Println("Starting consumer...")
	
	// arrange
	broker_host := "kafka:9092"
	topic := "go-lang-topic"
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", broker_host, topic, partition)
	lifetime := 10

	// read messages	
	conn.SetReadDeadline(time.Now().Add(time.Duration(lifetime)*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}

