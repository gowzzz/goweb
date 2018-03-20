package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("hello world 1!\n")
	err := ioutil.WriteFile("f1.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	r1,_:=ioutil.ReadFile("f1.txt"
		fmt.Print(string(r1)))
}
