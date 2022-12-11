package main

import "fmt"

func Trace(name string) func() {
	fmt.Println("Enter: ", name)
	return func() {
		fmt.Println("Exit: ", name)
	}
}

func foo() {
	defer Trace("foo")()
	bar()
}

func bar() {
	defer Trace("bar")()
}

func main() {
	defer Trace("main")()
	foo()
}
