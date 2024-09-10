package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/drizlye0/expensetracker/util"
	"github.com/spf13/cobra"
)

var month int

var SummaryCommand = &cobra.Command{
	Use:   "summary",
	Short: "calculate the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		if month != 0 {
			s, err := monthSummary(month)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(s)
			return
		}

		s, err := summary()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(s)
	},
}

func init() {
	SummaryCommand.Flags().IntVarP(&month, "month", "m", 0, "Specify the total amount per month")
}

func summary() (s string, err error) {
	data, err := util.GetData()
	if err != nil {
		return "", err
	}

	var sum float64

	for i := range data {
		sum += data[i].Amount
	}

	s = fmt.Sprintf("Total expenses: $%f", sum)

	return s, nil
}

func monthSummary(month int) (s string, err error) {
	if month < 1 {
		return "", errors.New("The month must be between 1 and 12")
	}
	if month > 12 {
		return "", errors.New("The month must be between 1 and 12")
	}

	data, err := util.GetData()
	if err != nil {
		return "", err
	}

	var sum float64

	for i := range data {
		if data[i].Month == month {
			sum += data[i].Amount
		}
	}
	m := time.Month(month)

	s = fmt.Sprintf("Total expenses for %s: $%f", m, sum)

	return s, nil
}
