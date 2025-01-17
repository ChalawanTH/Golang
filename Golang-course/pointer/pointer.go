package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)
	zeroval(i)
	fmt.Println("i form function zeroval =", i)
	zeropointer(&i)
	fmt.Println("i val form function zeropointer =", i)
	fmt.Println("i address form function zeropointer =", &i)
}
