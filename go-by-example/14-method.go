package main

import "fmt"

// 14-method.go
// go 方法的声明形式
// func (t *T或T) MethodNamd(参数列表) (返回值列表){
//	方法体
//}

// go 方法只用用于同一包内的方法

// go 不能为原生的数据类型定义方法
//func (i int) Foo() string {
//	return fmt.Sprintf("%d", i)
//}

// go 不能跨越包，为其他包的类型声明方法
//func (s http.Server) Bar() {
//
//}

type T struct{}

func (t T) M(n int) {
	fmt.Println("M input ", n)
}

func show_basic_method() {
	var t T
	t.M(8) // 通过类型 T 的变量实例调用方法

	p := &T{} // 通过类型 *T 的变量实例调用方法
	p.M(4)
}

func main() {
	show_basic_method()
}
