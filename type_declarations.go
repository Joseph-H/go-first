package main

// TYPE DECLARATIONS

// Alias Declaration
// This acts like a string it has all the methods and fields of a string (could be any type)
// You cannot extend this, that is, you cannot add additional functionality to it
// Here TwitterHandler is a string
type TwitterHandler = string

type Person struct {
	handler TwitterHandler
}

// Type Defintion
// This is a brand new type it has all of the fields of string but NOT the methods
// Here GoogleHandler is not a string
type GoogleHandler string

func (gh *GoogleHandler) Redirect() {

}

type OtherPerson struct {
	handler GoogleHandler
}
