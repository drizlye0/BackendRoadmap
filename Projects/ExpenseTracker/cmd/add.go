package cmd

import (
	"fmt"

	"github.com/drizlye0/expensetracker/types"
	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var (
	description string
	amount      float32
)

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add item to Expense",
	Run: func(cmd *cobra.Command, args []string) {
		err := addItem(description, amount)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	AddCommand.Flags().StringVarP(&description, "description", "d", "", "description for a expense")
	AddCommand.Flags().Float32VarP(&amount, "amount", "a", 0, "amount for a expense")
}

func addItem(description string, amount float32) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	item := &types.Expense{
		ID:          0,
		Description: description,
		Amount:      amount,
	}

	data = append(data, *item)

	for i := range data {
		data[i].AddID(i + 1)
	}

	err = util.SaveData(data)
	if err != nil {
		return err
	}
	fmt.Println("Item added succesfully")
	return nil
}
