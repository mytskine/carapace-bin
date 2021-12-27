package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_renameCmd).Standalone()

	addGlobalOptions(context_renameCmd)

	contextCmd.AddCommand(context_renameCmd)
}
