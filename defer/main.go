package main

import "fmt"

func main() {
	language := "GO"
	defer fmt.Println(" to " + language)
	language = "Python"
	defer fmt.Println(" from " + language)
	fmt.Println("Hello")
}
