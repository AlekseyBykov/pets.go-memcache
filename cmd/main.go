package main

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-memcache/internal/cache/storage"
	"time"
)

func main() {
	c := storage.NewCache(5 * time.Second)

	err := c.Set("key1", "value1", 5*time.Second)
	if err != nil {
		fmt.Printf("Error adding key1: %v\n", err)
	} else {
		fmt.Println("Added key1 -> value1")
	}

	val, err := c.Get("key1")
	if err != nil {
		fmt.Printf("Error retrieving key1: %v\n", err)
	} else {
		fmt.Printf("Found key1: %v\n", val)
	}

	fmt.Println("Waiting 6 seconds to check TTL expiration...")
	time.Sleep(6 * time.Second)

	val, err = c.Get("key1")
	if err != nil {
		fmt.Printf("Error retrieving key1 after TTL expired: %v\n", err)
	} else {
		fmt.Printf("Found key1: %v\n", val)
	}

	err = c.Set("key2", "value2", 5*time.Second)
	if err != nil {
		fmt.Printf("Error adding key2: %v\n", err)
	} else {
		fmt.Println("Added key2 -> value2")
	}

	err = c.Delete("key2")
	if err != nil {
		fmt.Printf("Error deleting key2: %v\n", err)
	} else {
		fmt.Println("Deleted key2")
	}

	val, err = c.Get("key2")
	if err != nil {
		fmt.Printf("key2 successfully deleted: %v\n", err)
	} else {
		fmt.Printf("Found key2: %v\n", val)
	}
}
