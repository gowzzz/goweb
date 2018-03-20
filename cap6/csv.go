package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvfile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "aa", Author: "wz1"},
		Post{Id: 2, Content: "bb", Author: "wz2"},
		Post{Id: 3, Content: "cc", Author: "wz3"},
		Post{Id: 4, Content: "dd", Author: "wz4"},
		Post{Id: 5, Content: "ee", Author: "wz5"},
	}

	w := csv.NewWriter(csvfile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := w.Write(line)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()

	fr, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	reader := csv.NewReader(fr)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts)

}
