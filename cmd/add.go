/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/backlog"
	"github.com/tomasvalettini/latte/tasks/controller"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  `Command to add task with auto generated id.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskPath := &backlog.LocalTaskPath{}
		taskController := controller.NewTaskController(taskPath)

		text := args[0]

		taskController.AddTask(text)
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
}
