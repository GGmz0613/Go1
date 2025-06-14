package main

import (
	"fmt"
	json "github.com/json-iterator/go"
)

func main() {
	aa := struct {
		Name string `json:"name"`
	}{
		Name: "test",
	}
	ss, err := json.Marshal(aa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(ss))
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")

}
