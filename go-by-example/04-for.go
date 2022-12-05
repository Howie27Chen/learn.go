package main

import "fmt"

func show_basic() {

	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i)
	}

	j := 0
	for j < 3 {
		fmt.Println("j = ", j)
		j += 1
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n = ", n)
	}

	for {
		fmt.Println("loop break")
		break
	}
}

func main() {

	show_basic()

}
