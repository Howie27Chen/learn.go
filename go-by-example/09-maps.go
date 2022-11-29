package main

import (
	"fmt"
)

func show_map_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: ", r)
			m := map[string]int{}
			m["key"] = 1
			fmt.Println("map:", m)
		}
	}()
	// 声明和初始化
	var m map[string]int // m map[string]int = nil
	fmt.Println("map 无法像 slice 一样，零值可用，直接操作零值会导致程序 panic")
	m["key"] = 1
	fmt.Println(m)
}

func show_map_init() {
	// 简单的初始化
	m0 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("simple map m0:", m0)
	// 复杂 value 的初始化
	m1 := map[int][]string{
		1: []string{"val_11", "val_12"},
		2: []string{"val_21", "val_22"},
		7: []string{"val_71", "val_72"},
	}
	fmt.Println("complex value map m1:", m1)
	// 复杂 key 的初始化
	type Position struct {
		x float64
		y float64
	}
	m2 := map[Position]string{
		Position{1.0, 2.0}:  "school",
		Position{4.0, 5.0}:  "shopping-mall",
		Position{8.0, 10.0}: "hospital",
	}
	fmt.Println("complex key map m2:", m2)
}

func basic_opr() {
	show_map_panic()
	show_map_init()
}

func main() {

	// s1 := make([]int, 1)
	// s2 := make([]int, 2)
	// fmt.Println(s1 == s2)
	// f1 := func() {}
	// f2 := func() {}
	// fmt.Println(f1 == f2)
	// m1 := make(map[int]string)
	// m2 := make(map[int]string)
	// fmt.Println(m1 == m2)
	basic_opr()

	m := make(map[string]int, 3)

	m["k1"] = 7
	m["k2"] = 20

	fmt.Println("map: ", m)
	fmt.Println("len: ", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("map before delete: ", m)
	delete(m, "k2")
	fmt.Println("map after delete: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

}
