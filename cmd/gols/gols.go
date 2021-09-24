package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func listDirectory(path string) {
	//fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		if strings.HasPrefix(name, ".") {
			continue
		} else {
			fmt.Println(name)
		}
	}
}

func main() {
	//var listAllFile string

	app := &cli.App{
		Name:  "gols",
		Usage: "A fun ls command using go",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 1 {
				for i := 0; i < c.Args().Len(); i++ {
					path := c.Args().Get(i)
					listDirectory(path)
				}
			} else if c.Args().Get(0) != "" {
				path := c.Args().Get(0)
				listDirectory(path)
			} else {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal("Failed to get the directory name: %s", err)
				}
				listDirectory(path)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Show all files",
				Action: func(c *cli.Context) error {
					fmt.Println("joololo")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
