package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type Items struct {
	Items []Item `json:"items"`
}

func main() {
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

	fmt.Println("Todo items ", todoItems)

	for _, item := range todoItems.Items {
		fmt.Printf("Item %+v\n", item)
	}

	defer todosJsonFile.Close()
}
