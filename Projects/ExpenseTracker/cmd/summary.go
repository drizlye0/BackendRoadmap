package cmd

import (
	"fmt"

	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var SummaryCommand = &cobra.Command{
	Use:   "summary",
	Short: "calculate the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		sum, err := summary()
		if err != nil {
			fmt.Println(err)
			return
		}
		s := fmt.Sprintf("$%f", sum)
		fmt.Println(s)
	},
}

func summary() (sum float64, err error) {
	data, err := util.GetData()
	if err != nil {
		return 0, err
	}

	for i := range data {
		sum += data[i].Amount
	}

	return sum, nil
}
