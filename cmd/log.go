package cmd

import (
	"github.com/redradrat/hourglass/lib"
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
	entry, err := lib.ParseLogEntry(args)
	if err != nil {
		return err
	}
	err = lib.WriteToLib(entry)
	if err != nil {
		return err
	}
	return nil
}
