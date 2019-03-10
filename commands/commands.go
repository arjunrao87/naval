// Package commands is where all cli logic is
package commands

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool // Display logs or not

	// NavalCmd - Starter
	NavalCmd = &cobra.Command{
		Use:           "naval",
		Short:         "naval quotes",
		Long:          ``,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {

	// commands
	NavalCmd.AddCommand(randomCmd)
	NavalCmd.AddCommand(listCmd)

}
