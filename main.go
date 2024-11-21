package main

import "fmt"

type Item struct {
	Task   string
	Status string
}

func main() {

	todoItems := map[int]Item{}

	fmt.Println("Todo items ")

	for _, item := range todoItems {
		fmt.Printf("%s\n", item)
	}
}
