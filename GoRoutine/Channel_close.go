package main

import "fmt"

func main() {
	ch := make(chan int, 3) //創建通道
	ch <- 999999
	ch <- 12222
	//關閉通道
	close(ch)

	msg_01 := <-ch
	fmt.Println(msg_01)

	msg_02 := <-ch
	fmt.Println(msg_02)

	msg_03 := <-ch
	fmt.Println(msg_03) //因為通道沒有數據可獨，所以回報0

}
