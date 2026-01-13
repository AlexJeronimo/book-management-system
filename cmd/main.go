package main

import (
	"encoding/json"
	"fmt"

	"github.com/AlexJeronimo/book-management-system/internal/book"
)

const (
	bookshelfFile = "data/.bookshelf.json"
)

//TODO: add normal error handling, leave it as feature refactoring tasks

func main() {
	repo := book.Repository{}

	id := repo.GenerateID()
	book1 := book.NewBook("title", []string{"author"}, "ISBN", true, id)
	fmt.Printf("%+v\n", book1)
	data, err := json.MarshalIndent(book1, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	repo.Add(book1)

	err = repo.Save(bookshelfFile)
	if err != nil {
		fmt.Println("can't save data to json file. Error: ", err)
	}

	newRepo := book.NewRepository()

	newRepo.Load(bookshelfFile)

	fmt.Printf("%+v\n", newRepo)

}
