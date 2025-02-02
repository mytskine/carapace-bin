package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Use the Charm key value store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kvCmd).Standalone()

	rootCmd.AddCommand(kvCmd)
}
