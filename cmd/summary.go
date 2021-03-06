package cmd

import (
	"github.com/redradrat/hourglass/lib"
	"github.com/redradrat/hourglass/lib/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().BoolVar(&showIDs, "id", false, "display IDs for entries in summary")
	summaryCmd.Flags().StringSliceVarP(&projectFilter, "project", "p", []string{}, "filter for specific project")
	summaryCmd.Flags().StringVarP(&weekFilter, "week", "w", "", "filter for specific weeks [W22 = week 22, -2 (current week + last 2 weeks]")
}

var (
	summaryCmd = &cobra.Command{
		Use:     "summary",
		Aliases: []string{"sum"},
		Short:   "Sum up logs ",
		Long:    `Display a summary of stored logs`,
		RunE:    summary,
	}

	showIDs       bool
	projectFilter []string
	weekFilter    string
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

	if len(projectFilter) != 0 {
		lib.ProjectFilterLog(projectFilter, log)
	}

	if weekFilter != "" {
		lib.WeekFilterLog(weekFilter, log)
	}

	// Print log to stdout
	lib.PrintLogToStdOut(log, showIDs)
	return nil
}
