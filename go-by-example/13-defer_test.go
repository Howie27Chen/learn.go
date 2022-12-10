// defer_test.go
package main

import "testing"

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}
	return total
}

func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}

func fooWithoutDefer() {
	sum(10)
}

func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

func BenchmarkFooWithOutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}
}

// 在终端使用以下命令对并函数内部是否使用 defer 的运行时间
// go test -bench . 13-defer_test.go
