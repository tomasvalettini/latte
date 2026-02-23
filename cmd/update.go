/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/backlog"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task",
	Long:  `Command to update task with specific id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Id entered is not a number.")
		}

		newText := args[1]

		bl := backlog.NewBacklog(taskFilePath())
		tasks := bl.Load()

		index := backlog.FindIndexFromId(tasks, id)

		tasks[index].Text = newText
		bl.Save(tasks)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
