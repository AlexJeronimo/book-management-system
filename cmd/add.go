/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AlexJeronimo/book-management-system/internal/shelf"
	"github.com/spf13/cobra"
)

var name, isbn, storelink, genre string
var authors []string
var read, whishlist bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the book to the shelf",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")

		id := application.Repo.GenerateID()
		book := shelf.NewBook(id, name, authors, isbn, read, genre)

		if whishlist {
			book.AddToWhishlist(storelink)
		}

		application.Repo.Add(book)
		application.Repo.Save(bookshelfFile)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "title", "t", "", "Book name")
	addCmd.Flags().StringSliceVarP(&authors, "author", "a", []string{}, "Book author")
	addCmd.Flags().StringVarP(&isbn, "isbn", "s", "", "ISBN")
	//addCmd.Flags().StringVarP(&read, "read", "r", "no", "is book read")
	addCmd.Flags().BoolVarP(&read, "read", "r", false, "is book read")
	addCmd.Flags().BoolVarP(&whishlist, "whishlist", "w", false, "add book to whishlist")
	addCmd.Flags().StringVarP(&storelink, "storelink", "l", "", "add storelink")
	addCmd.Flags().StringVarP(&genre, "genre", "g", "", "add book genre")

}
