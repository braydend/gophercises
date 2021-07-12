package main

import (
	"fmt"
	"log"
	"time"
)

func buildDuration(seconds int) time.Duration {
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", seconds))

	if err != nil {
		log.Fatal(err)
	}

	return duration
}

func startTimer(seconds int) *time.Timer {
	duration := buildDuration(seconds)
	return time.NewTimer(duration)
}