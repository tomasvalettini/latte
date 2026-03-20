/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/coffeeshop/controller"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add drip",
	Long:  `Command to add drip with auto generated id.`,
	Run: func(cmd *cobra.Command, args []string) {
		cPath := &carafepath.LocalCarafePath{}
		coffeeShopController := controller.NewCoffeeShopController(cPath)
		text := args[0]

		coffeeShopController.AddToBlends(&controller.BlendIdentifier{
			Id:    flagBlendId,
			Title: flagBlendName,
		}, text)
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

	addCmd.Flags().StringVar(&flagBlendName, "blend", DEFAULT_FLAG_BLEND_NAME, "The blend to assign this drip to.")
	addCmd.Flags().IntVar(&flagBlendId, "blendId", DEFAULT_FLAG_ID, "The ID of the blend to assign this drip to.")
}
