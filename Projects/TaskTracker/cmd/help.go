package cmd

import "fmt"

func Help() {
	fmt.Println("Available Commands:")
	fmt.Println("add 'this is a task' -> Add a Task")
	fmt.Println("list                 -> List all the Tasks")
	fmt.Println("delete id            -> Delete a Task with id")
	fmt.Println("mark-in-progress     -> Mark a task as in progress")
	fmt.Println("mark-done            -> Mark a task done")
}
