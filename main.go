package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Rioba-Ian/task-cli/helpers"
)

type Item struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type Items struct {
	Items []Item `json:"items"`
}

type Param struct {
	s []string
}

func main() {
	var userArgs Param
	todosJsonFile, err := os.Open("items.json")

	if err != nil {
		log.Fatalf("error in loading to do tasks %s", err)
	}

	byteValues, _ := io.ReadAll(todosJsonFile)

	var todoItems Items

	err = json.Unmarshal(byteValues, &todoItems)

	if err != nil {
		log.Fatalf("error in decoding json values from bytes %s", err)
	}

	fmt.Println("Todo items ", todoItems.Items)

	for _, item := range todoItems.Items {
		fmt.Printf("Item %+v\n", item)
	}

	defer todosJsonFile.Close()

	userArgs.s = os.Args[1:]

	fmt.Println(userArgs)
	switch len(userArgs.s) {
	case 1:
		fmt.Println("You set in one argument:", userArgs)
		cmd, err := userArgs.ParseArgs()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Println(cmd)
	// case 2:
	// 	fmt.Println("You entered two arguments", userArgs)
	// 	fmt.Printf("first one: %s, second %s\n", userArgs.s[0], userArgs.s[1])
	default:
		GiveCommands()
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("The time now is ", currentTime)

}

func (items *Items) GetItems(id int) (*Item, error) {
	for _, item := range items.Items {
		if item.ID == id {
			return &item, nil
		}
	}

	return nil, errors.New("failed to get the todo item")
}

func GiveCommands() {
	cmds := "\nYou haven't entered any commands: List of commands\n" +
		"--------------------------------------------------\n" +
		"1. Add a task\n  task-cli add 'Buy groceries' --use this to add an item\n" +
		"  #Output: Task added successfully (ID:1)\n\n"
	fmt.Println(cmds)
}

func (args *Param) ParseArgs() (string, error) {
	commands := []string{"list", "add", "update", "delete", "mark-in-progress", "mark-done", "in-progress"}
	var trimmedArgs Param

	if len(args.s) == 3 {
		trimmedArgs.s = args.s[:1]
		trimmedArgs.ParseArgs()
	}

	if len(args.s) == 1 {
		if helpers.Contains(commands, args.s[0]) {
			return args.s[0], nil

		}
		return args.s[0], errors.New("you didn't enter a valid command")
	} else if len(args.s) == 2 {
		for k := range args.s {
			for n := range commands {
				if args.s[k] != commands[n] {
					return args.s[k], errors.New("you didn't enter a valid command")
				}
			}
		}
	}
	return "", nil
}
