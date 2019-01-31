package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/jwalterweatherman"
)

// LogEntry represents an entry inside the Log
type LogEntry struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Project   string    `json:"project"`
	Message   string    `json:"message"`
}

// Log represents the collection of all LogEntry.
type Log struct {
	LogEntries map[uuid.UUID]LogEntry `json:"logEntries"`
}

// ParseLogEntry parses an array of strings for a log entry definition
func ParseLogEntry(args []string) (LogEntry, error) {
	var entry LogEntry
	durs := strings.Split(args[0], "-")
	if len(durs) != 2 {
		return LogEntry{}, fmt.Errorf("(%s) is not a valid time format", args[0])
	}

	// UGLY AF Time parsing... Let's redo that.
	// Probably we should start processing elements in reverse and strip message and project, leaving only the time
	// definition. Then parse time formats.
	year, month, day := time.Now().Date()
	starttime, err := time.Parse(ShortDate+" 15:04", strconv.Itoa(year)+"/"+fmt.Sprintf("%02d", month)+"/"+fmt.Sprintf("%02d", day)+" "+durs[0])
	if err != nil {
		starttime, err = time.Parse(ShortDate+" 15:04", durs[0])
		if err != nil {
			return LogEntry{}, err
		}
	}

	endtime, err := time.Parse(ShortDate+" 15:04", strconv.Itoa(year)+"/"+fmt.Sprintf("%02d", month)+"/"+fmt.Sprintf("%02d", day)+" "+durs[1])
	if err != nil {
		endtime, err = time.Parse(ShortDate+" 15:04", durs[1])
		if err != nil {
			return LogEntry{}, err
		}
	}

	switch len(args) {
	case 2:
		message := args[1]

		entry = LogEntry{StartTime: starttime, EndTime: endtime, Project: "default", Message: message}

	case 3:
		message := args[2]
		project := args[1]

		entry = LogEntry{StartTime: starttime, EndTime: endtime, Project: project, Message: message}
	}

	return entry, nil
}

// WeekFilterLog filters a given log for a specific week
func WeekFilterLog(filter string, log *Log) {
	jwalterweatherman.ERROR.Println("Sorry man! :( Week period filtering is not yet implemented!")
	os.Exit(1)
	/*processedEntries := make(map[uuid.UUID]LogEntry)
	finalLog := &Log{}
	*finalLog = *log
	for i, entry := range processedEntries {
		...
	}*/
}

// ProjectFilterLog filters a given log for specific projects
func ProjectFilterLog(filter []string, log *Log) *Log {
	processedEntries := log.LogEntries
	finalLog := &Log{}
	*finalLog = *log
	for i, entry := range processedEntries {
		hit := false
		for _, projectFilter := range filter {
			if entry.Project == projectFilter {
				hit = true
			}
		}
		if !hit {
			delete(processedEntries, i)
		}
	}
	finalLog.LogEntries = processedEntries
	return finalLog
}
