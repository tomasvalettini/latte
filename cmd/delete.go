/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/drips/controller"
	drippath "github.com/tomasvalettini/latte/drips/path"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete drip",
	Long:  `Command to delete drip with specific id.`,
	Run: func(cmd *cobra.Command, args []string) {
		dripPath := &drippath.LocalDripPath{}
		dripController := controller.NewDripController(dripPath)

		dripController.DeleteDrip(args[0])
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
