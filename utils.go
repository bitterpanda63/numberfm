package main

import "time"

func bitDepthToIntegerRange(bitDepth int) float64 {
	// 2^(bitDepth - 1) (the -1 is because it's signed lol)
	return float64(int(1) << (bitDepth - 1))
}

func frequencyToPeriod(frequency float64) time.Duration {
	periodInSeconds := 1.0 / frequency
	return time.Duration(periodInSeconds * float64(time.Second))
}
