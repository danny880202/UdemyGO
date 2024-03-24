package main

import (
	"fmt"
	"time"
)

// 建立子協程
func Process(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Hello World"
}
func main() {
	ch := make(chan string) //創建通道
	go Process(ch)          //啟用協程
	for {
		time.Sleep(1 * time.Second)
		select {
		case v := <-ch:
			fmt.Printf("接受到數據!\n%s\n", v)
			break //&跳出for循環，返回main函數，目前的協程寫在main函數，所以相當於終止main函數。
			//# 不能使用break，因為break只能跳出select
		default:
			fmt.Printf("沒有接收到數據!\n")
		}
	}
}
