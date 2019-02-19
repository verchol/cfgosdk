package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//ExportedCommand
var ExportedCommand = &cobra.Command{
	Use:   "exported",
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("exported command")
	},
}
