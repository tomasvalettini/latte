/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/backlog"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `Command to list all tasks with corresponding id.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: move this to a controller type component
		bl := backlog.NewBacklog(itemFilePath())
		items := bl.Load()
		itemsCount := len(items)

		if itemsCount <= 0 {
			cmd.Println("No items yet.")
			return
		}

		cmd.Println("===========")
		cmd.Println(" ITEM LIST ")
		cmd.Println("===========")

		w := backlog.MaxIdWidth(items)
		for _, t := range items {
			cmd.Printf("  [%*d]  %s\n", w, t.Id, t.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
