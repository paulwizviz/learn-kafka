package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func batchReader(batch *kafka.Batch) <-chan string {
	signal := make(chan string)
	// We use this to read data as soon as it
	// is available and sends it to channels
	go func() {
		b := make([]byte, 10e3)
		for {
			n, err := batch.Read(b)
			if err != nil {
				break
			}
			signal <- string(b[:n])
		}
		close(signal)
	}()
	return signal
}

func main() {
	topic := "quickstart"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	defer func() {
		err = batch.Close()
		if err != nil {
			log.Fatal("failed to close batch:", err)
		}
	}()

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal("failed to close connection:", err)
		}
	}()

	sig := batchReader(batch)
	for s := range sig {
		fmt.Println("-->", s)
	}

}
