package main

import "fmt"

func main() {
	createMap()
	createMapWithMake()
}

func createMap() {
	statePopulations := map[string]int{
		"California": 40000000,
		"Texas":      30000000,
		"Florida":    20000000,
	} // these orders won't be always the same
	fmt.Println(statePopulations)
	fmt.Printf("Population of California: %v\n", statePopulations["California"]) // access variable

	statePopulations["Other State"] = 10000000 // add key-value pair
	fmt.Println(statePopulations)

	delete(statePopulations, "Other State") // delete a key-value pair
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Other State"]) // 0 as falsy value

	population, okOne := statePopulations["California"]
	fakePopulation, okTwo := statePopulations["Other State"]
	fmt.Println(population, okOne)
	fmt.Println(fakePopulation, okTwo)

	fmt.Println("\nArray as key")
	// m := map[[]int]string{} <- slices can't be key
	n := map[[1]int]string{
		{2}: "Array as a key",
	} // but arrays can!
	// fmt.Println(m)
	fmt.Println(n)
}

func createMapWithMake() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 40000000,
		"Texas":      30000000,
		"Florida":    20000000,
	}
	fmt.Println("\nCreated with make")
	fmt.Println(statePopulations)
	fmt.Printf("Population of California: %v", statePopulations["California"])
}
