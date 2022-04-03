package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	doneCh <- struct{}{} // just sending a message
}

func logger() {
	for {
		// blocks until some channel receive a message
		select {
		case entry := <-logCh:
			fmt.Printf(
				"%v - [%v]%v\n",
				entry.time.Format("02-01-2006 | 15:04:05"),
				entry.severity,
				entry.message,
			)
		case <-doneCh:
			break
		}
	}
}
