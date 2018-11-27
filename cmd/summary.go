package cmd

import (
	"github.com/redradrat/hourglass/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Aliases: []string{"sum"},
	Short:   "Sum up logs ",
	Args:    cobra.NoArgs,
	Long:    `Display a summary of stored logs`,
	RunE:    summary,
}

func summary(cmd *cobra.Command, args []string) error {
	log, err := lib.ReadFromLib()
	if err != nil {
		return err
	}
	lib.PrintLogToStdOut(log)
	return nil
}
