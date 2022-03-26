package main

import (
	"fmt"
	"net/http"
)

func main() {
	// a, b := 1, 0
	// ans := a / b // 1 divided by 0
	// fmt.Println(ans)

	// customPanic()

	requestExample()
}

func customPanic() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func requestExample() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err.Error())
	}
}
