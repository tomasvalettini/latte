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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: move this to a controller type component
		bl := backlog.NewBacklog(itemFilePath())
		items := bl.Load()
		itemsCount := len(items)

		if itemsCount <= 0 {
			cmd.Println("No items yet.")
		}

		cmd.Println("===========")
		cmd.Println(" TASK LIST ")
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
