package lib

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	// ShortTime represents a short time notation
	ShortTime = "15:04:05"
	// ShortDate represents a short date notation
	ShortDate = "2006/01/02"
)

// SumEntry represents a single log entry in the summary
type SumEntry struct {
	Date      time.Time
	StartTime string
	EndTime   string
	Duration  string
	Project   string
	Message   string
	ID        string
}

// Day returns the day of a SumEntry
func (se SumEntry) Day() int {
	return se.Date.Day()
}

// Week returns the week of a SumEntry
func (se SumEntry) Week() int {
	_, week := se.Date.ISOWeek()
	return week
}

// Month returns the month of a SumEntry
func (se SumEntry) Month() int {
	return int(se.Month())
}

// Year returns the year of a SumEntry
func (se SumEntry) Year() int {
	return se.Year()
}

// ShortDate returns a string with the short date notation of a SumEntry
func (se SumEntry) ShortDate() string {
	return strconv.Itoa(se.Year()) + "-" + strconv.Itoa(se.Month()) + "-" + strconv.Itoa(se.Day())
}

// Summary represents the collection of SumEntries, comprising the entirety of log entries for a given period of time
type Summary struct {
	Entries []SumEntry
}

// PrintLogToStdOut compiles a Summary from a given Log, and prints it to STDOUT
func PrintLogToStdOut(l *Log, showIDs bool) {
	output := Summary{}
	for id, v := range l.LogEntries {
		d := v.EndTime.Sub(v.StartTime).Round(time.Second)
		h := d / time.Hour
		d -= h * time.Hour
		m := d / time.Minute
		d -= m * time.Minute
		s := d / time.Second
		prettyDuration := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

		se := SumEntry{Date: v.StartTime, StartTime: v.StartTime.Format(ShortTime), EndTime: v.EndTime.Format(ShortTime), Duration: prettyDuration, Project: v.Project, Message: v.Message, ID: id.String()}
		output.Entries = append(output.Entries, se)
	}

	table := compileSortedOutput(output.Entries, showIDs)

	table.Render()
}

func compileSortedOutput(in []SumEntry, showIDs bool) *tablewriter.Table {
	var data [][]string
	sort.SliceStable(in, func(i, j int) bool { return in[i].Date.Before(in[j].Date) })

	// TODO: Here, after sort, we could split into Days/Weeks for Summary Total Accumulation...
	// For now we're just using Automerge by tablewriter. We could use WXX - WYY.
	// Or: IF day.changes, inject summary line...

	for _, v := range in {
		date := v.Date.Format(ShortDate)
		week := "W" + strconv.Itoa(v.Week())
		line := []string{week, date, v.Project, v.Message, v.StartTime, v.EndTime, v.Duration}
		if showIDs {
			line = append(line, v.ID)
		}
		data = append(data, line)
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"Week", "Date", "Project", "Message", "Start", "End", "Time"}
	if showIDs {
		header = append(header, "ID")
	}
	table.SetHeader(header)
	// table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(false) // Set Border to false
	table.SetAutoMergeCells(true)
	table.AppendBulk(data) // Add Bulk Data

	return table
}
