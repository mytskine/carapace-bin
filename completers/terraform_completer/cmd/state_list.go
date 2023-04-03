package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_listCmd = &cobra.Command{
	Use:   "list [options] [address...]",
	Short: "List resources in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_listCmd).Standalone()

	state_listCmd.Flags().StringS("id", "id", "", "Filters the results by id")
	state_listCmd.Flags().StringS("state", "state", "", "Path to a Terraform state file")
	stateCmd.AddCommand(state_listCmd)

	// TODO id completion
	carapace.Gen(state_listCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(state_listCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources(state_listCmd).Invoke(c).Filter(c.Args).ToMultiPartsA("_")
		}),
	)
}
