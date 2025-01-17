package main

import "fmt"

func add(val1, val2 float64) {
	result := val1 + val2
	fmt.Println("Result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("I =", i)
	}
}

func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("J =", i)
	}
}
func main() {
	// fmt.Println("Welcome to Calculator")
	// defer fmt.Println("End")
	// defer add(20, 10)
	// defer add(15, 10)
	// defer add(12, 12)
	loop()
	deferloop()
}
