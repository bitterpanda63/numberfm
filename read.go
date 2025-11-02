package main

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

func ReadAndBuffer(ws *websocket.Conn, scaleFactor float64, frequency float64) []int {
	// Buffer for reconstructed samples
	//var samples []int

	// Period for sample timing
	period := frequencyToPeriod(frequency)

	formerInt := readWs(ws)
	log.Print(formerInt)
	for {
		currentInt := readWs(ws)
		diff := formerInt - currentInt
		log.Printf("Current int: %d, formerInt: %d diff: %d", currentInt, formerInt, diff)

		formerInt = currentInt
		time.Sleep(period)
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
