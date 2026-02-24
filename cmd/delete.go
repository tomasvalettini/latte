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
			cmd.PrintErrln("Error reading task id.")
			return
		}

		// duplicated code from list.go
		taskPath := backlog.LocalTaskPath{}
		bl := backlog.NewBacklog(taskPath.GetTaskPath())
		tasks := bl.Load()
		tasksCount := len(tasks)

		if tasksCount <= 0 {
			cmd.Println("No tasks yet.")
			return
		}

		idx := backlog.FindIndexFromId(tasks, id)
		tasks = append(tasks[:idx], tasks[idx+1:]...)
		bl.Save(tasks)

		cmd.Printf("Task with id: %d was successfully removed!!\n", id)
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
