package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "Macbook"
	productName[1] = "iPad"
	productName[2] = "iPhone"
	productName[3] = "AirPods"

	price := [4]float32{40000, 19999, 39990, 3990}
	fmt.Println(productName)
	fmt.Println(price)
}
