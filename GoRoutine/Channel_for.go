package main

import "fmt"

// 定義函數，產生數據，並且發送到通道
func Producer(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	// 創建 int 通道
	ch := make(chan int)

	go Producer(ch) //創建協程
	for v := range ch {
		fmt.Println("接收", v)
	}
}
