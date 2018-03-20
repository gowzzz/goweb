package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(e interface{}, filename string) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(e)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, w.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(e interface{}, filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	r := bytes.NewBuffer(buf)
	dec := gob.NewDecoder(r)
	err = dec.Decode(e)
	if err != nil {
		panic(err)
	}

}

func main() {
	e := Post{Id: 1, Content: "hhhhh", Author: "wzzz"}
	store(e, "p1")

	var postRead Post
	load(&postRead, "p1")
	fmt.Println(postRead)
}
