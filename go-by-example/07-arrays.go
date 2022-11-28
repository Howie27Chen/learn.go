package main

import (
	"fmt"
	"unsafe"
)

// 数组定义：Go 语言的数组是一个长度固定，由同构类型元素组成的连续序列
// var arr [N]T

func foo(arr [5]int) {
}

func main() {
	var a [5]int

	var arr1 [5]int
	// var arr2 [6]int
	//var arr3 [5]string
	foo(arr1)
	// foo(arr2) // 长度属性不一致
	//foo(arr3) // 元素类型属性不同
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度: ", len(arr))
	fmt.Println("数组大小: ", unsafe.Sizeof(arr))

	// ... 忽略数组长度
	var arr_init = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("type: %T\n", arr_init) // [6]int

	fmt.Println("empty 5 int arrays: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD: ", twoD)
}
