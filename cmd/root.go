/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "latte",
	Short: "Little Actions Toward Tangible Éndings",
	Long: `latte is a small, terminal-based task tracker focused on steady progress through simple, repeatable actions.

    LATTÉ — Little Actions Toward Tangible Éndings

No dashboards.
No notifications.
No productivity theater.

Just a quiet tool to keep track of what matters next.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.latte.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: extract this to be able to unit test with a tmp location
func itemFilePath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".latte", "items.json")
}
