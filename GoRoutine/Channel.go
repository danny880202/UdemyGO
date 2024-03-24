package main

import (
	"fmt"
	"time"
)

func display(teststring chan string) {
	time.Sleep(2 * time.Second) //子協成休眠2秒
	teststring <- "Hello World"
}

func main() {
	message := make(chan string) //創建通道
	go display(message)          //啟用協程

	//在主協成從通道中讀取數據
	msg := <-message
	fmt.Println(msg)

}
