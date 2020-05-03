package main

type Name struct {
	First string
	Last  string
}

func (n *Name) FullName() {
	// Implementation...
}

// By placing Name inside of Person it now has the fields/signature of Name.
/*
 * Person actually looks like this
 * type Person struct {
 *	first string
 *	last string
 *	TwitterHandler string
 * }
 */

// You can access the fields like so:
// p.First or p.Name.First and p.FullName or p.Name.FullName
type Person struct {
	Name
	TwitterHandler string
}

// You can do the same thing with an interface
type Foo interface {
	Hello() string
}

// OtherPerson can be created and an object that implements Foo can be passed in
type OtherPerson struct {
	Name
	TwitterHandler string
	Foo
}
