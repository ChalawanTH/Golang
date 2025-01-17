package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput(val string) float64 {
	val = strings.TrimSpace(val) // Trim whitespace, including newlines
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("%v must be a number only\n", val)
		panic(err)
	}
	return num
}

func main() {
	var input int
	fmt.Print("คุณต้องการคำนวนกี่ตัวเลข: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการรับค่า:", err)
	} else {
		var result float64

		for i := 0; i < input; i++ {
			countNumber := i + 1
			var inputValue string

			if countNumber == 1 {
				fmt.Print("Number", countNumber, ": ")
				fmt.Scan(&inputValue) // Use fmt.Scanf to read input
				value := getInput(inputValue)
				result = value
			} else {

				fmt.Print("Operation", countNumber-1, ": ")
				var op string
				fmt.Scan(&op) // Get the operator
				opval := op
				fmt.Print("Number", countNumber, ": ")
				fmt.Scan(&inputValue) // Use fmt.Scanf to read input
				value := getInput(inputValue)
				switch opval { // Perform calculation based on the operator
				case "+":
					result += value
				case "-":
					result -= value
				case "*":
					result *= value
				case "/":
					if value == 0 {
						panic("division by zero")
					}
					result /= value
				default:
					panic("invalid operator")
				}
			}
		}

		fmt.Println("ผลลัพธ์:", result)
	}
}
