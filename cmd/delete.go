package cmd

import (
	"github.com/redradrat/hourglass/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var (
	deleteCmd = &cobra.Command{
		Use:     "delete [ID]",
		Aliases: []string{"del"},
		Short:   "Delete a log",
		Args:    cobra.ExactArgs(1),
		Long:    `Deletes the log corresponding to the given log ID`,
		RunE:    deleteLog,
	}
)

func deleteLog(cmd *cobra.Command, args []string) error {
	if err := lib.DeleteFromLib(args[0]); err != nil {
		return err
	}
	return nil
}
