package main

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}
type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}
type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "helloword",
		Author: Author{
			Id:   "2",
			Name: "wz",
		},
		Comments: []Comment{
			Comment{
				Id:      "1",
				Content: "hellow1",
				Author: Author{
					Id:   "3",
					Name: "wz3",
				},
			},
			Comment{
				Id:      "2",
				Content: "hellow2",
				Author: Author{
					Id:   "4",
					Name: "wz4",
				},
			},
			Comment{
				Id:      "3",
				Content: "hellow3",
				Author: Author{
					Id:   "5",
					Name: "wz5",
				},
			},
		},
	}
	//method1
	/*	// output, err := xml.Marshal(&post)
		output, err := xml.MarshalIndent(&post, "", "\t") //格式化
		if err != nil {
			fmt.Println("err", err)
			return
		}
		//err = ioutil.WriteFile("create.xml", output, 0644)
		err = ioutil.WriteFile("create.xml", []byte(xml.Header+string(output)), 0644) //加头信息
		if err != nil {
			fmt.Println("err", err)
			return
		}*/
	//method2
	xmlFile, err := os.Create("create2.xml")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
