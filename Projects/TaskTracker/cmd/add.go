package cmd

import (
	"encoding/json"
	"os"

	"github.com/drizlye0/tasktracker/types"
	"github.com/drizlye0/tasktracker/util"
)

func Add(description string) error {
	filename := "Tasks.json"
	err := util.CheckFile(filename)
	if err != nil {
		return err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	data := []types.Task{}

	json.Unmarshal(file, &data)


	task := &types.Task{
		Description: description,
		Status: "Todo",
	}

	data = append(data, *task)
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

