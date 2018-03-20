package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

/*
 * DESC: 读取文件的数据已结构体的方式返回
 * PRE : 文件名称
 * POST: 数据和错误码
 */
func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		panic(err)
	}

	return
}
func encoder(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(post)
}
