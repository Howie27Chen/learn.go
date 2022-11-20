package main

import "fmt"

func main() {
	var (
		name   string
		gender string
		age    int
		height float64
		weight float64
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
		fmt.Printf("BMI 指数是: %.2f\n", bmi)

		var gender_weight float64
		if gender == "男" {
			gender_weight = 1
		} else {
			gender_weight = 0
		}
		bodyfat := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*gender_weight
		fmt.Printf("体脂率是: %.2f%s\n", bodyfat, "%")

		if age >= 18 && age <= 39 {
			if bodyfat <= 10 {
				fmt.Println("状态: 偏瘦")
			} else if bodyfat > 10 && bodyfat <= 16 {
				fmt.Println("状态: 标准")
			} else if bodyfat > 16 && bodyfat <= 24 {
				fmt.Println("状态: 偏重")
			} else if bodyfat > 24 && bodyfat <= 26 {
				fmt.Println("状态: 肥胖")
			} else {
				fmt.Println("状态: 严重肥胖")
			}

		} else if age >= 40 && age <= 59 {
			// todo
			fmt.Println("暂时不支持计算: ", age, "岁的人群")
		} else if age >= 60 {
			// todo
			fmt.Println("暂时不支持计算: ", age, "岁的人群")
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
