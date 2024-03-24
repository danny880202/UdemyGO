package main

import "fmt"

func main() {
	myChannel := make(<-chan string) //創建只接收數據的通道

	myChanne2 := make(chan<- string) //創建只發送數據的通道

	myChanne3 := make(chan string) //創建雙向通道
	fmt.Println(myChannel)
	fmt.Printf("myChanne1 :%T\n", myChannel)
	fmt.Println(myChanne2)
	fmt.Printf("myChanne2 :%T\n", myChanne2)
	fmt.Println(myChanne3)
	fmt.Printf("雙向通道 myChanne3 : %T\n", myChanne3)

}
