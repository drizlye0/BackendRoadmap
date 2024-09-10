package cmd

import (
	"fmt"
	"os"

	"github.com/drizlye0/expensetracker/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all Expenses",
	Run: func(cmd *cobra.Command, args []string) {
		err := list()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func list() error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	if len(data) == 0 {
		fmt.Println("Expenses are empty :)")
		return nil
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Date", "Description", "Amount"})
	for _, v := range data {
		t.AppendRow(table.Row{v.ID, v.Date, v.Description, fmt.Sprint("$", v.Amount)})
	}
	t.Render()

	return nil
}
