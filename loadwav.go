package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func LoadWav(filename string) (*audio.IntBuffer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening audio file: %v", err)
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	if !decoder.IsValidFile() {
		return nil, fmt.Errorf("not a valid WAV file")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, fmt.Errorf("error reading PCM buffer: %v", err)
	}
	return buf, nil
}
