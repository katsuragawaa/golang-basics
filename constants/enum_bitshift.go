package main

import "fmt"

const (
	_  = iota // ignore fist value (enum will start at 1)
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Printf("%.2fMB", fileSize/MB)
}
