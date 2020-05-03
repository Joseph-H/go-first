package main

func main() {
	// loop till contidion
	var i int
	for i < 5 {
		println(i)
		i++
	}

	// with post clause
	for i := 0; i < 5; i++ {
		println(i)
	}

	// infinite loop
	var j int
	for {
		if j == 5 {
			break
		}

		j++
	}

	// iterate through slice
	people := []string{"Joseph", "Josh"}
	for index, value := range people {
		println(index, value)
	}

	// iterate through map
	ports := map[string]int{"http": 80, "https": 443}
	for key, value := range ports {
		println(key, value)
	}

	// key only
	for key := range ports {
		println(key)
	}

	// value only
	for _, value := range ports {
		println(value)
	}

	// when your application hits an error that it cannot recover from
	// ie. can't connect to the database you throw a panic
	// panic("something went incredibly wrong")

	foo := 10
	if foo == 1 {
		// do somthing
	} else if foo == 2 {
		// do something else
	} else {
		// good luck
	}

	/*
	 * Switch
	 */

	type HTTPRequest struct {
		Method string
	}

	req := HTTPRequest{"GET"}

	switch req.Method {
	case "GET":
		println("Get")
	default:
		println("default case")
	}

	// you can also switch on type
	switch req.(Type) {
	case HTTPRequest:
		return "Hi"
	}
}
