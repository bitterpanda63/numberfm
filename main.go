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
			// scale factor 100 at 2 kHz
			Send(ws, 100, 2*1000)

			// keep sending out the audio
			time.Sleep(50 * time.Millisecond)
		}
	case "recv":
		fmt.Println("Receiving...")
		ReadAndBuffer(ws, 100, 2*1000)
	default:
		log.Fatal("Illegal argument")
	}
}
