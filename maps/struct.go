package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	createStruct()
	anonymousStruct()
}

func createStruct() {
	// struct let you have fields with different type of datas
	aDoctor := Doctor{
		number:     3,
		actorName:  "John",
		companions: []string{"Liz", "Jo"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	// can declare without naming, positional syntax (but not recommended)
	bDoctor := Doctor{
		1,
		"João",
		[]string{"Maria", "Teresa"},
	}
	fmt.Println(bDoctor)
	fmt.Println(bDoctor.actorName)
}

func anonymousStruct() {
	doc := struct{ name string }{name: "Andrew"}
	anotherDoc := doc // copy
	otherDoc := &doc  // pointer
	anotherDoc.name = "Jo"
	otherDoc.name = "Faker"
	fmt.Println("\n", doc.name)
	fmt.Println(anotherDoc.name)
	fmt.Println(otherDoc.name)
}
