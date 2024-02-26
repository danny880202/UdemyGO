package main

import "fmt"

const pi = 3.14159

type Shape interface {
	area() float64
	perrimeter() float64
}

// 定義矩形結構體
type Rectangle struct {
	width, height float64
}

// 定義圓形結構體
type Circle struct {
	radius float64
}

// 聲明矩形結構體面積接口方法
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// 聲明矩形結構體周長接口方法
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// 聲明圓形結構體面積接口方法
func (c Circle) area() float64 {
	return c.radius * c.radius * pi
}

// 聲明圓形結構體周長接口方法
func (c Circle) perimeter() float64 {
	return 2 * c.radius * pi
}

func main() {
	r := Rectangle{width: 10.0, height: 5.0}
	circle := Circle{radius: 5.0}
	fmt.Printf("矩形面積: %0.2f, 矩形周長: %0.2f\n", r.area(), r.perimeter())
	fmt.Printf("圓形面積: %0.2f, 圓形周長: %0.2f\n", circle.area(), circle.perimeter())
}
