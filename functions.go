package main

import "fmt"

func main() {
	bestFinish := bestLeagueFinishes(12, 10, 14, 13, 11, 15)
	fmt.Println(bestFinish)
}

// This is a variadic function taks a variable amount of arguments
func bestLeagueFinishes(finishes ...int) int {
	for _, finish := range finishes {
		// code goes here
	}
}

// Pointer Based Receiver
type Person struct {
	Name string
}

// When you are editing state like the function below, you need a pointer based
// receiver. If you don't do this you have to return a whole new object so that the
// changes are reflected. Since you do a pointer based receiver for this use case,
// for consistency you do it for all the methods
func (p *Person) AddName(n string) {
	p.Name = n
}
