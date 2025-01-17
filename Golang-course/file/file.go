package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/zaram/Golang-course/file/customers100.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		//fmt.Println("Line %d : %s\n", count, line)
		fmt.Println(item[4], item[6])
		count++
	}
}
