package main

import (
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/websocket"
)

func ReadAndBuffer(ws *websocket.Conn, scaleFactor float64, frequency float64) []int {
	var latestValue int64 // Use int64 for atomic operations

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Atomically update the latest value
			atomic.StoreInt64(&latestValue, int64(readWs(ws)))
		}
	}()

	// Buffer for reconstructed samples
	var samples []int

	frequencyTicker := time.NewTicker(frequencyToPeriod(frequency))
	defer frequencyTicker.Stop()

	formerInt := int(atomic.LoadInt64(&latestValue))
	log.Printf("Initial value: %d", formerInt)

	for {
		select {
		case <-frequencyTicker.C:
			currentInt := int(atomic.LoadInt64(&latestValue))
			diff := formerInt - currentInt

			log.Printf("Current int: %d, formerInt: %d, diff: %d", currentInt, formerInt, diff)

			samples = append(samples, diff)
			formerInt = currentInt
		}
	}
}

func readWs(ws *websocket.Conn) int {
	msg := make([]byte, 512)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	msgString := string(msg[:n])
	msgInt, err := strconv.Atoi(msgString)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	return msgInt
}
