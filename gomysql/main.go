package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Tag struct
type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil { 
		panic(err.Error())
	}
	for results.Next() {
		var tag Tag 
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
	}

	var tag Tag
	err = db.QueryRow("SELECT id, name FROM tags WHERE id = ?", 1).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error())
	}
	log.Println(tag.ID)
	log.Println(tag.Name)
}