package main

import "fmt"

func main() {
	channel := make(chan int, 2) //有緩衝區通道，他容量大小是 2

	channel <- 123 //向緩衝區通道輸入數據

	channel <- 456

	//channel <- 789
	//因為緩衝區通道只能容納2個數據，所以這裡會報錯，回報deadlock

	fmt.Println(<-channel) //從緩衝區通道讀取數據
	fmt.Println(<-channel)

	//fmt.Println(<-channel)
	//因為緩衝區通道沒有數據可讀，所以這裡會報錯，回報deadlock

}
