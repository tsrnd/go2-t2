package model

import (
	"database/sql"
	"log"
)

//DBCon dbcon
var DBCon *sql.DB
var err error

// Blog type struct
type Blog struct {
	ID      int
	Title   string
	Content string
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//SetDatabase return DBCon
func SetDatabase(database *sql.DB) {
	DBCon = database
}

// GetAllPosts return all posts
func GetAllPosts() []Blog {

	rows, e := DBCon.Query(
		"SELECT id, title, content FROM blogs ORDER BY id DESC;")
	checkErr(e)

	blogs := []Blog{}
	for rows.Next() {
		blg := Blog{}
		rows.Scan(&blg.ID, &blg.Title, &blg.Content)
		blogs = append(blogs, blg)
	}

	return blogs
}
