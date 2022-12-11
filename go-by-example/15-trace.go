package main

import (
	"fmt"
	"runtime"
)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Println("enter:", name)
	return func() {
		fmt.Println("exit", name)
	}

}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main() {
	defer Trace()()
	foo()
}
