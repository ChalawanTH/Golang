package main

import "fmt"

var input string

func main() {
	// input := 4
	// switch input {
	// case 1:
	// 	fmt.Println("One")

	// case 2:
	// 	fmt.Println("Two")

	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("invalid value")
	// }
	fmt.Scan(&input)
	switch input {
	case "blue":
		fmt.Println("#0000FF")

	case "green":
		fmt.Println("#008000")

	case "pink":
		fmt.Println("#FFC0CB")

	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("No color !")

	}

}
