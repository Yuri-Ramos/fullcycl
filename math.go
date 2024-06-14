package main
import "fmt"


func main(){
	println("Hello World!")
	fmt.Println(Soma(1, 2))
}

func Soma(a int, b int) int {
	return a + b
}