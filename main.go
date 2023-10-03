package main

import (
	variable "UdemyGO/Variable"
	"fmt"
)

func main() {
	var a, b = 10, 5 //聲明變量

	f1 := variable.Calculate("+") //調用函數Calculate
	f2 := variable.Calculate("-")

	fmt.Println(f1(a, b)) //調用f1和f2
	fmt.Println(f2(a, b))
}
