package cmd

import (
	"github.com/redradrat/hourglass/lib"
	"github.com/redradrat/hourglass/lib/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log time ",
	Args:  cobra.RangeArgs(1, 4),
	Long:  `Log a time entry`,
	RunE:  log,
}

func log(cmd *cobra.Command, args []string) error {
	// Get backend as configured
	be, err := storage.GetBackend()
	if err != nil {
		return err
	}

	// Parse arguments to Log Entry
	entry, err := lib.ParseLogEntry(args)
	if err != nil {
		return err
	}

	// Write parsed entry to Log backend
	err = be.WriteToLog(entry)
	if err != nil {
		return err
	}
	return nil
}
