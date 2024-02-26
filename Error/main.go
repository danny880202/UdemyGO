package main

import "fmt"

func main() {
	var number, divisor int
	fmt.Println("請輸入一個整數作為分子: ")
	fmt.Scan(&number)
	fmt.Println("請輸入一個整數作為分母: ")
	fmt.Scan(&divisor)

	result := divide(number, divisor)
	fmt.Println("result: ", result)

}

func divide(number, divisor int) float32 {

	return float32(number / divisor)
}
