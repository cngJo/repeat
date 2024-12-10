package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// Ensure there are at least two arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <number> <command> [args...]")
		return
	}

	// Retrieve the first argument and convert it to a number
	firstArg := os.Args[1]
	count, err := strconv.Atoi(firstArg)
	if err != nil {
		fmt.Printf("Error: First argument must be a number. Got '%s'\n", firstArg)
		return
	}

	// Extract the command and its arguments
	command := os.Args[2]
	args := os.Args[3:]

	// Execute the command 'count' times
	for i := 0; i < count; i++ {
		// fmt.Printf("Execution %d:\n", i+1)
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}
