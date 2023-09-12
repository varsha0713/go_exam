package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var postid int
	fmt.Print("enter the post id:")
	fmt.Scan(&postid)

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postid)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed %s", err)
		return
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("failed to get parse %s", err)
		return
	}
	fmt.Printf("userid:%d\n", post.UserId)
	fmt.Printf("id:%d\n", post.Id)
	fmt.Printf("title:%s\n", post.Title)
	fmt.Printf("body:%s\n", post.Body)

}
