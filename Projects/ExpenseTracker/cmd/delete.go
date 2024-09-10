package cmd

import (
	"errors"
	"fmt"

	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var id int

var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "delet a expense with id",
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteFunction(id)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	DeleteCommand.Flags().IntVarP(&id, "id", "i", 0, "Indicate the id")
}

func deleteFunction(id int) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("ID must be start with 1")
	}

	if id > len(data) {
		return errors.New("The ID is greather")
	}

	realID := id - 1
	data = append(data[:realID], data[realID+1:]...)

	for i := range data {
		data[i].AddID(i + 1)
	}

	err = util.SaveData(data)
	if err != nil {
		return err
	}

	fmt.Printf("Expense delete succesfully (ID: %d)\n", id)
	return nil
}
