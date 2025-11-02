package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://the-number.site/"
	url := "ws://the-number.site/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	period := 3000.
	val := 20 * math.Sin(float64(time.Now().Unix())/period)
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

/*
	msg := make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recieved: %s\n", msg[:n])
*/
