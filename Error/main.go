package main

import (
	"errors" //&  errors包，記得加上s
	"fmt"
)

func main() {
	var number, divisor int
	fmt.Println("請輸入一個整數作為分子: ")
	fmt.Scan(&number)
	fmt.Println("請輸入一個整數作為分母: ")
	fmt.Scan(&divisor)

	result, err := divide(number, divisor)
	if err != nil {
		fmt.Println("發生錯誤: ", err)
	} else {
		fmt.Println("result: ", result)
	}
}

func divide(number, divisor int) (float32, error) {
	//判斷分母為0的情況
	if divisor == 0 {
		//發現error，返回錯誤
		return 0.0, errors.New("分母不能為0!") //& 語法是 errors.New => 還要import errors包
	}
	//返回正常結果
	return float32(number / divisor), nil //& 沒有報錯，就是nil
}
