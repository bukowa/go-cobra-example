package frontend

import (
	"errors"
	"github.com/spf13/cobra"
)

var (
	Opt1 string
	Opt2 bool
)

var Cmd = &cobra.Command{
	Use:   "cmd2",
	Short: "cmd2",
	RunE: func(cmd *cobra.Command, args []string) error {
		if Opt1 == "" {
			return errors.New("invalid")
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVarP(&Opt1, "opt1", "a", "", "")
	Cmd.Flags().BoolVarP(&Opt2, "opt2", "b", false, "")
}
