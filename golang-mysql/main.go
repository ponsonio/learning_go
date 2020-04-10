package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

func main() {
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Printf("error:" + err.Error())
		panic(err.Error())
	}
	fmt.Printf("connected")
	defer db.Close()
	/**
	insert, err := db.Query("INSERT INTO posts(title) VALUES('My post')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	**/

	posts, err := db.Query("SELECT id, title, body FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for posts.Next() {
		var post Post
		err := posts.Scan(&post.ID, &post.Title, &post.Body)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(post)

	}

}
