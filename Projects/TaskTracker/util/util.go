package util

import (
	"encoding/json"
	"os"

	"github.com/drizlye0/tasktracker/types"
)

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetData() (data []types.Task, err error) {	
	filename := "Tasks.json"
	err = CheckFile(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data = []types.Task{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

