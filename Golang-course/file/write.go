package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("Hello\n Dev")
	err := os.WriteFile("C:/Users/zaram/Golang-course/file/data.txt", data1, 0644)
	check(err)

	f, err := os.Create("employeeName")
	check(err)
	defer f.Close()

	data2 := []byte("Kiattisak\n Pichaporn")
	os.WriteFile("file/employeeName.txt", data2, 0644)

}
