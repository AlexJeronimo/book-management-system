/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var wl bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all books in the shelf",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		//application.Repo.List()
		application.Repo.PrintTable(wl)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&wl, "whishlist", "w", false, "show books from whishlist")
}
