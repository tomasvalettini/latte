/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/backlog"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	Long:  `Command to delete task with specific id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErrln("Error reading item id.")
			return
		}

		// duplicated code from list.go
		bl := backlog.NewBacklog(itemFilePath())
		items := bl.Load()
		itemsCount := len(items)

		if itemsCount <= 0 {
			cmd.Println("No items yet.")
			return
		}

		idx := backlog.FindIndexFromId(items, id)
		items = append(items[:idx], items[idx+1:]...)
		bl.Save(items)

		cmd.Printf("Item with id: %d was successfully removed!!\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
