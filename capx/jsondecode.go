package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Id      int    `json:"id,attr"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer jsonFile.Close()

	//method 1
	/*	jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		var post Post
		json.Unmarshal(jsonData, &post)
		fmt.Println(post.Comments)*/

	//method2 stream
	/*
		因为手动解码 XML 文件需要做更多工作，所以这种方法并不适用于处理小型的 XML 文件。
		但如果程序面对的是流式 XML 数据，或者体积非常庞大的 XML 文件，那么解码将是从 XML
		里提取数据唯一可行的办法
	*/
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("err", err)
			return
		}
		fmt.Println(post)
	}

}
