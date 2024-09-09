package util

import (
	"encoding/json"
	"os"

	"github.com/drizlye0/expensetracker/types"
)

var filename = "Expenses.json"

func CheckFile() error {
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetData() (data []types.Expense, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}


	json.Unmarshal(file,  &data)

	return data, nil
}

func SaveData(data []types.Expense) error {
	bytes, err := json.MarshalIndent(data,  "" , "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename,  bytes,  0644)
	if err != nil {
		return err
	}

	return nil
}
