package main

import (
	"fmt"
	"log"
	"math"
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
			time.Sleep(50 * time.Millisecond)
			send(ws)
			fmt.Printf("Got: %v\n", read(ws))
		}
	case "recv":
		fmt.Println("Receiving...")
	default:
		log.Fatal("Illegal argument")
	}
}

func send(ws *websocket.Conn) {
	var err error
	period := 1000. / (2 * math.Pi)
	val := 20 * math.Sin(float64(time.Now().UnixMilli())/period)
	fmt.Println(val)
	for range int(math.Abs(val)) {
		if val > 0 {
			_, err = ws.Write([]byte("+"))
		} else {
			_, err = ws.Write([]byte("-"))
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func read(ws *websocket.Conn) string {
	msg := make([]byte, 512)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	return string(msg[:n])
}
