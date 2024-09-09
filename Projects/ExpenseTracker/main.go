package main

import (
	"os"

	"github.com/drizlye0/expensetracker/cmd"
	"github.com/drizlye0/expensetracker/util"
)

func main() {
	err := util.CheckFile()
	if err != nil {
		os.Exit(1)
		return
	}
	cmd.Execute()
}
