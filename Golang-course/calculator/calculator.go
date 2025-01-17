package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Println("%v", promt)
	input, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		msg, _ := fmt.Scanf("%v must number only", promt)
		panic(msg)
	}
	return val
}

func getOperator() string {
	fmt.Print("operator is ( + - * / ) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(val1, val2 float64) float64 {
	return val1 + val2
}
func subtract(val1, val2 float64) float64 {
	return val1 - val2
}
func multiply(val1, val2 float64) float64 {
	return val1 * val2
}
func divide(val1, val2 float64) float64 {
	return val1 / val2
}
func main() {
	val1 := getInput("Value1 =")
	val2 := getInput("Value2 =")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(val1, val2)
	case "-":
		result = subtract(val1, val2)
	case "*":
		result = multiply(val1, val2)
	case "/":
		result = divide(val1, val2)
	default:
		panic("wrong operator")
	}
	fmt.Printf("result is %v", result)
}
