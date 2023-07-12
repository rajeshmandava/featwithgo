package main

import (
	"fmt"
	"os"
)

func getPassArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func main() {
	minArgs := 3
	args := getPassArgs(minArgs)
	fmt.Println("Arguments:", args)
}

