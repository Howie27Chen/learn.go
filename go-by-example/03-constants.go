package main

import (
	"fmt"
	"math"
)

type myInt int

const nn myInt = 13

// const m int = nn + 15
// 无类型常量
const m = nn + 15

// iota 实现枚举
const (
	_ = iota
	one
	two
	three
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	// 常量具有严格的数据类型
	//var a int = 5
	//fmt.Println("const:", nn+a)
	fmt.Println(one, two, three)

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(d))
}
