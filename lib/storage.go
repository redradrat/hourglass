package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	uuid "github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
)

const (
	defaultStorageDirName  = ".hourglass/"
	defaultStorageFileName = "store"
	defaultStorageFile     = defaultStorageDirName + defaultStorageFileName
)

func ReadFromLib() (Log, error) {
	err := ensureStorageFile()
	if err != nil {
		return Log{}, err
	}

	home, err := homedir.Dir()
	if err != nil {
		return Log{}, err
	}

	dbFile := home + "/" + defaultStorageFile
	in, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return Log{}, err
	}

	log := Log{}
	json.Unmarshal(in, &log)

	return log, nil
}

func DeleteFromLib(id string) error {
	err := ensureStorageFile()
	if err != nil {
		return err
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	dbFile := home + "/" + defaultStorageFile
	oldStoreCont, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return err
	}

	var log Log
	json.Unmarshal(oldStoreCont, &log)

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if _, ok := log.LogEntries[uid]; !ok {
		return fmt.Errorf("No entry found with given log id [%s]", id)
	}

	delete(log.LogEntries, uid)

	output, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(dbFile, output, 0644)

	return nil
}

func WriteToLib(entry LogEntry) error {
	err := ensureStorageFile()
	if err != nil {
		return err
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	dbFile := home + "/" + defaultStorageFile
	oldStoreCont, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return err
	}

	var log Log
	json.Unmarshal(oldStoreCont, &log)

	uid := uuid.New()
	for _, hasKey := log.LogEntries[uid]; hasKey; {
		uid = uuid.New()
		if err != nil {
			return err
		}
	}

	if log.LogEntries == nil {
		log.LogEntries = make(map[uuid.UUID]LogEntry)
	}
	log.LogEntries[uid] = entry
	output, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(dbFile, output, 0644)

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
