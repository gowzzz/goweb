package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("hello world 1!\n")

	//方式1 使用ioutil读写文件  简单
	err := ioutil.WriteFile("f1.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	r1, _ := ioutil.ReadFile("f1.txt")
	fmt.Print(string(r1))

	//方式2 使用os读写文件  灵活，可以选择读取的指定位置
	fw, _ := os.Create("f2.txt")
	defer fw.Close()
	b, _ := fw.Write(data)
	fmt.Printf("write num:%d", b)

	fr, _ := os.Open("f2.txt")
	defer fr.Close()
	r2 := make([]byte, len(data))
	b2, _ := fr.Read(r2)
	fmt.Printf("read num:%d\n", b2)
	fmt.Println(string(r2))
}
