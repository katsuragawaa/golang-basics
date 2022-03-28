package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{} // connect interface with a struct
	w.Write([]byte("hellu~"))
}

// good practice to name the single method interface with methodName+er
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
