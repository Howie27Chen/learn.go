package main

import "fmt"

func main() {

	for {
		name, age, gender, height, weight := getInputINfo()
		bmi := calBMI(height, weight)
		bodyfat := calBodyFat(bmi, age, gender)
		fmt.Printf("%s 体脂率是 %.2f%s\n", name, bodyfat, "%")
		evaluateBodyFat(name, age, gender, bodyfat)
		if inputAgain() {
			break
		}
	}
}

func getInputINfo() (name string, age int, gender string, height float32, weight float32) {

	fmt.Print("姓名: ")
	fmt.Scanln(&name)
	fmt.Print("年龄: ")
	fmt.Scanln(&age)
	fmt.Print("性别（男/女）: ")
	fmt.Scanln(&gender)
	fmt.Print("身高（米）: ")
	fmt.Scanln(&height)
	fmt.Print("体重（公斤）: ")
	fmt.Scanln(&weight)
	return name, age, gender, height, weight

}

func calBMI(height float32, weight float32) float32 {
	return weight / (height * height)
}

func calBodyFat(bmi float32, age int, gender string) float32 {
	var powerGender float32
	if gender == "男" {
		powerGender = 1
	} else {
		powerGender = 0
	}
	return 1.2*bmi + 0.23*float32(age) - 5.4 - 10.8*powerGender
}

func evaluateBodyFat(name string, age int, gender string, bodyfat float32) {
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
}

func inputAgain() bool {
	cal_continue := "是"
	fmt.Print("是否继续输入(是/否): ")
	fmt.Scanln(&cal_continue)
	if cal_continue != "否" {
		return false
	} else {
		return true
	}
}
