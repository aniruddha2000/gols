package main

import (
	"fmt"
	"io/ioutil"
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

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
