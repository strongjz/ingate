package cmd

import (
	"github.com/kubernetes-sigs/ingate/internal/cmd/version"
	"github.com/spf13/cobra"
)

func GetDataPlaneCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"start", "s"},
		Short:   "Start InGate controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			return version.Print(cmd.OutOrStdout())
		},
	}

	return cmd
}
