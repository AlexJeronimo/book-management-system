/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AlexJeronimo/book-management-system/internal/shelf"
	"github.com/spf13/cobra"
)

var (
	query        string
	searchTitle  bool
	searchAuthor bool
	searchISBN   bool

	readOnly       bool
	unreadOnly     bool
	wishlistOnly   bool
	noWishlistOnly bool
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a book by title, author, isbn, from read list or from whishlist",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("search called")

		filter := shelf.SearchFilter{
			Query:  query,
			Title:  searchTitle,
			Author: searchAuthor,
			ISBN:   searchISBN,
		}

		// Read / Unread
		if readOnly && !unreadOnly {
			v := true
			filter.Read = &v
		}
		if unreadOnly && !readOnly {
			v := false
			filter.Read = &v
		}

		// Wishlist / No wishlist
		if wishlistOnly && !noWishlistOnly {
			v := true
			filter.Whishlist = &v
		}
		if noWishlistOnly && !wishlistOnly {
			v := false
			filter.Whishlist = &v
		}

		results := application.Repo.Search(filter)

		shelf.PrintResult(results)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&query, "query", "q", "", "Search text")

	searchCmd.Flags().BoolVar(&searchTitle, "title", true, "Search by title")
	searchCmd.Flags().BoolVar(&searchAuthor, "author", false, "Search by author")
	searchCmd.Flags().BoolVar(&searchISBN, "isbn", false, "Search by ISBN")

	searchCmd.Flags().BoolVar(&readOnly, "read", false, "Show only read books")
	searchCmd.Flags().BoolVar(&unreadOnly, "unread", false, "Show only unread books")

	searchCmd.Flags().BoolVar(&wishlistOnly, "wishlist", false, "Only wishlist books")
	searchCmd.Flags().BoolVar(&noWishlistOnly, "no-wishlist", false, "Exclude wishlist books")
}
