package lib

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func PrintLogToStdOut(l Log) {
	data := [][]string{}

	var prevDay int
	var prevWeek int
	var prevMonth int
	var prevYear int
	for _, v := range l.LogEntries {
		day := ""
		week := ""
		month := ""
		year := ""

		entryTime := v.StartTime.Truncate(24 * time.Hour)
		curDay := entryTime.Day()
		_, curWeek := entryTime.ISOWeek()
		curMonth := int(entryTime.Month())
		curYear := entryTime.Year()

		if curDay != prevDay {
			day = strconv.Itoa(curDay)
		}
		if curWeek != prevWeek {
			week = strconv.Itoa(curWeek)
		}
		if curMonth != prevMonth {
			month = strconv.Itoa(curMonth)
		}
		if curYear != prevYear {
			year = strconv.Itoa(curYear)
		}

		data = writeDataLine(data, year, month, week, day, v.StartTime.Format("15:04:05"), v.EndTime.Format("15:04:05"), v.Project, v.Message)

		prevDay = entryTime.Day()
		_, prevWeek = entryTime.ISOWeek()
		prevMonth = int(entryTime.Month())
		prevYear = entryTime.Year()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Year", "Month", "Week", "Day", "Start", "End", "Project", "Message"})
	// table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(false) // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func writeDataLine(data [][]string, year string, month string, week string, day string, starttime string, endtime string, project string, message string) [][]string {
	newLine := []string{year, month, week, day, starttime, endtime, project, message}
	return append(data, newLine)
}
