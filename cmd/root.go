package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(hogeCmd)
	rootCmd.AddCommand(fugaCmd)
}
