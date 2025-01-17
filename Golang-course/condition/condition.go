package main

import (
	"fmt"
)

func Grade(val int) string {

	gradeVal := ""
	if val >= 80 {
		gradeVal = "A"
	} else if val >= 70 {
		gradeVal = "B"
	} else if val >= 60 {
		gradeVal = "C"
	} else if val >= 50 {
		gradeVal = "C"
	} else {
		gradeVal = "F"
	}
	return gradeVal

}

var score int

func main() {
	// myMoney := 77
	// if myMoney > 100 {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }
	fmt.Println("Calculator Grade ")
	fmt.Scanf("%d", &score)
	fmt.Println("scoree =", score)

	Result := Grade(score)
	fmt.Println("Result Grade =", Result)
}
