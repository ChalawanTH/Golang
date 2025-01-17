package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println(courseName)
	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)
	courseWeb = courseName[:4]
	fmt.Println(courseWeb)

}
