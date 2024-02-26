package main

import "fmt"

type Student struct {
	//結構體成員，也稱為字段
	id     int
	name   string
	age    int
	city   string
	gender string
}

func main() {

	var stru1 Student //stru1實例化
	stru1.id = 101    //訪問結構體成員
	stru1.name = "小明"
	stru1.age = 18
	stru1.city = "台北"
	stru1.gender = "男"

	stru2 := &Student{102, "小白", 17, "新竹", "女"} //#   &符號表示 stru2 是一個指向 Student 結構體的指針。
	fmt.Printf("stru2姓名: %s\n", stru2.name)     // 訪問成員
	fmt.Printf("stru2年齡: %d\n", stru2.age)

	stru3 := new(Student) //#   new 函數會創建一個新的實例並返回一個指向它的指針。
	fmt.Println(stru3)
	fmt.Printf("stru3 ID: %d\n", stru3.id)
	fmt.Printf("stru3年齡: %d\n", stru3.age) // 訪問成員

}
