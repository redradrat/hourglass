package cmd

import (
	"github.com/redradrat/hourglass/lib/storage"
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
	// Get backend as configured
	be, err := storage.GetBackend()
	if err != nil {
		return err
	}

	// Delete Log Entry
	if err := be.DeleteFromLog(args[0]); err != nil {
		return err
	}
	return nil
}
