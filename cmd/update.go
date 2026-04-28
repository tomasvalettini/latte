/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/coffeeshop/controller"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update drip",
	Long:  `Command to update drip with specific id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !ShowHelpCmd(cmd, args) {
			cPath := &carafepath.LocalCarafePath{}
			coffeeShopController := controller.NewCoffeeShopController(cPath)
			text := args[0]

			coffeeShopController.UpdateDripInBlend(&controller.BlendIdentifier{
				Id:    flagBlendId,
				Title: flagBlendName,
			}, flagDripId, text)
		}
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
	updateCmd.Flags().StringVar(&flagBlendName, "blend", DEFAULT_FLAG_BLEND_NAME, "The blend to assign this drip to.")
	updateCmd.Flags().IntVar(&flagBlendId, "blendId", DEFAULT_FLAG_ID, "The ID of the blend to assign this drip to.")
	updateCmd.Flags().IntVar(&flagDripId, "dripId", DEFAULT_FLAG_ID, "The ID of the drip to update.")
}
