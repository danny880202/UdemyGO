package main

import variable "UdemyGO/Variable"

type Student struct {
	//結構體成員，也稱為字段
	id     int
	name   string
	age    int
	city   string
	gender string
}

func main() {
	// var a, b = 10, 5 //聲明變量

	// f1 := variable.Calculate("+") //調用函數Calculate
	// f2 := variable.Calculate("-")

	// fmt.Println(f1(a, b)) //調用f1和f2
	// fmt.Println(f2(a, b))

	variable.StringUse()

}
