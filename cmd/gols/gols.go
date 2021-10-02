package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func listFileDirectory(path string, all bool) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	for _, name := range list {
		if name.IsDir() {
			dirColor := color.New(color.FgCyan, color.Bold)
			dirColor.Printf("%s  ", name.Name())
		} else {
			fmt.Printf("%s  ", name.Name())
		}
	}
	fmt.Println()
}

func printPermissions(path string) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to open directory: %s", err)
	}
	for _, name := range list {
		if name.IsDir() {
			dirColor := color.New(color.FgCyan, color.Bold)
			dirColor.Printf("%s %04o %s\n", name.Mode(), name.Mode().Perm(), name.Name())
		} else {
			fmt.Printf("%s %04o %s\n", name.Mode(), name.Mode().Perm(), name.Name())
		}
	}
	fmt.Println()
}

func main() {
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
			{
				Name:    "long",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 1 {
						for i := 0; i < c.Args().Len(); i++ {
							path := c.Args().Get(i)
							printPermissions(path)
						}
					} else if c.Args().Get(0) != "" {
						path := c.Args().Get(0)
						printPermissions(path)
					} else {
						path, err := os.Getwd()
						if err != nil {
							log.Fatal("Failed to get the directory name: %s", err)
						}
						printPermissions(path)
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
