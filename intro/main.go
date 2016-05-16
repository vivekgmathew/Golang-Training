package main

import "fmt"

type Person struct {
	name string
	age int
}

// For constants types are automatically inffered.
const  (
	PI = 3.14
	Language = "Go"
	A = iota
	B = iota
	C = iota
)



func main() {
	var name string
	var a, b, c int = 1, 3, 4
	name = "hello"

	age := 10

	fmt.Println(name)
	fmt.Println(a, b, c, age)

	// Pointer example
	var greeting *string
	greeting = &name
	fmt.Println(*greeting)

	// Type example
	var person = Person {"vivek", 20}
	fmt.Println(person.name)
	fmt.Println(person.age)

	var newPerson = Person { age : 20, name : "vivek"}
	fmt.Println(newPerson.name)
	fmt.Println(newPerson.age)

	// Constants examples
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)

}