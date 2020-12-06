package command

import "github.com/spf13/cobra"

type root struct{}

func newRoot() *cobra.Command {
	return &cobra.Command{
		Use:     "Api Box",
		Short:   "Core business logic of api box",
		Example: "api_box",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
