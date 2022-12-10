package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

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

// 函数可以作为一个变量
var myPrintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
	fmt.Println("myPrint")
	return fmt.Fprintf(w, format, a...)
}

func show_func() {
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
	// 函数可以当作变量使用
	fmt.Printf("func type: %T\n", myPrintf)
	myPrintf(os.Stdout, "我的名字: %s\n", "Heiko")
}

func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func partialTimes(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func show_create_error(i int) {
	err := errors.New("your first demo error")
	errWithCtx := fmt.Errorf("index %d is out of bounds", i)
	fmt.Println(err.Error())
	fmt.Println(errWithCtx.Error())
}

func compare_err_link() {
	errSentinel := errors.New("the underlying sentinel error")

	err1 := fmt.Errorf("wrap sentinel: %w", errSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)

	fmt.Println("errSentinel:", errSentinel)
	fmt.Println("err1:       ", err1)
	fmt.Println("err2:       ", err2)
	println("err2 == errSentinel? ->", err2 == errSentinel)

	if errors.Is(err2, errSentinel) {
		println("err2 is sentinel")
	} else {
		println("err2 is sentinel")
	}
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func find_err_link() {

	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e *MyError
	fmt.Println("err:  ", err)
	fmt.Println("err1: ", err1)
	fmt.Println("err2: ", err2)
	fmt.Println("e:    ", e)

	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		fmt.Println("err == e?", err == e)
		fmt.Println("e:    ", e)
		return
	}
	println("MyError is not on the chain of err2")
}

func main() {
	show_func()

	teardown := setup("demo")
	defer teardown()
	println("do some business stuff")

	timeTwo := partialTimes(2)
	timeThree := partialTimes(3)
	fmt.Println("time two:", timeTwo(5))
	fmt.Println("time two:", timeTwo(6))
	fmt.Println("time three:", timeThree(5))
	fmt.Println("time three:", timeThree(6))

	show_create_error(8)
	compare_err_link()
	find_err_link()
}
