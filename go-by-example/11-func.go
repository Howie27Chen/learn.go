package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// show_var_param 展示函数的可变参数
func show_var_param(sl []int, elements ...int) []int {
	// 可变参数是基于切片实现的
	fmt.Printf("%T\n", elements)
	if len(elements) == 0 {
		println("no element to append")
		return sl
	}
	sl = append(sl, elements...)
	return sl
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", plusPlus(1, 2, 3))

	sl := []int{1, 2, 3}
	fmt.Println("before: sl = ", sl)
	// 可变参数可以是 0  个
	sl = show_var_param(sl)
	fmt.Println("after var none: sl = ", sl)
	// 可变参数可以是多个
	sl = show_var_param(sl, 1, 2, 3)
	fmt.Println("after var 1, 2, 3: sl = ", sl)

}
