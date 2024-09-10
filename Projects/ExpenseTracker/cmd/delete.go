package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "delet a expense with id",
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteFunction(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func deleteFunction(id string) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ID must be a integer")
	}

	if ID == 0 {
		return errors.New("ID must be start with 1")
	}

	if ID > len(data) {
		return errors.New("The ID is greather")
	}

	realID := ID - 1
	data = append(data[:realID], data[realID+1:]...)

	for i := range data {
		data[i].AddID(i + 1)
	}

	err = util.SaveData(data)
	if err != nil {
		return err
	}

	fmt.Printf("Expense delete succesfully (ID: %s)\n", id)
	return nil
}
