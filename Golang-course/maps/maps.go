package main

import "fmt"

var product = make(map[string]float64)

func main() {

	//Add Maps
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 19990

	//Delete Maps
	delete(product, "iPad")

	// Update Maps
	product["iPhone"] = 29990

	fmt.Println("Product =", product)

	value1 := product["iPhone"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName =", courseName)

}
