package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var blntCmd = &cobra.Command{
	Use:   "blnt",
	Short: "Bolinette cli is the best way to begin a Bolinette project",
	Long:  "",
}

func Execute() {
	if err := blntCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
