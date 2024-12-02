package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

var listSecondArgs = []string{"done", "todo", "in-progress"}

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

	for _, item := range todoItems.Items {
		fmt.Printf("Items %+v\n", item)
	}

	defer todosJsonFile.Close()

	userArgs.s = os.Args[1:]

	switch len(userArgs.s) {
	case 1:
		fmt.Println("You set in one argument:", userArgs)
		cmd, err := userArgs.ParseArgs()

		if err != nil {
			fmt.Printf("%s", err)
		}

		if helpers.CompareStrings(cmd[0], "list") {

			results, err := userArgs.ListItems(&todoItems)

			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println(results, "results here")
		}
	case 2:
		// fmt.Println("You entered two arguments", userArgs)
		_, err := userArgs.ParseArgs()

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		results, err := userArgs.ListItems(&todoItems)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(results, "results here")

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

func (items *Items) GetAllItems() Items {
	return *items
}

func GiveCommands() {
	cmds := "\nYou haven't entered any commands: List of commands\n" +
		"--------------------------------------------------\n" +
		"1. Add a task\n  task-cli add 'Buy groceries' --use this to add an item\n" +
		"  #Output: Task added successfully (ID:1)\n\n"
	fmt.Println(cmds)
}

func (args *Param) ParseArgs() ([]string, error) {
	allCommands := make(map[string][]string)

	allCommands["list"] = append(allCommands["list"], "list", "list done", "list todo", "list in-progress")

	commands := []string{"list", "add", "update", "delete", "mark-in-progress", "mark-done", "in-progress", "done", "todo", "in-progress"}
	var trimmedArgs Param

	if len(args.s) == 3 {
		trimmedArgs.s = args.s[:1]
		trimmedArgs.ParseArgs()
	}

	if len(args.s) == 1 {
		if helpers.Contains(commands, args.s[0]) {
			firstCmd := strings.Split(args.s[0], " ")
			return firstCmd, nil
		}
		return args.s, errors.New("you didn't enter a valid command")
	} else if len(args.s) == 2 {
		first, second := args.s[0], args.s[1]
		var combined []string

		if helpers.Contains(commands, first) && helpers.Contains(commands, second) {
			combined = append(combined, first, second)
			return combined, nil
		}

		return nil, errors.New("you might have entered one or two invalid commands ->")

	}
	return nil, nil
}

func (args *Param) ListItems(items *Items) (*Items, error) {
	list, secondListArg := args.s[0], ""

	if len(args.s) > 1 {
		secondListArg = args.s[1]
	}

	if list == "list" && len(args.s) == 1 {
		// do the read all operation
		return items, nil
	}

	if !helpers.ExistsInListCmds(listSecondArgs, secondListArg) {
		return nil, errors.New("second argument used does not exist, use todo, done or in-progress")
	}

	if helpers.CompareStrings("done", secondListArg) {
		fmt.Println("should return all done items")
		return items, nil
	}
	return nil, nil
}
