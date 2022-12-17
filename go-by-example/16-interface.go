package main

import "fmt"

// Todo 1: 接口的基本使用

type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func interface_basic() {
	var t T
	var i interface{} = t

	v1, ok := i.(MyInterface)

	if !ok {
		panic("the value of i is not MyInterface")
	}
	v1.M1()
	fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T

	i = int64(13)
	v2, ok := i.(MyInterface)
	fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>

	// v2 = 13 // will panic for v2 is nil, not the type of MyInterface
}

// Todo 2: 接口的静态(类型)和动态(调用)特性

type QuackaleAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(animal QuackaleAnimal) {
	animal.Quack()
}

func interface_dynamic() {
	println("================ Here is interface_dynamic() ================")
	animals := []QuackaleAnimal{new(Duck), new(Bird), new(Dog)} // Duck, Bird 和 Dog 没什么联系，但是都表现出了 QuackableAnimal 的特性
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}

// Todo 3: nil 接口类型变量

func printNilInterface() {
	println("================ Here is printnilinterface() ================")
	// nil 接口变量
	var i interface{}      // 空接口类型
	var err error          // 非空接口类型，-> 未赋初值的接口，也是 nil
	println("i = ", i)     // (e._type, e.data)
	println("err = ", err) // (i.tab, i.data)
	println("i == nil", i == nil)
	println("err == nil", err == nil)
	println("i == err", i == err)
}

// Todo 4: 空接口类型变量
func printEmptyInterface() {
	println("================ Here is printEmptyInterface() ================")
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2", eif1 == eif2) // false

	eif2 = 17 // eif1 var -> 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2", eif1 == eif2) // true

	eif2 = int64(17) // eif1 type -> int
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2", eif1 == eif2) // false
}

// Todo 5: 非空接口类型变量

// type T int
func (t T) Error() string {
	return "bad errror"
}

func printNonEmptyInterface() {
	println("================ Here is printNonEmptyInterface() ================")
	var err1 error
	var err2 error
	err1 = (*T)(nil)
	println("err1: ", err1)
	println("err1 = nil", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1: ", err1)
	println("err2: ", err2)
	println("err1 = err2", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1: ", err1)
	println("err2: ", err2)
	println("err1 = err2", err1 == err2)
}

// Todo 6: 空接口类型与非空接口类型变量的等值比较

func printEmptyInterfaceAndNoneEmptyInterface() {
	println("================ Here is printEmptyInterfaceAndNoneEmptyInterface() ================")
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif == err", eif == err)

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("err == eif", err == eif)
}

func main() {
	interface_basic()
	interface_dynamic()

	printNilInterface()
	printEmptyInterface()
	printNonEmptyInterface()
	printEmptyInterfaceAndNoneEmptyInterface()
}
