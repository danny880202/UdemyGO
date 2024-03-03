package main

import (
	"fmt"
	"math"
)

// 判斷一個數字是否為質數
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 計算在 x 和 y 之間的質數的數量
func countPrimes(x, y int) int {
	count := 0
	for i := x; i <= y; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	x := 1
	y := 30
	fmt.Printf("在 %d 和 %d 之間的質數的數量是 %d\n", x, y, countPrimes(x, y))
}
