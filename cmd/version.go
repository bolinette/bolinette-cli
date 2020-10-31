package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	blntCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Bolinette and Bolinette cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
	},
}
