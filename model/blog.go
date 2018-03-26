package model

import "log"

// Blog struct blog
type Blog struct {
	Model
	Title   string
	Content string
}

//CreateBlog create blog
func CreateBlog(title string, content string) (blog Blog, err error) {
	err = DBCon.QueryRow("INSERT INTO blogs(title, content) VALUES($1,$2) returning id;", title, content).Scan(&blog.ID)
	if err != nil {
		log.Fatal(err)
	}
	return blog, err
}
