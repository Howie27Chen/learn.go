package main

import (
	"fmt"
	"reflect"
)

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

// go 方法的本质是 以 recerver 参数作为第一个参数的普通函数
// func (t T)Method() <=> func Func(t T)
// func (t *T)Method() <=> func Func(t *T)
type Receiver struct {
	i int
}

// Receiver 类型的接口并不会改变原有的实例
func (r Receiver) FakeSet(i int) {
	fmt.Println("FakeSet i=", i)
	r.i = i
}

// *Receiver 类型的接口的方法会改变原有的实例
func (r *Receiver) RealSet(i int) {
	fmt.Println("RealSet i=", i)
	r.i = i
}

// 一般不需要改变原有实例的，我们都采用 T 方式而不是 *T 指针的方式进行传递
func (r Receiver) Get() int {
	return r.i
}

func show_method_ptr() {
	var instance Receiver

	fmt.Println("receiver i = ", instance.i)
	instance.FakeSet(20)
	fmt.Println("receiver i = ", instance.i)
	instance.RealSet(30)
	fmt.Println("receiver i = ", instance.i)
	fmt.Println("receiver Get ", instance.Get())
}

// go 中任何一个类型都有自己的方法集合，或者说方法集合是 Go 类型的一个属性

// dumpMethodSet 检查类型的方法集合
func dumpMethodSet(i interface{}) {

	dynamicType := reflect.TypeOf(i)

	if dynamicType == nil {
		fmt.Println("there is no dynamic type")
		return
	}

	n := dynamicType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty.\n", dynamicType)
		return
	}

	fmt.Printf("%s's method set: \n", dynamicType)
	for i := 0; i < n; i++ {
		fmt.Println("-", dynamicType.Method(i).Name)
	}
	fmt.Printf("\n")

}

type MethodSet struct{}

func (m MethodSet) Method1()  {}
func (m MethodSet) Method2()  {}
func (m *MethodSet) Method3() {}
func (m *MethodSet) Method4() {}

func show_method_set() {

	// Go 原生类型没有定义方法，其方法集合都是空的
	var a int
	dumpMethodSet(a)
	dumpMethodSet(&a)

	var set MethodSet
	dumpMethodSet(set)
	dumpMethodSet(&set)

}

func main() {
	show_basic_method()
	show_method_ptr()
	show_method_set()
}
