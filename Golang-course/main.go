package main

import "fmt"

var numInt2, numInt1 int = 1000, 2000
var msg string = "Hello"

func main() {
	fmt.Println("Hello Devdddddd")
	numFloat := 25.4
	fmt.Println(numInt1)
	fmt.Println(numInt2)
	fmt.Println(numFloat)
	fmt.Println(numInt1 + numInt2)
	fmt.Println(float64(numInt1+numInt2) + (numFloat))
	fmt.Println(msg + " World")
	fmt.Println(msg, numInt1)

}
