package main

import (
	"cfsdk/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(args)
	},
}
var justCmd = &cobra.Command{
	Use:   "just",
	Short: "it's just ",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("%s %s", "just", args)
	},
}

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	rootCmd.AddCommand(justCmd)
	rootCmd.AddCommand(cmd.ExportedCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
