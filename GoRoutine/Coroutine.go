package main

import (
	"fmt"
	"time"
)

func display(str string) {
	fmt.Println(str)

}

func main() {
	go func(str string) {
		time.Sleep(5 * time.Second) // 匿名協程休眠5秒
		fmt.Println(str)
	}("匿名函數練習:歡迎來到台灣") // 匿名函數

	time.Sleep(2 * time.Second) // 主協程休眠2秒
	display("Hello World")      // 在主協程中調用

	//go display("歡迎來到台灣") // 在子協程中調用

}
