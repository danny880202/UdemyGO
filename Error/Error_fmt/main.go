package main

import (
	"fmt"
)

type DivisionByZero struct {
	//錯誤訊息
	message string
}

// 實現 Error接口中的Error()函數
func (z DivisionByZero) Error() string {
	return "除法錯誤，分母發生錯誤"
}

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
	//& 要搭配defer調用函數
	defer handlePanic()

	//判斷分母為0的情況
	if divisor == 0 {
		panic("分母不能為0!")
		//# 有panic，就不會執行下面的代碼
		//獲得錯誤 error變數
		// err := DivisionByZero{message: "分母不能為0!"}
		// return 0.0, err
	} //# 因為沒有用到errors.New()，所以不用import errors包

	//返回正常結果
	return float32(number / divisor), nil //& 沒有報錯，就是nil
}
func handlePanic() {

	fmt.Println("調用handlePanic函數")
	//從panic狀態中恢復
	state := recover()
	// 判斷狀態
	if state != nil {
		fmt.Println("恢復狀態 ", state)
	}
}
