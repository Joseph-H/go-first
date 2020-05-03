package main

import "fmt"

func main() {
	/*
	 * Variables
	 */

	// one way to create a variable
	// var number int = 42 this is implied
	var number = 42

	// another way to create a variable
	anotherNumber := 42

	fmt.Println(number)
	fmt.Println(anotherNumber)

	// Go has a complex datatype
	expression := complex(3, 4)
	fmt.Println(expression)

	// You can pull out the real and imaginary numbers out
	real, imaginary := real(expression), imag(expression)
	fmt.Println(real, imaginary)

	/*
	 * Pointers
	 */

	// ways of creating a pointer variable
	//var firstName *string = new(string)
	// does the new keyword always create an empty pointer?
	var firstName = new(string)

	// to dereference a pointer and assign it a value
	*firstName = "Joseph"

	fmt.Println(*firstName)

	lastName := "Henriquez"
	// address of operator. This let's you create a pointer that points at another variables memory location
	pointer := &lastName
	fmt.Println(pointer, *pointer)

	/*
	 * Constants
	 */

	 const pi = 3.1415
	 fmt.Println(pi);

	 // Notice there is not type specified for this constant
	 const c = 3;
	 // Since c is not given an explicit type it is able to adapt
	 // to where ever it is being used.

	 // Here c is an int type
	 fmt.Println(c + 3); // result is 6
	 // Here c is a float type
	 fmt.Println(c + 1.2); // result is 4.2

	 /*
	 * Iota and Constant Expressions
	 * Look for more practical examples of this
	 * contstant blocks, iota, expressions
	 */
}
