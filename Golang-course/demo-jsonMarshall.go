package main

import (
	"encoding/json"
	"fmt"
)

type emplyee struct {
	Id          int
	EmplyeeName string
	Tel         string
	Email       string
}

func main() {
	data, _ := json.Marshal(&emplyee{101, "Pasit Win", "+66", "Pasit@gmail.com"})
	fmt.Println(string(data))
}
