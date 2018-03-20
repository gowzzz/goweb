package main

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	"io"
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
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer xmlFile.Close()

	//method 1
	/*	xmlData, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		var post Post
		xml.Unmarshal(xmlData, &post)
		fmt.Println(post.Comments)*/

	//method2 stream
	/*
		因为手动解码 XML 文件需要做更多工作，所以这种方法并不适用于处理小型的 XML 文件。
		但如果程序面对的是流式 XML 数据，或者体积非常庞大的 XML 文件，那么解码将是从 XML
		里提取数据唯一可行的办法
	*/
	decoder := xml.NewDecoder(xmlFile)
	for {
		//token 实际上就是一个表示 XML 元素的接口
		t, err := decoder.Token()
		//读到了文件的尾部
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		switch se := t.(type) {
		//判断该 token 是否为 XML 元素的起始标签
		case xml.StartElement:
			//看它是否就是 XML 中的 comment 元素
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}

	}

}
