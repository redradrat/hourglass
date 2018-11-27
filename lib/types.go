package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type LogEntry struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Project   string    `json:"project"`
	Message   string    `json:"message"`
}

type Log struct {
	LogEntries map[uuid.UUID]LogEntry `json:"logEntries"`
}

// Parses an array of strings for a log entry definition
func ParseLogEntry(args []string) (LogEntry, error) {
	var entry LogEntry
	durs := strings.Split(args[0], "-")
	if len(durs) != 2 {
		return LogEntry{}, fmt.Errorf("(%s) is not a valid time format!", args[0])
	}

	// UGLY AF Time parsing... Let's redo that.
	// Probably we should start processing elements in reverse and strip message and project, leaving only the time
	// definition. Then parse time formats.
	year, month, day := time.Now().Date()
	starttime, err := time.Parse("2006/01/02 15:04", strconv.Itoa(year)+"/"+strconv.Itoa(int(month))+"/"+strconv.Itoa(day)+" "+durs[0])
	if err != nil {
		starttime, err = time.Parse("2006/01/02 15:04", durs[0])
		if err != nil {
			return LogEntry{}, err
		}
	}

	endtime, err := time.Parse("2006/01/02 15:04", strconv.Itoa(year)+"/"+strconv.Itoa(int(month))+"/"+strconv.Itoa(day)+" "+durs[1])
	if err != nil {
		endtime, err = time.Parse("2006/01/02 15:04", durs[1])
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
