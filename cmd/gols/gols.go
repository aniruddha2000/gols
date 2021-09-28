package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func listFileDirectory(path string, all bool) {
	//fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		if all != true && strings.HasPrefix(name, ".") {
			continue
		} else {
			dirOrFile, err := os.Stat(name)
			if err != nil {
				log.Fatal(err)
			}
			if dirOrFile.IsDir() {
				//	color.Blue("%s ", name)
				dirColor := color.New(color.FgCyan, color.Bold)
				dirColor.Printf("%s  ", name)
			} else {
				fmt.Printf("%s  ", name)
			}
		}
	}
	fmt.Println()
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
					listFileDirectory(path, false)
				}
			} else if c.Args().Get(0) != "" {
				path := c.Args().Get(0)
				listFileDirectory(path, false)
			} else {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal("Failed to get the directory name: %s", err)
				}
				listFileDirectory(path, false)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Show all files",
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 1 {
						for i := 0; i < c.Args().Len(); i++ {
							path := c.Args().Get(i)
							listFileDirectory(path, true)
						}
					} else if c.Args().Get(0) != "" {
						path := c.Args().Get(0)
						listFileDirectory(path, true)
					} else {
						path, err := os.Getwd()
						if err != nil {
							log.Fatal("Failed to get the directory name: %s", err)
						}
						listFileDirectory(path, true)
					}
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
