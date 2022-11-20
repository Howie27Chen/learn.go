package main

import "fmt"

func main() {
	var (
		age    int     = 30
		height float64 = 1.70
		weight float64 = 70.0
		gender string  = "男"
	)

	bmi := weight / (height * height)
	fmt.Println("bmi: ", bmi)

	var gender_weight float64
	if gender == "男" {
		gender_weight = 1
	} else {
		gender_weight = 0
	}
	bodyfat := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*gender_weight
	fmt.Println("body fat: ", bodyfat)

	if age >= 18 && age <= 39 {
		if bodyfat <= 10 {
			fmt.Println("偏瘦")
		} else if bodyfat > 10 && bodyfat <= 16 {
			fmt.Println("标准")
		} else if bodyfat > 16 && bodyfat <= 24 {
			fmt.Println("偏重")
		} else if bodyfat > 24 && bodyfat <= 26 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}

	} else if age >= 40 && age <= 59 {
		// todo
	} else if age >= 60 {
		// todo
	} else {
		fmt.Println("超出计算范围")
	}
}
