package cmd

import (
	"github.com/redradrat/hourglass/lib"
	"github.com/redradrat/hourglass/lib/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().BoolVar(&showIDs, "id", false, "Display IDs for entries in summary")
}

var (
	summaryCmd = &cobra.Command{
		Use:     "summary",
		Aliases: []string{"sum"},
		Short:   "Sum up logs ",
		Long:    `Display a summary of stored logs`,
		RunE:    summary,
	}

	showIDs bool
)

func summary(cmd *cobra.Command, args []string) error {
	// Get backend as configured
	be, err := storage.GetBackend()
	if err != nil {
		return err
	}

	// Get log from backend
	log, err := be.GetLog()
	if err != nil {
		return err
	}

	// Print log to stdout
	lib.PrintLogToStdOut(log, showIDs)
	return nil
}
