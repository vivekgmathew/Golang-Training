package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Print(a)
	}

	if a:= 20; a == 10 {
		fmt.Println(a)
	}

	if a == 20 {
		fmt.Println("is 10")
	} else {
		fmt.Println("not 10")
	}

	val := switchBasic(10)
	fmt.Println(val)
	switchCondition()
	fmt.Println(switchType(1))
	fmt.Println(switchType("sdfsdf"))



}


func switchBasic(number int) (val string) {
	switch number {
		case 10: val ="Ten"
		case 4: val = "Four"
		default: val = "Not a number"
	}

	return
}

func switchCondition() {
	switch {
		case 10 == 10: fmt.Println("It's Ten")
	}
}

func switchType(x interface{}) string {
	var a string
	switch x.(type) {
		case int: a = "It's int"
		case string: a = "It's string"
	}
	return  a
}
