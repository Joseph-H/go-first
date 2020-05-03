package main

import "fmt"

func main() {
	name := "Joseph"
	fmt.Println("Your name is", name)

	changeName(&name)

	fmt.Println("Your middle name is", name)
}

func changeName(name *string) string {
	*name = "Alexander"
	return *name
}
