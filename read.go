package main

import (
	"log"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

func ReadAndBuffer(ws *websocket.Conn, scaleFactor float64, frequency float64) []int {
	var (
		latestValue int64
		mu          sync.RWMutex
		wg          sync.WaitGroup
	)

	// Goroutine to continuously read from the WebSocket
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Lock for writing
			mu.Lock()
			latestValue = readWs(ws)
			mu.Unlock()
		}
	}()

	// Buffer for reconstructed samples
	var samples []int64
	frequencyTicker := time.NewTicker(frequencyToPeriod(frequency))
	defer frequencyTicker.Stop()

	// Get initial value
	mu.RLock()
	formerInt := latestValue
	mu.RUnlock()
	log.Printf("Initial value: %d", formerInt)

	for {
		select {
		case <-frequencyTicker.C:
			mu.RLock()
			currentInt := latestValue
			mu.RUnlock()

			diff := formerInt - currentInt
			log.Printf("diff: %d", diff)
			samples = append(samples, diff)
			formerInt = currentInt
		}
	}
}

func readWs(ws *websocket.Conn) int64 {
	msg := make([]byte, 512)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	msgString := string(msg[:n])
	msgInt, err := strconv.ParseInt(msgString, 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	return msgInt
}
