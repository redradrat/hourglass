package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/redradrat/hourglass/lib"
)

const (
	defaultStorageDirName  = ".hourglass/"
	defaultStorageFileName = "store"
	defaultStorageFile     = defaultStorageDirName + defaultStorageFileName
)

type LocalStorage struct{}

func (ls LocalStorage) GetLog() (*lib.Log, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	dbFile := home + "/" + defaultStorageFile
	log, err := getLogFromDb(dbFile)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (ls LocalStorage) WriteToLog(l lib.LogEntry) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	dbFile := home + "/" + defaultStorageFile
	log, err := getLogFromDb(dbFile)
	if err != nil {
		return err
	}

	uid := uuid.New()
	for _, hasKey := log.LogEntries[uid]; hasKey; {
		uid = uuid.New()
		if err != nil {
			return err
		}
	}

	if log.LogEntries == nil {
		log.LogEntries = make(map[uuid.UUID]lib.LogEntry)
	}
	log.LogEntries[uid] = l
	if err = writeLogToFile(log, dbFile); err != nil {
		return err
	}

	return nil
}

func (ls LocalStorage) DeleteFromLog(id string) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	dbFile := home + "/" + defaultStorageFile
	log, err := getLogFromDb(dbFile)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if _, ok := log.LogEntries[uid]; !ok {
		return fmt.Errorf("No entry found with given log id [%s]", id)
	}

	fmt.Println()
	delete(log.LogEntries, uid)

	if err = writeLogToFile(log, dbFile); err != nil {
		return err
	}
	return nil
}

func ensureStorageFile() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	_, err = os.Stat(home + "/" + defaultStorageFile)
	if os.IsNotExist(err) {
		err := os.MkdirAll(home+"/"+defaultStorageDirName, 0755)
		if err != nil {
			return err
		}
		_, err = os.Create(home + "/" + defaultStorageFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func getLogFromDb(dbFile string) (*lib.Log, error) {
	err := ensureStorageFile()
	if err != nil {
		return nil, err
	}

	in, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return nil, err
	}

	log := lib.Log{}
	json.Unmarshal(in, &log)
	return &log, nil
}

func writeLogToFile(log *lib.Log, dbFile string) error {
	output, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(dbFile, output, 0644)

	return nil
}
