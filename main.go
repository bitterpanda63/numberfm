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
			// scale factor 100 at 2 kHz
			send(ws, 100, 2*1000)
			fmt.Printf("Got: %v\n", read(ws))
		}
	case "recv":
		fmt.Println("Receiving...")
	default:
		log.Fatal("Illegal argument")
	}
}

func send(ws *websocket.Conn, scaleFactor float64, frequency float64) {
	wav, err := LoadWav("input.wav")
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i < len(wav.Data); i++ {
		sample := float64(wav.Data[i])
		normalizedSample := sample / bitDepthToIntegerRange(wav.SourceBitDepth)

		// Calculate the number of "+" or "-" to send (simplified PWM)
		// Here, we use the absolute value and scale it to a reasonable range
		numPulses := int(math.Abs(normalizedSample * scaleFactor))

		// Send the PWM signal
		for j := 0; j < numPulses; j++ {
			if normalizedSample > 0 {
				_, err = ws.Write([]byte("+"))
			} else {
				_, err = ws.Write([]byte("-"))
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(frequencyToPeriod(frequency))
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

func bitDepthToIntegerRange(bitDepth int) float64 {
	// 2^(bitDepth - 1) (the -1 is because it's signed lol)
	return 1 << (bitDepth - 1)
}

func frequencyToPeriod(frequency float64) time.Duration {
	periodInSeconds := 1.0 / frequency
	return time.Duration(periodInSeconds * float64(time.Second))
}
