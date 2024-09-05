package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drizlye0/GithubAPI-CLI/cmd"
)

func main() {
	helpCommand := flag.Bool("help" , false , "a bool")
	flag.Parse()
	if *helpCommand{
		cmd.Help()
		return
	}

	argsLen := len(os.Args[1:])
	if argsLen != 1 {
		fmt.Println("username arg is neccesary")
		os.Exit(1)
	}
	user := os.Args[1]

	githubEndpoint := "https://api.github.com/users/" + user + "/events"
	err := cmd.RequestAPI(githubEndpoint)
	if err != nil {
		log.Fatal(err)
	}
}
