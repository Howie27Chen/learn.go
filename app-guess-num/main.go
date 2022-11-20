package main

import "fmt"

func guessNum(begin, end int) {

	target := (begin + end) / 2
	fmt.Println("参数字，猜大了输入 1， 猜小了输入 2，猜中了输入 9: ", target)

	command := 0
	fmt.Scanln(&command)

	switch command {
	case 1:
		if begin == target {
			fmt.Println("我猜出的数字是:", target, "你是不是改变主意了?")
			return
		}
		guessNum(begin, target-1)
	case 2:
		if end == target {
			fmt.Println("我猜出的数字是:", target, "你是不是改变主意了?")
			return
		}
		guessNum(target+1, end)
	case 9:
		fmt.Println("你心里想的数字是:", target, "我猜对了吗?")
		return
	}
}

func main() {

	guessNum(0, 100)

}
