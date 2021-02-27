package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, _ := sql.Open("sqlite3", "./books.db")

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully created table books")
	}
	statement.Exec()
	dbOperations(db)
}

func dbOperations(db *sql.DB) {
	statement, _ := db.Prepare("INSERT INTO books (name,author,isbn) VALUES (?,?,?)")
	statement.Exec("A Tale of Two Cities", "Charles Dickes", 140430547)
	fmt.Println("Inserted book into the database")

	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities", 1)
	fmt.Println("Succesfullu updated the book on the database")

	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)
	fmt.Println("Successfully delete book from the database")
}
