package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Id          int
	EmplyeeName string
	Tel         string
	Email       string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmplyeeName" : "Pasit Win","Tel":"+66","Email":"Pasit@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.EmplyeeName)
}
