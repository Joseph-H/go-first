package main

import "fmt"

func main() {
	/*
	 * Array
	 */

	// create an array the long way
	var people [3]string
	people[0] = "joseph"
	people[1] = "josh"
	people[2] = "jim"

	// shorthand
	otherPeople := [3]string{"joseph", "josh", "jim"}

	fmt.Println(people)
	fmt.Println(otherPeople)

	/*
	 * Slice
	 */

	 // Can make them with make
	 test := make([]string, 5, 5)

	// this is  a "slice" it seems to work like a list
	jobs := []string{"engineer", "chef", "carpenter"}
	fmt.Println(jobs)

	// append creates a new array and move the data over
	// time appears to be linear
	jobs = append(jobs, "nurse")
	fmt.Println(jobs)

	// slice many ways
	s1 := jobs[:2] // from beginning to index 2
	s2 := jobs[2:] // from 2 to the end
	s3 := jobs[1:2] // from 1 till 2 not inclusive

	fmt.Println(s1, s2, s3)

	// You can make a slice and "append" a slice to another slice
	mySlice := int[]{1,2,3,4,5}
	another := int[]{6,7,8,9}
	mySlice = append(mySlice, another...)

	/*
	 * Map
	 */

	// shorthand (map literal) for creating a map and adding things to it
	// pairs := map[string]int{"foo": 42}

	// if you want to create a map and then add things later you can do this
	// foo := map[string]int{} this is virtually the same as the "make" function

	// this is a nil value, it doesn't point to a map.
	// Attempting to assign a value to this will result in an error
	var pairs map[string]int
	// use this to assing a map to the value above
	pairs = make(map[string]int)
	pairs["foo"] = 42
	pairs["bar"] = 52

	// declare and initialize a map
	things := map[string]string{
		"foo": "bar",
		"baz": "qux"
	}

	fmt.Println(pairs)
	fmt.Println(pairs["foo"])

	delete(pairs, "bar")
	fmt.Println(pairs)

	/*
	 * Structs
	 */

	type User struct {
		ID int
		firstname string
		lastname string
	}

	// you can also use dot notation to create a struct
	// mainUser.ID = 1 and so on
	mainUser := User{ ID: 1, firstname: "Arthur", lastname: "Morgan" }
	fmt.Println(mainUser)
}
