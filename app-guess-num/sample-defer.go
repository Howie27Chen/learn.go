package main

import (
	"fmt"
	"time"
)

func main() {
	deferGuess()
	panicAndRecover()
}

func deferGuess() {
	startTime := time.Now()
	defer func() {
		fmt.Println("duartion: ", time.Now().Sub(startTime))
	}()
	fmt.Println("start time: ", startTime)
	time.Sleep(5 * time.Second)
}

func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("不能给空的 map 添加元素")
		}
	}()
	var my_may map[string]int = nil
	my_may["Howie"] = 1

}
