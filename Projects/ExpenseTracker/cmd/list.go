package cmd

import (
	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
}

func list() error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	return nil
}
