/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomasvalettini/latte/coffeeshop/controller"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list drips",
	Long:  `Command to list all drips with corresponding id.`,
	Run: func(cmd *cobra.Command, args []string) {
		cPath := &carafepath.LocalCarafePath{}
		coffeeShopController := controller.NewCoffeeShopController(cPath)
		bi := &controller.BlendIdentifier{
			Id:    flagBlendId,
			Title: flagBlendName,
		}

		coffeeShopController.ListBlends(bi)
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

	listCmd.Flags().StringVar(&flagBlendName, "blend", DEFAULT_FLAG_BLEND_NAME, "The blend to assign this drip to.")
	listCmd.Flags().IntVar(&flagBlendId, "blendId", DEFAULT_FLAG_ID, "The ID of the blend to assign this drip to.")
}
