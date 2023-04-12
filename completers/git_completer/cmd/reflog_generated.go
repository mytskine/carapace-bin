package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:     "reflog",
	Short:   "Manage reflog information",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(reflogCmd).Standalone()

	rootCmd.AddCommand(reflogCmd)
}
