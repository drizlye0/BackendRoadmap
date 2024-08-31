package main

import (
	"fmt"
	"os"

	"github.com/drizlye0/tasktracker/cmd"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong Flag")
		cmd.Help()
		os.Exit(1)
	}
	flag := os.Args[1]
	validate := validateAction(flag)
	if validate != 1 {
		fmt.Println("Wrong command")
		cmd.Help()
		os.Exit(1)
	}

	switch flag {
	case "help":
		{
			cmd.Help()
		}
	case "add":
		{
			task := os.Args[2]
			err := cmd.Add(task)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Task added succesfully")
		}
	case "list":
		{
			lenArgs := len(os.Args[1:])
			if lenArgs != 2 {
				err := cmd.List()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return 
			} else if lenArgs >= 3 {
				cmd.Help()
				return 
			}
			args := os.Args[2]
			allowedValues := map[string]bool {
				"done": true,
				"todo": true,
				"progress": true,
			}

			if !allowedValues[args] {
				cmd.Help()
				return
			}
			err := cmd.List(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
	}
	case "mark-done":
		{
			taskId := os.Args[2]
			err := cmd.Mark(taskId, "Done")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Task mark-done succesfully")
		}

	case "mark-in-progress":
		{
			taskID := os.Args[2]
			err := cmd.Mark(taskID, "Progress")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Task mark-in-progress succesfully")
		}
	case "delete":
		{
			taskID := os.Args[2]
			err := cmd.Delete(taskID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Task delete succesfully")
		}
	case "update":
		{
			taskID := os.Args[2]
			description := os.Args[3]
			err := cmd.Update(taskID, description)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Task update succesfully")
		}
	}
}

func validateAction(flag string) int {
	possibleFlags := []string{
		"add",
		"list",
		"delete",
		"mark-done",
		"mark-in-progress",
		"help",
		"update",
	}
	for _, action := range possibleFlags {
		if flag == action {
			return 1
		}
	}
	return 0
}
