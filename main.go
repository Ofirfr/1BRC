package main

import (
	"log"
	"time"

	"1BRC/first"
	"1BRC/second"
)

func main() {
	log.Println("First solution: ", measureTime(first.CalculateStatistics))
	log.Println("Second solution: ", measureTime(second.CalculateStatistics))
}

func measureTime(f func()) time.Duration {
	startTime := time.Now()

	f()

	endTime := time.Now()

	duration := endTime.Sub(startTime)

	return duration
}
