/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	_ "image/jpeg"
	_ "image/png"

	imageutil "github.com/boxy-pug/stegotool/pkg/imageutil"
	"github.com/spf13/cobra"
)

// hideCmd represents the hide command
var hideCmd = &cobra.Command{
	Use:   "hide",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hide called")
		imagePath := args[0]
		message := args[1]

		fmt.Printf("Parsing file: %s\n", imagePath)
		fmt.Printf("Writing message: %s\n", message)
		//for _, arg := range args {
		//	fmt.Println(arg)
		//}

		img := imageutil.StegoImage{
			FilePath: args[0],
			Message:  args[1],
		}

		err := img.LoadImage()
		if err != nil {
			fmt.Printf("error loading image: %s\n", err)
		}

		if verbose {
			fmt.Printf("Loaded image %s\nFormat is %s\n", img.FilePath, img.Format)
			fmt.Printf("Message is: %s\n", img.Message)
		}

		err = img.WriteMessageToImage()
		if err != nil {
			fmt.Printf("error writing to image: %s", err)
		}

		//fmt.Println(img.Image.Bounds())

		err = img.SavePNG()
		if err != nil {
			fmt.Printf("couldn't save image: %s", err)
		}

	},

	//byteHead := img.Pix[:7]

	//fmt.Printf("the first bytes are: %x\n", byteHead)
}

func init() {
	rootCmd.AddCommand(hideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
