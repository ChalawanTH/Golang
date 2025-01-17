package main

import "fmt"

func hello() {
	fmt.Println("Hello Dev")
}

func plus(val1 int, val2 int) int {
	result := val1 + val2
	return result
}

// Return หลาย ๆ ค่า
func plus3val(val1 int, val2 int, val3 int) int {
	return val1 + val2 + val3
}
func main() {
	hello()

	result := plus(2, 3)
	result2 := plus3val(2, 3, 5)
	fmt.Println("Result =", result)
	fmt.Println("Result =", result2)
}
