package main

import (
	"fmt"
	"time"
)

func main() {
	//建立兩個通道
	c1 := make(chan string)
	c2 := make(chan string)

	//建立匿名協程_1
	go func() {
		//休眠2秒
		time.Sleep(2 * time.Second)

		//發送數據
		c1 <- "One"
	}()

	//建立匿名協程_2
	go func() {
		//休眠2秒
		time.Sleep(4 * time.Second)

		//發送數據
		c2 <- "Two"
	}()

	//建立for循環走訪通道
	for i := 0; i < 2; i++ {
		select {
		//從通道c1讀取數據
		case msg_1 := <-c1:
			fmt.Printf("接收: %v\n", msg_1)

		//從通道c2讀取數據
		case msg_2 := <-c2:
			fmt.Printf("接收: %v\n", msg_2)
		}
	}
}
