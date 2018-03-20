package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
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
	post := Post{
		Id:      1,
		Content: "helloword",
		Author: Author{
			Id:   2,
			Name: "wz",
		},
		Comments: []Comment{
			Comment{
				Id:      1,
				Content: "hellow1",
				Author: Author{
					Id:   3,
					Name: "wz3",
				},
			},
			Comment{
				Id:      2,
				Content: "hellow2",
				Author: Author{
					Id:   4,
					Name: "wz4",
				},
			},
			Comment{
				Id:      3,
				Content: "hellow3",
				Author: Author{
					Id:   5,
					Name: "wz5",
				},
			},
		},
	}
	//method1
	/*	// output, err := json.Marshal(&post)
		output, err := json.MarshalIndent(&post, "", "\t\t") //格式化
		if err != nil {
			fmt.Println("err", err)
			return
		}
		err = ioutil.WriteFile("create.json", output, 0644)
		if err != nil {
			fmt.Println("err", err)
			return
		}*/
	//method2
	jsonFile, err := os.Create("create2.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
