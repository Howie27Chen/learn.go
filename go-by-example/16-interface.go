package main

import "fmt"

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
	animals := []QuackaleAnimal{new(Duck), new(Bird), new(Dog)} // Duck, Bird 和 Dog 没什么联系，但是都表现出了 QuackableAnimal 的特性
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}

func main() {
	interface_basic()
	interface_dynamic()
}
