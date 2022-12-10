package main

import "fmt"

func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from the panic:", err)
		}
	}()
	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func show_panic_recover() {
	println("call show_panic_recover")
	foo()
	println("exit show_panic_recover")
}

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func show_defer_order() {
	fmt.Println("foo1 result: ")
	foo1()
	fmt.Println("foo2 result: ")
	foo2()
	fmt.Println("foo3 result: ")
	foo3()
}

func main() {
	show_panic_recover()
	show_defer_order()
}
