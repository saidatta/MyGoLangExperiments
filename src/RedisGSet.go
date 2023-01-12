package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Add an element to the G-Set
	err := addToGSet(client, 1)
	if err != nil {
		fmt.Println(err)
	}

	// Retrieve the current set of elements from the G-Set
	elements, err := getGSet(client)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(elements)
}
