package cmd

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/drizlye0/tasktracker/util"
)

func Delete(index string) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	taskID, err := strconv.Atoi(index)
	if err != nil {
		return err
	}

	data = append(data[:taskID], data[taskID+1:]...)

	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile("Tasks.json", bytes, 0644)
	if err != nil {
		return err
	}

	err = List()
	if err != nil {
		return err
	}

	return nil
}
