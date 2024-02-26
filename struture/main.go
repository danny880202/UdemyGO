package main

import "fmt"

//*結構體裡面存在其他的結構體
type Books struct { //定義圖書結構體
	number         string
	title          string
	price          float32
	authers        []Author //# 定義作者結構體陣列，因為作者有很多位，所以是陣列型態
	publisher_data Publisher
}
type Author struct { // 定義作者結構體
	name string
	id   int
}

type Publisher struct { //定義出版商結構體
	name  string
	email string
}

func main() {
	book_01 := Books{
		"23132132131", "go lang練習", 15200, []Author{
			{"小明", 1}, // 一組Author結構體{"name", id}
			{"小白", 2}, // 第二組Author結構體{"name", id}
		},
		Publisher{"台灣出版社", "123456@gmail.com"},
	}
	fmt.Printf("書號: %s, 書名: %s, 第二作者: %s, 出版社: %s\n", book_01.number,
		book_01.title, book_01.authers[1].name,
		book_01.publisher_data.name)
}
