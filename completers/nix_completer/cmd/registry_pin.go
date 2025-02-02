package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "pin a flake to its current version or to the current version of a flake URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_pinCmd).Standalone()

	registry_pinCmd.Flags().String("registry", "", "The registry to operate on")
	registryCmd.AddCommand(registry_pinCmd)

	addEvaluationFlags(registry_pinCmd)
	addLoggingFlags(registry_pinCmd)

	carapace.Gen(registry_pinCmd).FlagCompletion(carapace.ActionMap{
		"registry": carapace.ActionFiles(),
	})

	// TODO positional completion
}
