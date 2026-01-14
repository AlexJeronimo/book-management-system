/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlexJeronimo/book-management-system/internal/shelf"
	"github.com/spf13/cobra"
)

var name, isbn, read string
var authors []string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the book to the shelf",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		// add book to the shelf
		// save shelf to the file
		// add command should reveive subcommands that will match book data
		// such as title, autors, isbn if exists (if no it should be empty), is book read (bool value, yes or no)
		id := application.Repo.GenerateID()
		book := shelf.NewBook(id, name, authors, isbn, read)
		application.Repo.Add(book)
		application.Repo.Save(bookshelfFile)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().StringVarP(&name, "title", "t", "book title", "Book name")
	addCmd.Flags().StringSliceVarP(&authors, "author", "a", []string{}, "Book author")
	addCmd.Flags().StringVarP(&isbn, "isbn", "i", "book isbn", "ISBN")
	addCmd.Flags().StringVarP(&read, "read", "r", "book read", "is book read")

}
