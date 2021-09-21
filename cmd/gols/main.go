package main

import (
	"fmt"
	"log"
	"os"
	//	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//	exPath := filepath.Dir(ex)
	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, f := range list {
		fmt.Println(f)
	}
}
