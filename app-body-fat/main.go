package main

import "fmt"

func main() {
	var (
		name   string  //= "陈豪坤"
		gender string  //= "男"
		age    int     //= 30
		height float64 //= 1.70
		weight float64 //= 70.0
	)

	for {
		fmt.Print("姓名: ")
		fmt.Scanln(&name)
		fmt.Print("性别（男/女）: ")
		fmt.Scanln(&gender)
		fmt.Print("年龄: ")
		fmt.Scanln(&age)
		fmt.Print("身高（米）: ")
		fmt.Scanln(&height)
		fmt.Print("体重（公斤）: ")
		fmt.Scanln(&weight)

		bmi := weight / (height * height)
		fmt.Println("BMI 指数是: ", bmi)

		var gender_weight float64
		if gender == "男" {
			gender_weight = 1
		} else {
			gender_weight = 0
		}
		bodyfat := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*gender_weight
		fmt.Println("体脂率是: ", bodyfat, "%")

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

		cal_continue := "是"
		fmt.Print("是否继续输入(是/否): ")
		fmt.Scanln(&cal_continue)
		if cal_continue != "否" {
			continue
		} else {
			break
		}
	}
}
