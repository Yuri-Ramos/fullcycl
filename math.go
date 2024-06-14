package main

import "fmt"

func main() {
	println("Hello World!")
	fmt.Println(soma(1, 2))
}

func soma(a int, b int) int {
	return a + b
}
