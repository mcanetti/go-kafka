package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/segmentio/kafka-go"
)

func handlerIco(w http.ResponseWriter, r *http.Request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	// arrange
	broker_host := "kafka:9092"
	topic := "go-lang-topic"
	partition := 0

	// messages
	message_1, err := json.Marshal(map[string]string{"text": "one!"})
	message_2, err := json.Marshal(map[string]string{"text": "two!"})
	message_3, err := json.Marshal(map[string]string{"text": "three!"})

	// for linter sake
	fmt.Println(err, "attempt")
	
	// connect to kafka and send messages
	conn, _ := kafka.DialLeader(context.Background(), "tcp", broker_host, topic, partition)
	
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	conn.WriteMessages(
		kafka.Message{Value: message_1},
		kafka.Message{Value: message_2},
		kafka.Message{Value: message_3},
	)
	conn.Close()

	// simple endpoint response
	fmt.Fprintf(w, "Messages pushed to kafka %s", broker_host)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", handlerIco)
	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
