package main

import "fmt"

// 定義結構體
type Rectangle struct {
	height, width int
}

// 計算面積結構體方法
func (rect Rectangle) Area() int {
	return rect.height * rect.width //Rectangle是結構體的名稱, rect是結構體的實例

}

// 計算雙倍面積結構體方法
func (rect Rectangle) DoubleArea() int {
	return rect.Area() * 2
}
func main() {
	//實力化結構體
	rect1 := Rectangle{4, 3}
	fmt.Println("結構體實例: ", rect1)
	fmt.Println("面積: ", rect1.Area())         // 調用結構體方法
	fmt.Println("雙倍面積: ", rect1.DoubleArea()) // 調用結構體方法
}
