package cmd

import "fmt"

func Help() string {
	fmt.Println("Usage: ")
	fmt.Println("github-activity <username>")
	fmt.Println(" ")
	fmt.Println("Example: ")
	fmt.Println("github-activity drizlye0")
	return ""
}
