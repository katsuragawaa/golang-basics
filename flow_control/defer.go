package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	basic()
	println("---")
	requestExample()
}

func basic() {
	a := "hi"
	defer fmt.Println(a) // will print "hi"
	a = "bye"

	fmt.Println("start")
	defer fmt.Println("middle") // run after basic func, before return
	defer fmt.Println("end")
}

func requestExample() {
	res, err := http.Get("http://www.google.com/robots.txt") // open
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // associate open and close resource

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
