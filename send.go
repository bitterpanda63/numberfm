package main

import (
	"log"
	"math"
	"time"

	"golang.org/x/net/websocket"
)

func Send(ws *websocket.Conn, scaleFactor float64, frequency float64) {
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
		if numPulses == 0 {
			continue
		}

		// Send the PWM signal
		log.Printf("Positive: %t, Pulses: %d", normalizedSample > 0, numPulses)
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
