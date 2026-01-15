/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var id string

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Set book as read",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read called")

		bookId, _ := strconv.Atoi(strings.TrimSpace(id))
		application.Repo.SetRead(bookId)
		application.Repo.Save(bookshelfFile)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&id, "id", "i", "", "set book id")
}
