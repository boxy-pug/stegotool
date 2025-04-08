/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stegotool",
	Short: "A CLI tool for hiding and reading messages in images using steganography",
	Long: `StegoTool is a CLI tool for hiding and reading messages in images using steganography.

Features:
-  Hide messages in PNG images using symmetric and asymmetric encryption.
-  Read hidden messages from PNG images.
-  Support for various image formats (PNG, JPEG, etc.) in future updates.

Usage:
  stegotool hide <image_path> <message> <output_path>
  stegotool read <image_path>

Examples:
  Hide a message in an image:
    $ stegotool hide input.png "Secret message" output.png

  Read a hidden message from an image:
    $ stegotool read output.png`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
