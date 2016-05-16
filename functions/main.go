package main

import "fmt"

type Person struct {
	name string
	age int
}

func Greet(s Person)  {
	fmt.Println(createMessage(s.name, s.age))
}

// Function that returns single value
func createMessage(name string, age int) string {
	return name + " " + string(age)
}

// Function that returns multiple values
func multiReturn() (a string, b string) {
	a = "hello"
	b = "hi"
	return
}

// Variadic functions (variable number of arguments)
func variadic(names ... string) {
	fmt.Println(names[0])
	fmt.Println(names[1])
}

// Function types (function that takes functions or known as higher order functions)
func higher(name string, rock func(string)) {
	rock(name)
}

// Function type declaration and usage
type Printer func(string) ()
type NewPrinter func(string) (string)

func higherOrder(name string, rock Printer) {
	rock(name)
}


func do(name string) {
	fmt.Println(name + "!!!!")
}

// Closures
func getFunc() func() () {
	name := "vivek"
	return func() {
		fmt.Println(name)
	}
}


func main() {
	s := Person{"vivek", 56}
	Greet(s)

	c, d := multiReturn()
	fmt.Println(c)
	fmt.Println(d)
	variadic("hi", "hello")
	higher("vivek", do)
	higherOrder("vivek", do)

	function := getFunc()
	function()
}





