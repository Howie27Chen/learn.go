package main

import (
	"fmt"
	"unsafe"
)

// slice 内部实现
type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int            // 切片的长度
	cap   int            // 切片的容量，即底层数组的长度，cap 永远大于等于 len
}

func main() {

	s1 := make([]byte, 6, 10) // len = 6, cap = 10
	fmt.Println(s1)

	s2 := make([]byte, 6) // len = cap = 6
	fmt.Println(s2)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := arr[3:7:9] // 采用 arr[low:high:max] 语法基于一个已存在的数组创建切片->数组的切片化
	fmt.Println("create slice from arr:", sl)
	fmt.Println("len:", len(sl), "cap:", cap(sl))
	sl[0] += 10
	fmt.Println("arr[3]:", sl[0])

	var sd []int
	for i := 0; i < 5; i++ {
		sd = append(sd, i)
		fmt.Println("len:", len(sd), "cap:", cap(sd)) // 因为初始化是 int 切片，当其增长超过 4 时，就会触发切片的自动扩容
	}

	u := [3]int{1, 2, 3}
	su := u[0:1]
	su[0] = 8
	fmt.Println("before u[0]", u[0])
	for i := 0; i < len(u)+1; i++ {
		su = append(su, i)
		fmt.Println("su len:", len(su), "su cap:", cap(su)) // 当 slice 增长超过其基于创建的 arr 的最大容量时，会创建并重新绑定自动扩容形成的 arr2
	}
	su[0] = 9 // 基于内部新创建的 arr 的 slice 切片的修改，不会影响到之前绑定的 arr 的内存数值
	fmt.Println("after u[0]", u[0])

	s := make([]string, 3)
	fmt.Println("empty slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[2:]
	fmt.Println("sl2: ", l)

	l = s[:5]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "t"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}
