package main

import (
	"github.com/AlexJeronimo/book-management-system/cmd"
	"github.com/AlexJeronimo/book-management-system/internal/app"
	"github.com/AlexJeronimo/book-management-system/internal/shelf"
)

const (
	bookshelfFile = "data/.bookshelf.json"
)

//TODO: add normal error handling, leave it as feature refactoring tasks

func main() {
	repo := shelf.NewRepository()
	_ = repo.Load(bookshelfFile)

	app := &app.App{
		Repo: repo,
	}

	cmd.Execute(app)
}
