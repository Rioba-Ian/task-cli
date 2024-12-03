package item

import (
	"errors"
	"fmt"

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

var listSecondArgs = []string{"done", "todo", "in-progress"}

func (args *Param) ListItems() (*Items, error) {
	list, secondListArg := args.s[0], ""

	if len(args.s) > 1 {
		secondListArg = args.s[1]
	}

	if list == "list" && len(args.s) == 1 {
		// do the read all operation
		return nil, nil
	}

	if !helpers.ExistsInListCmds(listSecondArgs, secondListArg) {
		return nil, errors.New("second argument used does not exist, use todo, done or in-progress")
	}

	return nil, nil
}

func (items *Items) AddItem(newTodo Item) (int, error) {
	items.Items = append(items.Items, newTodo)

	fmt.Println(newTodo, "item >>>>>>")

	return newTodo.ID, nil
}

/*
case len==1 : fmt.Println("You set in one argument:", userArgs)
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

case len == 2:
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

*/
