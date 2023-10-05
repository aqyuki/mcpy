package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "mcpy",
	Short: "copy file which is configuration of minecraft",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("2 arguments were needed")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
