package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://the-number.site/"
	url := "ws://the-number.site/ws"
	fmt.Printf("Connecting to %v (origin: %v)...\n", url, origin)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	switch os.Args[1] {
	case "send":
		fmt.Println("Sending...")
		for {
			// scale factor 1000 at 0.5 kHz
			Send(ws, 1000, 500)

			// keep sending out the audio
			time.Sleep(50 * time.Millisecond)
		}
	case "recv":
		fmt.Println("Receiving...")
		ReadAndBuffer(ws, 1000, 500)
	default:
		log.Fatal("Illegal argument")
	}
}
