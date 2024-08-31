package cmd

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/drizlye0/tasktracker/util"
)

func Mark(taskID, status string) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(taskID) 
	if err != nil {
		return err
	}

	data[id].Status = status

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
