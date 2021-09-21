package main

import (
	"fmt"
	"log"
	"os"
)

func listDirectory(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		fmt.Println(name)
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("3")
		log.Fatal("Can't be more than 1 argument")
	} else if len(os.Args) > 1 {
		fmt.Println("1")
		path := os.Args[1]
		listDirectory(path)
	} else {
		fmt.Println("2")
		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Failed to get the directory name: %s", err)
		}
		fmt.Println(path)
		listDirectory(path)
	}
}
