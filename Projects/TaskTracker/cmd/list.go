package cmd

import (
	"os"

	"github.com/drizlye0/tasktracker/types"
	"github.com/drizlye0/tasktracker/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func List(status ...string) error {
	data, err := util.GetData()
	if err != nil {
		return err
	}

	if status != nil {
		filterData := filterList(data, status)
		err = printTable(filterData)
		if err != nil {
			return err
		}
		return nil
	}

	hash := map[int]types.Task{}
	for i, v := range data {
		hash[i] = types.Task{v.Description, v.Status}
	}

	err = printTable(hash)
	if err != nil {
		return err
	} 
	
	return nil
}


func filterList(data []types.Task, status []string) (filterData map[int]types.Task) {
	capitalize := cases.Title(language.English, cases.NoLower).String(status[0])
	hash := make(map[int]types.Task)
	for i, v := range data {
		if v.Status == capitalize {
			hash[i] = types.Task{v.Description, v.Status}
		}
	}

	return hash
}

func printTable(data map[int]types.Task) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Task", "Status"})

		for i, task := range data {
		 t.AppendRows([]table.Row{
		 	{i, task.Description, task.Status},
		 })
	
	}
	t.SetStyle(table.StyleLight)
	t.Render()
	return nil
}
