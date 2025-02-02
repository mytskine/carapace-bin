package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var listfilesCmd = &cobra.Command{
	Use:   "listfiles",
	Short: "List files 'owned' by package(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listfilesCmd).Standalone()

	carapace.Gen(listfilesCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
