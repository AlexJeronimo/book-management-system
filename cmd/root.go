package cmd

import (
	"fmt"
	"os"

	"github.com/AlexJeronimo/book-management-system/internal/app"
	"github.com/spf13/cobra"
)

var application *app.App

const (
	bookshelfFile = "data/.bookshelf.json"
)

// Load list of books from the shelfFile
// If file does not esists create empty shelf
//Pass it to the cobra cli

var rootCmd = &cobra.Command{
	Use:   "bms",
	Short: "Book Management System",
	Long:  "CLI tool to manage your personal book library",
}

func Execute(app *app.App) {
	application = app
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
