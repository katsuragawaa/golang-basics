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

func file() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Printf("%.2fMB", fileSize/MB)
}

const (
	isAdmin = 1 << iota
	isHQ
	isFinance

	regionA
	regionB
	regionC
)

func rolesCheck() {
	// fast AF
	var roles byte = isAdmin | isFinance | regionB
	fmt.Printf("%b\n", roles)

	fmt.Printf("Is admin? %v\n", roles&isAdmin == isAdmin)
	fmt.Printf("Is regionA? %v", roles&regionA == regionA)
}

func main() {
	file()
	rolesCheck()
}
