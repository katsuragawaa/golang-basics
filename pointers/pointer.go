package main

import "fmt"

func main() {
	basic()
	pointerTypes()
}

func basic() {
	var a int = 42
	var b *int = &a    // * in front type -> pointer to a address with that type in memory
	fmt.Println(a, *b) // * in front variable -> dereference, find the memory location and pull the value
	fmt.Println(&a, b) // print both address

	a = 27
	fmt.Println(a, *b)
	*b = 9
	fmt.Println(a, *b)
}

func pointerTypes() {
	var ms *myStruct
	fmt.Println(ms) // nil
	ms = &myStruct{foo: 123}
	fmt.Println(ms)

	var ns *myStruct
	ns = new(myStruct)
	(*ns).foo = 312
	fmt.Println((*ns).foo)
	ns.foo = 111 // syntax sugar, pointer actually don't have foo
	fmt.Println(ns.foo)
}

type myStruct struct {
	foo int
}
