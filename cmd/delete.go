/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/coffeeshop/controller"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete drip",
	Long:  `Command to delete drip with specific id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !ShowHelpCmd(cmd, args) {
			cPath := &carafepath.LocalCarafePath{}
			coffeeShopController := controller.NewCoffeeShopController(cPath)

			coffeeShopController.DeleteFromBlends(&controller.BlendIdentifier{
				Id:    flagBlendId,
				Title: flagBlendName,
			}, flagDripId)
		}
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
	deleteCmd.Flags().StringVar(&flagBlendName, "blend", DEFAULT_FLAG_BLEND_NAME, "The blend to assign this drip to.")
	deleteCmd.Flags().IntVar(&flagBlendId, "blendId", DEFAULT_FLAG_ID, "The ID of the blend to assign this drip to.")
	deleteCmd.Flags().IntVar(&flagDripId, "dripId", DEFAULT_FLAG_ID, "The ID of the drip to delete.")
}
