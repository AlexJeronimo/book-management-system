package main

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJeronimo/book-management-system/internal/book"
)

//TODO: add normal error handling, leave it as feature refactoring tasks

func main() {
	repo := book.Repository{}

	id := repo.GenerateID()
	book := book.NewBook("title", []string{"author"}, " ISBN", true, id)
	fmt.Printf("%+v\n", book)
	data, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
